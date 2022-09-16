package loadBalancer

import (
	"github.com/ElrondNetwork/elrond-go-p2p/common"
)

func (oplb *OutgoingChannelLoadBalancer) Chans() []chan *common.SendableData {
	return oplb.chans
}

func (oplb *OutgoingChannelLoadBalancer) Names() []string {
	return oplb.names
}

func (oplb *OutgoingChannelLoadBalancer) NamesChans() map[string]chan *common.SendableData {
	return oplb.namesChans
}

func DefaultSendChannel() string {
	return defaultSendChannel
}
