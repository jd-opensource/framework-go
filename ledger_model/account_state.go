package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2021/09/10 下午3:28
 */

// 账户状态
type AccountState int8

const (
	// 正常
	NORMAL AccountState = iota + 1
	// 冻结
	FREEZE
	// 撤销/移除
	REVOKE
)

func init() {
	binary_proto.RegisterEnum(NORMAL)
}

var _ binary_proto.EnumContract = (*AccountState)(nil)

func (as AccountState) ContractCode() int32 {
	return binary_proto.ENUM_ACCOUNT_STATE
}

func (as AccountState) ContractType() string {
	return binary_proto.PRIMITIVETYPE_INT8
}

func (as AccountState) ContractName() string {
	return "AccountState"
}

func (as AccountState) Description() string {
	return ""
}

func (as AccountState) ContractVersion() int64 {
	return 0
}

func (as AccountState) GetValue(CODE int32) binary_proto.EnumContract {
	if CODE == int32(1) {
		return NORMAL
	} else if CODE == int32(2) {
		return FREEZE
	} else if CODE == int32(3) {
		return REVOKE
	}

	panic("no enum value founded")
}

func (as AccountState) GetValueByName(name string) binary_proto.EnumContract {
	if name == "NORMAL" {
		return NORMAL
	} else if name == "FREEZE" {
		return FREEZE
	} else if name == "REVOKE" {
		return REVOKE
	}

	panic("no enum value founded")
}
