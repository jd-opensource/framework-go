package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2021/9/10 下午4:24
 */

var _ binary_proto.DataContract = (*ContractStateUpdateOperation)(nil)

func init() {
	binary_proto.RegisterContract(ContractStateUpdateOperation{})
}

type ContractStateUpdateOperation struct {
	ContractAddress []byte       `primitiveType:"BYTES"`
	State           AccountState `refEnum:"788"`
}

func (u ContractStateUpdateOperation) ContractCode() int32 {
	return binary_proto.TX_OP_CONTRACT_STATE
}

func (u ContractStateUpdateOperation) ContractName() string {
	return OperationTypeContractStateUpdate
}

func (u ContractStateUpdateOperation) Description() string {
	return ""
}
