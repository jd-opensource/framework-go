package ledger_model

import (
	"github.com/blockchain-jd-com/framework-go/utils/ca"
)

/*
 * Author: imuge
 * Date: 2020/5/28 下午4:33
 */

type UserRegisterOperationBuilder struct {
	factory *BlockchainOperationFactory
}

func NewUserRegisterOperationBuilder(factory *BlockchainOperationFactory) *UserRegisterOperationBuilder {
	return &UserRegisterOperationBuilder{factory: factory}
}

func (urob *UserRegisterOperationBuilder) Register(userID BlockchainIdentity) UserRegisterOperation {
	operation := UserRegisterOperation{
		UserID: userID,
	}
	if urob.factory != nil {
		urob.factory.addOperation(operation)
	}

	return operation
}

func (urob *UserRegisterOperationBuilder) RegisterWithCA(cert ca.Certificate) UserRegisterOperation {
	operation := UserRegisterOperation{
		UserID: NewBlockchainIdentity(cert.ResolvePubKey()),
		//Certificate: ca.ToPEMString(cert),
	}
	if urob.factory != nil {
		urob.factory.addOperation(operation)
	}

	return operation
}
