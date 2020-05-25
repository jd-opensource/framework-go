package test

import (
	"framework-go/binary-proto"
	"framework-go/utils/bytes"
)

var _ binary_proto.DataContract = (*StructWithAllTypes)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(StructWithAllTypes{})
}

type StructWithAllTypes struct {
	I8    int8          `primitiveType:"INT8"`
	I16   int16         `primitiveType:"INT16"`
	I32   int32         `primitiveType:"INT32"`
	I64   int64         `primitiveType:"INT64"`
	I64m  int64         `primitiveType:"INT64" numberEncoding:"LONG"`
	Bool  bool          `primitiveType:"BOOLEAN"`
	Text  string        `primitiveType:"TEXT"`
	Bytes []byte        `primitiveType:"BYTES"`
	I8s   []int8        `primitiveType:"INT8" repeatable:"true"`
	I16s  []int16       `primitiveType:"INT16" repeatable:"true"`
	I32s  []int32       `primitiveType:"INT32" repeatable:"true"`
	I64s  []int64       `primitiveType:"INT64" repeatable:"true"`
	I64ms []int64       `primitiveType:"INT64" numberEncoding:"LONG" repeatable:"true" numberEncoding:"LONG"`
	Bools []bool        `primitiveType:"BOOLEAN" repeatable:"true"`
	Texts []string      `primitiveType:"TEXT" repeatable:"true"`
	Enum  RefEnum       `refEnum:"2"`
	Enums []RefEnum     `refEnum:"2" repeatable:"true"`
	JP    *RefContract  `refContract:"3"`
	JPs   []RefContract `refContract:"3" repeatable:"true"`
	JG    binary_proto.DataContract   `refContract:"4" genericContract:"true"`
	JGs   []RefContract `refContract:"4" genericContract:"true" repeatable:"true"`
}

func NewStructWithAllTypes() StructWithAllTypes {
	return StructWithAllTypes{
		8, 16, 32, 64,
		64,
		true,
		"text",
		bytes.StringToBytes("bytes"),
		[]int8{8, 8},
		[]int16{16, 16}, []int32{32, 32}, []int64{64, 64},
		[]int64{64, 64}, []bool{true, false}, []string{"text1", "text2"},
		ONE,
		[]RefEnum{ONE, TWO},
		nil,
		nil, //[]RefContract{NewRefContract(), NewRefContract()},
		NewRefContract(),
		[]RefContract{NewRefContract(), NewRefContract()},
	}
}

func (p StructWithAllTypes) Code() int32 {
	return 0x01
}

func (p StructWithAllTypes) Name() string {
	return ""
}

func (p StructWithAllTypes) Description() string {
	return ""
}

func (p StructWithAllTypes) Equals(contract StructWithAllTypes) bool {
	if p.I8 != contract.I8 {
		return false
	}
	if p.I16 != contract.I16 {
		return false
	}
	if p.I32 != contract.I32 {
		return false
	}
	if p.I64 != contract.I64 {
		return false
	}
	if p.I64m != contract.I64m {
		return false
	}
	if p.Bool != contract.Bool {
		return false
	}
	if p.Text != contract.Text {
		return false
	}
	if !bytes.Equals(p.Bytes, contract.Bytes) {
		return false
	}
	if len(p.I8s) != len(contract.I8s) {
		return false
	}
	for i := 0; i < len(p.I8s); i++ {
		if p.I8s[i] != contract.I8s[i] {
			return false
		}
	}
	if len(p.I16s) != len(contract.I16s) {
		return false
	}
	for i := 0; i < len(p.I16s); i++ {
		if p.I16s[i] != contract.I16s[i] {
			return false
		}
	}
	if len(p.I32s) != len(contract.I32s) {
		return false
	}
	for i := 0; i < len(p.I32s); i++ {
		if p.I32s[i] != contract.I32s[i] {
			return false
		}
	}
	if len(p.I64s) != len(contract.I64s) {
		return false
	}
	for i := 0; i < len(p.I64s); i++ {
		if p.I64s[i] != contract.I64s[i] {
			return false
		}
	}
	if len(p.I64ms) != len(contract.I64ms) {
		return false
	}
	for i := 0; i < len(p.I64ms); i++ {
		if p.I64ms[i] != contract.I64ms[i] {
			return false
		}
	}
	if len(p.Bools) != len(contract.Bools) {
		return false
	}
	for i := 0; i < len(p.Bools); i++ {
		if p.Bools[i] != contract.Bools[i] {
			return false
		}
	}
	if len(p.Texts) != len(contract.Texts) {
		return false
	}
	for i := 0; i < len(p.Texts); i++ {
		if p.Texts[i] != contract.Texts[i] {
			return false
		}
	}
	if p.Enum != contract.Enum {
		return false
	}
	if len(p.Enums) != len(contract.Enums) {
		return false
	}
	for i := 0; i < len(p.Enums); i++ {
		if p.Enums[i] != contract.Enums[i] {
			return false
		}
	}
	if p.JP == nil && contract.JP != nil {
		return false
	} else if p.JP != nil && contract.JP == nil {
		return false
	} else if p.JP != nil && contract.JP != nil && !p.JP.Equals(*contract.JP) {
		return false
	}
	if len(p.JPs) != len(contract.JPs) {
		return false
	}
	for i := 0; i < len(p.JPs); i++ {
		if !p.JPs[i].Equals(contract.JPs[i]) {
			return false
		}
	}
	if !p.JG.(RefContract).Equals(contract.JG.(RefContract)) {
		return false
	}
	for i := 0; i < len(p.JGs); i++ {
		if !p.JGs[i].Equals(contract.JGs[i]) {
			return false
		}
	}

	return true
}
