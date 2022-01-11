package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

var _ binary_proto.DataContract = (*CryptoHashAlgoUpdateOperation)(nil)

func init() {
	binary_proto.RegisterContract(CryptoHashAlgoUpdateOperation{})
}

type CryptoHashAlgoUpdateOperation struct {
	HashAlgoName string `primitiveType:"TEXT"`
}

func (c CryptoHashAlgoUpdateOperation) ContractCode() int32 {
	return binary_proto.TX_OP_CRYPTO_ALGO_UPDATE
}

func (c CryptoHashAlgoUpdateOperation) ContractName() string {
	return OperationTypeCryptoHashAlgoUpdateOperation
}

func (c CryptoHashAlgoUpdateOperation) Description() string {
	return ""
}
