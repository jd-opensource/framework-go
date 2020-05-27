package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午3:01
 */

// 键值操作的数据类型
type DataType uint8

const (
	NIL              = DataType(binary_proto.NIL)
	BOOLEAN          = DataType(binary_proto.BOOLEAN)
	INT8             = DataType(binary_proto.INT8)
	INT16            = DataType(binary_proto.INT16)
	INT32            = DataType(binary_proto.INT32)
	INT64            = DataType(binary_proto.INT64)
	TEXT             = DataType(binary_proto.TEXT)
	BYTES            = DataType(binary_proto.BYTES)
	TIMESTAMP        = DataType(binary_proto.BASE_TYPE_INTEGER | 0x08)
	JSON             = DataType(binary_proto.BASE_TYPE_TEXT | 0x01)
	XML              = DataType(binary_proto.BASE_TYPE_TEXT | 0x02)
	BIG_INT          = DataType(binary_proto.BASE_TYPE_BYTES | 0x01)
	IMG              = DataType(binary_proto.BASE_TYPE_BYTES | 0x02)
	VIDEO            = DataType(binary_proto.BASE_TYPE_BYTES | 0x03)
	LOCATION         = DataType(binary_proto.BASE_TYPE_BYTES | 0x04)
	PUB_KEY          = DataType(binary_proto.BASE_TYPE_BYTES | 0x05)
	SIGNATURE_DIGEST = DataType(binary_proto.BASE_TYPE_BYTES | 0x06)
	HASH_DIGEST      = DataType(binary_proto.BASE_TYPE_BYTES | 0x07)
	ENCRYPTED_DATA   = DataType(binary_proto.BASE_TYPE_BYTES | 0x08)
	DATA_CONTRACT    = DataType(binary_proto.BASE_TYPE_EXT | 0x01)
)

func init() {
	binary_proto.Cdc.RegisterEnum(NIL)
}

var _ binary_proto.EnumContract = (*DataType)(nil)

func (d DataType) ContractCode() int32 {
	return binary_proto.ENUM_TYPE_BYTES_VALUE_TYPE
}

func (d DataType) ContractType() string {
	return binary_proto.PRIMITIVETYPE_INT8
}

func (d DataType) ContractName() string {
	return "DataType"
}

func (d DataType) Description() string {
	return ""
}

func (d DataType) ContractVersion() int64 {
	return 0
}

func (d DataType) GetValue(CODE int32) binary_proto.EnumContract {
	switch CODE {
	case int32(binary_proto.NIL):
		return NIL
	case int32(binary_proto.BOOLEAN):
		return BOOLEAN
	case int32(binary_proto.INT8):
		return INT8
	case int32(binary_proto.INT16):
		return INT16
	case int32(binary_proto.INT32):
		return INT32
	case int32(binary_proto.INT64):
		return INT64
	case int32(binary_proto.TEXT):
		return TEXT
	case int32(binary_proto.BYTES):
		return BYTES
	case int32(binary_proto.BASE_TYPE_INTEGER | 0x08):
		return TIMESTAMP
	case int32(binary_proto.BASE_TYPE_TEXT | 0x01):
		return JSON
	case int32(binary_proto.BASE_TYPE_TEXT | 0x02):
		return XML
	case int32(binary_proto.BASE_TYPE_BYTES | 0x01):
		return BIG_INT
	case int32(binary_proto.BASE_TYPE_BYTES | 0x02):
		return IMG
	case int32(binary_proto.BASE_TYPE_BYTES | 0x03):
		return VIDEO
	case int32(binary_proto.BASE_TYPE_BYTES | 0x04):
		return LOCATION
	case int32(binary_proto.BASE_TYPE_BYTES | 0x05):
		return PUB_KEY
	case int32(binary_proto.BASE_TYPE_BYTES | 0x06):
		return SIGNATURE_DIGEST
	case int32(binary_proto.BASE_TYPE_BYTES | 0x07):
		return HASH_DIGEST
	case int32(binary_proto.BASE_TYPE_BYTES | 0x08):
		return ENCRYPTED_DATA
	case int32((byte)(binary_proto.BASE_TYPE_EXT | 0x01)):
		return DATA_CONTRACT
	}

	panic("no enum value founded")
}
