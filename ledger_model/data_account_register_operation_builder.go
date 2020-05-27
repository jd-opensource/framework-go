package ledger_model

/*
 * Author: imuge
 * Date: 2020/5/28 下午4:45
 */

type DataAccountRegisterOperationBuilder struct {
	factory *BlockchainOperationFactory
}

func NewDataAccountRegisterOperationBuilder(factory *BlockchainOperationFactory) *DataAccountRegisterOperationBuilder {
	return &DataAccountRegisterOperationBuilder{factory: factory}
}

func (drob *DataAccountRegisterOperationBuilder) Register(id BlockchainIdentity) DataAccountRegisterOperation {
	operation := DataAccountRegisterOperation{
		AccountID: id,
	}
	if drob.factory != nil {
		drob.factory.addOperation(operation)
	}

	return operation
}
