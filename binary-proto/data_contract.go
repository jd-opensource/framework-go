package binary_proto

/**
 * @Author: imuge
 * @Date: 2020/5/21 下午
 */

/**
数据契约，所有需要序列化/反序列化的Struct必须实现该接口
*/
type DataContract interface {
	// 唯一标识
	Code() int32
	// 版本号
	Version() int64
	// 标识名称
	Name() string
	// 描述信息
	Description() string
}
