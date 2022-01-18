package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

var _ binary_proto.DataContract = (*LedgerInitSetting)(nil)

func init() {
	binary_proto.RegisterContract(LedgerInitSetting{})
}

// 账本初始配置
type LedgerInitSetting struct {
	LedgerSeed             []byte              `primitiveType:"BYTES"`
	ConsensusParticipants  []*ParticipantNode  `refContract:"1569" list:"true"`
	CryptoSetting          CryptoSetting       `refContract:"1602"`
	ConsensusProvider      string              `primitiveType:"TEXT"`
	ConsensusSettings      []byte              `primitiveType:"BYTES"`
	CreatedTime            int64               `primitiveType:"INT64"`
	LedgerStructureVersion int64               `primitiveType:"INT64"`
	IdentityMode           IdentityMode        `refEnum:"1604"`
	LedgerCertificates     []string            `primitiveType:"TEXT" list:"true"`
	GenesisUsers           []GenesisUser       `refContract:"1605" list:"true"`
	LedgerDataStructure    LedgerDataStructure `refEnum:"1606"`
}

func (ls LedgerInitSetting) ContractCode() int32 {
	return binary_proto.METADATA_INIT_SETTING
}

func (ls LedgerInitSetting) ContractName() string {
	return OperationTypeRootCAUpdateOperation
}

func (ls LedgerInitSetting) Description() string {
	return ""
}
