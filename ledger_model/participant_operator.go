package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/3 下午5:12
 */

type ParticipantOperator interface {

	// 注册参与方操作
	Participants() *ParticipantRegisterOperationBuilder

	// 参与方状态更新操作
	States() *ParticipantStateUpdateOperationBuilder
}
