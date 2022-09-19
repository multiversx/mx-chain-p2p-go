package factory

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go-p2p/common"
	"github.com/ElrondNetwork/elrond-go-p2p/config"
	"github.com/ElrondNetwork/elrond-go-p2p/libp2p/networksharding"
	"github.com/ElrondNetwork/elrond-go-p2p/mock"
	"github.com/stretchr/testify/assert"
)

func createMockArg() ArgsSharderFactory {
	return ArgsSharderFactory{

		PeerShardResolver:    &mock.PeerShardResolverStub{},
		Pid:                  "",
		PreferredPeersHolder: &mock.PeersHolderStub{},
		P2pConfig: config.P2PConfig{
			Sharding: config.ShardingConfig{
				Type:                    "unknown",
				TargetPeerCount:         6,
				MaxIntraShardValidators: 1,
				MaxCrossShardValidators: 1,
				MaxIntraShardObservers:  1,
				MaxCrossShardObservers:  1,
				AdditionalConnections: config.AdditionalConnectionsConfig{
					MaxFullHistoryObservers: 1,
				},
			},
		},
		NodeOperationMode: common.NormalOperation,
	}
}

func TestNewSharder_CreateListsSharderUnknownNodeOperationShouldError(t *testing.T) {
	t.Parallel()

	arg := createMockArg()
	arg.P2pConfig.Sharding.Type = common.ListsSharder
	arg.NodeOperationMode = ""
	sharder, err := NewSharder(arg)

	assert.True(t, errors.Is(err, common.ErrInvalidValue))
	assert.True(t, strings.Contains(err.Error(), "unknown node operation mode"))
	assert.True(t, check.IfNil(sharder))
}

func TestNewSharder_CreateListsSharderShouldWork(t *testing.T) {
	t.Parallel()

	arg := createMockArg()
	arg.P2pConfig.Sharding.Type = common.ListsSharder
	sharder, err := NewSharder(arg)
	maxPeerCount := uint32(5)
	maxValidators := uint32(1)
	maxObservers := uint32(1)

	argListsSharder := networksharding.ArgListsSharder{
		PeerResolver: &mock.PeerShardResolverStub{},
		SelfPeerId:   "",
		P2pConfig: config.P2PConfig{
			Sharding: config.ShardingConfig{
				TargetPeerCount:         maxPeerCount,
				MaxIntraShardObservers:  maxObservers,
				MaxIntraShardValidators: maxValidators,
				MaxCrossShardObservers:  maxObservers,
				MaxCrossShardValidators: maxValidators,
				MaxSeeders:              0,
			},
		},
	}
	expectedSharder, _ := networksharding.NewListsSharder(argListsSharder)
	assert.Nil(t, err)
	assert.IsType(t, reflect.TypeOf(expectedSharder), reflect.TypeOf(sharder))
}

func TestNewSharder_CreateOneListSharderShouldWork(t *testing.T) {
	t.Parallel()

	arg := createMockArg()
	arg.P2pConfig.Sharding.Type = common.OneListSharder
	sharder, err := NewSharder(arg)
	maxPeerCount := 2

	expectedSharder, _ := networksharding.NewOneListSharder("", maxPeerCount)
	assert.Nil(t, err)
	assert.IsType(t, reflect.TypeOf(expectedSharder), reflect.TypeOf(sharder))
}

func TestNewSharder_CreateNilListSharderShouldWork(t *testing.T) {
	t.Parallel()

	arg := createMockArg()
	arg.P2pConfig.Sharding.Type = common.NilListSharder
	sharder, err := NewSharder(arg)

	expectedSharder := networksharding.NewNilListSharder()
	assert.Nil(t, err)
	assert.IsType(t, reflect.TypeOf(expectedSharder), reflect.TypeOf(sharder))
}

func TestNewSharder_CreateWithUnknownVariantShouldErr(t *testing.T) {
	t.Parallel()

	arg := createMockArg()
	sharder, err := NewSharder(arg)

	assert.True(t, errors.Is(err, common.ErrInvalidValue))
	assert.True(t, check.IfNil(sharder))
}