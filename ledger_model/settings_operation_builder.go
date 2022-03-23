package ledger_model

type SettingsOperationBuilder struct {
	factory *BlockchainOperationFactory
}

func NewSettingsOperationBuilder(factory *BlockchainOperationFactory) *SettingsOperationBuilder {
	return &SettingsOperationBuilder{factory: factory}
}

func (ctuob *SettingsOperationBuilder) HashAlgorithm(hashAlgoName string) HashAlgorithmUpdateOperation {
	operation := HashAlgorithmUpdateOperation{
		Algorithm: hashAlgoName,
	}
	if ctuob.factory != nil {
		ctuob.factory.addOperation(operation)
	}

	return operation
}
