package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/6/9 下午6:31
 */

var _ binary_proto.DataContract = (*EventAccountRegisterOperation)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(EventAccountRegisterOperation{})
}

type EventAccountRegisterOperation struct {
	EventAccountID BlockchainIdentity `refContract:"144"`
}

func (e EventAccountRegisterOperation) ContractCode() int32 {
	return binary_proto.TX_OP_EVENT_ACC_REG
}

func (e EventAccountRegisterOperation) ContractName() string {
	return "EventAccountRegisterOperation"
}

func (e EventAccountRegisterOperation) Description() string {
	return ""
}
