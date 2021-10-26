package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

var _ binary_proto.DataContract = (*AccountPermissionSetOperation)(nil)

func init() {
	binary_proto.RegisterContract(AccountPermissionSetOperation{})
}

type AccountPermissionSetOperation struct {
	Address     []byte      `primitiveType:"BYTES"`
	AccountType AccountType `refEnum:"3330"`
	Mode        int32       `primitiveType:"INT32"`
	Role        string      `primitiveType:"TEXT"`
}

func (o AccountPermissionSetOperation) ContractCode() int32 {
	return binary_proto.TX_OP_ACC_PERMISSION_SET
}

func (o AccountPermissionSetOperation) ContractName() string {
	return OperationTypeAccountPermissionSetOperation
}

func (o AccountPermissionSetOperation) Description() string {
	return ""
}
