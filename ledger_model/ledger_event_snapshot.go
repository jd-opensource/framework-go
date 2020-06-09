package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/6/9 下午7:09
 */

var _ binary_proto.DataContract = (*LedgerEventSnapshot)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(LedgerEventSnapshot{})
}

type LedgerEventSnapshot struct {
	SystemEventSetHash []byte `primitiveType:"BYTES"`
	UserEventSetHash   []byte `primitiveType:"BYTES"`
}

func (l LedgerEventSnapshot) ContractCode() int32 {
	return binary_proto.EVENT_SNAPSHOT
}

func (l LedgerEventSnapshot) ContractName() string {
	return "LedgerEventSnapshot"
}

func (l LedgerEventSnapshot) Description() string {
	return ""
}
