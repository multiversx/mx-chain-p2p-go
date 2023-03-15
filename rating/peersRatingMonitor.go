package rating

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	p2p "github.com/multiversx/mx-chain-p2p-go"
	"github.com/multiversx/mx-chain-storage-go/types"
)

const unknownRating = "unknown"

// ArgPeersRatingMonitor is the DTO used to create a new peers rating monitor
type ArgPeersRatingMonitor struct {
	TopRatedCache       types.Cacher
	BadRatedCache       types.Cacher
	ConnectionsProvider connectionsProvider
}

type peersRatingMonitor struct {
	topRatedCache       types.Cacher
	badRatedCache       types.Cacher
	connectionsProvider connectionsProvider
}

// NewPeersRatingMonitor returns a new peers rating monitor
func NewPeersRatingMonitor(args ArgPeersRatingMonitor) (*peersRatingMonitor, error) {
	err := checkMonitorArgs(args)
	if err != nil {
		return nil, err
	}

	monitor := &peersRatingMonitor{
		topRatedCache:       args.TopRatedCache,
		badRatedCache:       args.BadRatedCache,
		connectionsProvider: args.ConnectionsProvider,
	}

	go monitor.processTestLogs(context.Background())

	return monitor, nil
}

func checkMonitorArgs(args ArgPeersRatingMonitor) error {
	if check.IfNil(args.TopRatedCache) {
		return fmt.Errorf("%w for TopRatedCache", p2p.ErrNilCacher)
	}
	if check.IfNil(args.BadRatedCache) {
		return fmt.Errorf("%w for BadRatedCache", p2p.ErrNilCacher)
	}
	if check.IfNil(args.ConnectionsProvider) {
		return p2p.ErrNilConnectionsProvider
	}

	return nil
}

func (monitor *peersRatingMonitor) processTestLogs(ctx context.Context) {
	displayCachersTime := time.Second * 10
	timerDisplayCachers := time.NewTimer(displayCachersTime)
	defer timerDisplayCachers.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Debug("testing- closing processTestLogs")
			return
		case <-timerDisplayCachers.C:
			monitor.displayCachers()
			timerDisplayCachers.Reset(displayCachersTime)
		}
	}
}

func (monitor *peersRatingMonitor) displayCachers() {
	displayMsg := fmt.Sprintf("testing- Ratings cachers\nTop rated:%s\nBad rated:%s", getPrintableRatings(monitor.topRatedCache), getPrintableRatings(monitor.badRatedCache))
	displayMsg += fmt.Sprintf("\ntesting- Connected peers ratings:\n%s", monitor.GetConnectedPeersRatings())
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

// GetConnectedPeersRatings returns the ratings of the current connected peers
func (monitor *peersRatingMonitor) GetConnectedPeersRatings() string {
	connectedPeersRatings := monitor.extractConnectedPeersRatings()

	jsonMap, err := json.Marshal(&connectedPeersRatings)
	if err != nil {
		return ""
	}

	return string(jsonMap)
}

func (monitor *peersRatingMonitor) extractConnectedPeersRatings() map[string]string {
	connectedPeers := monitor.connectionsProvider.ConnectedPeers()
	connectedPeersRatings := make(map[string]string, len(connectedPeers))
	for _, connectedPeer := range connectedPeers {
		connectedPeersRatings[connectedPeer.Pretty()] = monitor.fetchRating(connectedPeer)
	}

	return connectedPeersRatings
}

func (monitor *peersRatingMonitor) fetchRating(pid core.PeerID) string {
	rating, found := monitor.topRatedCache.Get(pid.Bytes())
	if found {
		return fmt.Sprintf("%d", rating)
	}

	rating, found = monitor.badRatedCache.Get(pid.Bytes())
	if found {
		return fmt.Sprintf("%d", rating)
	}

	return unknownRating
}

// IsInterfaceNil returns true if there is no value under the interface
func (monitor *peersRatingMonitor) IsInterfaceNil() bool {
	return monitor == nil
}
