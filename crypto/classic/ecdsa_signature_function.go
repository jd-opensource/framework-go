package classic

import (
	ecdsa2 "crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/blockchain-jd-com/framework-go/crypto/ca"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/utils/base64"
	"github.com/blockchain-jd-com/framework-go/utils/ecdsa"
	"strings"
)

/**
 * @Author: imuge
 * @Date: 2020/4/28 3:01 下午
 */

var (
	ECDSA_PUBKEY_SIZE            = 65
	ECDSA_PRIVKEY_SIZE           = 32
	ECDSA_SIGNATUREDIGEST_SIZE   = 64
	ECDSA_PUBKEY_LENGTH          = framework.ALGORYTHM_CODE_SIZE + framework.KEY_TYPE_BYTES + ECDSA_PUBKEY_SIZE
	ECDSA_PRIVKEY_LENGTH         = framework.ALGORYTHM_CODE_SIZE + framework.KEY_TYPE_BYTES + ECDSA_PRIVKEY_SIZE
	ECDSA_SIGNATUREDIGEST_LENGTH = framework.ALGORYTHM_CODE_SIZE + ECDSA_SIGNATUREDIGEST_SIZE
)

var _ framework.SignatureFunction = (*ECDSASignatureFunction)(nil)

type ECDSASignatureFunction struct {
}

func (E ECDSASignatureFunction) RetrievePrivKey(privateKey string) (*framework.PrivKey, error) {
	index := strings.Index(privateKey, "END EC PARAMETERS-----")
	if index > 0 && strings.Contains(privateKey[:index], "BggqhkjOPQMBBw==") {
		privateKey = privateKey[index+22:]
	}
	privateKey = strings.ReplaceAll(privateKey, "-----BEGIN EC PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "-----END EC PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "-----BEGIN PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "-----END PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "\n", "")
	privateKey = strings.Trim(privateKey, "")
	encoded := base64.MustDecode(privateKey)
	key, err := x509.ParsePKCS8PrivateKey(encoded)
	if err != nil {
		key, err = x509.ParsePKCS1PrivateKey(encoded)
		if err != nil {
			key, err = x509.ParseECPrivateKey(encoded)
			if err != nil {
				return nil, errors.New("not ecdsa private key")
			}
		}
	}
	ecdsaKey, ok := key.(*ecdsa2.PrivateKey)
	if ok {
		return framework.NewPrivKey(E.GetAlgorithm(), ecdsaKey.D.Bytes()), nil
	}

	return nil, errors.New("not ecdsa private key")
}

func (E ECDSASignatureFunction) RetrieveEncrypedPrivKey(privateKey string, pwd []byte) (*framework.PrivKey, error) {
	index := strings.Index(privateKey, "END EC PARAMETERS-----")
	if index > 0 && strings.Contains(privateKey[:index], "BgUrgQQACg==") {
		privateKey = privateKey[index+22:]
	}
	privateKey = strings.ReplaceAll(privateKey, "-----BEGIN EC PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "-----END EC PRIVATE KEY-----", "")
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
		if err != nil {
			return nil, errors.New("not ecdsa private key")
		}
	}
	ecdsaKey, ok := key.(*ecdsa2.PrivateKey)
	if ok {
		return framework.NewPrivKey(E.GetAlgorithm(), ecdsaKey.D.Bytes()), nil
	}

	return nil, errors.New("not ecdsa private key")
}

func (E ECDSASignatureFunction) RetrievePubKeyFromCA(cert *ca.Certificate) (*framework.PubKey, error) {
	key, ok := cert.ClassicCert.PublicKey.(*ecdsa2.PublicKey)
	if !ok {
		return nil, errors.New("not ecdsa public key")
	}
	x := key.X.Bytes()
	y := key.Y.Bytes()
	return framework.NewPubKey(E.GetAlgorithm(), append(append([]byte{0x04}, x...), y...)), nil
}

func (E ECDSASignatureFunction) GenerateKeypair() (*framework.AsymmetricKeypair, error) {
	priv := ecdsa.GenerateKeyPair()
	return framework.NewAsymmetricKeypair(framework.NewPubKey(E.GetAlgorithm(), ecdsa.PubKeyToBytes(&priv.PublicKey)), framework.NewPrivKey(E.GetAlgorithm(), ecdsa.PrivKeyToBytes(priv)))
}

func (E ECDSASignatureFunction) GenerateKeypairWithSeed(seed []byte) (keypair *framework.AsymmetricKeypair, err error) {
	return nil, errors.New("not support yet")
}

func (E ECDSASignatureFunction) GetAlgorithm() framework.CryptoAlgorithm {
	return ECDSA_ALGORITHM
}

func (E ECDSASignatureFunction) Sign(privKey *framework.PrivKey, data []byte) (*framework.SignatureDigest, error) {

	rawPrivKeyBytes, err := privKey.GetRawKeyBytes()
	if err != nil {
		return nil, err
	}

	// 验证原始私钥长度为256比特，即32字节
	if len(rawPrivKeyBytes) != ECDSA_PRIVKEY_SIZE {
		return nil, errors.New("This key has wrong format!")
	}

	// 验证密钥数据的算法标识对应ECDSA签名算法
	if privKey.GetAlgorithm() != E.GetAlgorithm().Code {
		return nil, errors.New("This key is not ECDSA private key!")
	}

	// 调用ECDSA签名算法计算签名结果
	return framework.NewSignatureDigest(E.GetAlgorithm(), ecdsa.Sign(ecdsa.BytesToPrivKey(rawPrivKeyBytes), data)), nil
}

func (E ECDSASignatureFunction) Verify(pubKey *framework.PubKey, data []byte, digest *framework.SignatureDigest) bool {
	rawPubKeyBytes := pubKey.GetRawKeyBytes()
	rawDigestBytes := digest.GetRawDigest()

	// 验证原始公钥长度为256比特，即32字节
	if len(rawPubKeyBytes) != ECDSA_PUBKEY_SIZE {
		return false
	}

	// 验证密钥数据的算法标识对应ECDSA签名算法
	if pubKey.GetAlgorithm() != E.GetAlgorithm().Code {
		return false
	}

	// 验证签名数据的算法标识对应ECDSA签名算法，并且原始摘要长度为64字节
	if digest.GetAlgorithm() != E.GetAlgorithm().Code || len(rawDigestBytes) != ECDSA_SIGNATUREDIGEST_SIZE {
		return false
	}

	// 调用ECDSA验签算法验证签名结果
	return ecdsa.Verify(ecdsa.BytesToPubKey(rawPubKeyBytes), data, rawDigestBytes)
}

func (E ECDSASignatureFunction) RetrievePubKey(privKey *framework.PrivKey) (*framework.PubKey, error) {
	bytes, err := privKey.GetRawKeyBytes()
	if err != nil {
		return nil, err
	}
	return framework.NewPubKey(E.GetAlgorithm(), ecdsa.PubKeyToBytes(&ecdsa.BytesToPrivKey(bytes).PublicKey)), nil
}

func (E ECDSASignatureFunction) SupportPrivKey(privKeyBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+密钥类型长度+密钥长度，密钥数据的算法标识对应ECDSA签名算法，并且密钥类型是私钥
	return len(privKeyBytes) == ECDSA_PRIVKEY_LENGTH && E.GetAlgorithm().Match(privKeyBytes, 0) && privKeyBytes[framework.ALGORYTHM_CODE_SIZE] == framework.PRIVATE.Code
}

func (E ECDSASignatureFunction) ParsePrivKey(privKeyBytes []byte) (*framework.PrivKey, error) {
	if E.SupportPrivKey(privKeyBytes) {
		return framework.ParsePrivKey(privKeyBytes)
	} else {
		return nil, errors.New("invalid privKeyBytes!")
	}
}

func (E ECDSASignatureFunction) SupportPubKey(pubKeyBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+密钥类型长度+密钥长度，密钥数据的算法标识对应ECDSA签名算法，并且密钥类型是公钥
	return len(pubKeyBytes) == ECDSA_PUBKEY_LENGTH && E.GetAlgorithm().Match(pubKeyBytes, 0) && pubKeyBytes[framework.ALGORYTHM_CODE_SIZE] == framework.PUBLIC.Code
}

func (E ECDSASignatureFunction) ParsePubKey(pubKeyBytes []byte) (*framework.PubKey, error) {
	if E.SupportPubKey(pubKeyBytes) {
		return framework.ParsePubKey(pubKeyBytes)
	} else {
		return nil, errors.New("invalid pubKeyBytes!")
	}
}

func (E ECDSASignatureFunction) SupportDigest(digestBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+摘要长度，字节数组的算法标识对应ECDSA算法
	return len(digestBytes) == ECDSA_SIGNATUREDIGEST_LENGTH && E.GetAlgorithm().Match(digestBytes, 0)
}

func (E ECDSASignatureFunction) ParseDigest(digestBytes []byte) (*framework.SignatureDigest, error) {
	if E.SupportDigest(digestBytes) {
		return framework.ParseSignatureDigest(digestBytes)
	} else {
		return nil, errors.New("invalid digestBytes!")
	}
}
