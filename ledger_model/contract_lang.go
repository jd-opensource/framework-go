package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

type ContractLang int8

const (
	Java ContractLang = iota + 1
	JavaScript
	Python
	Rust
)

func init() {
	binary_proto.RegisterEnum(Java)
}

var _ binary_proto.EnumContract = (*ContractLang)(nil)

func (r ContractLang) ContractCode() int32 {
	return binary_proto.CONTRACT_LANG
}

func (r ContractLang) ContractType() string {
	return binary_proto.PRIMITIVETYPE_INT8
}

func (r ContractLang) ContractName() string {
	return "AccountType"
}

func (r ContractLang) Description() string {
	return ""
}

func (r ContractLang) ContractVersion() int64 {
	return 0
}

func (r ContractLang) GetValue(CODE int32) binary_proto.EnumContract {
	if CODE == int32(1) {
		return Java
	} else if CODE == int32(2) {
		return JavaScript
	} else if CODE == int32(3) {
		return Python
	} else if CODE == int32(4) {
		return Rust
	}

	panic("no enum value founded")
}

func (r ContractLang) GetValueByName(name string) binary_proto.EnumContract {
	if name == "Java" {
		return Java
	} else if name == "JavaScript" {
		return JavaScript
	} else if name == "Python" {
		return Python
	} else if name == "Rust" {
		return Rust
	}

	panic("no enum value founded")
}
