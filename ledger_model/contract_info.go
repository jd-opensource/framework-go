package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午4:00
 */

var _ binary_proto.DataContract = (*ContractInfo)(nil)

func init() {
	binary_proto.RegisterContract(ContractInfo{})
}

type ContractInfo struct {
	BlockchainIdentity
	MerkleSnapshot
	ChainCode []byte       `primitiveType:"BYTES"`
	State     AccountState `refEnum:"788"`
}

func (c ContractInfo) ContractCode() int32 {
	return binary_proto.CONTRACT_INFO
}

func (c ContractInfo) ContractName() string {
	return "ContractInfo"
}

func (c ContractInfo) Description() string {
	return ""
}
