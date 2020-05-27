package ledger_model

/*
 * Author: imuge
 * Date: 2020/5/29 下午4:50
 */

type SecurityOperationBuilder struct {
	factory *BlockchainOperationFactory
}

func NewSecurityOperationBuilder(factory *BlockchainOperationFactory) *SecurityOperationBuilder {
	return &SecurityOperationBuilder{factory: factory}
}

func (sob *SecurityOperationBuilder) Roles() RolesConfigurer {
	opt := NewRolesConfigureOpTemplate()
	if sob.factory != nil {
		sob.factory.addOperation(opt.getOperation())
	}
	return opt
}

func (sob *SecurityOperationBuilder) Authorziations() UserAuthorizer {
	opt := NewUserAuthorizeOpTemplate()
	if sob.factory != nil {
		sob.factory.addOperation(opt.getOperation())
	}
	return opt
}
