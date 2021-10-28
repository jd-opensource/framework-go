package ledger_model

type MetaInfoUpdateOperationBuilder struct {
	factory *BlockchainOperationFactory
}

func NewMetaInfoUpdateOperationBuilder(factory *BlockchainOperationFactory) *MetaInfoUpdateOperationBuilder {
	return &MetaInfoUpdateOperationBuilder{
		factory: factory,
	}
}

func (uuob *MetaInfoUpdateOperationBuilder) CA() *RootCAUpdateOperationBuilder {
	return NewRootCAUpdateOperationBuilder(uuob.factory)
}
