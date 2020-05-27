package ledger_model

/*
 * Author: imuge
 * Date: 2020/5/29 下午4:57
 */

type RolesConfigure interface {
	Configure(roleName string) RolePrivilegeConfigurer
}
