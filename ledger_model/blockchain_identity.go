package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/25 下午6:56
 */

var _ binary_proto.DataContract = (*BlockchainIdentity)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(BlockchainIdentity{})
}

type BlockchainIdentity struct {
	Address []byte `primitiveType:"BYTES"`
	PubKey  []byte `primitiveType:"BYTES"`
}

func (b BlockchainIdentity) ContractCode() int32 {
	return binary_proto.BLOCK_CHAIN_IDENTITY
}

func (b BlockchainIdentity) ContractName() string {
	return "BlockchainIdentity"
}

func (b BlockchainIdentity) Description() string {
	return ""
}
