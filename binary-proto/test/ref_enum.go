package test

import "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/22 下午1:50
 */

type RefEnum int8

const (
	ONE RefEnum = iota + 1
	TWO
)

func init() {
	binary_proto.Cdc.RegisterEnum(ONE)
}

var _ binary_proto.EnumContract = (*RefEnum)(nil)

func (J RefEnum) Code() int32 {
	return 0x02
}

func (J RefEnum) Type() string {
	return binary_proto.PRIMITIVETYPE_INT8
}

func (J RefEnum) Name() string {
	return ""
}

func (J RefEnum) Description() string {
	return ""
}

func (J RefEnum) Version() int64 {
	return 0
}

func (J RefEnum) GetValue(CODE int32) binary_proto.EnumContract {
	if CODE == 1 {
		return ONE
	}
	if CODE == 2 {
		return TWO
	}

	panic("no enum value founded")
}
