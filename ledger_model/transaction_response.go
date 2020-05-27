package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/28 下午6:23
 */

var _ binary_proto.DataContract = (*TransactionResponse)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(TransactionResponse{})
}

type TransactionResponse struct {
	ContentHash []byte `primitiveType:"BYTES"`

	ExecutionState TransactionState `refEnum:"2850"`

	BlockHash []byte `primitiveType:"BYTES"`

	BlockHeight int64 `primitiveType:"INT64"`

	Success bool `primitiveType:"BOOLEAN"`

	OperationResults []OperationResult `refContract:"880" list:"true"`
}

func (t TransactionResponse) ContractCode() int32 {
	return binary_proto.TX_RESPONSE
}

func (t TransactionResponse) ContractName() string {
	return "TransactionResponse"
}

func (t TransactionResponse) Description() string {
	return ""
}
