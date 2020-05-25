## Binary proto

`JD Chain`序列化反序列化`Go`语言实现，针对`struct`进行序列化和反序列化。

### 实现原理

需要序列化/反序列化结构体需要实现`DataContract`接口，枚举类型需要实现`EnumContract`接口。
序列化/反序列化之前向`Codec`中注册数据契约/枚举契约，基于反射和`Tag`实现序列化结构体对象为与`JD Chain`一致的字节序列。

### 支持字段类型

- `int8`
- `[]int8`
- `int16`
- `[]int16`
- `int32`
- `[]int32`
- `int64`
- `[]int64`
- `string`
- `[]string`
- `[]byte`
- 引用的`EnumContract`类型（包括指针类型）
- 引用的`EnumContract`类型数组
- 引用的`DataContract`类型（包括指针类型）
- 引用的`DataContract`类型数组

### 标签

```go
TAG_NAME            = "name"            // 名称标识，string
TAG_ORDER           = "order"           // 序号，int
TAG_DESCRIPTION     = "description"     // 描述,string
TAG_PRIMITIVETYPE   = "primitiveType"   // 基础类型,string
TAG_REFCONTRACT     = "refContract"     // 引用契约类型，类型code
TAG_REFENUM         = "refEnum"         // 引用枚举类型，类型code
TAG_GENERICCONTRACT = "genericContract" // 是否泛型字段,true/false
TAG_MAXSIZE         = "maxSize"         // 最大长度，int
TAG_NUMBERENCODING  = "numberEncoding"  // 动态数值字段，true/false
TAG_REPEATABLE      = "repeatable"      // 是否列表，true/false
```

- `TAG_ORDER`序号字段暂时没处理（TODO），使用字段定义顺序顺序编解码
- `TAG_DESCRIPTION`可使用值：
`NIL`     // 空
`BOOLEAN` // 布尔
`INT8`    // int8
`INT16`   // int16
`INT32`   // int32
`INT64`   // int64
`TEXT`    // 字符串
`BYTES`   // 字节数组

### 快速上手

1. 类型定义

`EnumContract`:
```go
type RefEnum int8

const (
	ONE RefEnum = iota + 1
	TWO
)

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

```

`DataContract`:
```go
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

func (J RefContract) Name() string {
	return ""
}

func (J RefContract) Description() string {
	return ""
}
```

```go
var _ binary_proto.DataContract = (*StructWithAllTypes)(nil)

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
		[]RefContract{NewRefContract(), NewRefContract()},
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
```

2. 注册类型信息

```go
binary_proto.Cdc.RegisterEnum(ONE)

binary_proto.Cdc.RegisterContract(RefContract{})

binary_proto.Cdc.RegisterContract(StructWithAllTypes{})
```

3. 序列化/反序列化
```go
// 序列化
origin := NewStructWithAllTypes()
bytes, err := binary_proto.Cdc.Encode(origin)

// 反序列化
obj, err := binary_proto.Cdc.Decode(bytes)
contract := obj.(StructWithAllTypes)
```