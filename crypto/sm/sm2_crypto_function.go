package sm

import "framework-go/crypto/framework"

/**
 * @Author: imuge
 * @Date: 2020/4/28 3:00 下午
 */

var _ framework.AsymmetricEncryptionFunction = (*SM2CryptoFunction)(nil)
var _ framework.SignatureFunction = (*SM2CryptoFunction)(nil)

// TODO

type SM2CryptoFunction struct {

}

func (S SM2CryptoFunction) Encrypt(pubKey framework.PubKey, data []byte) framework.AsymmetricCiphertext {
	panic("implement me")
}

func (S SM2CryptoFunction) Decrypt(privKey framework.PrivKey, ciphertext framework.AsymmetricCiphertext) []byte {
	panic("implement me")
}

func (S SM2CryptoFunction) SupportCiphertext(ciphertextBytes []byte) bool {
	panic("implement me")
}

func (S SM2CryptoFunction) ParseCiphertext(ciphertextBytes []byte) framework.AsymmetricCiphertext {
	panic("implement me")
}

func (S SM2CryptoFunction) Sign(privKey framework.PrivKey, data []byte) framework.SignatureDigest {
	panic("implement me")
}

func (S SM2CryptoFunction) Verify(pubKey framework.PubKey, data []byte, digest framework.SignatureDigest) bool {
	panic("implement me")
}

func (S SM2CryptoFunction) RetrievePubKey(privKey framework.PrivKey) framework.PubKey {
	panic("implement me")
}

func (S SM2CryptoFunction) SupportPrivKey(privKeyBytes []byte) bool {
	panic("implement me")
}

func (S SM2CryptoFunction) ParsePrivKey(privKeyBytes []byte) framework.PrivKey {
	panic("implement me")
}

func (S SM2CryptoFunction) SupportPubKey(pubKeyBytes []byte) bool {
	panic("implement me")
}

func (S SM2CryptoFunction) ParsePubKey(pubKeyBytes []byte) framework.PubKey {
	panic("implement me")
}

func (S SM2CryptoFunction) SupportDigest(digestBytes []byte) bool {
	panic("implement me")
}

func (S SM2CryptoFunction) ParseDigest(digestBytes []byte) framework.SignatureDigest {
	panic("implement me")
}

func (S SM2CryptoFunction) GenerateKeypair() framework.AsymmetricKeypair {
	panic("implement me")
}

func (S SM2CryptoFunction) GetAlgorithm() framework.CryptoAlgorithm {
	panic("implement me")
}
