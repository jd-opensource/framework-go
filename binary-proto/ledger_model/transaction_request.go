package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/26 上午11:37
 */

var _ binary_proto.DataContract = (*TransactionRequest)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(TransactionRequest{})
}

type TransactionRequest struct {
	TransactionContent TransactionContent `refContract:"528"`
	EndpointSignatures []DigitalSignature `refContract:"2864" repeatable:"true"`
	NodeSignatures     []DigitalSignature `refContract:"2864" repeatable:"true"`
	Hash               []byte             `primitiveType:"BYTES"`
}

func (t TransactionRequest) Code() int32 {
	return binary_proto.REQUEST
}

func (t TransactionRequest) Name() string {
	return "TransactionRequest"
}

func (t TransactionRequest) Description() string {
	return ""
}
