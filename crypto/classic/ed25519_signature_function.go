package classic

import "framework-go/crypto/framework"

/**
 * @Author: imuge
 * @Date: 2020/4/28 3:02 下午
 */

var _ framework.SignatureFunction = (*ED25519SignatureFunction)(nil)

// TODO

type ED25519SignatureFunction struct {

}

func (E ED25519SignatureFunction) GenerateKeypair() framework.AsymmetricKeypair {
	panic("implement me")
}

func (E ED25519SignatureFunction) GetAlgorithm() framework.CryptoAlgorithm {
	panic("implement me")
}

func (E ED25519SignatureFunction) Sign(privKey framework.PrivKey, data []byte) framework.SignatureDigest {
	panic("implement me")
}

func (E ED25519SignatureFunction) Verify(pubKey framework.PubKey, data []byte, digest framework.SignatureDigest) bool {
	panic("implement me")
}

func (E ED25519SignatureFunction) RetrievePubKey(privKey framework.PrivKey) framework.PubKey {
	panic("implement me")
}

func (E ED25519SignatureFunction) SupportPrivKey(privKeyBytes []byte) bool {
	panic("implement me")
}

func (E ED25519SignatureFunction) ParsePrivKey(privKeyBytes []byte) framework.PrivKey {
	panic("implement me")
}

func (E ED25519SignatureFunction) SupportPubKey(pubKeyBytes []byte) bool {
	panic("implement me")
}

func (E ED25519SignatureFunction) ParsePubKey(pubKeyBytes []byte) framework.PubKey {
	panic("implement me")
}

func (E ED25519SignatureFunction) SupportDigest(digestBytes []byte) bool {
	panic("implement me")
}

func (E ED25519SignatureFunction) ParseDigest(digestBytes []byte) framework.SignatureDigest {
	panic("implement me")
}
