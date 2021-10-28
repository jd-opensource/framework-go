package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

// 底层数据结构
type LedgerDataStructure int8

const (
	// 正常
	MERKLE_TREE LedgerDataStructure = iota + 1
	// 冻结
	KV
)

func init() {
	binary_proto.RegisterEnum(MERKLE_TREE)
}

var _ binary_proto.EnumContract = (*LedgerDataStructure)(nil)

func (ls LedgerDataStructure) ContractCode() int32 {
	return binary_proto.METADATA_LEDGER_DATA_STRUCTURE
}

func (ls LedgerDataStructure) ContractType() string {
	return binary_proto.PRIMITIVETYPE_INT8
}

func (ls LedgerDataStructure) ContractName() string {
	return "LedgerDataStructure"
}

func (ls LedgerDataStructure) Description() string {
	return ""
}

func (ls LedgerDataStructure) ContractVersion() int64 {
	return 0
}

func (ls LedgerDataStructure) GetValue(CODE int32) binary_proto.EnumContract {
	if CODE == int32(1) {
		return MERKLE_TREE
	} else if CODE == int32(2) {
		return KV
	}

	panic("no enum value founded")
}

func (ls LedgerDataStructure) GetValueByName(name string) binary_proto.EnumContract {
	if name == "MERKLE_TREE" {
		return MERKLE_TREE
	} else if name == "KV" {
		return KV
	}

	panic("no enum value founded")
}
