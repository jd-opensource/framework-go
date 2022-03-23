package ledger_model

type ConsensusOperator interface {
	Consensus() *ConsensusSettingsUpdateOperationBuilder
}
