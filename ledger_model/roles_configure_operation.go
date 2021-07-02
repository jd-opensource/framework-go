package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/29 下午5:30
 */

var _ binary_proto.DataContract = (*RolesConfigureOperation)(nil)

func init() {
	binary_proto.RegisterContract(RolePrivilegeEntry{})
	binary_proto.RegisterContract(RolesConfigureOperation{})
}

type RolesConfigureOperation struct {
	Roles []RolePrivilegeEntry `refContract:"882" list:"true"`
}

func (r RolesConfigureOperation) ContractCode() int32 {
	return binary_proto.TX_OP_ROLE_CONFIGURE
}

func (r RolesConfigureOperation) ContractName() string {
	return OperationTypeRolesConfigureOperation
}

func (r RolesConfigureOperation) Description() string {
	return ""
}

var _ binary_proto.DataContract = (*RolePrivilegeEntry)(nil)

type RolePrivilegeEntry struct {
	RoleName string `primitiveType:"TEXT"`

	EnableLedgerPermissions []LedgerPermission `refEnum:"1026" list:"true"`

	DisableLedgerPermissions []LedgerPermission `refEnum:"1026" list:"true"`

	EnableTransactionPermissions []TransactionPermission `refEnum:"1025" list:"true"`

	DisableTransactionPermissions []TransactionPermission `refEnum:"1025" list:"true"`
}

func (r RolePrivilegeEntry) ContractCode() int32 {
	return binary_proto.TX_OP_ROLE_CONFIGURE_ENTRY
}

func (r RolePrivilegeEntry) ContractName() string {
	return "RolePrivilegeEntry"
}

func (r RolePrivilegeEntry) Description() string {
	return ""
}
