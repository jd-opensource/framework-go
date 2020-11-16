package ledger_model

import (
	"github.com/blockchain-jd-com/framework-go/crypto"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
)

/*
 * Author: imuge
 * Date: 2020/5/28 下午5:47
 */

func SignBytes(transactionHash []byte, keyPair framework.AsymmetricKeypair) DigitalSignature {
	signatureDigest := SignBytesWithPrivKey(transactionHash, keyPair.PrivKey)
	return DigitalSignature{
		DigitalSignatureBody{
			PubKey: keyPair.PubKey.ToBytes(),
			Digest: signatureDigest.ToBytes(),
		},
	}
}

func Sign(transactionHash framework.HashDigest, keyPair framework.AsymmetricKeypair) DigitalSignature {
	signatureDigest := SignWithPrivKey(transactionHash, keyPair.PrivKey)
	return DigitalSignature{
		DigitalSignatureBody{
			PubKey: keyPair.PubKey.ToBytes(),
			Digest: signatureDigest.ToBytes(),
		},
	}
}

func SignWithPrivKey(hash framework.HashDigest, privKey framework.PrivKey) framework.SignatureDigest {
	return crypto.GetSignatureFunctionByCode(privKey.GetAlgorithm()).Sign(privKey, hash.ToBytes())
}

func SignBytesWithPrivKey(hash []byte, privKey framework.PrivKey) framework.SignatureDigest {
	return crypto.GetSignatureFunctionByCode(privKey.GetAlgorithm()).Sign(privKey, hash)
}

func VerifySignature(hashAlgorithm int16, txContent TransactionContent, signDigest framework.SignatureDigest, pubKey framework.PubKey) bool {
	return VerifyHashSignature(ComputeTxContentHash(crypto.GetCryptoFunctionByCode(hashAlgorithm).GetAlgorithm(), txContent).ToBytes(), signDigest, pubKey)
}

func VerifyHashSignature(hash []byte, signDigest framework.SignatureDigest, pubKey framework.PubKey) bool {
	return crypto.GetSignatureFunctionByCode(pubKey.GetAlgorithm()).Verify(pubKey, hash, signDigest)
}
