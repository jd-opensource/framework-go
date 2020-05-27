package ledger_model

import (
	"framework-go/crypto"
	"framework-go/crypto/framework"
)

/*
 * Author: imuge
 * Date: 2020/5/28 下午5:47
 */

func Sign(txContent TransactionContent, keyPair framework.AsymmetricKeypair) DigitalSignature {
	signatureDigest := SignWithPrivKey(txContent, keyPair.PrivKey)
	return DigitalSignature{
		DigitalSignatureBody{
			PubKey: keyPair.PubKey.ToBytes(),
			Digest: signatureDigest.ToBytes(),
		},
	}
}

func SignWithPrivKey(txContent TransactionContent, privKey framework.PrivKey) framework.SignatureDigest {
	return crypto.GetSignatureFunctionByCode(privKey.GetAlgorithm()).Sign(privKey, txContent.Hash)
}

func VerifySignature(txContent TransactionContent, signDigest framework.SignatureDigest, pubKey framework.PubKey) bool {
	if !verifyTxContentHash(txContent.TransactionContentBody, framework.ParseHashDigest(txContent.Hash)) {
		return false
	}

	return VerifyHashSignature(txContent.Hash, signDigest, pubKey)
}

func VerifyHashSignature(hash []byte, signDigest framework.SignatureDigest, pubKey framework.PubKey) bool {
	return crypto.GetSignatureFunctionByCode(pubKey.GetAlgorithm()).Verify(pubKey, hash, signDigest)
}
