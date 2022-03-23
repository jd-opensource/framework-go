package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

var _ binary_proto.DataContract = (*HashAlgorithmUpdateOperation)(nil)

func init() {
	binary_proto.RegisterContract(HashAlgorithmUpdateOperation{})
}

type HashAlgorithmUpdateOperation struct {
	Algorithm string `primitiveType:"TEXT"`
}

func (c HashAlgorithmUpdateOperation) ContractCode() int32 {
	return binary_proto.TX_OP_HASH_ALGORITHM_UPDATE
}

func (c HashAlgorithmUpdateOperation) ContractName() string {
	return OperationTypeCryptoHashAlgoUpdateOperation
}

func (c HashAlgorithmUpdateOperation) Description() string {
	return ""
}
