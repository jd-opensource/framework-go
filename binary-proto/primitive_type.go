package binary_proto

import "fmt"

/*
 * Author: imuge
 * Date: 2020/5/25 上午10:50
 */

type PrimitiveType = byte

const (
	PRIMITIVETYPE_NIL     = "NIL"     // 空
	PRIMITIVETYPE_BOOLEAN = "BOOLEAN" // 布尔
	PRIMITIVETYPE_INT8    = "INT8"    // int8
	PRIMITIVETYPE_INT16   = "INT16"   // int16
	PRIMITIVETYPE_INT32   = "INT32"   // int32
	PRIMITIVETYPE_INT64   = "INT64"   // int64
	PRIMITIVETYPE_TEXT    = "TEXT"    // 字符串
	PRIMITIVETYPE_BYTES   = "BYTES"   // 字节数组

	BASE_TYPE_NIL     = 0x00
	BASE_TYPE_BOOLEAN = 0x01
	BASE_TYPE_INTEGER = 0x10
	BASE_TYPE_INT8    = BASE_TYPE_INTEGER | 0x01
	BASE_TYPE_INT16   = BASE_TYPE_INTEGER | 0x02
	BASE_TYPE_INT32   = BASE_TYPE_INTEGER | 0x03
	BASE_TYPE_INT64   = BASE_TYPE_INTEGER | 0x04
	BASE_TYPE_TEXT    = 0x20
	BASE_TYPE_BYTES   = 0x40
	BASE_TYPE_EXT     = -128

	NIL     PrimitiveType = BASE_TYPE_NIL
	BOOLEAN PrimitiveType = BASE_TYPE_BOOLEAN
	INT8    PrimitiveType = BASE_TYPE_INT8
	INT16   PrimitiveType = BASE_TYPE_INT16
	INT32   PrimitiveType = BASE_TYPE_INT32
	INT64   PrimitiveType = BASE_TYPE_INT64
	TEXT    PrimitiveType = BASE_TYPE_TEXT
	BYTES   PrimitiveType = BASE_TYPE_BYTES
)

func GetPrimitiveType(name string) PrimitiveType {
	switch name {
	case PRIMITIVETYPE_NIL:
		return NIL
	case PRIMITIVETYPE_BOOLEAN:
		return BOOLEAN
	case PRIMITIVETYPE_INT8:
		return INT8
	case PRIMITIVETYPE_INT16:
		return INT16
	case PRIMITIVETYPE_INT32:
		return INT32
	case PRIMITIVETYPE_INT64:
		return INT64
	case PRIMITIVETYPE_TEXT:
		return TEXT
	case PRIMITIVETYPE_BYTES:
		return BYTES
	default:
		panic(fmt.Sprintf("un support primitive type: %s", name))
	}
}
