package ledger_model

import (
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
)

/*
 * Author: imuge
 * Date: 2020/5/28 下午5:41
 */

var _ TransactionRequestBuilder = (*TxRequestBuilder)(nil)

type TxRequestBuilder struct {
	transactionHash    framework.HashDigest
	txContent          TransactionContent
	endpointSignatures []DigitalSignature
	nodeSignatures     []DigitalSignature
}

func NewTxRequestBuilder(transactionHash framework.HashDigest, txContent TransactionContent) *TxRequestBuilder {
	return &TxRequestBuilder{
		transactionHash: transactionHash,
		txContent:       txContent,
	}
}

func (t *TxRequestBuilder) GetTransactionHash() framework.HashDigest {
	return t.transactionHash
}

func (t *TxRequestBuilder) GetTransactionContent() TransactionContent {
	return t.txContent
}

func (t *TxRequestBuilder) SignAsEndpoint(keyPair framework.AsymmetricKeypair) DigitalSignature {
	signature := Sign(t.transactionHash, keyPair)
	t.AddEndpointSignature(signature)
	return signature
}

func (t *TxRequestBuilder) SignAsNode(keyPair framework.AsymmetricKeypair) DigitalSignature {
	signature := Sign(t.transactionHash, keyPair)
	t.AddNodeSignature(signature)
	return signature
}

func (t *TxRequestBuilder) AddEndpointSignature(signature DigitalSignature) {
	t.endpointSignatures = append(t.endpointSignatures, signature)
}

func (t *TxRequestBuilder) AddNodeSignature(signature DigitalSignature) {
	t.nodeSignatures = append(t.nodeSignatures, signature)
}

func (t *TxRequestBuilder) BuildRequest() TransactionRequest {
	txMessage := NewTransactionRequest(t.transactionHash.ToBytes(), t.txContent)
	txMessage.EndpointSignatures = append(txMessage.EndpointSignatures, t.endpointSignatures...)
	txMessage.NodeSignatures = append(txMessage.NodeSignatures, t.nodeSignatures...)

	return txMessage
}
