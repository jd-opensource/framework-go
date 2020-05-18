package binary_proto

import (
	"framework-go/utils/bytes"
)

var _ DataContract = (*JType)(nil)

type JType struct {
	I8    int8     `primitiveType:"INT8"`
	I16   int16    `primitiveType:"INT16"`
	I32   int32    `primitiveType:"INT32"`
	I64   int64    `primitiveType:"INT64"`
	I64m  int64    `primitiveType:"INT64" numberEncoding:"LONG"`
	Bool  bool     `primitiveType:"BOOLEAN"`
	Text  string   `primitiveType:"TEXT"`
	Bytes []byte   `primitiveType:"BYTES"`
	I8s   []int8   `primitiveType:"INT8" repeatable:"true"`
	I16s  []int16  `primitiveType:"INT16" repeatable:"true"`
	I32s  []int32  `primitiveType:"INT32" repeatable:"true"`
	I64s  []int64  `primitiveType:"INT64" repeatable:"true"`
	I64ms []int64  `primitiveType:"INT64" numberEncoding:"LONG" repeatable:"true" numberEncoding:"LONG"`
	Bools []bool   `primitiveType:"BOOLEAN" repeatable:"true"`
	Texts []string `primitiveType:"TEXT" repeatable:"true"`
}

func NewJType() JType {
	return JType{
		8, 16, 32, 64, 64, true, "text", bytes.StringToBytes("bytes"),
		[]int8{8, 8}, []int16{16, 16}, []int32{32, 32}, []int64{64, 64}, []int64{64, 64}, []bool{true, false}, []string{"text1", "text2"},
	}
}

func (p JType) Code() int32 {
	return 0x01
}

func (p JType) Version() int64 {
	return 6998934717933545896
}

func (p JType) Name() string {
	return ""
}

func (p JType) Description() string {
	return ""
}

func (p JType) Register() {
	Cdc.RegisterContract(p.Code(), p)
}
