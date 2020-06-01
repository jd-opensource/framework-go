package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/1 下午7:25
 */

type ContractCodeDeployOperationBuilder struct {
	factory *BlockchainOperationFactory
}

func NewContractCodeDeployOperationBuilder(factory *BlockchainOperationFactory) *ContractCodeDeployOperationBuilder {
	return &ContractCodeDeployOperationBuilder{factory: factory}
}

func (cdob *ContractCodeDeployOperationBuilder) Deploy(id BlockchainIdentity, chainCode []byte) ContractCodeDeployOperation {
	operation := ContractCodeDeployOperation{
		ContractID: id,
		ChainCode:  chainCode,
	}
	if cdob.factory != nil {
		cdob.factory.addOperation(operation)
	}

	return operation
}
