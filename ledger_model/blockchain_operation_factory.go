package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/29 上午10:52
 */

var _ ClientOperator = (*BlockchainOperationFactory)(nil)

type BlockchainOperationFactory struct {
	operationList []binary_proto.DataContract
}

func (b *BlockchainOperationFactory) Security() *SecurityOperationBuilder {
	return NewSecurityOperationBuilder(b)
}

func (b *BlockchainOperationFactory) Users() *UserRegisterOperationBuilder {
	return NewUserRegisterOperationBuilder(b)
}

func (b *BlockchainOperationFactory) DataAccounts() *DataAccountRegisterOperationBuilder {
	return NewDataAccountRegisterOperationBuilder(b)
}

func (b *BlockchainOperationFactory) DataAccount(accountAddress []byte) *DataAccountKVSetOperationBuilder {
	return NewDataAccountKVSetOperationBuilder(accountAddress, b)
}

func (b *BlockchainOperationFactory) addOperation(operation binary_proto.DataContract) {
	b.operationList = append(b.operationList, operation)
}
