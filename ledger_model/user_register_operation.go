package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/25 下午6:52
 */

var _ binary_proto.DataContract = (*UserRegisterOperation)(nil)

func init() {
	binary_proto.RegisterContract(UserRegisterOperation{})
}

type UserRegisterOperation struct {
	UserID BlockchainIdentity `refContract:"144"`
}

func (u UserRegisterOperation) ContractCode() int32 {
	return binary_proto.TX_OP_USER_REG
}

func (u UserRegisterOperation) ContractName() string {
	return OperationTypeUserRegister
}

func (u UserRegisterOperation) Description() string {
	return ""
}
