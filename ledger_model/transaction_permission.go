package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/29 下午5:03
 */

type TransactionPermission int8

const (
	// 交易中包含指令操作
	DIRECT_OPERATION = TransactionPermission(1)
	// 交易中包含合约操作
	CONTRACT_OPERATION = TransactionPermission(2)
)

func init() {
	binary_proto.RegisterEnum(DIRECT_OPERATION)
}

var _ binary_proto.EnumContract = (*TransactionPermission)(nil)

func (t TransactionPermission) ContractCode() int32 {
	return binary_proto.ENUM_TX_PERMISSION
}

func (t TransactionPermission) ContractType() string {
	return binary_proto.PRIMITIVETYPE_INT8
}

func (t TransactionPermission) ContractName() string {
	return "TransactionPermission"
}

func (t TransactionPermission) Description() string {
	return ""
}

func (t TransactionPermission) ContractVersion() int64 {
	return 0
}

func (t TransactionPermission) GetValue(CODE int32) binary_proto.EnumContract {
	if CODE == 1 {
		return DIRECT_OPERATION
	}
	if CODE == 2 {
		return CONTRACT_OPERATION
	}

	panic("no enum value founded")
}

func (t TransactionPermission) GetValueByName(name string) binary_proto.EnumContract {
	if name == "DIRECT_OPERATION" {
		return DIRECT_OPERATION
	}
	if name == "CONTRACT_OPERATION" {
		return CONTRACT_OPERATION
	}

	panic("no enum value founded")
}
