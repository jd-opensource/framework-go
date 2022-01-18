package ledger_model

import "github.com/blockchain-jd-com/framework-go/crypto/framework"

/*
 * Author: imuge
 * Date: 2020/5/28 下午6:20
 */

type PreparedTransaction interface {
	GetHash() *framework.HashDigest

	GetTransactionContent() *TransactionContent

	Sign(keyPair *framework.AsymmetricKeypair) (*DigitalSignature, error)

	AddSignature(signature *DigitalSignature)

	Commit() (*TransactionResponse, error)
}
