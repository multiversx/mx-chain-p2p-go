package memp2p

import (
	"github.com/ElrondNetwork/elrond-go-p2p/common"
)

func (messenger *Messenger) TopicValidator(name string) common.MessageProcessor {
	messenger.topicsMutex.RLock()
	processor := messenger.topicValidators[name]
	messenger.topicsMutex.RUnlock()

	return processor
}
