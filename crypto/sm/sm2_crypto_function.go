package classic

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

func (S SM2CryptoFunction) Sign(privKey framework.PrivKey, data []byte) framework.SignatureDigest {
	panic("implement me")
}

func (S SM2CryptoFunction) Verify(digest framework.SignatureDigest, pubKey framework.PubKey, data byte) bool {
	panic("implement me")
}

func (S SM2CryptoFunction) RetrievePubKey(privKey framework.PrivKey) framework.PubKey {
	panic("implement me")
}

func (S SM2CryptoFunction) SupportPrivKey(privKeyBytes []byte) bool {
	panic("implement me")
}

func (S SM2CryptoFunction) ResolvePrivKey(privKeyBytes []byte) framework.PrivKey {
	panic("implement me")
}

func (S SM2CryptoFunction) SupportPubKey(pubKeyBytes []byte) bool {
	panic("implement me")
}

func (S SM2CryptoFunction) ResolvePubKey(pubKeyBytes []byte) framework.PubKey {
	panic("implement me")
}

func (S SM2CryptoFunction) SupportDigest(digestBytes []byte) {
	panic("implement me")
}

func (S SM2CryptoFunction) ResolveDigest(digestBytes []byte) framework.SignatureDigest {
	panic("implement me")
}

func (S SM2CryptoFunction) GenerateKeypair() framework.AsymmetricKeypair {
	panic("implement me")
}

func (S SM2CryptoFunction) GetAlgorithm() framework.CryptoAlgorithm {
	panic("implement me")
}
