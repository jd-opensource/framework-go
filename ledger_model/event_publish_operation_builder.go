package ledger_model

import (
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
	"sync"
)

/*
 * Author: imuge
 * Date: 2020/6/9 下午6:20
 */

type EventPublishOperationBuilder struct {
	factory   *BlockchainOperationFactory
	Operation *EventPublishOperation

	added bool // 该构建是否已经加入
	mutex sync.Mutex
}

func NewEventPublishOperationBuilder(address []byte, factory *BlockchainOperationFactory) *EventPublishOperationBuilder {
	return &EventPublishOperationBuilder{
		factory: factory,
		Operation: &EventPublishOperation{
			EventAddress: address,
			Events:       nil,
		},
	}
}

func (epob *EventPublishOperationBuilder) addOperation() {
	epob.mutex.Lock()
	defer epob.mutex.Unlock()
	if !epob.added && epob.factory != nil && epob.Operation.Events != nil {
		epob.factory.addOperation(epob.Operation)
		epob.added = true
	}
}

func (epob *EventPublishOperationBuilder) PublishBytes(name string, content []byte, sequence int64) *EventPublishOperationBuilder {
	epob.Operation.Events = append(epob.Operation.Events,
		EventEntry{
			Name: name,
			Content: BytesValue{
				Type:  BYTES,
				Bytes: content,
			},
			Sequence: sequence,
		})
	epob.addOperation()

	return epob
}

func (epob *EventPublishOperationBuilder) PublishString(name string, content string, sequence int64) *EventPublishOperationBuilder {
	epob.Operation.Events = append(epob.Operation.Events,
		EventEntry{
			Name: name,
			Content: BytesValue{
				Type:  TEXT,
				Bytes: bytes.StringToBytes(content),
			},
			Sequence: sequence,
		})
	epob.addOperation()

	return epob
}

func (epob *EventPublishOperationBuilder) PublishInt64(name string, content int64, sequence int64) *EventPublishOperationBuilder {
	epob.Operation.Events = append(epob.Operation.Events,
		EventEntry{
			Name: name,
			Content: BytesValue{
				Type:  INT64,
				Bytes: bytes.Int64ToBytes(content),
			},
			Sequence: sequence,
		})
	epob.addOperation()

	return epob
}
