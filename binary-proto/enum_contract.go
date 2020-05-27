package binary_proto

/**
 * @Author: imuge
 * @Date: 2020/5/21 下午
 */

/**
枚举契约契约
*/
type EnumContract interface {

	// 唯一标识
	ContractCode() int32
	// 字段基础类型信息，只支持INT8,INT16,INT32
	ContractType() string
	// 标识名称
	ContractName() string
	// 描述信息
	Description() string
	// 版本
	ContractVersion() int64

	GetValue(CODE int32) EnumContract
	GetValueByName(name string) EnumContract
}
