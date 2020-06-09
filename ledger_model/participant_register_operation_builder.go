package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/3 下午5:13
 */

type ParticipantRegisterOperationBuilder struct {
	factory *BlockchainOperationFactory
}

func NewParticipantRegisterOperationBuilder(factory *BlockchainOperationFactory) *ParticipantRegisterOperationBuilder {
	return &ParticipantRegisterOperationBuilder{factory: factory}
}

func (prob *ParticipantRegisterOperationBuilder) Register(participantName string, participantPubKey BlockchainIdentity, networkAddress []byte) ParticipantRegisterOperation {
	operation := ParticipantRegisterOperation{
		ParticipantName:             participantName,
		ParticipantRegisterIdentity: participantPubKey,
		NetworkAddress:              networkAddress,
	}
	if prob.factory != nil {
		prob.factory.addOperation(operation)
	}

	return operation
}
