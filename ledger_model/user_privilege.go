package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/28 下午2:29
 */
type UserRolesPrivileges struct {
	UserAddress                 []byte
	UserRoles                   []string
	LedgerPrivilegesBitset      LedgerPrivilegeBitset
	TransactionPrivilegesBitset TransactionPrivilegeBitset
}
