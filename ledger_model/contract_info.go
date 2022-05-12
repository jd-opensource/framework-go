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
	*BlockchainIdentity
	MerkleSnapshot
	ChainCode        []byte         `primitiveType:"BYTES"`
	ChainCodeVersion int64          `primitiveType:"INT64"`
	State            AccountState   `refEnum:"788"`
	Permission       DataPermission `json:"permission"`
	Lang             ContractLang   `refEnum:"2561"`
}

func (c ContractInfo) ContractCode() int32 {
	return binary_proto.CONTRACT_ACCOUNT_HEADER
}

func (c ContractInfo) ContractName() string {
	return "ContractInfo"
}

func (c ContractInfo) Description() string {
	return ""
}
