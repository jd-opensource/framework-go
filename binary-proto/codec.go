package binary_proto

import (
	"crypto/sha256"
	"errors"
	"framework-go/utils/bytes"
	"reflect"
	"strconv"
)

/**
 * @Author: imuge
 * @Date: 2020/5/21 下午
 */

var Cdc = NewCodec()

type Codec struct {
	// 契约类型
	ContractMap map[int32]DataContract
	VersionMap  map[int32]int64
	// 枚举类型
	EnumMap map[int32]EnumContract
}

func NewCodec() *Codec {
	return &Codec{
		make(map[int32]DataContract),
		make(map[int32]int64),
		make(map[int32]EnumContract),
	}
}

/**
注册契约
*/
func (c *Codec) RegisterContract(contract DataContract) {
	c.ContractMap[contract.Code()] = contract
	c.calculateVersion(contract)
}

// 计算契约版本号
func (c *Codec) calculateVersion(contract DataContract) {
	var hasher = sha256.New()

	rt := reflect.TypeOf(contract)
	// 解析字段信息
	for i := 0; i < rt.NumField(); i++ {
		tField := rt.Field(i)
		_, _, _, primitiveType, refContract, refEnum, genericContract, _, _, repeatable, err := resolveTags(tField)
		if err != nil {
			panic(err)
		}
		array := byte(0)
		if repeatable {
			array = byte(1)
		}
		if genericContract { // 泛型编码
			bs := make([]byte, 14)
			bs[0] = array
			bs[1] = byte(3)
			refCon := (c.ContractMap[int32(refContract)]).(DataContract)
			copy(bs[2:6], bytes.Int32ToBytes(refCon.Code()))
			ver, ok := c.VersionMap[refCon.Code()]
			if !ok {
				c.calculateVersion(refCon)
				ver = c.VersionMap[refCon.Code()]
			}
			copy(bs[6:], bytes.Int64ToBytes(ver))
			hasher.Write(bs)
		} else if refContract != 0 { // 引用其他契约
			bs := make([]byte, 14)
			bs[0] = array
			bs[1] = byte(2)
			refCon := (c.ContractMap[int32(refContract)]).(DataContract)
			copy(bs[2:6], bytes.Int32ToBytes(refCon.Code()))
			ver, ok := c.VersionMap[refCon.Code()]
			if !ok {
				c.calculateVersion(refCon)
				ver = c.VersionMap[refCon.Code()]
			}
			copy(bs[6:], bytes.Int64ToBytes(ver))
			hasher.Write(bs)
		} else if refEnum != 0 { // 引用枚举
			bs := make([]byte, 14)
			bs[0] = array
			bs[1] = byte(1)
			enumCon := (c.EnumMap[int32(refEnum)]).(EnumContract)
			copy(bs[2:6], bytes.Int32ToBytes(enumCon.Code()))
			copy(bs[6:], bytes.Int64ToBytes(enumCon.Version()))
			hasher.Write(bs)
		} else { // 基础类型字段
			bs := make([]byte, 6)
			bs[0] = array
			bs[1] = byte(0)
			copy(bs[2:], bytes.Int16ToBytes(int16(GetPrimitiveType(primitiveType))))
			hasher.Write(bs)
		}
	}

	c.VersionMap[contract.Code()] = bytes.ToInt64(hasher.Sum(nil))
}

/**
注册枚举
*/
func (c *Codec) RegisterEnum(enum EnumContract) {
	c.EnumMap[enum.Code()] = enum
}

/**
  编码
*/
func (c *Codec) Encode(contract DataContract) ([]byte, error) {
	var err error
	rt := reflect.TypeOf(contract)
	rv := reflect.ValueOf(contract)
	if contract == nil || (rv.Kind() == reflect.Ptr && rv.IsNil()) {
		return nil, errors.New("nil value")
	}

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	// 编码头信息
	buf := bytes.Int32ToBytes(contract.Code())
	buf = append(buf, bytes.Int64ToBytes(c.VersionMap[contract.Code()])...)

	// 编码字段信息
	for i := 0; i < rt.NumField(); i++ {
		tField := rt.Field(i)
		vField := rv.Field(i)
		_, _, _, primitiveType, refContract, refEnum, genericContract, _, numberMask, repeatable, err := resolveTags(tField)
		if err != nil {
			return nil, err
		}

		if primitiveType == PRIMITIVETYPE_BYTES { // 字节数组
			buf = append(buf, encodeBytes(vField.Bytes())...)
		} else {
			repeat := 1
			if repeatable {
				repeat = rv.Field(i).Len()
				// 编码数组头信息
				buf = append(buf, encodeArrayHeader(repeat)...)
			}

			for j := 0; j < repeat; j++ {
				var value reflect.Value
				if !repeatable {
					value = vField
				} else {
					value = vField.Index(j)
				}
				if genericContract { // 泛型编码
					buf = append(buf, encodeGeneric(c, refContract, value.Interface())...)
				} else if refContract != 0 { // 引用其他契约
					buf = append(buf, encodeContract(c, refContract, value.Interface())...)
				} else if refEnum != 0 { // 引用枚举
					buf = append(buf, encodeEnum(c, value.Int(), refEnum)...)
				} else { // 基础类型字段
					buf = append(buf, encodePrimitiveType(value, primitiveType, numberMask)...)
				}
			}
		}
	}

	return buf, err
}

/**
解码
*/
func (c *Codec) Decode(data []byte) (interface{}, error) {
	// 解析头信息
	code, _ := decodeHeader(data)
	offset := int64(12)
	contract := c.ContractMap[code]
	rt := reflect.TypeOf(contract)
	obj := reflect.New(rt)
	rv := obj.Elem()

	// 解析字段信息
	for i := 0; i < rt.NumField(); i++ {
		tField := rt.Field(i)
		vField := rv.Field(i)
		_, _, _, primitiveType, refContract, refEnum, genericContract, _, numberMask, repeatable, err := resolveTags(tField)
		if err != nil {
			return nil, err
		}

		if primitiveType == PRIMITIVETYPE_BYTES { // 字节数组
			bs, size := decodeBytes(data[offset:])
			vField.SetBytes(bs)
			offset += size
		} else {
			repeat := 1
			size := int64(0)
			if repeatable {
				// 编码数组头信息
				repeat, size = decodeArrayHeader(data[offset:])
				offset += size
				// 初始化数组
				vField = reflect.MakeSlice(tField.Type, repeat, repeat)
			}

			for j := 0; j < repeat; j++ {
				var value reflect.Value
				if !repeatable {
					value = vField
				} else {
					value = vField.Index(j)
				}
				if genericContract || refContract != 0 { // 泛型/引用其他契约
					contract, size := decodeContract(c, data[offset:])
					if contract != nil {
						if value.Kind() == reflect.Ptr {
							value.Set(reflect.New(vField.Type().Elem()))
							value.Elem().Set(reflect.ValueOf(contract))
						} else {
							value.Set(reflect.ValueOf(contract))
						}
					}
					offset += size
				} else if refEnum != 0 { // 引用枚举
					enum, size := decodeEnum(c, data[offset:], refEnum)
					value.Set(reflect.ValueOf(enum))
					offset += size
				} else { // 基础类型字段
					size = decodePrimitiveType(data[offset:], value, primitiveType, numberMask)
					offset += size
				}
			}
			if repeatable {
				rv.Field(i).Set(vField)
			}
		}
	}

	return obj.Elem().Interface(), nil
}

// 解析Tag
func resolveTags(field reflect.StructField) (
	name string,
	order int,
	description string,
	primitiveType string,
	refContract int,
	refEnum int,
	genericContract bool,
	maxSize int,
	numberMask bytes.NumberMask,
	repeatable bool,
	err error) {
	name, ok := field.Tag.Lookup(TAG_NAME)
	if !ok {
		name = ""
	}
	orderStr, ok := field.Tag.Lookup(TAG_ORDER)
	if ok {
		order, err = strconv.Atoi(orderStr)
		if err != nil {
			return
		}
	}
	description, ok = field.Tag.Lookup(TAG_DESCRIPTION)
	if !ok {
		description = ""
	}
	primitiveType = field.Tag.Get(TAG_PRIMITIVETYPE)
	refContractStr, ok := field.Tag.Lookup(TAG_REFCONTRACT)
	if ok {
		refContract, err = strconv.Atoi(refContractStr)
		if err != nil {
			return
		}
	}
	refEnumStr, ok := field.Tag.Lookup(TAG_REFENUM)
	if ok {
		refEnum, err = strconv.Atoi(refEnumStr)
		if err != nil {
			return
		}
	}
	genericContractStr, ok := field.Tag.Lookup(TAG_GENERICCONTRACT)
	if ok && genericContractStr == "true" {
		genericContract = true
	}
	maxSizeStr, ok := field.Tag.Lookup(TAG_MAXSIZE)
	if ok {
		maxSize, err = strconv.Atoi(maxSizeStr)
		if err != nil {
			return
		}
	} else {
		maxSize = -1
	}
	numberEncodingStr, ok := field.Tag.Lookup(TAG_NUMBERENCODING)
	if !ok {
		numberMask = bytes.NUMBERMASK_NONE
	} else {
		numberMask = bytes.GetNumberMask(numberEncodingStr)
	}
	repeatableStr, ok := field.Tag.Lookup(TAG_REPEATABLE)
	if ok && repeatableStr == "true" {
		repeatable = true
	}

	return
}
