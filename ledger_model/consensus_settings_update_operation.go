package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

var _ binary_proto.DataContract = (*ConsensusSettingsUpdateOperation)(nil)

type ConsensusSettingsUpdateOperation struct {
	Properties []Property
}

func (c ConsensusSettingsUpdateOperation) ContractCode() int32 {
	return binary_proto.TX_OP_CONSENSUS_SETTINGS_UPDATE
}

func (c ConsensusSettingsUpdateOperation) ContractName() string {
	return OperationTypeConsensusSettingsUpdateOperation
}

func (c ConsensusSettingsUpdateOperation) Description() string {
	return ""
}

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
