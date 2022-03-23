package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

type AccountType int8

const (
	DATA AccountType = iota + 1
	EVENT
	CONTRACT
)

func init() {
	binary_proto.RegisterEnum(DATA)
}

var _ binary_proto.EnumContract = (*AccountType)(nil)

func (r AccountType) ContractCode() int32 {
	return binary_proto.ACCOUNT_TYPE
}

func (r AccountType) ContractType() string {
	return binary_proto.PRIMITIVETYPE_INT8
}

func (r AccountType) ContractName() string {
	return "AccountType"
}

func (r AccountType) Description() string {
	return ""
}

func (r AccountType) ContractVersion() int64 {
	return 0
}

func (r AccountType) GetValue(CODE int32) binary_proto.EnumContract {
	if CODE == int32(1) {
		return DATA
	} else if CODE == int32(2) {
		return EVENT
	} else if CODE == int32(3) {
		return CONTRACT
	}

	panic("no enum value founded")
}

func (r AccountType) GetValueByName(name string) binary_proto.EnumContract {
	if name == "DATA" {
		return DATA
	} else if name == "EVENT" {
		return EVENT
	} else if name == "CONTRACT" {
		return CONTRACT
	}

	panic("no enum value founded")
}
