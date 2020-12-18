package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午1:46
 */

var _ binary_proto.DataContract = (*LedgerSettings)(nil)

func init() {
	binary_proto.RegisterContract(LedgerSettings{})
}

type LedgerSettings struct {
	ConsensusProvider string        `primitiveType:"TEXT"`
	ConsensusSetting  []byte        `primitiveType:"BYTES"`
	CryptoSetting     CryptoSetting `refContract:"1602"`
}

func (l LedgerSettings) ContractCode() int32 {
	return binary_proto.METADATA_LEDGER_SETTING
}

func (l LedgerSettings) ContractName() string {
	return "LedgerSettings"
}

func (l LedgerSettings) Description() string {
	return ""
}
