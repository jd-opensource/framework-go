package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/28 下午4:51
 */

var _ binary_proto.DataContract = (*DigitalSignatureBody)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(DigitalSignatureBody{})
}

type DigitalSignatureBody struct {
	PubKey []byte `primitiveType:"BYTES"`
	Digest []byte `primitiveType:"BYTES"`
}

func (d DigitalSignatureBody) ContractCode() int32 {
	return binary_proto.DIGITALSIGNATURE_BODY
}

func (d DigitalSignatureBody) ContractName() string {
	return "DigitalSignatureBody"
}

func (d DigitalSignatureBody) Description() string {
	return ""
}
