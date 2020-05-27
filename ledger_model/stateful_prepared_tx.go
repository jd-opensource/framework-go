package ledger_model

/*
 * Author: imuge
 * Date: 2020/5/29 下午4:13
 */

var _ PreparedTransaction = (*StatefulPreparedTx)(nil)

type StatefulPreparedTx struct {
	*PreparedTx

	stateManager *TxStateManager
}

func NewStatefulPreparedTx(tx *PreparedTx, stateManager *TxStateManager) *StatefulPreparedTx {
	return &StatefulPreparedTx{
		PreparedTx:   tx,
		stateManager: stateManager,
	}
}

func (s *StatefulPreparedTx) Commit() (TransactionResponse, error) {
	s.stateManager.commit()
	return s.PreparedTx.Commit()
}
