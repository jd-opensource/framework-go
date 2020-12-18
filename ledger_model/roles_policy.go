package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午3:52
 */

// 多角色策略
// 表示如何处理一个对象被赋予多个角色时的综合权限
type RolesPolicy int8

const (
	UNION RolesPolicy = iota
	INTERSECT
)

func init() {
	binary_proto.RegisterEnum(UNION)
}

var _ binary_proto.EnumContract = (*RolesPolicy)(nil)

func (r RolesPolicy) ContractCode() int32 {
	return binary_proto.ENUM_MULTI_ROLES_POLICY
}

func (r RolesPolicy) ContractType() string {
	return binary_proto.PRIMITIVETYPE_INT8
}

func (r RolesPolicy) ContractName() string {
	return "USER-ROLE-POLICY"
}

func (r RolesPolicy) Description() string {
	return ""
}

func (r RolesPolicy) ContractVersion() int64 {
	return 0
}

func (r RolesPolicy) GetValue(CODE int32) binary_proto.EnumContract {
	if CODE == int32(0) {
		return UNION
	} else if CODE == int32(1) {
		return INTERSECT
	}

	panic("no enum value founded")
}

func (r RolesPolicy) GetValueByName(name string) binary_proto.EnumContract {
	if name == "UNION" {
		return UNION
	} else if name == "INTERSECT" {
		return INTERSECT
	}

	panic("no enum value founded")
}
