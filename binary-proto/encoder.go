package binary_proto

import (
	"fmt"
	"framework-go/utils/bytes"
	"reflect"
)

/**
 * @Author: imuge
 * @Date: 2020/5/21 下午
 */

func encodePrimitiveType(v reflect.Value, primitiveType string, numberMask bytes.NumberMask) []byte {
	switch primitiveType {
	case PRIMITIVETYPE_NIL:
		return []byte{}
	case PRIMITIVETYPE_BOOLEAN:
		return []byte{encodeBool(v.Bool())}
	case PRIMITIVETYPE_INT8:
		if numberMask.Equals(bytes.NUMBERMASK_NONE) {
			return []byte{encodeInt8(int8(v.Int()))}
		} else {
			return encodeNumberMask(v.Int(), numberMask)
		}
	case PRIMITIVETYPE_INT16:
		if numberMask.Equals(bytes.NUMBERMASK_NONE) {
			return encodeInt16(int16(v.Int()))
		} else {
			return encodeNumberMask(v.Int(), numberMask)
		}
	case PRIMITIVETYPE_INT32:
		if numberMask.Equals(bytes.NUMBERMASK_NONE) {
			return encodeInt32(int32(v.Int()))
		} else {
			return encodeNumberMask(v.Int(), numberMask)
		}
	case PRIMITIVETYPE_INT64:
		if numberMask.Equals(bytes.NUMBERMASK_NONE) {
			return encodeInt64(v.Int())
		} else {
			return encodeNumberMask(v.Int(), numberMask)
		}
	case PRIMITIVETYPE_TEXT:
		return encodeString(v.String())
	case PRIMITIVETYPE_BYTES: // 字节数组
		return encodeBytes(v.Bytes())
	default:
		panic("un support primitive type")
	}
}

func encodeBytes(data []byte) []byte {
	buf := encodeSize(len(data))
	return append(buf, data...)
}

func encodeBool(data bool) byte {
	return bytes.BoolToBytes(data)
}

func encodeInt8(data int8) byte {
	return bytes.Int8ToBytes(data)
}

func encodeInt16(data int16) []byte {
	return bytes.Int16ToBytes(data)
}

func encodeInt32(data int32) []byte {
	return bytes.Int32ToBytes(data)
}

func encodeInt64(data int64) []byte {
	return bytes.Int64ToBytes(data)
}

func encodeNumberMask(data int64, mask bytes.NumberMask) []byte {
	return mask.WriteMask(data)
}

func encodeSize(data int) []byte {
	return bytes.NUMBERMASK_NORMAL.WriteMask(int64(data))
}

func encodeString(data string) []byte {
	return encodeBytes(bytes.StringToBytes(data))
}

func encodeEnum(c *Codec, value int64, refEnum int) []byte {
	contract := (c.EnumMap[int32(refEnum)]).(EnumContract)
	switch contract.ContractType() {
	case PRIMITIVETYPE_INT8:
		return []byte{encodeInt8(int8(value))}
	case PRIMITIVETYPE_INT16:
		return encodeInt16(int16(value))
	case PRIMITIVETYPE_INT32:
		return encodeInt32(int32(value))
	default:
		panic("un support enum value type, int8,int16,int32 only")
	}
}

func encodeArrayHeader(count int) []byte {
	return bytes.NUMBERMASK_NORMAL.WriteMask(int64(count))
}

func encodeGeneric(c *Codec, refContract int, v interface{}) []byte {
	// 与非泛型引用无差别
	return encodeContract(c, refContract, v)
}

func encodeContract(c *Codec, reflectContract int, v interface{}) []byte {
	if v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil()) {
		// 空值，仅编码头信息
		// 编码头信息
		contract, ok := c.ContractMap[int32(reflectContract)]
		if !ok {
			panic(fmt.Sprintf("contract %d not exists", reflectContract))
		}
		buf := bytes.Int32ToBytes(contract.ContractCode())
		buf = append(buf, bytes.Int64ToBytes(c.VersionMap[contract.ContractCode()])...)
		buf = append(encodeSize(len(buf)), buf...)
		return buf
	}
	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	buf, err := c.Encode(v.(DataContract))
	if err != nil {
		panic(err)
	}
	buf = append(encodeSize(len(buf)), buf...)
	return buf
}
