package ledger_model

import "github.com/blockchain-jd-com/framework-go/crypto/ca"

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

func (uuob *UserUpdateOperationBuilder) CA(cert *ca.Certificate) UserCAUpdateOperation {
	operation := UserCAUpdateOperation{
		UserAddress: uuob.address,
		Certificate: cert.ToPEMString(),
	}
	if uuob.factory != nil {
		uuob.factory.addOperation(operation)
	}

	return operation
}
