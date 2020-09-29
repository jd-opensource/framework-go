package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午1:36
 */

var _ binary_proto.DataContract = (*LedgerMetadata_V2)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(LedgerMetadata_V2{})
}

type LedgerMetadata_V2 struct {
	LedgerMetadata
	// 角色权限集合的根哈希
	RolePrivilegesHash []byte `primitiveType:"BYTES"`
	// 用户角色授权集合的根哈希
	UserRolesHash      []byte `primitiveType:"BYTES"`
}

func (l LedgerMetadata_V2) ContractCode() int32 {
	return binary_proto.METADATA_V2
}

func (l LedgerMetadata_V2) ContractName() string {
	return "LEDGER-METADATA-V2"
}

func (l LedgerMetadata_V2) Description() string {
	return ""
}
