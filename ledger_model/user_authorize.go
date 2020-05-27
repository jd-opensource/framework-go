package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/1 下午1:26
 */

type UserAuthorize interface {

	 ForUser(users [][]byte) UserRolesAuthorizer

}
