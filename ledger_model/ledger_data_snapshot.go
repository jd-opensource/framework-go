package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午2:28
 */

var _ binary_proto.DataContract = (*LedgerDataSnapshot)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(LedgerDataSnapshot{})
}

type LedgerDataSnapshot struct {
	AdminAccountHash       []byte `primitiveType:"BYTES"`
	UserAccountSetHash     []byte `primitiveType:"BYTES"`
	DataAccountSetHash     []byte `primitiveType:"BYTES"`
	ContractAccountSetHash []byte `primitiveType:"BYTES"`
}

func (l LedgerDataSnapshot) ContractCode() int32 {
	return binary_proto.DATA_SNAPSHOT
}

func (l LedgerDataSnapshot) ContractName() string {
	return "LedgerDataSnapshot"
}

func (l LedgerDataSnapshot) Description() string {
	return ""
}
