package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午2:28
 */

var _ binary_proto.DataContract = (*BlockBody)(nil)

func init() {
	binary_proto.RegisterContract(BlockBody{})
}

type BlockBody struct {
	LedgerDataSnapshot
	LedgerEventSnapshot
	PreviousHash       []byte `primitiveType:"BYTES"`
	LedgerHash         []byte `primitiveType:"BYTES"`
	Height             int64  `primitiveType:"INT64"`
	TransactionSetHash []byte `primitiveType:"BYTES"`
	Timestamp          int64  `primitiveType:"INT64"`
}

func (b BlockBody) ContractCode() int32 {
	return binary_proto.BLOCK_BODY
}

func (b BlockBody) ContractName() string {
	return "BlockBody"
}

func (b BlockBody) Description() string {
	return ""
}
