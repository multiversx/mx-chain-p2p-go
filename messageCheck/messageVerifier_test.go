package messagecheck_test

import (
	"errors"
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go-p2p/common"
	"github.com/ElrondNetwork/elrond-go-p2p/message"
	messagecheck "github.com/ElrondNetwork/elrond-go-p2p/messageCheck"
	"github.com/ElrondNetwork/elrond-go-p2p/mock"
	"github.com/stretchr/testify/require"
)

func createMessageVerifierArgs() messagecheck.ArgsMessageVerifier {
	return messagecheck.ArgsMessageVerifier{
		Marshaller: &mock.MarshallerStub{},
		P2PSigner:  &mock.P2PSignerStub{},
	}
}

func TestNewMessageVerifier(t *testing.T) {
	t.Parallel()

	t.Run("nil marshaller", func(t *testing.T) {
		t.Parallel()

		args := createMessageVerifierArgs()
		args.Marshaller = nil

		mv, err := messagecheck.NewMessageVerifier(args)
		require.Nil(t, mv)
		require.Equal(t, common.ErrNilMarshalizer, err)
	})

	t.Run("nil p2p signer", func(t *testing.T) {
		t.Parallel()

		args := createMessageVerifierArgs()
		args.P2PSigner = nil

		mv, err := messagecheck.NewMessageVerifier(args)
		require.Nil(t, mv)
		require.Equal(t, common.ErrNilP2PSigner, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		args := createMessageVerifierArgs()
		mv, err := messagecheck.NewMessageVerifier(args)
		require.Nil(t, err)
		require.False(t, check.IfNil(mv))
	})
}

func TestSerializeDeserialize(t *testing.T) {
	t.Parallel()

	t.Run("serialize, marshal should err", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("expected error")

		args := createMessageVerifierArgs()
		args.Marshaller = &mock.MarshallerStub{
			MarshalCalled: func(obj interface{}) ([]byte, error) {
				return nil, expectedErr
			},
		}

		messages := []common.MessageP2P{
			&message.Message{
				FromField:    []byte("from1"),
				PayloadField: []byte("payload1"),
			},
		}

		mv, err := messagecheck.NewMessageVerifier(args)
		require.Nil(t, err)

		messagesBytes, err := mv.Serialize(messages)
		require.Nil(t, messagesBytes)
		require.Equal(t, expectedErr, err)
	})

	t.Run("deserialize, unmarshal should err", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("expected error")

		args := createMessageVerifierArgs()
		args.Marshaller = &mock.MarshallerStub{
			UnmarshalCalled: func(obj interface{}, buff []byte) error {
				return expectedErr
			},
		}

		mv, err := messagecheck.NewMessageVerifier(args)
		require.Nil(t, err)

		messages, err := mv.Deserialize([]byte("messages data"))
		require.Nil(t, messages)
		require.Equal(t, expectedErr, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		args := createMessageVerifierArgs()
		args.Marshaller = &mock.MarshallerMock{}

		expectedMessages := []common.MessageP2P{
			&message.Message{
				FromField:      []byte("from1"),
				PayloadField:   []byte("payload1"), // it is used as data field for pubsub
				SeqNoField:     []byte("seq"),
				TopicField:     string("topic"),
				SignatureField: []byte("sig"),
				KeyField:       []byte("key"),
			},
			&message.Message{
				FromField:      []byte("from2"),
				PayloadField:   []byte("payload2"),
				SeqNoField:     []byte("seq"),
				TopicField:     string("topic"),
				SignatureField: []byte("sig"),
				KeyField:       []byte("key"),
			},
		}

		mv, err := messagecheck.NewMessageVerifier(args)
		require.Nil(t, err)

		messagesBytes, err := mv.Serialize(expectedMessages)
		require.Nil(t, err)

		messages, err := mv.Deserialize(messagesBytes)
		require.Nil(t, err)

		require.Equal(t, expectedMessages, messages)
	})
}

func TestVerify(t *testing.T) {
	t.Parallel()

	t.Run("nil p2p message", func(t *testing.T) {
		t.Parallel()

		args := createMessageVerifierArgs()
		mv, err := messagecheck.NewMessageVerifier(args)
		require.Nil(t, err)

		err = mv.Verify(nil)
		require.Equal(t, common.ErrNilMessage, err)
	})

	t.Run("p2p signer verify should fail", func(t *testing.T) {
		t.Parallel()

		args := createMessageVerifierArgs()

		expectedErr := errors.New("expected err")
		args.P2PSigner = &mock.P2PSignerStub{
			VerifyCalled: func(payload []byte, pid core.PeerID, signature []byte) error {
				return expectedErr
			},
		}
		mv, err := messagecheck.NewMessageVerifier(args)
		require.Nil(t, err)

		msg := &message.Message{
			FromField:      []byte("from1"),
			PayloadField:   []byte("payload1"),
			SeqNoField:     []byte("seq"),
			TopicField:     string("topic"),
			SignatureField: []byte("sig"),
			KeyField:       []byte("key"),
		}

		err = mv.Verify(msg)
		require.Equal(t, expectedErr, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		args := createMessageVerifierArgs()

		wasCalled := false
		args.P2PSigner = &mock.P2PSignerStub{
			VerifyCalled: func(payload []byte, pid core.PeerID, signature []byte) error {
				wasCalled = true

				return nil
			},
		}
		mv, err := messagecheck.NewMessageVerifier(args)
		require.Nil(t, err)

		msg := &message.Message{
			FromField:      []byte("from1"),
			PayloadField:   []byte("payload1"),
			SeqNoField:     []byte("seq"),
			TopicField:     string("topic"),
			SignatureField: []byte("sig"),
			KeyField:       []byte("key"),
		}

		err = mv.Verify(msg)
		require.Nil(t, err)

		require.True(t, wasCalled)
	})
}