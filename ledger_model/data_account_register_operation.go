package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/28 下午4:47
 */

var _ binary_proto.DataContract = (*DataAccountRegisterOperation)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(DataAccountRegisterOperation{})
}

type DataAccountRegisterOperation struct {
	AccountID        BlockchainIdentity `refContract:"144"`
	AddressSignature DigitalSignature   `refContract:"2864"`
}

func (d DataAccountRegisterOperation) ContractCode() int32 {
	return binary_proto.TX_OP_DATA_ACC_REG
}

func (d DataAccountRegisterOperation) ContractName() string {
	return "DataAccountRegisterOperation"
}

func (d DataAccountRegisterOperation) Description() string {
	return ""
}
