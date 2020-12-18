package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午1:38
 */

var _ binary_proto.DataContract = (*LedgerMetadata)(nil)

func init() {
	binary_proto.RegisterContract(LedgerMetadata{})
}

// 账本的元数据
type LedgerMetadata struct {
	// 账本的初始化种子
	Seed []byte `primitiveType:"BYTES"`
	// 共识参与方的默克尔树的根
	ParticipantsHash []byte `primitiveType:"BYTES"`
	// 账本配置的哈希
	SettingsHash []byte `primitiveType:"BYTES"`
}

func (l LedgerMetadata) ContractCode() int32 {
	return binary_proto.METADATA
}

func (l LedgerMetadata) ContractName() string {
	return "LEDGER-METADATA"
}

func (l LedgerMetadata) Description() string {
	return ""
}
