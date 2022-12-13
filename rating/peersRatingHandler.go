package rating

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/core/check"
	logger "github.com/ElrondNetwork/elrond-go-logger"
	"github.com/ElrondNetwork/elrond-go-p2p"
	"github.com/ElrondNetwork/elrond-go-storage/types"
)

const (
	topRatedTier   = "top rated tier"
	badRatedTier   = "bad rated tier"
	defaultRating  = int32(0)
	minRating      = -100
	maxRating      = 100
	increaseFactor = 2
	decreaseFactor = -1
	minNumOfPeers  = 1
	int32Size      = 4
	minDuration    = time.Second
)

var log = logger.GetOrCreate("p2p/peersRatingHandler")

// ArgPeersRatingHandler is the DTO used to create a new peers rating handler
type ArgPeersRatingHandler struct {
	TopRatedCache              types.Cacher
	BadRatedCache              types.Cacher
	AppStatusHandler           core.AppStatusHandler
	TimeWaitingForReconnection time.Duration
	TimeBetweenMetricsUpdate   time.Duration
	TimeBetweenCachersSweep    time.Duration
}

type ratingInfo struct {
	Rating                       int32 `json:"rating"`
	TimestampLastRequestToPid    int64 `json:"timestampLastRequestToPid"`
	TimestampLastResponseFromPid int64 `json:"timestampLastResponseFromPid"`
}

type peersRatingHandler struct {
	topRatedCache              types.Cacher
	badRatedCache              types.Cacher
	removalTimestampsMap       map[string]int64
	mut                        sync.RWMutex
	ratingsMap                 map[string]*ratingInfo
	appStatusHandler           core.AppStatusHandler
	timeWaitingForReconnection time.Duration
	timeBetweenMetricsUpdate   time.Duration
	timeBetweenCachersSweep    time.Duration
	getTimeHandler             func() time.Time
	cancel                     func()
}

// NewPeersRatingHandler returns a new peers rating handler
func NewPeersRatingHandler(args ArgPeersRatingHandler) (*peersRatingHandler, error) {
	err := checkArgs(args)
	if err != nil {
		return nil, err
	}

	prh := &peersRatingHandler{
		topRatedCache:              args.TopRatedCache,
		badRatedCache:              args.BadRatedCache,
		removalTimestampsMap:       make(map[string]int64),
		appStatusHandler:           args.AppStatusHandler,
		timeWaitingForReconnection: args.TimeWaitingForReconnection,
		timeBetweenMetricsUpdate:   args.TimeBetweenMetricsUpdate,
		timeBetweenCachersSweep:    args.TimeBetweenCachersSweep,
		ratingsMap:                 make(map[string]*ratingInfo),
		getTimeHandler:             time.Now,
	}

	var ctx context.Context
	ctx, prh.cancel = context.WithCancel(context.Background())
	go prh.processLoop(ctx)

	return prh, nil
}

func checkArgs(args ArgPeersRatingHandler) error {
	if check.IfNil(args.TopRatedCache) {
		return fmt.Errorf("%w for TopRatedCache", p2p.ErrNilCacher)
	}
	if check.IfNil(args.BadRatedCache) {
		return fmt.Errorf("%w for BadRatedCache", p2p.ErrNilCacher)
	}
	if check.IfNil(args.AppStatusHandler) {
		return p2p.ErrNilAppStatusHandler
	}
	if args.TimeWaitingForReconnection < minDuration {
		return fmt.Errorf("%w for TimeWaitingForReconnection, received %f, min expected %d",
			p2p.ErrInvalidValue, args.TimeWaitingForReconnection.Seconds(), minDuration)
	}
	if args.TimeBetweenMetricsUpdate < minDuration {
		return fmt.Errorf("%w for TimeBetweenMetricsUpdate, received %f, min expected %d",
			p2p.ErrInvalidValue, args.TimeBetweenMetricsUpdate.Seconds(), minDuration)
	}
	if args.TimeBetweenCachersSweep < minDuration {
		return fmt.Errorf("%w for TimeBetweenCachersSweep, received %f, min expected %d",
			p2p.ErrInvalidValue, args.TimeBetweenCachersSweep.Seconds(), minDuration)
	}

	return nil
}

// AddPeers adds a new list of peers to the cache with rating 0
// this is called when peers list is refreshed
func (prh *peersRatingHandler) AddPeers(pids []core.PeerID) {
	if len(pids) == 0 {
		return
	}

	prh.mut.Lock()
	defer prh.mut.Unlock()

	receivedPIDsMap := make(map[string]struct{}, len(pids))
	for _, pid := range pids {
		pidBytes := pid.Bytes()
		pidString := string(pidBytes)
		receivedPIDsMap[pidString] = struct{}{}

		_, isMarkedForRemoval := prh.removalTimestampsMap[pidString]
		if isMarkedForRemoval {
			delete(prh.removalTimestampsMap, pidString)
		}

		_, found := prh.getOldRating(pidBytes)
		if found {
			continue
		}

		prh.topRatedCache.Put(pidBytes, defaultRating, int32Size)
		prh.updateRatingsMap(pid, defaultRating, 0)
	}

	prh.markInactivePIDsForRemoval(receivedPIDsMap)
}

// IncreaseRating increases the rating of a peer with the increase factor
func (prh *peersRatingHandler) IncreaseRating(pid core.PeerID) {
	prh.mut.Lock()
	defer prh.mut.Unlock()

	prh.updateRatingIfNeeded(pid, increaseFactor)
}

// DecreaseRating decreases the rating of a peer with the decrease factor
func (prh *peersRatingHandler) DecreaseRating(pid core.PeerID) {
	prh.mut.Lock()
	defer prh.mut.Unlock()

	prh.updateRatingIfNeeded(pid, decreaseFactor)
}

func (prh *peersRatingHandler) getOldRating(pid []byte) (int32, bool) {
	oldRating, found := prh.topRatedCache.Get(pid)
	if found {
		oldRatingInt, _ := oldRating.(int32)
		return oldRatingInt, found
	}

	oldRating, found = prh.badRatedCache.Get(pid)
	if found {
		oldRatingInt, _ := oldRating.(int32)
		return oldRatingInt, found
	}

	return defaultRating, found
}

func (prh *peersRatingHandler) markInactivePIDsForRemoval(receivedPIDs map[string]struct{}) {
	prh.markKeysForRemoval(prh.topRatedCache.Keys(), receivedPIDs)
	prh.markKeysForRemoval(prh.badRatedCache.Keys(), receivedPIDs)
}

func (prh *peersRatingHandler) markKeysForRemoval(cachedPIDs [][]byte, receivedPIDs map[string]struct{}) {
	removalTimestamp := prh.getTimeHandler().Add(prh.timeWaitingForReconnection).Unix()

	for _, cachedPID := range cachedPIDs {
		pidString := string(cachedPID)
		_, isPIDStillActive := receivedPIDs[pidString]
		_, alreadyMarked := prh.removalTimestampsMap[pidString]
		if !isPIDStillActive && !alreadyMarked {
			prh.removalTimestampsMap[pidString] = removalTimestamp
		}
	}
}

func (prh *peersRatingHandler) processLoop(ctx context.Context) {
	timerCachersSweep := time.NewTimer(prh.timeBetweenCachersSweep)
	timerMetricsUpdate := time.NewTimer(prh.timeBetweenMetricsUpdate)
	displayCachersTime := time.Second * 15
	timerDisplayCachers := time.NewTimer(displayCachersTime)

	defer timerCachersSweep.Stop()
	defer timerMetricsUpdate.Stop()
	defer timerDisplayCachers.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timerCachersSweep.C:
			prh.sweepCachers()
			timerCachersSweep.Reset(prh.timeBetweenCachersSweep)
		case <-timerMetricsUpdate.C:
			prh.updateMetrics()
			timerMetricsUpdate.Reset(prh.timeBetweenMetricsUpdate)
		case <-timerDisplayCachers.C:
			prh.displayCachers()
			timerDisplayCachers.Reset(displayCachersTime)
		}
	}
}

func (prh *peersRatingHandler) sweepCachers() {
	prh.mut.Lock()
	defer prh.mut.Unlock()

	for pidString, removalTimestamp := range prh.removalTimestampsMap {
		if removalTimestamp <= prh.getTimeHandler().Unix() {
			prh.removePIDFromCachers(pidString)
		}
	}
}

func (prh *peersRatingHandler) removePIDFromCachers(pidString string) {
	delete(prh.removalTimestampsMap, pidString)
	pidBytes := []byte(pidString)
	prh.topRatedCache.Remove(pidBytes)
	prh.badRatedCache.Remove(pidBytes)
	delete(prh.ratingsMap, core.PeerID(pidBytes).Pretty())
}

func (prh *peersRatingHandler) updateRatingIfNeeded(pid core.PeerID, updateFactor int32) {
	oldRating, found := prh.getOldRating(pid.Bytes())
	if !found {
		// new pid, add it with default rating
		prh.topRatedCache.Put(pid.Bytes(), defaultRating, int32Size)
		prh.updateRatingsMap(pid, defaultRating, updateFactor)
	}

	decreasingUnderMin := oldRating == minRating && updateFactor == decreaseFactor
	increasingOverMax := oldRating == maxRating && updateFactor == increaseFactor
	shouldSkipUpdate := decreasingUnderMin || increasingOverMax
	if shouldSkipUpdate {
		prh.updateRatingsMap(pid, oldRating, updateFactor)
		return
	}

	newRating := oldRating + updateFactor
	if newRating > maxRating {
		newRating = maxRating
	}

	if newRating < minRating {
		newRating = minRating
	}

	prh.updateRating(pid, oldRating, newRating)
	prh.updateRatingsMap(pid, newRating, updateFactor)
}

func (prh *peersRatingHandler) updateRating(pid core.PeerID, oldRating, newRating int32) {
	oldTier := computeRatingTier(oldRating)
	newTier := computeRatingTier(newRating)
	if newTier == oldTier {
		if newTier == topRatedTier {
			prh.topRatedCache.Put(pid.Bytes(), newRating, int32Size)
		} else {
			prh.badRatedCache.Put(pid.Bytes(), newRating, int32Size)
		}

		return
	}

	prh.movePeerToNewTier(newRating, pid)
}

func computeRatingTier(peerRating int32) string {
	if peerRating >= defaultRating {
		return topRatedTier
	}

	return badRatedTier
}

func (prh *peersRatingHandler) movePeerToNewTier(newRating int32, pid core.PeerID) {
	newTier := computeRatingTier(newRating)
	if newTier == topRatedTier {
		prh.badRatedCache.Remove(pid.Bytes())
		prh.topRatedCache.Put(pid.Bytes(), newRating, int32Size)
	} else {
		prh.topRatedCache.Remove(pid.Bytes())
		prh.badRatedCache.Put(pid.Bytes(), newRating, int32Size)
	}
}

// GetTopRatedPeersFromList returns a list of peers, searching them in the order of rating tiers
func (prh *peersRatingHandler) GetTopRatedPeersFromList(peers []core.PeerID, minNumOfPeersExpected int) []core.PeerID {
	prh.mut.Lock()
	defer prh.mut.Unlock()

	peersTopRated := make([]core.PeerID, 0)
	defer prh.displayPeersRating(&peersTopRated, minNumOfPeersExpected)

	isListEmpty := len(peers) == 0
	if minNumOfPeersExpected < minNumOfPeers || isListEmpty {
		return make([]core.PeerID, 0)
	}

	peersTopRated, peersBadRated := prh.splitPeersByTiers(peers)
	if len(peersTopRated) < minNumOfPeersExpected {
		peersTopRated = append(peersTopRated, peersBadRated...)
	}

	return peersTopRated
}

func (prh *peersRatingHandler) displayPeersRating(peers *[]core.PeerID, minNumOfPeersExpected int) {
	if log.GetLevel() != logger.LogTrace {
		return
	}

	strPeersRatings := ""
	for _, peer := range *peers {
		rating, ok := prh.topRatedCache.Get(peer.Bytes())
		if !ok {
			rating, _ = prh.badRatedCache.Get(peer.Bytes())
		}

		ratingInt, ok := rating.(int32)
		if ok {
			strPeersRatings += fmt.Sprintf("\n peerID: %s, rating: %d", peer.Pretty(), ratingInt)
		} else {
			strPeersRatings += fmt.Sprintf("\n peerID: %s, rating: invalid", peer.Pretty())
		}
	}

	log.Trace("Best peers to request from", "min requested", minNumOfPeersExpected, "peers ratings", strPeersRatings)
}

func (prh *peersRatingHandler) splitPeersByTiers(peers []core.PeerID) ([]core.PeerID, []core.PeerID) {
	topRated := make([]core.PeerID, 0)
	badRated := make([]core.PeerID, 0)

	for _, peer := range peers {
		if prh.topRatedCache.Has(peer.Bytes()) {
			topRated = append(topRated, peer)
		}

		if prh.badRatedCache.Has(peer.Bytes()) {
			badRated = append(badRated, peer)
		}
	}

	return topRated, badRated
}

func (prh *peersRatingHandler) updateRatingsMap(pid core.PeerID, newRating int32, updateFactor int32) {
	prettyPID := pid.Pretty()
	peerRatingInfo, exists := prh.ratingsMap[prettyPID]
	if !exists {
		prh.ratingsMap[prettyPID] = &ratingInfo{
			Rating:                       newRating,
			TimestampLastRequestToPid:    0,
			TimestampLastResponseFromPid: 0,
		}
		return
	}

	peerRatingInfo.Rating = newRating

	newTimeStamp := prh.getTimeHandler().Unix()
	if updateFactor == decreaseFactor {
		peerRatingInfo.TimestampLastRequestToPid = newTimeStamp
		return
	}

	peerRatingInfo.TimestampLastResponseFromPid = newTimeStamp
}

func (prh *peersRatingHandler) updateMetrics() {
	prh.mut.RLock()
	defer prh.mut.RUnlock()

	jsonMap, err := json.Marshal(&prh.ratingsMap)
	if err != nil {
		log.Debug("could not update metrics", "metric", p2p.MetricP2PPeersRating, "error", err.Error())
		return
	}

	prh.appStatusHandler.SetStringValue(p2p.MetricP2PPeersRating, string(jsonMap))
}

func (prh *peersRatingHandler) displayCachers() {
	prh.mut.RLock()
	defer prh.mut.RUnlock()

	displayMsg := fmt.Sprintf("testing- Ratings cachers\nTop rated:%s\nBad rated:%s", getPrintableRatings(prh.topRatedCache), getPrintableRatings(prh.badRatedCache))
	displayMsg += fmt.Sprintf("\nCurrent timestamp: %d", prh.getTimeHandler().Unix())
	displayMsg += "\nMarked for removal:"
	if len(prh.removalTimestampsMap) == 0 {
		displayMsg += " none"
	} else {
		for pid, timestamp := range prh.removalTimestampsMap {
			displayMsg += fmt.Sprintf("\npeerID: %s, removal timestamp: %d", pid, timestamp)
		}
	}
	log.Debug(displayMsg)
}

func getPrintableRatings(cache types.Cacher) string {
	keys := cache.Keys()
	ratings := ""
	for _, key := range keys {
		rating, ok := cache.Get(key)
		if !ok {
			continue
		}

		ratingInt, ok := rating.(int32)
		if !ok {
			log.Error("testing- could not cast to int32")
			continue
		}

		ratings += fmt.Sprintf("\npeerID: %s, rating: %d", core.PeerID(key).Pretty(), ratingInt)
	}

	return ratings
}

// Close stops the go routines started by this instance
func (prh *peersRatingHandler) Close() error {
	prh.cancel()
	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (prh *peersRatingHandler) IsInterfaceNil() bool {
	return prh == nil
}
