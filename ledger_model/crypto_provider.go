package ledger_model

import (
	binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
)

/*
 * Author: imuge
 * Date: 2020/5/27 下午1:53
 */

var _ binary_proto.DataContract = (*CryptoProvider)(nil)

func init() {
	binary_proto.RegisterContract(CryptoProvider{})
}

type CryptoProvider struct {
	Name       string                      `primitiveType:"TEXT"`
	Algorithms []framework.CryptoAlgorithm `refContract:"2849" list:"true"`
}

func (c CryptoProvider) ContractCode() int32 {
	return binary_proto.METADATA_CRYPTO_SETTING_PROVIDER
}

func (c CryptoProvider) ContractName() string {
	return "CryptoProvider"
}

func (c CryptoProvider) Description() string {
	return ""
}
