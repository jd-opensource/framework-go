package classic

import (
	ed255192 "crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/blockchain-jd-com/framework-go/crypto/ca"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/utils/base64"
	"github.com/blockchain-jd-com/framework-go/utils/ed25519"
	"strings"
)

/**
 * @Author: imuge
 * @Date: 2020/4/28 3:02 下午
 */

var (
	ED25519_PUBKEY_SIZE            = 32
	ED25519_PRIVKEY_SIZE           = 32
	ED25519_SIGNATUREDIGEST_SIZE   = 64
	ED25519_PUBKEY_LENGTH          = framework.ALGORYTHM_CODE_SIZE + framework.KEY_TYPE_BYTES + ED25519_PUBKEY_SIZE
	ED25519_PRIVKEY_LENGTH         = framework.ALGORYTHM_CODE_SIZE + framework.KEY_TYPE_BYTES + ED25519_PRIVKEY_SIZE
	ED25519_SIGNATUREDIGEST_LENGTH = framework.ALGORYTHM_CODE_SIZE + ED25519_SIGNATUREDIGEST_SIZE
)

var _ framework.SignatureFunction = (*ED25519SignatureFunction)(nil)

type ED25519SignatureFunction struct {
}

func (E ED25519SignatureFunction) RetrievePrivKey(privateKey string) (*framework.PrivKey, error) {
	privateKey = strings.ReplaceAll(privateKey, "-----BEGIN PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "-----END PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "\n", "")
	privateKey = strings.Trim(privateKey, "")
	encoded := base64.MustDecode(privateKey)
	key, err := x509.ParsePKCS8PrivateKey(encoded)
	if err != nil {
		key, err = x509.ParsePKCS1PrivateKey(encoded)
	}
	ed25519Key, ok := key.(ed255192.PrivateKey)
	if ok {
		return framework.NewPrivKey(E.GetAlgorithm(), ed25519Key[:32]), nil
	}

	return nil, errors.New("not ed25519 private key")
}

func (E ED25519SignatureFunction) RetrieveEncrypedPrivKey(privateKey string, pwd []byte) (*framework.PrivKey, error) {
	privateKey = strings.ReplaceAll(privateKey, "-----BEGIN PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "-----END PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "\n", "")
	privateKey = strings.Trim(privateKey, "")
	encoded := base64.MustDecode(privateKey)
	block, rest := pem.Decode(encoded)
	if len(rest) > 0 {
		return nil, errors.New("extra data included in key")
	}
	encoded, err := x509.DecryptPEMBlock(block, pwd)
	if err != nil {
		return nil, errors.New("decrypt error")
	}
	key, err := x509.ParsePKCS8PrivateKey(encoded)
	if err != nil {
		key, err = x509.ParsePKCS1PrivateKey(encoded)
	}
	ed25519Key, ok := key.(ed255192.PrivateKey)
	if ok {
		return framework.NewPrivKey(E.GetAlgorithm(), ed25519Key[:32]), nil
	}

	return nil, errors.New("not ed25519 private key")
}

func (E ED25519SignatureFunction) RetrievePubKeyFromCA(cert *ca.Certificate) (*framework.PubKey, error) {
	return framework.NewPubKey(E.GetAlgorithm(), cert.ClassicCert.PublicKey.(ed255192.PublicKey)), nil
}

func (E ED25519SignatureFunction) GenerateKeypair() (*framework.AsymmetricKeypair, error) {
	pub, seed := ed25519.GenerateKeyPair()

	return framework.NewAsymmetricKeypair(framework.NewPubKey(E.GetAlgorithm(), pub), framework.NewPrivKey(E.GetAlgorithm(), seed))
}

func (E ED25519SignatureFunction) GenerateKeypairWithSeed(seed []byte) (keypair *framework.AsymmetricKeypair, err error) {
	if len(seed) < 32 {
		return nil, errors.New("seed length must gte 32")
	}
	pub, seed := ed25519.GenerateKeyPairWithSeed(seed)
	return framework.NewAsymmetricKeypair(framework.NewPubKey(E.GetAlgorithm(), pub), framework.NewPrivKey(E.GetAlgorithm(), seed))
}

func (E ED25519SignatureFunction) GetAlgorithm() framework.CryptoAlgorithm {
	return ED25519_ALGORITHM
}

func (E ED25519SignatureFunction) Sign(privKey *framework.PrivKey, data []byte) (*framework.SignatureDigest, error) {
	rawPrivKeyBytes, err := privKey.GetRawKeyBytes()
	if err != nil {
		return nil, err
	}

	// 验证原始私钥长度为256比特，即32字节
	if len(rawPrivKeyBytes) != ED25519_PRIVKEY_SIZE {
		return nil, errors.New("This key has wrong format!")
	}

	// 验证密钥数据的算法标识对应ED25519签名算法
	if privKey.GetAlgorithm() != E.GetAlgorithm().Code {
		return nil, errors.New("This key is not ED25519 private key!")
	}

	// 调用ED25519签名算法计算签名结果
	return framework.NewSignatureDigest(E.GetAlgorithm(), ed25519.Sign(rawPrivKeyBytes, data)), nil
}

func (E ED25519SignatureFunction) Verify(pubKey *framework.PubKey, data []byte, digest *framework.SignatureDigest) bool {
	return ed25519.Verify(pubKey.GetRawKeyBytes(), data, digest.GetRawDigest())
}

func (E ED25519SignatureFunction) RetrievePubKey(privKey *framework.PrivKey) (*framework.PubKey, error) {
	bytes, err := privKey.GetRawKeyBytes()
	if err != nil {
		return nil, err
	}
	return framework.NewPubKey(E.GetAlgorithm(), ed25519.RetrievePubKey(bytes)), nil
}

func (E ED25519SignatureFunction) SupportPrivKey(privKeyBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+密钥类型长度+密钥长度，密钥数据的算法标识对应ED25519签名算法，并且密钥类型是私钥
	return len(privKeyBytes) == ED25519_PRIVKEY_LENGTH && E.GetAlgorithm().Match(privKeyBytes, 0) && privKeyBytes[framework.ALGORYTHM_CODE_SIZE] == framework.PRIVATE.Code
}

func (E ED25519SignatureFunction) ParsePrivKey(privKeyBytes []byte) (*framework.PrivKey, error) {
	if E.SupportPrivKey(privKeyBytes) {
		return framework.ParsePrivKey(privKeyBytes)
	} else {
		return nil, errors.New("invalid privKeyBytes!")
	}
}

func (E ED25519SignatureFunction) SupportPubKey(pubKeyBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+密钥类型长度+密钥长度，密钥数据的算法标识对应ED25519签名算法，并且密钥类型是公钥
	return len(pubKeyBytes) == ED25519_PUBKEY_LENGTH && E.GetAlgorithm().Match(pubKeyBytes, 0) && pubKeyBytes[framework.ALGORYTHM_CODE_SIZE] == framework.PUBLIC.Code
}

func (E ED25519SignatureFunction) ParsePubKey(pubKeyBytes []byte) (*framework.PubKey, error) {
	if E.SupportPubKey(pubKeyBytes) {
		return framework.ParsePubKey(pubKeyBytes)
	} else {
		return nil, errors.New("invalid pubKeyBytes!")
	}
}

func (E ED25519SignatureFunction) SupportDigest(digestBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+摘要长度，字节数组的算法标识对应ED25519算法
	return len(digestBytes) == ED25519_SIGNATUREDIGEST_LENGTH && E.GetAlgorithm().Match(digestBytes, 0)
}

func (E ED25519SignatureFunction) ParseDigest(digestBytes []byte) (*framework.SignatureDigest, error) {
	if E.SupportDigest(digestBytes) {
		return framework.ParseSignatureDigest(digestBytes)
	} else {
		return nil, errors.New("invalid digestBytes!")
	}
}
