package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/28 下午2:18
 */

type PrivilegeSetVO struct {
	RoleName string

	LedgerPrivilege LedgerPrivilegeVO

	TransactionPrivilege TransactionPrivilegeVO
}
