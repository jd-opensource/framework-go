package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/26 下午1:45
 */

var _ binary_proto.DataContract = (*DigitalSignature)(nil)

func init() {
	binary_proto.RegisterContract(DigitalSignature{})
}

type DigitalSignature struct {
	DigitalSignatureBody
}

func (d DigitalSignature) ContractCode() int32 {
	return binary_proto.DIGITALSIGNATURE
}

func (d DigitalSignature) ContractName() string {
	return "DigitalSignature"
}

func (d DigitalSignature) Description() string {
	return ""
}
