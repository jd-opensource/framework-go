package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/9 下午4:54
 */

type ParticipantStateUpdateOperationBuilder struct {
	factory *BlockchainOperationFactory
}

func NewParticipantStateUpdateOperationBuilder(factory *BlockchainOperationFactory) *ParticipantStateUpdateOperationBuilder {
	return &ParticipantStateUpdateOperationBuilder{factory: factory}
}

func (psuob *ParticipantStateUpdateOperationBuilder) Update(blockchainIdentity BlockchainIdentity, networkAddress []byte, participantNodeState ParticipantNodeState) ParticipantStateUpdateOperation {
	operation := ParticipantStateUpdateOperation{
		StateUpdateIdentity: blockchainIdentity,
		NetworkAddress:      networkAddress,
		State:               participantNodeState,
	}
	if psuob.factory != nil {
		psuob.factory.addOperation(operation)
	}

	return operation
}
