package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/1 下午1:28
 */

type UserAuthorizer interface {
	UserAuthorize

	getOperation() *UserAuthorizeOperation
}
