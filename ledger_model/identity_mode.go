package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2021/09/10 下午3:45
 */

// 身份认证模式
type IdentityMode int8

const (
	// 公私钥对模式
	KEYPAIR IdentityMode = iota + 1
	// 证书模式
	CA
)

func init() {
	binary_proto.RegisterEnum(NORMAL)
}

var _ binary_proto.EnumContract = (*IdentityMode)(nil)

func (m IdentityMode) ContractCode() int32 {
	return binary_proto.METADATA_IDENTITY_MODE
}

func (m IdentityMode) ContractType() string {
	return binary_proto.PRIMITIVETYPE_INT8
}

func (m IdentityMode) ContractName() string {
	return "IdentityMode"
}

func (m IdentityMode) Description() string {
	return ""
}

func (m IdentityMode) ContractVersion() int64 {
	return 0
}

func (m IdentityMode) GetValue(CODE int32) binary_proto.EnumContract {
	if CODE == int32(1) {
		return KEYPAIR
	} else if CODE == int32(2) {
		return CA
	}

	panic("no enum value founded")
}

func (m IdentityMode) GetValueByName(name string) binary_proto.EnumContract {
	if name == "KEYPAIR" {
		return KEYPAIR
	} else if name == "CA" {
		return CA
	}

	panic("no enum value founded")
}
