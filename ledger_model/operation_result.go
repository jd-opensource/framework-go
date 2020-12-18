package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午2:54
 */

var _ binary_proto.DataContract = (*OperationResult)(nil)

func init() {
	binary_proto.RegisterContract(OperationResult{})
}

type OperationResult struct {
	Index  int32      `primitiveType:"INT32"`
	Result BytesValue `refContract:"128"`
}

func (o OperationResult) ContractCode() int32 {
	return binary_proto.TX_OP_RESULT
}

func (o OperationResult) ContractName() string {
	return "OperationResult"
}

func (o OperationResult) Description() string {
	return ""
}
