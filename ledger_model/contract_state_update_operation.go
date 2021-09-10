package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2021/9/10 下午4:24
 */

var _ binary_proto.DataContract = (*UserStateUpdateOperation)(nil)

func init() {
	binary_proto.RegisterContract(UserStateUpdateOperation{})
}

type UserStateUpdateOperation struct {
	UserAddress []byte       `primitiveType:"BYTES"`
	State       AccountState `refEnum:"788"`
}

func (u UserStateUpdateOperation) ContractCode() int32 {
	return binary_proto.TX_OP_USER_STATE
}

func (u UserStateUpdateOperation) ContractName() string {
	return OperationTypeUserStateUpdate
}

func (u UserStateUpdateOperation) Description() string {
	return ""
}
