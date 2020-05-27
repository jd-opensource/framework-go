package ledger_model

/*
 * Author: imuge
 * Date: 2020/5/27 下午3:43
 */

type TypedKVEntry struct {
	// 键名
	Key string
	// 版本
	Version int64
	// 数据类型
	Type DataType
	// 值
	Value interface{}
}
