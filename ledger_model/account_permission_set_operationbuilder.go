package ledger_model

import "sync"

type AccountPermissionSetOperationBuilder struct {
	operation *AccountPermissionSetOperation
	factory   *BlockchainOperationFactory

	added bool // 该构建是否已经加入
	mutex sync.Mutex
}

func NewAccountPermissionSetOperationBuilder(accountType AccountType, address []byte, factory *BlockchainOperationFactory) *AccountPermissionSetOperationBuilder {
	return &AccountPermissionSetOperationBuilder{
		factory: factory,
		operation: &AccountPermissionSetOperation{
			Address:     address,
			AccountType: accountType,
			Mode:        -1,
		},
	}
}

func (cuob *AccountPermissionSetOperationBuilder) addOperation() {
	cuob.mutex.Lock()
	defer cuob.mutex.Unlock()
	if !cuob.added && cuob.factory != nil {
		cuob.factory.addOperation(cuob.operation)
		cuob.added = true
	}
}

func (cuob *AccountPermissionSetOperationBuilder) Mode(mode int32) *AccountPermissionSetOperationBuilder {
	cuob.operation.Mode = mode
	cuob.addOperation()

	return cuob
}

func (cuob *AccountPermissionSetOperationBuilder) Role(role string) *AccountPermissionSetOperationBuilder {
	cuob.operation.Role = role
	cuob.addOperation()

	return cuob
}
