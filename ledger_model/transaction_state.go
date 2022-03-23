package ledger_model

import (
	binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"
)

/*
 * Author: imuge
 * Date: 2020/5/27 下午3:18
 */

// 交易（事务）执行状态
type TransactionState int8

const (
	// 成功
	SUCCESS = TransactionState(0)
	// 账本的未知错误
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
	// 事件账户不存在
	EVENT_ACCOUNT_DOES_NOT_EXIST = TransactionState(0x07)
	// 合约部署时版本冲突
	CONTRACT_VERSION_CONFLICT = TransactionState(0x08)
	// 账户状态错误
	ILLEGAL_ACCOUNT_STATE = TransactionState(0x09)
	// 合约执行错误
	CONTRACT_EXECUTE_ERROR = TransactionState(0x10)
	// 被安全策略拒绝
	REJECTED_BY_SECURITY_POLICY = TransactionState(0x11)
	// 账户注册冲突
	ACCOUNT_REGISTER_CONFLICT = TransactionState(0x12)
	// 角色不存在
	ROLE_DOES_NOT_EXIST = TransactionState(0x13)
	// 不支持的HASH算法
	UNSUPPORTED_HASH_ALGORITHM = TransactionState(0x14)
	// 合约方法不存在
	CONTRACT_METHOD_NOT_FOUND = TransactionState(0x15)
	//合约参数错误
	CONTRACT_PARAMETER_ERROR = TransactionState(0x16)
	// 由于在错误的账本上执行交易而被忽略
	IGNORED_BY_WRONG_LEDGER = TransactionState(0x40)
	// 由于交易内容的验签失败而忽略
	IGNORED_BY_ILLEGAL_CONTENT_SIGNATURE = TransactionState(0x41)
	// 由于交易内容哈希不一致而忽略
	IGNORED_BY_INCONSISTENT_CONTENT_HASH = TransactionState(0x42)
	// 由于交易的整体回滚而丢弃
	IGNORED_BY_TX_FULL_ROLLBACK = TransactionState(0x43)
	// 由于区块的整体回滚而丢弃
	IGNORED_BY_BLOCK_FULL_ROLLBACK = TransactionState(0x44)
	// 系统错误
	SYSTEM_ERROR = TransactionState(-128)
	// 超时
	TIMEOUT = TransactionState(-127)
	// 共识错误
	CONSENSUS_ERROR = TransactionState(-126)
	// 未收到共识网络响应的错误
	CONSENSUS_NO_REPLY_ERROR = TransactionState(-125)
	// 创建共识的代理客户端错误
	CONSENSUS_PROXY_CLIENT_ERROR = TransactionState(-124)
	// 空区块错误
	EMPTY_BLOCK_ERROR = TransactionState(-123)
	// 共识时间戳错误
	CONSENSUS_TIMESTAMP_ERROR = TransactionState(-122)
	// 账本参数丢失
	LEDGER_HASH_EMPTY = TransactionState(-48)
	// 不合法的合约包
	ILLEGAL_CONTRACT_CAR = TransactionState(-47)
	// 不合法的节点签名
	ILLEGAL_NODE_SIGNATURE = TransactionState(-46)
	// 没有终端签名
	NO_ENDPOINT_SIGNATURE = TransactionState(-45)
	// 终端签名验证不通过
	INVALID_ENDPOINT_SIGNATURE = TransactionState(-44)
)

func init() {
	binary_proto.RegisterEnum(SUCCESS)
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
	case int32(0x07):
		return EVENT_ACCOUNT_DOES_NOT_EXIST
	case int32(0x08):
		return CONTRACT_VERSION_CONFLICT
	case int32(0x09):
		return ILLEGAL_ACCOUNT_STATE
	case int32(0x10):
		return CONTRACT_EXECUTE_ERROR
	case int32(0x11):
		return REJECTED_BY_SECURITY_POLICY
	case int32(0x12):
		return ACCOUNT_REGISTER_CONFLICT
	case int32(0x13):
		return ROLE_DOES_NOT_EXIST
	case int32(0x14):
		return UNSUPPORTED_HASH_ALGORITHM
	case int32(0x15):
		return CONTRACT_METHOD_NOT_FOUND
	case int32(0x16):
		return CONTRACT_PARAMETER_ERROR
	case int32(0x40):
		return IGNORED_BY_WRONG_LEDGER
	case int32(0x41):
		return IGNORED_BY_ILLEGAL_CONTENT_SIGNATURE
	case int32(0x42):
		return IGNORED_BY_INCONSISTENT_CONTENT_HASH
	case int32(0x43):
		return IGNORED_BY_TX_FULL_ROLLBACK
	case int32(0x44):
		return IGNORED_BY_BLOCK_FULL_ROLLBACK
	case -int32(0x80):
		return SYSTEM_ERROR
	case -int32(0x81):
		return TIMEOUT
	case -int32(0x82):
		return CONSENSUS_ERROR
	case -int32(0x83):
		return CONSENSUS_NO_REPLY_ERROR
	case -int32(0x84):
		return CONSENSUS_PROXY_CLIENT_ERROR
	case -int32(0x85):
		return EMPTY_BLOCK_ERROR
	case -int32(0x86):
		return CONSENSUS_TIMESTAMP_ERROR
	case -int32(0xd0):
		return LEDGER_HASH_EMPTY
	case -int32(0xd1):
		return ILLEGAL_CONTRACT_CAR
	case -int32(0xd2):
		return ILLEGAL_NODE_SIGNATURE
	case -int32(0xd3):
		return NO_ENDPOINT_SIGNATURE
	case -int32(0xd4):
		return INVALID_ENDPOINT_SIGNATURE
	default:
		return SYSTEM_ERROR
	}
}

func (t TransactionState) GetValueByName(name string) binary_proto.EnumContract {
	switch name {
	case "SUCCESS":
		return SUCCESS
	case "LEDGER_ERROR":
		return LEDGER_ERROR
	case "DATA_ACCOUNT_DOES_NOT_EXIST":
		return DATA_ACCOUNT_DOES_NOT_EXIST
	case "USER_DOES_NOT_EXIST":
		return USER_DOES_NOT_EXIST
	case "CONTRACT_DOES_NOT_EXIST":
		return CONTRACT_DOES_NOT_EXIST
	case "DATA_VERSION_CONFLICT":
		return DATA_VERSION_CONFLICT
	case "PARTICIPANT_DOES_NOT_EXIST":
		return PARTICIPANT_DOES_NOT_EXIST
	case "EVENT_ACCOUNT_DOES_NOT_EXIST":
		return EVENT_ACCOUNT_DOES_NOT_EXIST
	case "CONTRACT_VERSION_CONFLICT":
		return CONTRACT_VERSION_CONFLICT
	case "ILLEGAL_ACCOUNT_STATE":
		return ILLEGAL_ACCOUNT_STATE
	case "CONTRACT_EXECUTE_ERROR":
		return CONTRACT_EXECUTE_ERROR
	case "REJECTED_BY_SECURITY_POLICY":
		return REJECTED_BY_SECURITY_POLICY
	case "ACCOUNT_REGISTER_CONFLICT":
		return ACCOUNT_REGISTER_CONFLICT
	case "ROLE_DOES_NOT_EXIST":
		return ROLE_DOES_NOT_EXIST
	case "UNSUPPORTED_HASH_ALGORITHM":
		return UNSUPPORTED_HASH_ALGORITHM
	case "CONTRACT_METHOD_NOT_FOUND":
		return CONTRACT_METHOD_NOT_FOUND
	case "CONTRACT_PARAMETER_ERROR":
		return CONTRACT_PARAMETER_ERROR
	case "IGNORED_BY_WRONG_LEDGER":
		return IGNORED_BY_WRONG_LEDGER
	case "IGNORED_BY_ILLEGAL_CONTENT_SIGNATURE":
		return IGNORED_BY_ILLEGAL_CONTENT_SIGNATURE
	case "IGNORED_BY_INCONSISTENT_CONTENT_HASH":
		return IGNORED_BY_INCONSISTENT_CONTENT_HASH
	case "IGNORED_BY_TX_FULL_ROLLBACK":
		return IGNORED_BY_TX_FULL_ROLLBACK
	case "IGNORED_BY_BLOCK_FULL_ROLLBACK":
		return IGNORED_BY_BLOCK_FULL_ROLLBACK
	case "SYSTEM_ERROR":
		return SYSTEM_ERROR
	case "TIMEOUT":
		return TIMEOUT
	case "CONSENSUS_ERROR":
		return CONSENSUS_ERROR
	case "CONSENSUS_NO_REPLY_ERROR":
		return CONSENSUS_NO_REPLY_ERROR
	case "CONSENSUS_PROXY_CLIENT_ERROR":
		return CONSENSUS_PROXY_CLIENT_ERROR
	case "EMPTY_BLOCK_ERROR":
		return EMPTY_BLOCK_ERROR
	case "CONSENSUS_TIMESTAMP_ERROR":
		return CONSENSUS_TIMESTAMP_ERROR
	case "LEDGER_HASH_EMPTY":
		return LEDGER_HASH_EMPTY
	case "ILLEGAL_CONTRACT_CAR":
		return ILLEGAL_CONTRACT_CAR
	case "ILLEGAL_NODE_SIGNATURE":
		return ILLEGAL_NODE_SIGNATURE
	case "NO_ENDPOINT_SIGNATURE":
		return NO_ENDPOINT_SIGNATURE
	case "INVALID_ENDPOINT_SIGNATURE":
		return INVALID_ENDPOINT_SIGNATURE
	default:
		return SYSTEM_ERROR
	}
}

func (t TransactionState) ToString() string {
	switch t {
	case SUCCESS:
		return "SUCCESS"
	case LEDGER_ERROR:
		return "LEDGER_ERROR"
	case DATA_ACCOUNT_DOES_NOT_EXIST:
		return "DATA_ACCOUNT_DOES_NOT_EXIST"
	case USER_DOES_NOT_EXIST:
		return "USER_DOES_NOT_EXIST"
	case CONTRACT_DOES_NOT_EXIST:
		return "CONTRACT_DOES_NOT_EXIST"
	case DATA_VERSION_CONFLICT:
		return "DATA_VERSION_CONFLICT"
	case PARTICIPANT_DOES_NOT_EXIST:
		return "PARTICIPANT_DOES_NOT_EXIST"
	case EVENT_ACCOUNT_DOES_NOT_EXIST:
		return "EVENT_ACCOUNT_DOES_NOT_EXIST"
	case CONTRACT_VERSION_CONFLICT:
		return "CONTRACT_VERSION_CONFLICT"
	case ILLEGAL_ACCOUNT_STATE:
		return "ILLEGAL_ACCOUNT_STATE"
	case CONTRACT_EXECUTE_ERROR:
		return "CONTRACT_EXECUTE_ERROR"
	case REJECTED_BY_SECURITY_POLICY:
		return "REJECTED_BY_SECURITY_POLICY"
	case ACCOUNT_REGISTER_CONFLICT:
		return "ACCOUNT_REGISTER_CONFLICT"
	case ROLE_DOES_NOT_EXIST:
		return "ROLE_DOES_NOT_EXIST"
	case IGNORED_BY_WRONG_LEDGER:
		return "IGNORED_BY_WRONG_LEDGER"
	case IGNORED_BY_ILLEGAL_CONTENT_SIGNATURE:
		return "IGNORED_BY_ILLEGAL_CONTENT_SIGNATURE"
	case IGNORED_BY_INCONSISTENT_CONTENT_HASH:
		return "IGNORED_BY_INCONSISTENT_CONTENT_HASH"
	case IGNORED_BY_TX_FULL_ROLLBACK:
		return "IGNORED_BY_TX_FULL_ROLLBACK"
	case IGNORED_BY_BLOCK_FULL_ROLLBACK:
		return "IGNORED_BY_BLOCK_FULL_ROLLBACK"
	case SYSTEM_ERROR:
		return "SYSTEM_ERROR"
	case TIMEOUT:
		return "TIMEOUT"
	case CONSENSUS_ERROR:
		return "CONSENSUS_ERROR"
	case CONSENSUS_NO_REPLY_ERROR:
		return "CONSENSUS_NO_REPLY_ERROR"
	case CONSENSUS_PROXY_CLIENT_ERROR:
		return "CONSENSUS_PROXY_CLIENT_ERROR"
	case EMPTY_BLOCK_ERROR:
		return "EMPTY_BLOCK_ERROR"
	case CONSENSUS_TIMESTAMP_ERROR:
		return "CONSENSUS_TIMESTAMP_ERROR"
	case LEDGER_HASH_EMPTY:
		return "LEDGER_HASH_EMPTY"
	case ILLEGAL_CONTRACT_CAR:
		return "ILLEGAL_CONTRACT_CAR"
	case ILLEGAL_NODE_SIGNATURE:
		return "ILLEGAL_NODE_SIGNATURE"
	case NO_ENDPOINT_SIGNATURE:
		return "NO_ENDPOINT_SIGNATURE"
	case INVALID_ENDPOINT_SIGNATURE:
		return "INVALID_ENDPOINT_SIGNATURE"
	default:
		return "SYSTEM_ERROR"
	}
}
