package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/6/1 下午7:28
 */

var _ binary_proto.DataContract = (*ContractCodeDeployOperation)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(ContractCodeDeployOperation{})
}

type ContractCodeDeployOperation struct {
	ContractID BlockchainIdentity `refContract:"144"`

	ChainCode []byte `primitiveType:"BYTES"`
	// 地址签名
	// 这是合约账户身份 使用对应的私钥对地址做出的签名
	// 在注册时将校验此签名与账户地址、公钥是否相匹配，以此保证只有私钥的持有者才能注册相应的合约账户，确保合约账户的唯一性
	AddressSignature DigitalSignature `refContract:"2864"`
}

func (c ContractCodeDeployOperation) ContractCode() int32 {
	return binary_proto.TX_OP_CONTRACT_DEPLOY
}

func (c ContractCodeDeployOperation) ContractName() string {
	return "ContractCodeDeployOperation"
}

func (c ContractCodeDeployOperation) Description() string {
	return ""
}
