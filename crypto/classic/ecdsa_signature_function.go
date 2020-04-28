package classic

import "framework-go/crypto/framework"

/**
 * @Author: imuge
 * @Date: 2020/4/28 3:01 下午
 */
var _ framework.SignatureFunction = (*ECDSASignatureFunction)(nil)

// TODO
type ECDSASignatureFunction struct {

}

func (E ECDSASignatureFunction) GenerateKeypair() framework.AsymmetricKeypair {
	panic("implement me")
}

func (E ECDSASignatureFunction) GetAlgorithm() framework.CryptoAlgorithm {
	panic("implement me")
}

func (E ECDSASignatureFunction) Sign(privKey framework.PrivKey, data []byte) framework.SignatureDigest {
	panic("implement me")
}

func (E ECDSASignatureFunction) Verify(digest framework.SignatureDigest, pubKey framework.PubKey, data byte) bool {
	panic("implement me")
}

func (E ECDSASignatureFunction) RetrievePubKey(privKey framework.PrivKey) framework.PubKey {
	panic("implement me")
}

func (E ECDSASignatureFunction) SupportPrivKey(privKeyBytes []byte) bool {
	panic("implement me")
}

func (E ECDSASignatureFunction) ResolvePrivKey(privKeyBytes []byte) framework.PrivKey {
	panic("implement me")
}

func (E ECDSASignatureFunction) SupportPubKey(pubKeyBytes []byte) bool {
	panic("implement me")
}

func (E ECDSASignatureFunction) ResolvePubKey(pubKeyBytes []byte) framework.PubKey {
	panic("implement me")
}

func (E ECDSASignatureFunction) SupportDigest(digestBytes []byte) {
	panic("implement me")
}

func (E ECDSASignatureFunction) ResolveDigest(digestBytes []byte) framework.SignatureDigest {
	panic("implement me")
}
