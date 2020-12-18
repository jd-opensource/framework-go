package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午3:59
 */

var _ binary_proto.DataContract = (*MerkleSnapshot)(nil)

func init() {
	binary_proto.RegisterContract(MerkleSnapshot{})
}

type MerkleSnapshot struct {
	RootHash []byte `primitiveType:"BYTES"`
}

func (m MerkleSnapshot) ContractCode() int32 {
	return binary_proto.MERKLE_SNAPSHOT
}

func (m MerkleSnapshot) ContractName() string {
	return "MerkleSnapshot"
}

func (m MerkleSnapshot) Description() string {
	return ""
}
