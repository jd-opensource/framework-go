package ledger_model

/*
 * Author: imuge
 * Date: 2020/5/29 下午5:29
 */

type RolesConfigurer interface {
	RolesConfigure

	getOperation() *RolesConfigureOperation
}