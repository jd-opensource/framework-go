package binary_proto

import (
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
	"reflect"
)

/**
 * @Author: imuge
 * @Date: 2020/5/21 下午
 */

func decodeHeader(data []byte) (code int32, version int64) {
	code = bytes.ToInt32(data[0:4])
	version = bytes.ToInt64(data[4:12])

	return
}

func decodeBytes(data []byte) ([]byte, int64) {
	size := decodeSize(data)
	len := bytes.NUMBERMASK_NORMAL.GetMaskLength(size)
	return data[len : size+int64(len)], size + int64(len)
}

func decodeSize(data []byte) int64 {
	return bytes.NUMBERMASK_NORMAL.ResolveMaskedNumber(data)
}

func decodePrimitiveType(data []byte, v reflect.Value, primitiveType string, numberMask bytes.NumberMask) int64 {
	switch primitiveType {
	case PRIMITIVETYPE_NIL:
		return 0
	case PRIMITIVETYPE_BOOLEAN:
		v.SetBool(data[0] == 1)
		return 1
	case PRIMITIVETYPE_INT8:
		if numberMask.Equals(bytes.NUMBERMASK_NONE) {
			v.Set(reflect.ValueOf(bytes.ToInt8(data[0])))
			return 1
		} else {
			iv, size := decodeInt8NumberMask(data, numberMask)
			v.Set(reflect.ValueOf(iv))
			return size
		}
	case PRIMITIVETYPE_INT16:
		if numberMask.Equals(bytes.NUMBERMASK_NONE) {
			v.Set(reflect.ValueOf(bytes.ToInt16(data[:2])))
			return 2
		} else {
			iv, size := decodeInt16NumberMask(data, numberMask)
			v.Set(reflect.ValueOf(iv))
			return size
		}
	case PRIMITIVETYPE_INT32:
		if numberMask.Equals(bytes.NUMBERMASK_NONE) {
			v.Set(reflect.ValueOf(bytes.ToInt32(data[:4])))
			return 4
		} else {
			iv, size := decodeInt32NumberMask(data, numberMask)
			v.Set(reflect.ValueOf(iv))
			return size
		}
	case PRIMITIVETYPE_INT64:
		if numberMask.Equals(bytes.NUMBERMASK_NONE) {
			v.Set(reflect.ValueOf(bytes.ToInt64(data[:8])))
			return 8
		} else {
			iv, size := decodeInt64NumberMask(data, numberMask)
			v.Set(reflect.ValueOf(iv))
			return size
		}
	case PRIMITIVETYPE_TEXT:
		s, size := decodeString(data)
		v.SetString(s)
		return size
	case PRIMITIVETYPE_BYTES: // 字节数组
		bs, size := decodeBytes(data)
		v.SetBytes(bs)
		return size
	default:
		panic("un support primitive type")
	}
}

func decodeInt8NumberMask(data []byte, mask bytes.NumberMask) (int8, int64) {
	v := mask.ResolveMaskedNumber(data)
	len := mask.GetMaskLength(v)
	return int8(v), int64(len)
}

func decodeInt16NumberMask(data []byte, mask bytes.NumberMask) (int16, int64) {
	v := mask.ResolveMaskedNumber(data)
	len := mask.GetMaskLength(v)
	return int16(v), int64(len)
}

func decodeInt32NumberMask(data []byte, mask bytes.NumberMask) (int32, int64) {
	v := mask.ResolveMaskedNumber(data)
	len := mask.GetMaskLength(v)
	return int32(v), int64(len)
}

func decodeInt64NumberMask(data []byte, mask bytes.NumberMask) (int64, int64) {
	v := mask.ResolveMaskedNumber(data)
	len := mask.GetMaskLength(v)
	return v, int64(len)
}

func decodeString(data []byte) (string, int64) {
	len := decodeSize(data)
	size := bytes.NUMBERMASK_NORMAL.GetMaskLength(len)
	return bytes.ToString(data[size : len+int64(size)]), len + int64(size)
}

func decodeArrayHeader(data []byte) (int, int64) {
	len := decodeSize(data)
	size := bytes.NUMBERMASK_NORMAL.GetMaskLength(len)
	return int(len), int64(size)
}

func decodeEnum(c *Codec, data []byte, refEnum int) (EnumContract, int64) {
	mapLocker.RLock()
	contract := (enumMap[int32(refEnum)]).(EnumContract)
	mapLocker.RUnlock()

	switch contract.ContractType() {
	case PRIMITIVETYPE_INT8:
		return contract.GetValue(int32(bytes.ToInt8(data[0]))), 1
	case PRIMITIVETYPE_INT16:
		return contract.GetValue(int32(bytes.ToInt16(data))), 1
	case PRIMITIVETYPE_INT32:
		return contract.GetValue(bytes.ToInt32(data)), 1
	default:
		panic("un support enum value type, int8,int16,int32 only")
	}
}

func decodeContract(c *Codec, data []byte) (interface{}, int64) {
	len := decodeSize(data)
	size := bytes.NUMBERMASK_NORMAL.GetMaskLength(len)
	if len == HEAD_BYTES {
		// 只有头信息
		return nil, len + int64(size)
	}
	contract, err := c.Decode(data[size : len+int64(size)])
	if err != nil {
		panic(err)
	}
	return contract, len + int64(size)
}
