package ledger_model

import (
	"github.com/blockchain-jd-com/framework-go/crypto/ca"
	"sync"
)

type RootCAUpdateOperationBuilder struct {
	factory   *BlockchainOperationFactory
	operation *RootCAUpdateOperation

	added bool // 该构建是否已经加入
	mutex sync.Mutex
}

func NewRootCAUpdateOperationBuilder(factory *BlockchainOperationFactory) *RootCAUpdateOperationBuilder {
	return &RootCAUpdateOperationBuilder{
		factory:   factory,
		operation: &RootCAUpdateOperation{},
	}
}

func (raob *RootCAUpdateOperationBuilder) addOperation() {
	raob.mutex.Lock()
	defer raob.mutex.Unlock()
	if !raob.added && raob.factory != nil {
		raob.factory.addOperation(raob.operation)
		raob.added = true
	}
}

func (raob *RootCAUpdateOperationBuilder) Add(cert *ca.Certificate) *RootCAUpdateOperationBuilder {
	raob.operation.CertificatesAdd = append(raob.operation.CertificatesAdd, cert.ToPEMString())
	raob.addOperation()

	return raob
}

func (raob *RootCAUpdateOperationBuilder) Update(cert *ca.Certificate) *RootCAUpdateOperationBuilder {
	raob.operation.CertificatesUpdate = append(raob.operation.CertificatesUpdate, cert.ToPEMString())
	raob.addOperation()

	return raob
}

func (raob *RootCAUpdateOperationBuilder) Remove(cert *ca.Certificate) *RootCAUpdateOperationBuilder {
	raob.operation.CertificatesRemove = append(raob.operation.CertificatesRemove, cert.ToPEMString())
	raob.addOperation()

	return raob
}
