package libp2p

import (
	"context"

	"github.com/ElrondNetwork/elrond-go-p2p/common"
	"github.com/ElrondNetwork/elrond-go-storage/types"
	"github.com/ElrondNetwork/go-libp2p-pubsub"
	pb "github.com/ElrondNetwork/go-libp2p-pubsub/pb"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/whyrusleeping/timecache"
)

var MaxSendBuffSize = maxSendBuffSize
var BroadcastGoRoutines = broadcastGoRoutines
var PubsubTimeCacheDuration = pubsubTimeCacheDuration
var AcceptMessagesInAdvanceDuration = acceptMessagesInAdvanceDuration
var SequenceNumberSize = sequenceNumberSize

const CurrentTopicMessageVersion = currentTopicMessageVersion
const PollWaitForConnectionsInterval = pollWaitForConnectionsInterval

// SetHost -
func (netMes *networkMessenger) SetHost(newHost ConnectableHost) {
	netMes.p2pHost = newHost
}

// SetLoadBalancer -
func (netMes *networkMessenger) SetLoadBalancer(outgoingPLB common.ChannelLoadBalancer) {
	netMes.outgoingPLB = outgoingPLB
}

// SetPeerDiscoverer -
func (netMes *networkMessenger) SetPeerDiscoverer(discoverer common.PeerDiscoverer) {
	netMes.peerDiscoverer = discoverer
}

// PubsubCallback -
func (netMes *networkMessenger) PubsubCallback(handler common.MessageProcessor, topic string) func(ctx context.Context, pid peer.ID, message *pubsub.Message) bool {
	topicProcs := newTopicProcessors()
	_ = topicProcs.addTopicProcessor("identifier", handler)

	return netMes.pubsubCallback(topicProcs, topic)
}

// ValidMessageByTimestamp -
func (netMes *networkMessenger) ValidMessageByTimestamp(msg common.MessageP2P) error {
	return netMes.validMessageByTimestamp(msg)
}

// MapHistogram -
func (netMes *networkMessenger) MapHistogram(input map[uint32]int) string {
	return netMes.mapHistogram(input)
}

// PubsubHasTopic -
func (netMes *networkMessenger) PubsubHasTopic(expectedTopic string) bool {
	netMes.mutTopics.RLock()
	topics := netMes.pb.GetTopics()
	netMes.mutTopics.RUnlock()

	for _, topic := range topics {
		if topic == expectedTopic {
			return true
		}
	}
	return false
}

// HasProcessorForTopic -
func (netMes *networkMessenger) HasProcessorForTopic(expectedTopic string) bool {
	processor, found := netMes.processors[expectedTopic]

	return found && processor != nil
}

// ProcessReceivedDirectMessage -
func (ds *directSender) ProcessReceivedDirectMessage(message *pb.Message, fromConnectedPeer peer.ID) error {
	return ds.processReceivedDirectMessage(message, fromConnectedPeer)
}

// SeenMessages -
func (ds *directSender) SeenMessages() *timecache.TimeCache {
	return ds.seenMessages
}

// Counter -
func (ds *directSender) Counter() uint64 {
	return ds.counter
}

// Mutexes -
func (mh *MutexHolder) Mutexes() types.Cacher {
	return mh.mutexes
}

// SetSignerInDirectSender sets the signer in the direct sender
func (netMes *networkMessenger) SetSignerInDirectSender(signer common.SignerVerifier) {
	netMes.ds.(*directSender).signer = signer
}