package ledger_model

type CryptoOperator interface {
	SwitchHashAlgo() *CryptoHashAlgoUpdateOperationBuilder
}
