package ledger_model

/*
 * Author: imuge
 * Date: 2021/9/10 下午4:40
 */

type ContractUpdateOperationBuilder struct {
	factory *BlockchainOperationFactory
	address []byte
}

func NewContractUpdateOperationBuilder(address []byte, factory *BlockchainOperationFactory) *ContractUpdateOperationBuilder {
	return &ContractUpdateOperationBuilder{
		address: address,
		factory: factory,
	}
}

func (cuob *ContractUpdateOperationBuilder) State(state AccountState) ContractStateUpdateOperation {
	operation := ContractStateUpdateOperation{
		ContractAddress: cuob.address,
		State:           state,
	}
	if cuob.factory != nil {
		cuob.factory.addOperation(operation)
	}

	return operation
}
