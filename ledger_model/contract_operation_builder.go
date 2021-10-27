package ledger_model

type ContractOperationBuilder struct {
	invokeBuilder     *ContractEventSendOperationBuilder
	updateBuilder     *ContractUpdateOperationBuilder
	permissionBuilder *AccountPermissionSetOperationBuilder
}

func NewContractOperationBuilder(address []byte, factory *BlockchainOperationFactory) *ContractOperationBuilder {
	return &ContractOperationBuilder{
		invokeBuilder:     NewContractEventSendOperationBuilder(address, factory),
		updateBuilder:     NewContractUpdateOperationBuilder(address, factory),
		permissionBuilder: NewAccountPermissionSetOperationBuilder(CONTRACT, address, factory),
	}
}

func (cob *ContractOperationBuilder) Invoke(event string, args ...interface{}) error {
	return cob.invokeBuilder.Send(-1, event, args)
}

func (cob *ContractOperationBuilder) State(state AccountState) {
	cob.updateBuilder.State(state)
}

func (eaob *ContractOperationBuilder) Permission() *AccountPermissionSetOperationBuilder {
	return eaob.permissionBuilder
}
