package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/6/1 下午1:27
 */

var _ binary_proto.DataContract = (*UserAuthorizeOperation)(nil)

func init() {
	binary_proto.RegisterContract(UserRolesEntry{})
	binary_proto.RegisterContract(UserAuthorizeOperation{})
}

// 角色配置操作
type UserAuthorizeOperation struct {
	UserRolesAuthorizations []UserRolesEntry `refContract:"884" list:"true"`
}

func (u UserAuthorizeOperation) ContractCode() int32 {
	return binary_proto.TX_OP_USER_ROLES_AUTHORIZE
}

func (u UserAuthorizeOperation) ContractName() string {
	return OperationTypeUserAuthorizeOperation
}

func (u UserAuthorizeOperation) Description() string {
	return ""
}

var _ binary_proto.DataContract = (*UserRolesEntry)(nil)

type UserRolesEntry struct {
	// 用户地址
	Addresses [][]byte `primitiveType:"BYTES" list:"true"`

	// 要更新的多角色权限策略
	Policy RolesPolicy `refEnum:"1027"`

	// 授权的角色清单
	AuthorizedRoles []string `primitiveType:"TEXT" list:"true"`

	// 取消授权的角色清单
	UnauthorizedRoles []string `primitiveType:"TEXT" list:"true"`
}

func (u UserRolesEntry) ContractCode() int32 {
	return binary_proto.TX_OP_USER_ROLE_AUTHORIZE_ENTRY
}

func (u UserRolesEntry) ContractName() string {
	return "UserRolesEntry"
}

func (u UserRolesEntry) Description() string {
	return ""
}
