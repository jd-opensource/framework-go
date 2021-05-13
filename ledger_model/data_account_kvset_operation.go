package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/28 下午5:07
 */

var _ binary_proto.DataContract = (*DataAccountKVSetOperation)(nil)

func init() {
	binary_proto.RegisterContract(DataAccountKVSetOperation{})
}

type DataAccountKVSetOperation struct {
	AccountAddress []byte         `primitiveType:"BYTES"`
	WriteSet       []KVWriteEntry `refContract:"802" list:"true"`
}

func (d DataAccountKVSetOperation) ContractCode() int32 {
	return binary_proto.TX_OP_DATA_ACC_SET
}

func (d DataAccountKVSetOperation) ContractName() string {
	return OperationTypeDataAccountKVSetOperation
}

func (d DataAccountKVSetOperation) Description() string {
	return ""
}
