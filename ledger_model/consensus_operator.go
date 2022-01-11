package ledger_model

type ConsensusOperator interface {
	SwitchSettings() *ConsensusTypeUpdateOperationBuilder
}
