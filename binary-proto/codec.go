package binary_proto

import (
	"errors"
	"fmt"
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
	"github.com/blockchain-jd-com/framework-go/utils/sha"
	"reflect"
	"strconv"
	"sync"
)

/**
 * @Author: imuge
 * @Date: 2020/5/21 下午
 */

//var Cdc = NewCodec()

var (
	// 契约类型
	contractMap map[int32]DataContract

	// 枚举类型
	enumMap map[int32]EnumContract

	// 锁
	mapLocker sync.RWMutex

	// once
	once sync.Once
)

func init() {
	once.Do(func() {
		contractMap = make(map[int32]DataContract)
		enumMap = make(map[int32]EnumContract)
	})
}

/**
注册契约
*/
func RegisterContract(contract DataContract) {
	mapLocker.Lock()
	defer mapLocker.Unlock()

	contractMap[contract.ContractCode()] = contract
}

/**
注册枚举
*/
func RegisterEnum(enum EnumContract) {
	mapLocker.Lock()
	defer mapLocker.Unlock()

	enumMap[enum.ContractCode()] = enum
}

type Codec struct {
	VersionMap map[int32]int64
}

func NewCodec() *Codec {
	return &Codec{
		make(map[int32]int64),
	}
}

// 计算契约版本号`
func (c *Codec) CalculateVersion(contract DataContract) error {
	rt := reflect.TypeOf(contract)
	buf, err := c.calculateFieldVersion(rt)
	if err != nil {
		return err
	}
	c.VersionMap[contract.ContractCode()] = bytes.ToInt64(sha.Sha256(buf))
	return nil
}

func (c *Codec) calculateFieldVersion(rt reflect.Type) ([]byte, error) {
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	var buf []byte
	for i := 0; i < rt.NumField(); i++ {
		tField := rt.Field(i)
		if tField.Anonymous { // 嵌套契约
			bs, err := c.calculateFieldVersion(tField.Type)
			if err != nil {
				return nil, err
			}
			buf = append(buf, bs...)
		} else {
			_, _, _, primitiveType, refContract, refEnum, genericContract, _, _, repeatable, err := resolveTags(tField)
			if err != nil {
				return nil, err
			}
			array := byte(0)
			if repeatable {
				array = byte(1)
			}
			if genericContract { // 泛型编码
				bs := make([]byte, 14)
				bs[0] = array
				bs[1] = byte(3)

				mapLocker.RLock()
				refCon := (contractMap[int32(refContract)]).(DataContract)
				mapLocker.RUnlock()

				copy(bs[2:6], bytes.Int32ToBytes(refCon.ContractCode()))
				ver, ok := c.VersionMap[refCon.ContractCode()]
				if !ok {
					err = c.CalculateVersion(refCon)
					if err != nil {
						return nil, err
					}
					ver = c.VersionMap[refCon.ContractCode()]
				}
				copy(bs[6:], bytes.Int64ToBytes(ver))
				buf = append(buf, bs...)
			} else if refContract != 0 { // 引用其他契约
				bs := make([]byte, 14)
				bs[0] = array
				bs[1] = byte(2)

				mapLocker.RLock()
				refCon := (contractMap[int32(refContract)]).(DataContract)
				mapLocker.RUnlock()

				copy(bs[2:6], bytes.Int32ToBytes(refCon.ContractCode()))
				ver, ok := c.VersionMap[refCon.ContractCode()]
				if !ok {
					err = c.CalculateVersion(refCon)
					if err != nil {
						return nil, err
					}
					ver = c.VersionMap[refCon.ContractCode()]
				}
				copy(bs[6:], bytes.Int64ToBytes(ver))
				buf = append(buf, bs...)
			} else if refEnum != 0 { // 引用枚举
				bs := make([]byte, 14)
				bs[0] = array
				bs[1] = byte(1)

				mapLocker.RLock()
				enumCon := (enumMap[int32(refEnum)]).(EnumContract)
				mapLocker.RUnlock()

				copy(bs[2:6], bytes.Int32ToBytes(enumCon.ContractCode()))
				copy(bs[6:], bytes.Int64ToBytes(enumCon.ContractVersion()))
				buf = append(buf, bs...)
			} else { // 基础类型字段
				bs := make([]byte, 6)
				bs[0] = array
				bs[1] = byte(0)
				copy(bs[2:], bytes.Int16ToBytes(int16(GetPrimitiveType(primitiveType))))
				buf = append(buf, bs...)
			}
		}
	}

	return buf, nil
}

func (c *Codec) Encode(contract DataContract) ([]byte, error) {
	return c.encode(contract, true)
}

/**
  编码
*/
func (c *Codec) encode(contract DataContract, withHead bool) ([]byte, error) {
	var err error
	_, ok := c.VersionMap[contract.ContractCode()]
	if !ok {
		err = c.CalculateVersion(contract)
		if err != nil {
			return nil, err
		}
	}
	rt := reflect.TypeOf(contract)
	rv := reflect.ValueOf(contract)
	if contract == nil || (rv.Kind() == reflect.Ptr && rv.IsNil()) {
		return nil, errors.New("nil value")
	}

	if rv.Kind() == reflect.Ptr {
		rt = rt.Elem()
		rv = rv.Elem()
	}

	var buf []byte
	// 编码头信息
	if withHead {
		buf = append(bytes.Int32ToBytes(contract.ContractCode()), bytes.Int64ToBytes(c.VersionMap[contract.ContractCode()])...)
	}

	// 编码字段信息
	for i := 0; i < rt.NumField(); i++ {
		tField := rt.Field(i)
		vField := rv.Field(i)

		if tField.Anonymous { // 匿名契约字段
			field, err := c.encode(vField.Interface().(DataContract), false)
			if err != nil {
				return nil, err
			}
			buf = append(buf, field...)
		} else {
			field, err := c.encodeField(tField, vField)
			if err != nil {
				return nil, err
			}
			buf = append(buf, field...)
		}
	}

	return buf, err
}

func (c *Codec) encodeField(tField reflect.StructField, vField reflect.Value) ([]byte, error) {
	_, _, _, primitiveType, refContract, refEnum, genericContract, _, numberMask, repeatable, err := resolveTags(tField)
	if err != nil {
		return nil, err
	}
	var buf []byte

	repeat := 1
	if repeatable {
		repeat = vField.Len()
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

	return buf, nil
}

/**
解码
*/
func (c *Codec) Decode(data []byte) (val interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic in Decode %s", r)
		}
	}()
	// 解析头信息
	code, _ := decodeHeader(data)

	mapLocker.RLock()
	contract := contractMap[code]
	mapLocker.RUnlock()

	value, _, err := c.decode(data[HEAD_BYTES:], contract)
	if err != nil {
		return nil, err
	}

	return value.Elem().Interface(), err
}

func (c *Codec) decode(data []byte, contract DataContract) (reflect.Value, int64, error) {
	var err error
	// 解析头信息
	offset := int64(0)
	rt := reflect.TypeOf(contract)
	obj := reflect.New(rt)
	rv := obj.Elem()

	// 解析字段信息
	for i := 0; i < rt.NumField(); i++ {
		tField := rt.Field(i)
		vField := rv.Field(i)
		size := int64(0)
		if tField.Anonymous { // 匿名契约字段
			var value reflect.Value
			value, size, err = c.decode(data[offset:], vField.Interface().(DataContract))
			if err == nil {
				vField.Set(value.Elem())
			}
		} else {
			size, err = c.decodeField(tField, vField, data[offset:])
		}
		if err != nil {
			return obj, offset, err
		}
		offset += size
	}

	return obj, offset, nil
}

func (c *Codec) decodeField(tField reflect.StructField, vField reflect.Value, data []byte) (int64, error) {
	var offset = int64(0)
	_, _, _, primitiveType, refContract, refEnum, genericContract, _, numberMask, repeatable, err := resolveTags(tField)
	if err != nil {
		return offset, err
	}

	vFieldOrigin := vField

	repeat := 1
	size := int64(0)
	if repeatable {
		// 编码数组头信息
		repeat, size = decodeArrayHeader(data[offset:])
		offset += size
		// 初始化数组
		if repeat > 0 {
			vField = reflect.MakeSlice(tField.Type, repeat, repeat)
		}
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
		vFieldOrigin.Set(vField)
	}

	return offset, nil
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
