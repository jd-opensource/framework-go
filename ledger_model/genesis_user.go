package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2021/9/10 下午4:00
 */

var _ binary_proto.DataContract = (*ContractInfo)(nil)

func init() {
	binary_proto.RegisterContract(ContractInfo{})
}

// 创世用户信息
type GenesisUser struct {
	PubKey      []byte      `primitiveType:"BYTES"`
	Certificate string      `primitiveType:"TEXT"`
	Roles       []string    `primitiveType:"TEXT" list:"true"`
	RolesPolicy RolesPolicy `refEnum:"1027"`
}

func (u GenesisUser) ContractCode() int32 {
	return binary_proto.METADATA_GENESIS_USER
}

func (u GenesisUser) ContractName() string {
	return "GenesisUser"
}

func (u GenesisUser) Description() string {
	return ""
}
