package test

import (
	"framework-go/binary-proto"
	"framework-go/utils/bytes"
)

var _ binary_proto.DataContract = (*StructWithAllTypes)(nil)

type StructWithAllTypes struct {
	I8      int8          `primitiveType:"INT8"`
	I16     int16         `primitiveType:"INT16"`
	I32     int32         `primitiveType:"INT32"`
	I64     int64         `primitiveType:"INT64"`
	I64m    int64         `primitiveType:"INT64" numberEncoding:"LONG"`
	Bool    bool          `primitiveType:"BOOLEAN"`
	Text    string        `primitiveType:"TEXT"`
	Bytes   []byte        `primitiveType:"BYTES"`
	I8s     []int8        `primitiveType:"INT8" repeatable:"true"`
	I16s    []int16       `primitiveType:"INT16" repeatable:"true"`
	I32s    []int32       `primitiveType:"INT32" repeatable:"true"`
	I64s    []int64       `primitiveType:"INT64" repeatable:"true"`
	I64ms   []int64       `primitiveType:"INT64" numberEncoding:"LONG" repeatable:"true" numberEncoding:"LONG"`
	Bools   []bool        `primitiveType:"BOOLEAN" repeatable:"true"`
	Texts   []string      `primitiveType:"TEXT" repeatable:"true"`
	enum    RefEnum       `refEnum:"2"`
	enums   []RefEnum     `refEnum:"2" repeatable:"true"`
	JP      RefContract   `refContract:"3"`
	JPs     []RefContract `refContract:"3" repeatable:"true"`
}

func NewStructWithAllTypes() StructWithAllTypes {
	return StructWithAllTypes{
		8, 16, 32, 64, 64, true, "text", bytes.StringToBytes("bytes"),
		[]int8{8, 8}, []int16{16, 16}, []int32{32, 32}, []int64{64, 64}, []int64{64, 64}, []bool{true, false}, []string{"text1", "text2"},
		ONE, []RefEnum{ONE, TWO},
		NewRefContract(), []RefContract{NewRefContract(), NewRefContract()},
	}
}

func (p StructWithAllTypes) Code() int32 {
	return 0x01
}

func (p StructWithAllTypes) Version() int64 {
	return 4719784014675470359
}

func (p StructWithAllTypes) Name() string {
	return ""
}

func (p StructWithAllTypes) Description() string {
	return ""
}
