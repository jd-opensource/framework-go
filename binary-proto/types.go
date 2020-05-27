package binary_proto

/**
 * @Author: imuge
 * @Date: 2020/5/21 下午
 */

const (
	TAG_NAME            = "name"            // 名称标识，string
	TAG_ORDER           = "order"           // 序号，int
	TAG_DESCRIPTION     = "description"     // 描述,string
	TAG_PRIMITIVETYPE   = "primitiveType"   // 基础类型,string
	TAG_REFCONTRACT     = "refContract"     // 引用契约类型，类型code
	TAG_REFENUM         = "refEnum"         // 引用枚举类型，类型code
	TAG_GENERICCONTRACT = "genericContract" // 是否泛型字段,true/false
	TAG_MAXSIZE         = "maxSize"         // 最大长度，int
	TAG_NUMBERENCODING  = "numberEncoding"  // 动态数值字段，true/false
	TAG_REPEATABLE      = "list"            // 是否列表，true/false

	HEAD_BYTES = 12 // 头信息长度
)
