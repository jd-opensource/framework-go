package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

var _ binary_proto.DataContract = (*ConsensusTypeUpdateOperation)(nil)

func init() {
	binary_proto.RegisterContract(ConsensusTypeUpdateOperation{})
}

type ConsensusTypeUpdateOperation struct {
	ProviderName string   `primitiveType:"TEXT"`
	Properties   [][]byte `primitiveType:"BYTES" list:"true"`
}

func (c ConsensusTypeUpdateOperation) ContractCode() int32 {
	return binary_proto.TX_OP_CONSENSUS_TYPE_UPDATE
}

func (c ConsensusTypeUpdateOperation) ContractName() string {
	return OperationTypeConsensusTypeUpdateOperation
}

func (c ConsensusTypeUpdateOperation) Description() string {
	return ""
}
