package ledger_model

type MetaInfoOperator interface {
	MetaInfo() *MetaInfoUpdateOperationBuilder
}
