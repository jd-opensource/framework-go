package classic

import "framework-go/crypto/framework"

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:52 下午
 */

var _ framework.AsymmetricEncryptionFunction = (*RSACryptoFunction)(nil)
var _ framework.SignatureFunction = (*RSACryptoFunction)(nil)

// TODO

type RSACryptoFunction struct {

}

func (R RSACryptoFunction) Sign(privKey framework.PrivKey, data []byte) framework.SignatureDigest {
	panic("implement me")
}

func (R RSACryptoFunction) Verify(digest framework.SignatureDigest, pubKey framework.PubKey, data byte) bool {
	panic("implement me")
}

func (R RSACryptoFunction) RetrievePubKey(privKey framework.PrivKey) framework.PubKey {
	panic("implement me")
}

func (R RSACryptoFunction) SupportPrivKey(privKeyBytes []byte) bool {
	panic("implement me")
}

func (R RSACryptoFunction) ResolvePrivKey(privKeyBytes []byte) framework.PrivKey {
	panic("implement me")
}

func (R RSACryptoFunction) SupportPubKey(pubKeyBytes []byte) bool {
	panic("implement me")
}

func (R RSACryptoFunction) ResolvePubKey(pubKeyBytes []byte) framework.PubKey {
	panic("implement me")
}

func (R RSACryptoFunction) SupportDigest(digestBytes []byte) {
	panic("implement me")
}

func (R RSACryptoFunction) ResolveDigest(digestBytes []byte) framework.SignatureDigest {
	panic("implement me")
}

func (R RSACryptoFunction) GenerateKeypair() framework.AsymmetricKeypair {
	panic("implement me")
}

func (R RSACryptoFunction) GetAlgorithm() framework.CryptoAlgorithm {
	panic("implement me")
}
