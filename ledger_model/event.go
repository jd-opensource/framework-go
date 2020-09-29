package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/6/10 下午4:30
 */

var _ binary_proto.DataContract = (*Event)(nil)

// 事件
type Event struct {
	// 事件名
	Name string `primitiveType:"TEXT"`

	// 事件序号
	Sequence int64 `primitiveType:"INT64"`

	// 事件内容
	Content BytesValue `refContract:"128"`

	// 产生事件的交易哈希
	TransactionSource []byte `primitiveType:"BYTES"`

	// 产生事件的合约地址
	ContractSource string `primitiveType:"TEXT"`

	// 产生事件的区块高度
	BlockHeight int64 `primitiveType:"INT64"`

	// 事件账户地址
	EventAccount []byte `primitiveType:"BYTES"`
}

func (e Event) ContractCode() int32 {
	return binary_proto.EVENT_MESSAGE
}

func (e Event) ContractName() string {
	return "Event"
}

func (e Event) Description() string {
	return ""
}
