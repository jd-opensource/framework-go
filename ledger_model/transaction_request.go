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
	NodeRequest
	Hash []byte `primitiveType:"BYTES"`
}

func NewTransactionRequest(content TransactionContent) TransactionRequest {
	return TransactionRequest{
		NodeRequest: NodeRequest{
			EndpointRequest: EndpointRequest{
				TransactionContent: content,
			},
		},
	}
}

func (t TransactionRequest) ContractCode() int32 {
	return binary_proto.REQUEST
}

func (t TransactionRequest) ContractName() string {
	return "TransactionRequest"
}

func (t TransactionRequest) Description() string {
	return ""
}
