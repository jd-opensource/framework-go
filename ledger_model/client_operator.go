package ledger_model

/*
 * Author: imuge
 * Date: 2020/5/28 下午5:31
 */

type ClientOperator interface {
	UserOperator
	DataAccountOperator
	SecurityOperator
	ContractOperator
	ParticipantOperator
	EventOperator
	MetaInfoOperator
	ConsensusOperator
	CryptoOperator
}
