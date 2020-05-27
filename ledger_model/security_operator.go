package ledger_model

/*
 * Author: imuge
 * Date: 2020/5/29 下午4:48
 */

type SecurityOperator interface {
	Security() *SecurityOperationBuilder
}