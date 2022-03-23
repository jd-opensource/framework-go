package ledger_model

import (
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
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

func (t *TxTemplate) SwitchHashAlgo() *CryptoHashAlgoUpdateOperationBuilder {
	t.stateManager.operate()
	return t.txBuilder.SwitchHashAlgo()
}

func (t *TxTemplate) SwitchSettings() *ConsensusTypeUpdateOperationBuilder {
	t.stateManager.operate()
	return t.txBuilder.SwitchSettings()
}

func (t *TxTemplate) MetaInfo() *MetaInfoUpdateOperationBuilder {
	t.stateManager.operate()
	return t.txBuilder.MetaInfo()
}

func (t *TxTemplate) User(address []byte) *UserUpdateOperationBuilder {
	t.stateManager.operate()
	return t.txBuilder.User(address)
}

func (t *TxTemplate) Contract(address []byte) *ContractOperationBuilder {
	t.stateManager.operate()
	return t.txBuilder.Contract(address)
}

func (t *TxTemplate) ContractEvents() *ContractEventSendOperationBuilder {
	t.stateManager.operate()
	return t.txBuilder.ContractEvents()
}

func (t *TxTemplate) EventAccounts() *EventAccountRegisterOperationBuilder {
	t.stateManager.operate()
	return t.txBuilder.EventAccounts()
}

func (t *TxTemplate) EventAccount(accountAddress []byte) *EventAccountOperationBuilder {
	t.stateManager.operate()
	return t.txBuilder.EventAccount(accountAddress)
}

func (t *TxTemplate) Participants() *ParticipantRegisterOperationBuilder {
	t.stateManager.operate()
	return t.txBuilder.Participants()
}

func (t *TxTemplate) States() *ParticipantStateUpdateOperationBuilder {
	t.stateManager.operate()
	return t.txBuilder.States()
}

func NewTxTemplate(ledgerHash *framework.HashDigest, hashAlgorithm framework.CryptoAlgorithm, txService TransactionService) *TxTemplate {
	return &TxTemplate{
		txBuilder:    NewTxBuilder(ledgerHash, hashAlgorithm),
		stateManager: NewTxStateManager(),
		txService:    txService,
	}
}

func (t *TxTemplate) Contracts() *ContractCodeDeployOperationBuilder {
	t.stateManager.operate()
	return t.txBuilder.Contracts()
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

func (t *TxTemplate) DataAccount(accountAddress []byte) *DataAccountOperationBuilder {
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
