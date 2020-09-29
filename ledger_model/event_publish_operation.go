package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/6/9 下午6:37
 */

var _ binary_proto.DataContract = (*EventPublishOperation)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(EventPublishOperation{})
}

type EventPublishOperation struct {
	EventAddress []byte `primitiveType:"BYTES"`

	Events []EventEntry `refContract:"898" list:"true"`
}

func (e EventPublishOperation) ContractCode() int32 {
	return binary_proto.TX_OP_EVENT_PUBLISH
}

func (e EventPublishOperation) ContractName() string {
	return "EventPublishOperation"
}

func (e EventPublishOperation) Description() string {
	return ""
}
