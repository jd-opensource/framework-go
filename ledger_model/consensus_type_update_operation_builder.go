package ledger_model

type ConsensusTypeUpdateOperationBuilder struct {
	factory *BlockchainOperationFactory
}

func NewConsensusTypeUpdateOperationBuilder(factory *BlockchainOperationFactory) *ConsensusTypeUpdateOperationBuilder {
	return &ConsensusTypeUpdateOperationBuilder{factory: factory}
}

func (ctuob *ConsensusTypeUpdateOperationBuilder) Update(providerName string, properties []Property) ConsensusTypeUpdateOperation {
	len := len(properties)
	pss := make([][]byte, len)
	for i := 0; i < len; i++ {
		pss[i] = properties[i].ToBytes()
	}
	operation := ConsensusTypeUpdateOperation{
		ProviderName: providerName,
		Properties:   pss,
	}
	if ctuob.factory != nil {
		ctuob.factory.addOperation(operation)
	}

	return operation
}

func (ctuob *ConsensusTypeUpdateOperationBuilder) UpdateWithConfigFile(providerName string, configFile string) ConsensusTypeUpdateOperation {
	return ctuob.Update(providerName, LoadProperties(configFile))
}
