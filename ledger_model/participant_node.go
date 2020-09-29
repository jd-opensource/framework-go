package ledger_model

import (
	binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"
)

/*
 * Author: imuge
 * Date: 2020/5/27 下午2:09
 */

var _ binary_proto.DataContract = (*ParticipantNode)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(ParticipantNode{})
}

// 参与方节点
type ParticipantNode struct {
	// 节点的顺序编号 TODO
	Id int32 `primitiveType:"INT32"`
	// 节点的虚拟地址，根据公钥生成
	Address []byte `primitiveType:"BYTES"`
	// 参与者名称
	Name string `primitiveType:"TEXT"`
	// 节点消息认证的公钥
	PubKey []byte `primitiveType:"BYTES"`
	// 节点的状态：已注册/已参与共识
	ParticipantNodeState ParticipantNodeState `refEnum:"2852"`
}

func (p ParticipantNode) ContractCode() int32 {
	return binary_proto.METADATA_CONSENSUS_PARTICIPANT
}

func (p ParticipantNode) ContractName() string {
	return "ParticipantNode"
}

func (p ParticipantNode) Description() string {
	return ""
}
