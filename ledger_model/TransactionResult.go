package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/11/16 下午3:19
 */

var _ binary_proto.DataContract = (*TransactionResult)(nil)

func init() {
	binary_proto.RegisterContract(TransactionResult{})
}

type TransactionResult struct {
	TransactionHash   []byte                      `primitiveType:"BYTES"`
	BlockHeight       int64                       `primitiveType:"INT64"`
	ExecutionState    TransactionState            `refEnum:"2850"`
	OperationResults  []OperationResult           `refContract:"880" list:"true"`
	DataSnapshot      LedgerDataSnapshot          `refContract:"304"`
	DerivedOperations []binary_proto.DataContract `refContract:"768" genericContract:"true" list:"true"`
}

func (t TransactionResult) ContractCode() int32 {
	return binary_proto.TX_RESULT
}

func (t TransactionResult) ContractName() string {
	return "TransactionResult"
}

func (t TransactionResult) Description() string {
	return ""
}
