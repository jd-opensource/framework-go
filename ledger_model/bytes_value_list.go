package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/9/17 下午2:51
 */

var _ binary_proto.DataContract = (*BytesValueList)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(BytesValueList{})
}

type BytesValueList struct {
	Values []BytesValue `refContract:"128" list:"true"`
}

func (b BytesValueList) ContractCode() int32 {
	return binary_proto.BYTES_VALUE_LIST
}

func (b BytesValueList) ContractName() string {
	return "BytesValueList"
}

func (b BytesValueList) Description() string {
	return ""
}
