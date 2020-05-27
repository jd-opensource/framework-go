package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午3:18
 */

// 交易（事务）执行状态
type TransactionState int8

const (
	// 成功
	SUCCESS = TransactionState(0)
	// 账本错误
	LEDGER_ERROR = TransactionState(0x01)
	// 数据账户不存在
	DATA_ACCOUNT_DOES_NOT_EXIST = TransactionState(0x02)
	// 用户不存在
	USER_DOES_NOT_EXIST = TransactionState(0x03)
	// 合约不存在
	CONTRACT_DOES_NOT_EXIST = TransactionState(0x04)
	// 数据写入时版本冲突
	DATA_VERSION_CONFLICT = TransactionState(0x05)
	// 参与方不存在
	PARTICIPANT_DOES_NOT_EXIST = TransactionState(0x06)
	// 被安全策略拒绝
	REJECTED_BY_SECURITY_POLICY = TransactionState(0x10)
	// 由于在错误的账本上执行交易而被丢弃
	IGNORED_BY_WRONG_LEDGER = TransactionState(0x40)
	// 由于交易内容的验签失败而丢弃
	IGNORED_BY_WRONG_CONTENT_SIGNATURE = TransactionState(0x41)
	// 由于交易内容的验签失败而丢弃
	IGNORED_BY_CONFLICTING_STATE = TransactionState(0x42)
	// 由于交易的整体回滚而丢弃
	IGNORED_BY_TX_FULL_ROLLBACK = TransactionState(0x43)
	// 由于区块的整体回滚而丢弃
	IGNORED_BY_BLOCK_FULL_ROLLBACK = TransactionState(0x44)
	// 共识阶段加入新区块哈希预计算功能, 如果来自其他Peer的新区块哈希值不一致，本批次整体回滚
	IGNORED_BY_CONSENSUS_PHASE_PRECOMPUTE_ROLLBACK = TransactionState(0x45)
	// 系统错误
	SYSTEM_ERROR = TransactionState(byte(0x80))
	// 超时
	TIMEOUT = TransactionState(byte(0x81))
	// 共识错误
	CONSENSUS_ERROR = TransactionState(byte(0x82))
)

func init() {
	binary_proto.Cdc.RegisterEnum(SUCCESS)
}

var _ binary_proto.EnumContract = (*TransactionState)(nil)

func (t TransactionState) ContractCode() int32 {
	return binary_proto.ENUM_TYPE_TRANSACTION_STATE
}

func (t TransactionState) ContractType() string {
	return binary_proto.PRIMITIVETYPE_INT8
}

func (t TransactionState) ContractName() string {
	return "TransactionState"
}

func (t TransactionState) Description() string {
	return ""
}

func (t TransactionState) ContractVersion() int64 {
	return 0
}

func (t TransactionState) GetValue(CODE int32) binary_proto.EnumContract {
	switch CODE {
	case int32(0):
		return SUCCESS
	case int32(0x01):
		return LEDGER_ERROR
	case int32(0x02):
		return DATA_ACCOUNT_DOES_NOT_EXIST
	case int32(0x03):
		return USER_DOES_NOT_EXIST
	case int32(0x04):
		return CONTRACT_DOES_NOT_EXIST
	case int32(0x05):
		return DATA_VERSION_CONFLICT
	case int32(0x06):
		return PARTICIPANT_DOES_NOT_EXIST
	case int32(0x10):
		return REJECTED_BY_SECURITY_POLICY
	case int32(0x40):
		return IGNORED_BY_WRONG_LEDGER
	case int32(0x41):
		return IGNORED_BY_WRONG_CONTENT_SIGNATURE
	case int32(0x42):
		return IGNORED_BY_CONFLICTING_STATE
	case int32(0x43):
		return IGNORED_BY_TX_FULL_ROLLBACK
	case int32(0x44):
		return IGNORED_BY_BLOCK_FULL_ROLLBACK
	case int32(0x45):
		return IGNORED_BY_CONSENSUS_PHASE_PRECOMPUTE_ROLLBACK
	case int32(byte(0x80)):
		return SYSTEM_ERROR
	case int32(byte(0x81)):
		return TIMEOUT
	case int32(byte(0x82)):
		return CONSENSUS_ERROR
	}

	panic("no enum value founded")
}
