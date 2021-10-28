package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

var _ binary_proto.DataContract = (*UserCAUpdateOperation)(nil)

func init() {
	binary_proto.RegisterContract(UserCAUpdateOperation{})
}

// 根证书更新
type UserCAUpdateOperation struct {
	UserAddress []byte `primitiveType:"BYTES"`
	Certificate string `primitiveType:"TEXT"`
}

func (u UserCAUpdateOperation) ContractCode() int32 {
	return binary_proto.TX_OP_USER_CA_UPDATE
}

func (u UserCAUpdateOperation) ContractName() string {
	return OperationTypeUserCAUpdateOperation
}

func (u UserCAUpdateOperation) Description() string {
	return ""
}
