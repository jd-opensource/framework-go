package ledger_model

import (
	"github.com/blockchain-jd-com/framework-go/crypto/ca"
	ca2 "github.com/blockchain-jd-com/framework-go/utils/ca"
)

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

func (prob *ParticipantRegisterOperationBuilder) Register(participantName string, participantPubKey BlockchainIdentity) ParticipantRegisterOperation {
	operation := ParticipantRegisterOperation{
		ParticipantName: participantName,
		ParticipantID:   participantPubKey,
	}
	if prob.factory != nil {
		prob.factory.addOperation(operation)
	}

	return operation
}

func (prob *ParticipantRegisterOperationBuilder) RegisterWithCA(participantName string, certificate *ca.Certificate) ParticipantRegisterOperation {
	operation := ParticipantRegisterOperation{
		ParticipantName: participantName,
		ParticipantID: NewBlockchainIdentity(ca2.RetrievePubKey(certificate)),
		Certificate: certificate.ToPEMString(),
	}
	if prob.factory != nil {
		prob.factory.addOperation(operation)
	}

	return operation
}
