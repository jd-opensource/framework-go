package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/28 下午2:29
 */
type UserPrivilege struct {
	RoleSet RoleSet

	RolePrivilege []PrivilegeSetVO
}
