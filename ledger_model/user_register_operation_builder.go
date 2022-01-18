package ledger_model

import (
	ca2 "github.com/blockchain-jd-com/framework-go/crypto/ca"
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

func (urob *UserRegisterOperationBuilder) Register(userID *BlockchainIdentity) UserRegisterOperation {
	operation := UserRegisterOperation{
		UserID: userID,
	}
	if urob.factory != nil {
		urob.factory.addOperation(operation)
	}

	return operation
}

func (urob *UserRegisterOperationBuilder) RegisterWithCA(cert *ca2.Certificate) (*UserRegisterOperation, error) {
	key, err := ca.RetrievePubKey(cert)
	if err != nil {
		return nil, err
	}
	operation := &UserRegisterOperation{
		UserID:      NewBlockchainIdentity(key),
		Certificate: cert.ToPEMString(),
	}
	if urob.factory != nil {
		urob.factory.addOperation(operation)
	}

	return operation, nil
}
