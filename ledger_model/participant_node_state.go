package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午2:21
 */

// 参与方节点状态
type ParticipantNodeState int8

const (
	// 已注册
	READY ParticipantNodeState = iota
	// 已激活
	CONSENSUS
	// 不可用
	DECONSENSUS
)

func init() {
	binary_proto.Cdc.RegisterEnum(READY)
}

var _ binary_proto.EnumContract = (*ParticipantNodeState)(nil)

func (p ParticipantNodeState) ContractCode() int32 {
	return binary_proto.ENUM_TYPE_PARTICIPANT_NODE_STATE
}

func (p ParticipantNodeState) ContractType() string {
	return binary_proto.PRIMITIVETYPE_INT8
}

func (p ParticipantNodeState) ContractName() string {
	return "ParticipantNodeState"
}

func (p ParticipantNodeState) Description() string {
	return ""
}

func (p ParticipantNodeState) ContractVersion() int64 {
	return 0
}

func (p ParticipantNodeState) GetValue(CODE int32) binary_proto.EnumContract {
	if CODE == 0 {
		return READY
	}
	if CODE == 1 {
		return CONSENSUS
	}
	if CODE == 2 {
		return DECONSENSUS
	}

	panic("no enum value founded")
}

func (p ParticipantNodeState) GetValueByName(name string) binary_proto.EnumContract {
	if name == "READY" {
		return READY
	}
	if name == "CONSENSUS" {
		return CONSENSUS
	}
	if name == "DECONSENSUS" {
		return DECONSENSUS
	}

	panic("no enum value founded")
}
