package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

var _ binary_proto.DataContract = (*RootCAUpdateOperation)(nil)

func init() {
	binary_proto.RegisterContract(RootCAUpdateOperation{})
}

// 用户证书更新
type RootCAUpdateOperation struct {
	CertificatesAdd    []string `primitiveType:"TEXT" list:"true" json:"certificatesAdd"`
	CertificatesUpdate []string `primitiveType:"TEXT" list:"true" json:"certificatesUpdate"`
	CertificatesRemove []string `primitiveType:"TEXT" list:"true" json:"certificatesRemove"`
}

func (u RootCAUpdateOperation) ContractCode() int32 {
	return binary_proto.TX_OP_META_CA_UPDATE
}

func (u RootCAUpdateOperation) ContractName() string {
	return OperationTypeRootCAUpdateOperation
}

func (u RootCAUpdateOperation) Description() string {
	return ""
}
