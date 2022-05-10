package adv

import (
	"errors"
	"github.com/blockchain-jd-com/framework-go/crypto/ca"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
)

var _ framework.AsymmetricEncryptionFunction = (*PaillierCryptoFunction)(nil)
var _ framework.SignatureFunction = (*PaillierCryptoFunction)(nil)

type PaillierCryptoFunction struct {
}

func (p PaillierCryptoFunction) Sign(privKey *framework.PrivKey, data []byte) (*framework.SignatureDigest, error) {
	return nil, errors.New("not implement")
}

func (p PaillierCryptoFunction) Verify(pubKey *framework.PubKey, data []byte, digest *framework.SignatureDigest) bool {
	return false
}

func (p PaillierCryptoFunction) SupportDigest(digestBytes []byte) bool {
	return false
}

func (p PaillierCryptoFunction) ParseDigest(digestBytes []byte) (*framework.SignatureDigest, error) {
	return nil, errors.New("not implement")
}

func (p PaillierCryptoFunction) RetrievePubKeyFromCA(cert *ca.Certificate) (*framework.PubKey, error) {
	return nil, errors.New("not implement")
}

func (p PaillierCryptoFunction) RetrievePrivKey(privateKey string) (*framework.PrivKey, error) {
	return nil, errors.New("not implement")
}

func (p PaillierCryptoFunction) RetrieveEncrypedPrivKey(privateKey string, pwd []byte) (*framework.PrivKey, error) {
	return nil, errors.New("not implement")
}

func (p PaillierCryptoFunction) GenerateKeypair() (*framework.AsymmetricKeypair, error) {
	return nil, errors.New("not implement")
}

func (p PaillierCryptoFunction) GenerateKeypairWithSeed(seed []byte) (*framework.AsymmetricKeypair, error) {
	return nil, errors.New("not implement")
}

func (p PaillierCryptoFunction) GetAlgorithm() framework.CryptoAlgorithm {
	return PAILLIER_ALGORITHM
}

func (p PaillierCryptoFunction) Encrypt(pubKey *framework.PubKey, data []byte) (*framework.AsymmetricCiphertext, error) {
	return nil, errors.New("not implement")
}

func (p PaillierCryptoFunction) Decrypt(privKey *framework.PrivKey, ciphertext *framework.AsymmetricCiphertext) ([]byte, error) {
	return nil, errors.New("not implement")
}

func (p PaillierCryptoFunction) RetrievePubKey(privKey *framework.PrivKey) (*framework.PubKey, error) {
	return nil, errors.New("not implement")
}

func (p PaillierCryptoFunction) SupportPrivKey(privKeyBytes []byte) bool {
	return false
}

func (p PaillierCryptoFunction) ParsePrivKey(privKeyBytes []byte) (*framework.PrivKey, error) {
	return nil, errors.New("not implement")
}

func (p PaillierCryptoFunction) SupportPubKey(pubKeyBytes []byte) bool {
	return false
}

func (p PaillierCryptoFunction) ParsePubKey(pubKeyBytes []byte) (*framework.PubKey, error) {
	return nil, errors.New("not implement")
}

func (p PaillierCryptoFunction) SupportCiphertext(ciphertextBytes []byte) bool {
	return false
}

func (p PaillierCryptoFunction) ParseCiphertext(ciphertextBytes []byte) (*framework.AsymmetricCiphertext, error) {
	return nil, errors.New("not implement")
}
