package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/6/3 下午5:15
 */

var _ binary_proto.DataContract = (*ParticipantRegisterOperation)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(ParticipantRegisterOperation{})
}

type ParticipantRegisterOperation struct {
	ParticipantName string `primitiveType:"TEXT"`

	ParticipantRegisterIdentity BlockchainIdentity `refContract:"144"`
}

func (p ParticipantRegisterOperation) ContractCode() int32 {
	return binary_proto.TX_OP_PARTICIPANT_REG
}

func (p ParticipantRegisterOperation) ContractName() string {
	return "ParticipantRegisterOperation"
}

func (p ParticipantRegisterOperation) Description() string {
	return ""
}
