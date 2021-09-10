package ledger_model

/*
 * Author: imuge
 * Date: 2021/9/10 下午4:33
 */

type UserUpdateOperationBuilder struct {
	factory *BlockchainOperationFactory
	address []byte
}

func NewUserUpdateOperationBuilder(address []byte, factory *BlockchainOperationFactory) *UserUpdateOperationBuilder {
	return &UserUpdateOperationBuilder{
		address: address,
		factory: factory,
	}
}

func (uuob *UserUpdateOperationBuilder) State(state AccountState) UserStateUpdateOperation {
	operation := UserStateUpdateOperation{
		UserAddress: uuob.address,
		State:       state,
	}
	if uuob.factory != nil {
		uuob.factory.addOperation(operation)
	}

	return operation
}
