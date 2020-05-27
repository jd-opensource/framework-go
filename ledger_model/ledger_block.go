package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午2:27
 */

var _ binary_proto.DataContract = (*LedgerBlock)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(LedgerBlock{})
}

type LedgerBlock struct {
	BlockBody
	Hash []byte `primitiveType:"BYTES"`
}

func (l LedgerBlock) ContractCode() int32 {
	return binary_proto.BLOCK
}

func (l LedgerBlock) ContractName() string {
	return "LedgerBlock"
}

func (l LedgerBlock) Description() string {
	return ""
}
