package ledger_model

type CryptoHashAlgoUpdateOperationBuilder struct {
	factory *BlockchainOperationFactory
}

func NewCryptoHashAlgoUpdateOperationBuilder(factory *BlockchainOperationFactory) *CryptoHashAlgoUpdateOperationBuilder {
	return &CryptoHashAlgoUpdateOperationBuilder{factory: factory}
}

func (ctuob *CryptoHashAlgoUpdateOperationBuilder) Update(hashAlgoName string) CryptoHashAlgoUpdateOperation {
	operation := CryptoHashAlgoUpdateOperation{
		HashAlgoName: hashAlgoName,
	}
	if ctuob.factory != nil {
		ctuob.factory.addOperation(operation)
	}

	return operation
}
