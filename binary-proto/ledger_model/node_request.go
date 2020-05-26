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
	Hash               []byte             `primitiveType:"BYTES"`
	TransactionContent TransactionContent `refContract:"528"`
	EndpointSignatures []DigitalSignature `refContract:"2864" repeatable:"true"`
	NodeSignatures     []DigitalSignature `refContract:"2864" repeatable:"true"`
}

func (n NodeRequest) Code() int32 {
	return binary_proto.REQUEST_NODE
}

func (n NodeRequest) Name() string {
	return "NodeRequest"
}

func (n NodeRequest) Description() string {
	return ""
}
