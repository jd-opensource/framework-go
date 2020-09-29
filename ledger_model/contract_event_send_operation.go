package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/9/17 下午6:52
 */

var _ binary_proto.DataContract = (*ContractEventSendOperation)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(ContractEventSendOperation{})
}

type ContractEventSendOperation struct {

	// 响应事件的合约地址
	ContractAddress []byte `primitiveType:"BYTES"`

	// 事件名
	Event string `primitiveType:"TEXT"`

	// 事件参数
	Args BytesValueList `refContract:"129"`

	// contract's version
	Version int64 `primitiveType:"INT64"`
}

func (c ContractEventSendOperation) ContractCode() int32 {
	return binary_proto.TX_OP_CONTRACT_EVENT_SEND
}

func (c ContractEventSendOperation) ContractName() string {
	return "ContractEventSendOperation"
}

func (c ContractEventSendOperation) Description() string {
	return ""
}
