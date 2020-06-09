package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/9 下午6:18
 */

type EventAccountRegisterOperationBuilder struct {
	factory *BlockchainOperationFactory
}

func NewEventAccountRegisterOperationBuilder(factory *BlockchainOperationFactory) *EventAccountRegisterOperationBuilder {
	return &EventAccountRegisterOperationBuilder{factory: factory}
}

func (earob *EventAccountRegisterOperationBuilder) Register(accountIdentity BlockchainIdentity) EventAccountRegisterOperation {
	operation := EventAccountRegisterOperation{
		EventAccountID: accountIdentity,
	}
	if earob.factory != nil {
		earob.factory.addOperation(operation)
	}

	return operation
}
