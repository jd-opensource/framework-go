package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/26 下午2:03
 */

var _ binary_proto.DataContract = (*NodeRequest)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(NodeRequest{})
}

type NodeRequest struct {
	EndpointRequest
	// 接入交易的节点的签名
	NodeSignatures []DigitalSignature `refContract:"2864" list:"true"`
}

func (n NodeRequest) ContractCode() int32 {
	return binary_proto.REQUEST_NODE
}

func (n NodeRequest) ContractName() string {
	return "NodeRequest"
}

func (n NodeRequest) Description() string {
	return ""
}
