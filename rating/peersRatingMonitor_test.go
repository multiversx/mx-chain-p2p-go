package rating

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	coreMocks "github.com/multiversx/mx-chain-core-go/data/mock"
	p2p "github.com/multiversx/mx-chain-p2p-go"
	"github.com/multiversx/mx-chain-p2p-go/mock"
	"github.com/stretchr/testify/assert"
)

func createMockArgPeersRatingMonitor() ArgPeersRatingMonitor {
	return ArgPeersRatingMonitor{
		TopRatedCache: &mock.CacherStub{},
		BadRatedCache: &mock.CacherStub{},
	}
}

func TestNewPeersRatingMonitor(t *testing.T) {
	t.Parallel()

	t.Run("nil top rated cache should error", func(t *testing.T) {
		t.Parallel()

		args := createMockArgPeersRatingMonitor()
		args.TopRatedCache = nil

		monitor, err := NewPeersRatingMonitor(args)
		assert.True(t, errors.Is(err, p2p.ErrNilCacher))
		assert.True(t, strings.Contains(err.Error(), "TopRatedCache"))
		assert.True(t, check.IfNil(monitor))
	})
	t.Run("nil bad rated cache should error", func(t *testing.T) {
		t.Parallel()

		args := createMockArgPeersRatingMonitor()
		args.BadRatedCache = nil

		monitor, err := NewPeersRatingMonitor(args)
		assert.True(t, errors.Is(err, p2p.ErrNilCacher))
		assert.True(t, strings.Contains(err.Error(), "BadRatedCache"))
		assert.True(t, check.IfNil(monitor))
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		args := ArgPeersRatingMonitor{
			TopRatedCache: coreMocks.NewCacherMock(),
			BadRatedCache: coreMocks.NewCacherMock(),
		}
		monitor, err := NewPeersRatingMonitor(args)
		assert.Nil(t, err)
		assert.False(t, check.IfNil(monitor))

		topPid1 := core.PeerID("top_pid1")
		args.TopRatedCache.Put(topPid1.Bytes(), int32(100), 32)
		topPid2 := core.PeerID("top_pid2")
		args.TopRatedCache.Put(topPid2.Bytes(), int32(50), 32)
		topPid3 := core.PeerID("top_pid3")
		args.TopRatedCache.Put(topPid3.Bytes(), int32(0), 32)
		commonPid := core.PeerID("common_pid")
		args.TopRatedCache.Put(commonPid.Bytes(), int32(10), 32)

		badPid1 := core.PeerID("bad_pid1")
		args.BadRatedCache.Put(badPid1.Bytes(), int32(-100), 32)
		badPid2 := core.PeerID("bad_pid2")
		args.BadRatedCache.Put(badPid2.Bytes(), int32(-50), 32)
		badPid3 := core.PeerID("bad_pid3")
		args.BadRatedCache.Put(badPid3.Bytes(), int32(-10), 32)
		args.TopRatedCache.Put(commonPid.Bytes(), int32(-10), 32) // should override

		expectedMap := map[string]int32{
			topPid1.Pretty():   int32(100),
			topPid2.Pretty():   int32(50),
			topPid3.Pretty():   int32(0),
			badPid1.Pretty():   int32(-100),
			badPid2.Pretty():   int32(-50),
			badPid3.Pretty():   int32(-10),
			commonPid.Pretty(): int32(-10),
		}
		expectedStr, _ := json.Marshal(&expectedMap)
		ratings := monitor.GetPeersRatings()
		assert.Equal(t, string(expectedStr), ratings)
	})
}
