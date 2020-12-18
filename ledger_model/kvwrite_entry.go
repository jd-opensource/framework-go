package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/28 下午5:10
 */

var _ binary_proto.DataContract = (*KVWriteEntry)(nil)

func init() {
	binary_proto.RegisterContract(KVWriteEntry{})
}

type KVWriteEntry struct {
	Key             string     `primitiveType:"TEXT"`
	Value           BytesValue `refContract:"128"`
	ExpectedVersion int64      `primitiveType:"INT64"`
}

func (K KVWriteEntry) ContractCode() int32 {
	return binary_proto.TX_OP_DATA_ACC_SET_KV
}

func (K KVWriteEntry) ContractName() string {
	return "KVWriteEntry"
}

func (K KVWriteEntry) Description() string {
	return ""
}
