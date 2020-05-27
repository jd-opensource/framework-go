package ledger_model

/*
 * Author: imuge
 * Date: 2020/5/29 下午4:58
 */

type RolePrivilegeConfigurer interface {
	RolesConfigure

	RoleName() string

	DisableTransactionPermission(permission TransactionPermission) RolePrivilegeConfigurer

	EnableTransactionPermission(permission TransactionPermission) RolePrivilegeConfigurer

	DisableLedgerPermission(permission LedgerPermission) RolePrivilegeConfigurer

	EnableLedgerPermission(permission LedgerPermission) RolePrivilegeConfigurer
}
