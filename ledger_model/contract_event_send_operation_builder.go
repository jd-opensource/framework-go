package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/1 下午7:25
 */

type ContractEventSendOperationBuilder struct {
	factory *BlockchainOperationFactory
}

func NewContractEventSendOperationBuilder(factory *BlockchainOperationFactory) *ContractEventSendOperationBuilder {
	return &ContractEventSendOperationBuilder{factory: factory}
}

/*
	address 合约地址
	event   合约方法
	args	参数列表
 */
func (cesob *ContractEventSendOperationBuilder) Send(address []byte, event string, args BytesValueList) ContractEventSendOperation {
	operation := ContractEventSendOperation{
		ContractAddress: address,
		Event:           event,
		Args:            args,
		Version:         0,
	}
	if cesob.factory != nil {
		cesob.factory.addOperation(operation)
	}

	return operation
}
