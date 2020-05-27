package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/1 下午1:33
 */

type UserRolesAuthorizer interface {
	UserAuthorize

	Authorize(role string) UserRolesAuthorizer

	Unauthorize(role string) UserRolesAuthorizer

	SetPolicy(rolePolicy RolesPolicy) UserRolesAuthorizer
}
