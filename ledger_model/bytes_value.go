package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午2:57
 */

var _ binary_proto.DataContract = (*BytesValue)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(BytesValue{})
}

type BytesValue struct {
	// 数据类型
	Type DataType `refEnum:"2851"`
	// 数据值的二进制序列
	Bytes []byte `primitiveType:"BYTES"`
}

func (b BytesValue) ContractCode() int32 {
	return binary_proto.BYTES_VALUE
}

func (b BytesValue) ContractName() string {
	return "BytesValue"
}

func (b BytesValue) Description() string {
	return ""
}
