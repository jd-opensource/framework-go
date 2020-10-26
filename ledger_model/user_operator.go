package ledger_model

/*
 * Author: imuge
 * Date: 2020/5/28 下午4:32
 */

type UserOperator interface {
	Users() *UserRegisterOperationBuilder
}
