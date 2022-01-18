package ledger_model

/*
 * Author: imuge
 * Date: 2020/5/28 下午6:39
 */

type TransactionService interface {
	Process(txRequest *TransactionRequest) (*TransactionResponse, error)
}
