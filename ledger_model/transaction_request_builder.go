package ledger_model

import "github.com/blockchain-jd-com/framework-go/crypto/framework"

/*
 * Author: imuge
 * Date: 2020/5/28 下午5:36
 */

// 已就绪的交易
type TransactionRequestBuilder interface {
	GetTransactionHash() *framework.HashDigest

	GetTransactionContent() *TransactionContent

	SignAsEndpoint(keyPair *framework.AsymmetricKeypair) (*DigitalSignature, error)

	SignAsNode(keyPair *framework.AsymmetricKeypair) (*DigitalSignature, error)

	AddEndpointSignature(signature *DigitalSignature)

	AddNodeSignature(signature *DigitalSignature)

	BuildRequest() *TransactionRequest
}
