package ledger_model

/*
 * Author: imuge
 * Date: 2020/5/28 下午4:40
 */

type DataAccountOperator interface {

	// 数据账户
	DataAccounts() *DataAccountRegisterOperationBuilder

	// 写入数据
	DataAccount(accountAddress []byte) *DataAccountOperationBuilder
}
