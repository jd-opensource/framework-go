package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/28 下午2:18
 */

type RolePrivileges struct {
	RoleName             string
	Version              int64
	LedgerPrivilege      LedgerPrivilegeBitset
	TransactionPrivilege TransactionPrivilegeBitset
}
