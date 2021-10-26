package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/9 下午6:16
 */

type EventOperator interface {
	// 事件账户
	EventAccounts() *EventAccountRegisterOperationBuilder
	// 发布消息
	EventAccount(accountAddress []byte) *EventAccountOperationBuilder
}
