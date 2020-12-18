package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/29 下午5:16
 */

type LedgerPermission int8

const (
	// 配置角色的权限
	CONFIGURE_ROLES = LedgerPermission(0x01)
	// 授权用户角色
	AUTHORIZE_USER_ROLES = LedgerPermission(0x02)
	// 设置共识协议
	SET_CONSENSUS = LedgerPermission(0x03)
	// 设置密码体系
	SET_CRYPTO = LedgerPermission(0x04)
	// 注册参与方
	REGISTER_PARTICIPANT = LedgerPermission(0x05)
	// 注册用户
	REGISTER_USER = LedgerPermission(0x06)
	// 注册数据账户
	REGISTER_DATA_ACCOUNT = LedgerPermission(0x07)
	// 注册合约
	REGISTER_CONTRACT = LedgerPermission(0x08)
	// 升级合约
	UPGRADE_CONTRACT = LedgerPermission(0x14)
	// 设置用户属性
	SET_USER_ATTRIBUTES = LedgerPermission(0x09)
	// 写入数据账户
	WRITE_DATA_ACCOUNT = LedgerPermission(0x0A)
	// 参与方核准交易
	// 如果不具备此项权限，则无法作为节点签署由终端提交的交易
	// 只对交易请求的节点签名列表{@link TransactionRequest#getNodeSignatures()}的用户产生影响
	APPROVE_TX = LedgerPermission(0x0B)
	// 参与方共识交易
	// 如果不具备此项权限，则无法作为共识节点接入并对交易进行共识
	CONSENSUS_TX = LedgerPermission(0x0C)
	// 注册事件账户
	REGISTER_EVENT_ACCOUNT = LedgerPermission(0x0D)
	// 发布事件
	WRITE_EVENT_ACCOUNT = LedgerPermission(0x0E)
)

func init() {
	binary_proto.RegisterEnum(CONFIGURE_ROLES)
}

var _ binary_proto.EnumContract = (*LedgerPermission)(nil)

func (l LedgerPermission) ContractCode() int32 {
	return binary_proto.ENUM_LEDGER_PERMISSION
}

func (l LedgerPermission) ContractType() string {
	return binary_proto.PRIMITIVETYPE_INT8
}

func (l LedgerPermission) ContractName() string {
	return "LedgerPermission"
}

func (l LedgerPermission) Description() string {
	return ""
}

func (l LedgerPermission) ContractVersion() int64 {
	return 0
}

func (l LedgerPermission) GetValue(CODE int32) binary_proto.EnumContract {
	if CODE == int32(CONFIGURE_ROLES) {
		return CONFIGURE_ROLES
	}
	if CODE == int32(AUTHORIZE_USER_ROLES) {
		return AUTHORIZE_USER_ROLES
	}
	if CODE == int32(SET_CONSENSUS) {
		return SET_CONSENSUS
	}
	if CODE == int32(SET_CRYPTO) {
		return SET_CRYPTO
	}
	if CODE == int32(REGISTER_PARTICIPANT) {
		return REGISTER_PARTICIPANT
	}
	if CODE == int32(REGISTER_USER) {
		return REGISTER_USER
	}
	if CODE == int32(REGISTER_DATA_ACCOUNT) {
		return REGISTER_DATA_ACCOUNT
	}
	if CODE == int32(REGISTER_CONTRACT) {
		return REGISTER_CONTRACT
	}
	if CODE == int32(UPGRADE_CONTRACT) {
		return UPGRADE_CONTRACT
	}
	if CODE == int32(SET_USER_ATTRIBUTES) {
		return SET_USER_ATTRIBUTES
	}
	if CODE == int32(WRITE_DATA_ACCOUNT) {
		return WRITE_DATA_ACCOUNT
	}
	if CODE == int32(APPROVE_TX) {
		return APPROVE_TX
	}
	if CODE == int32(CONSENSUS_TX) {
		return CONSENSUS_TX
	}
	if CODE == int32(REGISTER_EVENT_ACCOUNT) {
		return REGISTER_EVENT_ACCOUNT
	}
	if CODE == int32(WRITE_EVENT_ACCOUNT) {
		return WRITE_EVENT_ACCOUNT
	}

	panic("no enum value founded")
}

func (l LedgerPermission) GetValueByName(name string) binary_proto.EnumContract {
	if name == "CONFIGURE_ROLES" {
		return CONFIGURE_ROLES
	}
	if name == "AUTHORIZE_USER_ROLES" {
		return AUTHORIZE_USER_ROLES
	}
	if name == "SET_CONSENSUS" {
		return SET_CONSENSUS
	}
	if name == "SET_CRYPTO" {
		return SET_CRYPTO
	}
	if name == "REGISTER_PARTICIPANT" {
		return REGISTER_PARTICIPANT
	}
	if name == "REGISTER_USER" {
		return REGISTER_USER
	}
	if name == "REGISTER_DATA_ACCOUNT" {
		return REGISTER_DATA_ACCOUNT
	}
	if name == "REGISTER_CONTRACT" {
		return REGISTER_CONTRACT
	}
	if name == "UPGRADE_CONTRACT" {
		return UPGRADE_CONTRACT
	}
	if name == "SET_USER_ATTRIBUTES" {
		return SET_USER_ATTRIBUTES
	}
	if name == "WRITE_DATA_ACCOUNT" {
		return WRITE_DATA_ACCOUNT
	}
	if name == "APPROVE_TX" {
		return APPROVE_TX
	}
	if name == "CONSENSUS_TX" {
		return CONSENSUS_TX
	}
	if name == "REGISTER_EVENT_ACCOUNT" {
		return REGISTER_EVENT_ACCOUNT
	}
	if name == "WRITE_EVENT_ACCOUNT" {
		return WRITE_EVENT_ACCOUNT
	}

	panic("no enum value founded")
}
