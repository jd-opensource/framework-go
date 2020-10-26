package ledger_model

import "github.com/blockchain-jd-com/framework-go/crypto/framework"

/*
 * Author: imuge
 * Date: 2020/5/28 下午5:36
 */

// 已就绪的交易
type TransactionRequestBuilder interface {
	GetHash() framework.HashDigest

	GetTransactionContent() TransactionContent

	SignAsEndpoint(keyPair framework.AsymmetricKeypair) DigitalSignature

	SignAsNode(keyPair framework.AsymmetricKeypair) DigitalSignature

	AddEndpointSignature(signature DigitalSignature)

	AddNodeSignature(signature DigitalSignature)

	BuildRequest() TransactionRequest
}
