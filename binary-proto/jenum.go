package binary_proto

/*
 * Author: imuge
 * Date: 2020/5/22 下午1:50
 */

type JEnum int8

const (
	JEnumOne JEnum = iota + 1
	JEnumTwo
)

func init() {
	Cdc.RegisterEnum(JEnumOne.Code(), JEnumOne)
}

var _ EnumContract = (*JEnum)(nil)

func (J JEnum) Code() int32 {
	return 0x02
}

func (J JEnum) Type() string {
	return PRIMITIVETYPE_INT8
}

func (J JEnum) Name() string {
	return ""
}

func (J JEnum) Description() string {
	return ""
}

func (J JEnum) GetValue(CODE int32) EnumContract {
	if CODE == 1 {
		return JEnumOne
	}
	if CODE == 2 {
		return JEnumTwo
	}

	panic("no enum value founded")
}
