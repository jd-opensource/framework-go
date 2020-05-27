package ledger_model

import (
	"framework-go/utils/bytes"
	"sync"
)

/*
 * Author: imuge
 * Date: 2020/5/28 下午4:56
 */

type DataAccountKVSetOperationBuilder struct {
	Operation *DataAccountKVSetOperation
	factory   *BlockchainOperationFactory

	added bool // 该构建是否已经加入
	mutex sync.Mutex
}

func NewDataAccountKVSetOperationBuilder(address []byte, factory *BlockchainOperationFactory) *DataAccountKVSetOperationBuilder {
	return &DataAccountKVSetOperationBuilder{
		factory: factory,
		Operation: &DataAccountKVSetOperation{
			AccountAddress: address,
			WriteSet:       nil,
		},
	}
}

func (dakvob *DataAccountKVSetOperationBuilder) checkRepeat(key string) {
	for _, set := range dakvob.Operation.WriteSet {
		if set.Key == key {
			panic("Can't set the same key repeatedly!")
		}
	}
}

func (dakvob *DataAccountKVSetOperationBuilder) addOperation() {
	dakvob.mutex.Lock()
	defer dakvob.mutex.Unlock()
	if !dakvob.added && dakvob.factory != nil && dakvob.Operation.WriteSet != nil {
		dakvob.factory.addOperation(dakvob.Operation)
		dakvob.added = true
	}
}

func (dakvob *DataAccountKVSetOperationBuilder) SetBytes(key string, value []byte, expVersion int64) *DataAccountKVSetOperationBuilder {
	dakvob.checkRepeat(key)
	dakvob.Operation.WriteSet = append(dakvob.Operation.WriteSet, KVWriteEntry{
		Key: key,
		Value: BytesValue{
			Type:  BYTES,
			Bytes: value,
		},
		ExpectedVersion: expVersion,
	})
	dakvob.addOperation()

	return dakvob
}

func (dakvob *DataAccountKVSetOperationBuilder) SetImage(key string, value []byte, expVersion int64) *DataAccountKVSetOperationBuilder {
	dakvob.checkRepeat(key)
	dakvob.Operation.WriteSet = append(dakvob.Operation.WriteSet, KVWriteEntry{
		Key: key,
		Value: BytesValue{
			Type:  IMG,
			Bytes: value,
		},
		ExpectedVersion: expVersion,
	})
	dakvob.addOperation()

	return dakvob
}

func (dakvob *DataAccountKVSetOperationBuilder) SetText(key, value string, expVersion int64) *DataAccountKVSetOperationBuilder {
	dakvob.checkRepeat(key)
	dakvob.Operation.WriteSet = append(dakvob.Operation.WriteSet, KVWriteEntry{
		Key: key,
		Value: BytesValue{
			Type:  TEXT,
			Bytes: []byte(value),
		},
		ExpectedVersion: expVersion,
	})
	dakvob.addOperation()

	return dakvob
}

func (dakvob *DataAccountKVSetOperationBuilder) SetJSON(key, value string, expVersion int64) *DataAccountKVSetOperationBuilder {
	dakvob.checkRepeat(key)
	dakvob.Operation.WriteSet = append(dakvob.Operation.WriteSet, KVWriteEntry{
		Key: key,
		Value: BytesValue{
			Type:  JSON,
			Bytes: []byte(value),
		},
		ExpectedVersion: expVersion,
	})
	dakvob.addOperation()

	return dakvob
}

func (dakvob *DataAccountKVSetOperationBuilder) SetInt64(key string, value int64, expVersion int64) *DataAccountKVSetOperationBuilder {
	dakvob.checkRepeat(key)
	dakvob.Operation.WriteSet = append(dakvob.Operation.WriteSet, KVWriteEntry{
		Key: key,
		Value: BytesValue{
			Type:  INT64,
			Bytes: bytes.Int64ToBytes(value),
		},
		ExpectedVersion: expVersion,
	})
	dakvob.addOperation()

	return dakvob
}

func (dakvob *DataAccountKVSetOperationBuilder) SetTimestamp(key string, value int64, expVersion int64) *DataAccountKVSetOperationBuilder {
	dakvob.checkRepeat(key)
	dakvob.Operation.WriteSet = append(dakvob.Operation.WriteSet, KVWriteEntry{
		Key: key,
		Value: BytesValue{
			Type:  TIMESTAMP,
			Bytes: bytes.Int64ToBytes(value),
		},
		ExpectedVersion: expVersion,
	})
	dakvob.addOperation()

	return dakvob
}
