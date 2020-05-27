package ledger_model

import (
	"framework-go/crypto/framework"
)

/*
 * Author: imuge
 * Date: 2020/5/28 下午6:33
 */

var _ TransactionTemplate = (*TxTemplate)(nil)

type TxTemplate struct {
	txBuilder *TxBuilder

	txService TransactionService

	stateManager *TxStateManager
}

func NewTxTemplate(ledgerHash framework.HashDigest, txService TransactionService) *TxTemplate {
	return &TxTemplate{
		txBuilder:    NewTxBuilder(ledgerHash),
		stateManager: NewTxStateManager(),
		txService:    txService,
	}
}

func (t *TxTemplate) Security() *SecurityOperationBuilder {
	t.stateManager.operate()
	return t.txBuilder.Security()
}

func (t *TxTemplate) Users() *UserRegisterOperationBuilder {
	t.stateManager.operate()
	return t.txBuilder.Users()
}

func (t *TxTemplate) DataAccounts() *DataAccountRegisterOperationBuilder {
	t.stateManager.operate()
	return t.txBuilder.DataAccounts()
}

func (t *TxTemplate) DataAccount(accountAddress []byte) *DataAccountKVSetOperationBuilder {
	t.stateManager.operate()
	return t.txBuilder.DataAccount(accountAddress)
}

func (t *TxTemplate) GetLedgerHash() framework.HashDigest {
	return t.txBuilder.GetLedgerHash()
}

func (t *TxTemplate) Prepare() PreparedTransaction {
	t.stateManager.prepare()
	txReqBuilder := t.txBuilder.PrepareRequestNow()
	return NewStatefulPreparedTx(NewPreparedTx(txReqBuilder, t.txService), t.stateManager)
}
