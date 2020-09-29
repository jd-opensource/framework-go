package ledger_model

import (
	binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"
	"github.com/blockchain-jd-com/framework-go/crypto"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
)

/*
 * Author: imuge
 * Date: 2020/5/28 下午5:41
 */

var _ TransactionRequestBuilder = (*TxRequestBuilder)(nil)

const DEFAULT_HASH_ALGORITHM = "SHA256"

type TxRequestBuilder struct {
	txContent          TransactionContent
	endpointSignatures []DigitalSignature
	nodeSignatures     []DigitalSignature
}

func NewTxRequestBuilder(txContent TransactionContent) *TxRequestBuilder {
	return &TxRequestBuilder{
		txContent: txContent,
	}
}

func (t *TxRequestBuilder) GetHash() framework.HashDigest {
	return framework.ParseHashDigest(t.txContent.Hash)
}

func (t *TxRequestBuilder) GetTransactionContent() TransactionContent {
	return t.txContent
}

func (t *TxRequestBuilder) SignAsEndpoint(keyPair framework.AsymmetricKeypair) DigitalSignature {
	signature := Sign(t.txContent, keyPair)
	t.AddEndpointSignature(signature)
	return signature
}

func (t *TxRequestBuilder) SignAsNode(keyPair framework.AsymmetricKeypair) DigitalSignature {
	signature := Sign(t.txContent, keyPair)
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
	txMessage := NewTransactionRequest(t.txContent)
	txMessage.EndpointSignatures = append(txMessage.EndpointSignatures, t.endpointSignatures...)
	txMessage.NodeSignatures = append(txMessage.NodeSignatures, t.nodeSignatures...)
	reqBytes, err := binary_proto.Cdc.Encode(txMessage.NodeRequest)
	if err != nil {
		panic(err)
	}
	reqHash := crypto.GetHashFunctionByName(DEFAULT_HASH_ALGORITHM).Hash(reqBytes)
	txMessage.Hash = reqHash.ToBytes()

	return txMessage
}
