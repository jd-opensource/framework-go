package adv

import (
	"errors"
	"github.com/blockchain-jd-com/framework-go/crypto/ca"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
)

var _ framework.AsymmetricEncryptionFunction = (*ElgamalCryptoFunction)(nil)
var _ framework.SignatureFunction = (*ElgamalCryptoFunction)(nil)

type ElgamalCryptoFunction struct {
}

func (e ElgamalCryptoFunction) Sign(privKey *framework.PrivKey, data []byte) (*framework.SignatureDigest, error) {
	return nil, errors.New("not implement")
}

func (e ElgamalCryptoFunction) Verify(pubKey *framework.PubKey, data []byte, digest *framework.SignatureDigest) bool {
	return false
}

func (e ElgamalCryptoFunction) SupportDigest(digestBytes []byte) bool {
	return false
}

func (e ElgamalCryptoFunction) ParseDigest(digestBytes []byte) (*framework.SignatureDigest, error) {
	return nil, errors.New("not implement")
}

func (e ElgamalCryptoFunction) RetrievePubKeyFromCA(cert *ca.Certificate) (*framework.PubKey, error) {
	return nil, errors.New("not implement")
}

func (e ElgamalCryptoFunction) RetrievePrivKey(privateKey string) (*framework.PrivKey, error) {
	return nil, errors.New("not implement")
}

func (e ElgamalCryptoFunction) RetrieveEncrypedPrivKey(privateKey string, pwd []byte) (*framework.PrivKey, error) {
	return nil, errors.New("not implement")
}

func (e ElgamalCryptoFunction) GenerateKeypair() (*framework.AsymmetricKeypair, error) {
	return nil, errors.New("not implement")
}

func (e ElgamalCryptoFunction) GenerateKeypairWithSeed(seed []byte) (*framework.AsymmetricKeypair, error) {
	return nil, errors.New("not implement")
}

func (e ElgamalCryptoFunction) GetAlgorithm() framework.CryptoAlgorithm {
	return ELGAMAL_ALGORITHM
}

func (e ElgamalCryptoFunction) Encrypt(pubKey *framework.PubKey, data []byte) (*framework.AsymmetricCiphertext, error) {
	return nil, errors.New("not implement")
}

func (e ElgamalCryptoFunction) Decrypt(privKey *framework.PrivKey, ciphertext *framework.AsymmetricCiphertext) ([]byte, error) {
	return nil, errors.New("not implement")
}

func (e ElgamalCryptoFunction) RetrievePubKey(privKey *framework.PrivKey) (*framework.PubKey, error) {
	return nil, errors.New("not implement")
}

func (e ElgamalCryptoFunction) SupportPrivKey(privKeyBytes []byte) bool {
	return false
}

func (e ElgamalCryptoFunction) ParsePrivKey(privKeyBytes []byte) (*framework.PrivKey, error) {
	return nil, errors.New("not implement")
}

func (e ElgamalCryptoFunction) SupportPubKey(pubKeyBytes []byte) bool {
	return false
}

func (e ElgamalCryptoFunction) ParsePubKey(pubKeyBytes []byte) (*framework.PubKey, error) {
	return nil, errors.New("not implement")
}

func (e ElgamalCryptoFunction) SupportCiphertext(ciphertextBytes []byte) bool {
	return false
}

func (e ElgamalCryptoFunction) ParseCiphertext(ciphertextBytes []byte) (*framework.AsymmetricCiphertext, error) {
	return nil, errors.New("not implement")
}
