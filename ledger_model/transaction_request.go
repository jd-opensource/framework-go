package ledger_model

import (
	binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
)

/*
 * Author: imuge
 * Date: 2020/5/26 上午11:37
 */

var _ binary_proto.DataContract = (*TransactionRequest)(nil)

func init() {
	binary_proto.RegisterContract(TransactionRequest{})
}

type TransactionRequest struct {
	TransactionHash    []byte              `primitiveType:"BYTES"`
	TransactionContent *TransactionContent `refContract:"512"`
	EndpointSignatures []*DigitalSignature `refContract:"2864" list:"true"`
	NodeSignatures     []*DigitalSignature `refContract:"2864" list:"true"`
}

func NewTransactionRequest(transactionHash []byte, content *TransactionContent) *TransactionRequest {
	return &TransactionRequest{
		TransactionHash:    transactionHash,
		TransactionContent: content,
	}
}

func (t TransactionRequest) ContractCode() int32 {
	return binary_proto.TX_REQUEST
}

func (t TransactionRequest) ContractName() string {
	return "TransactionRequest"
}

func (t TransactionRequest) Description() string {
	return ""
}

func (t *TransactionRequest) ContainsEndpointSignature(pubKey []byte) bool {
	for _, s := range t.EndpointSignatures {
		if bytes.Equals(s.PubKey, pubKey) {
			return true
		}
	}

	return false
}

func (t *TransactionRequest) AddEndpointSignatures(signature *DigitalSignature) {
	t.EndpointSignatures = append(t.EndpointSignatures, signature)
}
