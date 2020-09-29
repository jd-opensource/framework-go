package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/6/9 下午6:44
 */

var _ binary_proto.DataContract = (*EventEntry)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(EventEntry{})
}

type EventEntry struct {
	Name     string     `primitiveType:"TEXT"`
	Content  BytesValue `refContract:"128"`
	Sequence int64      `primitiveType:"INT64"`
}

func (e EventEntry) ContractCode() int32 {
	return binary_proto.TX_OP_EVENT_PUBLISH_ENTITY
}

func (e EventEntry) ContractName() string {
	return "EventEntry"
}

func (e EventEntry) Description() string {
	return ""
}
