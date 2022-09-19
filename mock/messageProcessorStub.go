package mock

import (
	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-p2p/common"
)

// MessageProcessorStub -
type MessageProcessorStub struct {
	ProcessMessageCalled func(message common.MessageP2P, fromConnectedPeer core.PeerID) error
}

// ProcessReceivedMessage -
func (mps *MessageProcessorStub) ProcessReceivedMessage(message common.MessageP2P, fromConnectedPeer core.PeerID) error {
	if mps.ProcessMessageCalled != nil {
		return mps.ProcessMessageCalled(message, fromConnectedPeer)
	}

	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (mps *MessageProcessorStub) IsInterfaceNil() bool {
	return mps == nil
}