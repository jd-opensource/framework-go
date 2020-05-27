package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午2:42
 */

var _ binary_proto.DataContract = (*Transaction)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(Transaction{})
}

// 区块链交易，是被原子执行的操作集合
type Transaction struct {
	NodeRequest
	// 交易 Hash
	Hash []byte `primitiveType:"BYTES"`
	// 交易被包含的区块高度
	BlockHeight int64 `primitiveType:"INT64"`
	// 交易的执行结果
	ExecutionState TransactionState `refEnum:"2850"`
	// 交易的返回结果
	OperationResults []OperationResult `refContract:"880" list:"true"`
}

func (n Transaction) ContractCode() int32 {
	return binary_proto.TX
}

func (n Transaction) ContractName() string {
	return "Transaction"
}

func (n Transaction) Description() string {
	return ""
}
