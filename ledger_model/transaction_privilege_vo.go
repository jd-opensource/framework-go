package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/28 下午2:25
 */

type TransactionPrivilegeVO struct {
	Privilege       []TransactionPermission
	PermissionCount int32
}
