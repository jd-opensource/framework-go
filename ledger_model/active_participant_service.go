package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/29 下午2:54
 */

type ActiveParticipantService interface {
	ActivateParticipant(ledgerHash string) (TransactionResponse, error)
}
