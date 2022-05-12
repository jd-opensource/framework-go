package ledger_model

type ConsensusSettingsUpdateOperationBuilder struct {
	factory *BlockchainOperationFactory
}

func NewConsensusSettingsUpdateOperationBuilder(factory *BlockchainOperationFactory) *ConsensusSettingsUpdateOperationBuilder {
	return &ConsensusSettingsUpdateOperationBuilder{factory: factory}
}

func (ctuob *ConsensusSettingsUpdateOperationBuilder) UpdateWithConfigFile(provider string, configFile string) ConsensusSettingsUpdateOperation {
	properties := LoadProperties(configFile)
	len := len(properties)
	pss := make([][]byte, len)
	for i := 0; i < len; i++ {
		pss[i] = properties[i].ToBytes()
	}
	operation := ConsensusSettingsUpdateOperation{
		Properties: pss,
		Provider:   provider,
	}
	if ctuob.factory != nil {
		ctuob.factory.addOperation(operation)
	}

	return operation
}

func (ctuob *ConsensusSettingsUpdateOperationBuilder) Update(provider string, properties []Property) ConsensusSettingsUpdateOperation {
	len := len(properties)
	pss := make([][]byte, len)
	for i := 0; i < len; i++ {
		pss[i] = properties[i].ToBytes()
	}
	operation := ConsensusSettingsUpdateOperation{
		Provider:   provider,
		Properties: pss,
	}
	if ctuob.factory != nil {
		ctuob.factory.addOperation(operation)
	}

	return operation
}
