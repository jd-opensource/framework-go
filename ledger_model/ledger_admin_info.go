package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午1:24
 */

var _ binary_proto.DataContract = (*LedgerAdminInfo)(nil)

func init() {
	binary_proto.RegisterContract(LedgerAdminInfo{})
}

type LedgerAdminInfo struct {
	Metadata         LedgerMetadata_V2 `refContract:"1537"`
	Settings         LedgerSettings    `refContract:"1568"`
	Participants     []ParticipantNode `refContract:"1569" list:"true"`
	ParticipantCount int64             `primitiveType:"INT64"`
}

func (l LedgerAdminInfo) ContractCode() int32 {
	return binary_proto.LEDGER_ADMIN_INFO
}

func (l LedgerAdminInfo) ContractName() string {
	return "LedgerAdminInfo"
}

func (l LedgerAdminInfo) Description() string {
	return ""
}
