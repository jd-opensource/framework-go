package test

import "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/22 下午3:55
 */

var _ binary_proto.DataContract = (*RefContract)(nil)

type RefContract struct {
	I8 int8 `primitiveType:"INT8"`
}

func NewRefContract() RefContract {
	return RefContract{1}
}

func (J RefContract) Code() int32 {
	return 0x03
}

func (J RefContract) Version() int64 {
	return -4451409565821993051
}

func (J RefContract) Name() string {
	return ""
}

func (J RefContract) Description() string {
	return ""
}
