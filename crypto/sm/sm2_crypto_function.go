package sm

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"errors"
	"github.com/blockchain-jd-com/framework-go/crypto/ca"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/utils/base64"
	"github.com/blockchain-jd-com/framework-go/utils/sm2"
	x5092 "github.com/tjfoc/gmsm/x509"
	"strings"
)

/**
 * @Author: imuge
 * @Date: 2020/4/28 3:00 下午
 */

var (
	SM2_ECPOINT_SIZE           = 65
	SM2_PRIVKEY_SIZE           = 32
	SM2_SIGNATUREDIGEST_SIZE   = 64
	SM2_HASHDIGEST_SIZE        = 32
	SM2_PUBKEY_LENGTH          = framework.ALGORYTHM_CODE_SIZE + framework.KEY_TYPE_BYTES + SM2_ECPOINT_SIZE
	SM2_PRIVKEY_LENGTH         = framework.ALGORYTHM_CODE_SIZE + framework.KEY_TYPE_BYTES + SM2_PRIVKEY_SIZE
	SM2_SIGNATUREDIGEST_LENGTH = framework.ALGORYTHM_CODE_SIZE + SM2_SIGNATUREDIGEST_SIZE
)

var _ framework.AsymmetricEncryptionFunction = (*SM2CryptoFunction)(nil)
var _ framework.SignatureFunction = (*SM2CryptoFunction)(nil)

type SM2CryptoFunction struct {
}

func (S SM2CryptoFunction) RetrievePrivKey(privateKey string) (*framework.PrivKey, error) {
	index := strings.Index(privateKey, "END EC PARAMETERS-----")
	if index > 0 && strings.Contains(privateKey[:index], "BggqgRzPVQGCLQ==") {
		privateKey = privateKey[index+22:]
	}
	privateKey = strings.ReplaceAll(privateKey, "-----BEGIN EC PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "-----END EC PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "-----BEGIN PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "-----END PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "\n", "")
	privateKey = strings.Trim(privateKey, "")
	encoded := base64.MustDecode(privateKey)
	sm2Key, err := x5092.ParsePKCS8UnecryptedPrivateKey(encoded)
	if err != nil {
		sm2Key, err = x5092.ParseSm2PrivateKey(encoded)
		if err != nil {
			return nil, err
		}
	}
	return framework.NewPrivKey(S.GetAlgorithm(), sm2Key.D.Bytes()), nil
}

func (S SM2CryptoFunction) RetrieveEncrypedPrivKey(privateKey string, pwd []byte) (*framework.PrivKey, error) {
	index := strings.Index(privateKey, "END EC PARAMETERS-----")
	if index > 0 && strings.Contains(privateKey[:index], "BggqgRzPVQGCLQ==") {
		privateKey = privateKey[index+22:]
	}
	privateKey = strings.ReplaceAll(privateKey, "-----BEGIN EC PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "-----END EC PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "-----BEGIN PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "-----END PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "\n", "")
	privateKey = strings.Trim(privateKey, "")
	encoded := base64.MustDecode(privateKey)
	sm2Key, err := x5092.ParsePKCS8EcryptedPrivateKey(encoded, pwd)
	if err != nil {
		sm2Key, err = x5092.ParseSm2PrivateKey(encoded)
		if err != nil {
			return nil, err
		}
	}
	return framework.NewPrivKey(S.GetAlgorithm(), sm2Key.D.Bytes()), nil
}

func (S SM2CryptoFunction) RetrievePubKeyFromCA(cert *ca.Certificate) (*framework.PubKey, error) {
	key, ok := cert.SMCert.PublicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("not sm2 public key")
	}
	return framework.NewPubKey(S.GetAlgorithm(), elliptic.Marshal(key, key.X, key.Y)), nil
}

func (S SM2CryptoFunction) Encrypt(pubKey *framework.PubKey, data []byte) (*framework.AsymmetricCiphertext, error) {
	rawPubKeyBytes := pubKey.GetRawKeyBytes()

	// 验证原始公钥长度为65字节
	if len(rawPubKeyBytes) != SM2_ECPOINT_SIZE {
		return nil, errors.New("This key has wrong format!")
	}

	// 验证密钥数据的算法标识对应SM2算法
	if pubKey.GetAlgorithm() != S.GetAlgorithm().Code {
		return nil, errors.New("This key is not sm2 public key!")
	}

	key, err := sm2.BytesToPubKey(rawPubKeyBytes)
	if err != nil {
		return nil, err
	}
	encrypt, err := sm2.Encrypt(key, data)
	if err != nil {
		return nil, err
	}
	// 调用SM2加密算法计算密文
	return framework.NewAsymmetricCiphertext(S.GetAlgorithm(), encrypt), nil
}

func (S SM2CryptoFunction) Decrypt(privKey *framework.PrivKey, ciphertext *framework.AsymmetricCiphertext) ([]byte, error) {
	rawPrivKeyBytes, err := privKey.GetRawKeyBytes()
	if err != nil {
		return nil, err
	}
	rawCiphertextBytes := ciphertext.GetRawCiphertext()

	// 验证原始私钥长度为32字节
	if len(rawPrivKeyBytes) != SM2_PRIVKEY_SIZE {
		return nil, errors.New("This key has wrong format!")
	}

	// 验证密钥数据的算法标识对应SM2算法
	if privKey.GetAlgorithm() != S.GetAlgorithm().Code {
		return nil, errors.New("This key is not SM2 private key!")
	}

	// 验证密文数据的算法标识对应SM2算法，并且密文符合长度要求
	if ciphertext.GetAlgorithm() != S.GetAlgorithm().Code || len(rawCiphertextBytes) < SM2_ECPOINT_SIZE+SM2_HASHDIGEST_SIZE {
		return nil, errors.New("This is not SM2 ciphertext!")
	}

	key, err := sm2.BytesToPrivKey(rawPrivKeyBytes)
	if err != nil {
		return nil, err
	}
	// 调用SM2解密算法得到明文结果
	return sm2.Decrypt(key, rawCiphertextBytes)
}

func (S SM2CryptoFunction) SupportCiphertext(ciphertextBytes []byte) bool {
	// 验证输入字节数组长度>=算法标识长度+椭圆曲线点长度+哈希长度，字节数组的算法标识对应SM2算法
	return len(ciphertextBytes) >= framework.ALGORYTHM_CODE_SIZE+SM2_ECPOINT_SIZE+SM2_HASHDIGEST_SIZE && S.GetAlgorithm().Match(ciphertextBytes, 0)
}

func (S SM2CryptoFunction) ParseCiphertext(ciphertextBytes []byte) (*framework.AsymmetricCiphertext, error) {
	if S.SupportCiphertext(ciphertextBytes) {
		return framework.ParseAsymmetricCiphertext(ciphertextBytes)
	} else {
		return nil, errors.New("ciphertextBytes are invalid!")
	}
}

func (S SM2CryptoFunction) Sign(privKey *framework.PrivKey, data []byte) (*framework.SignatureDigest, error) {
	rawPrivKeyBytes, err := privKey.GetRawKeyBytes()
	if err != nil {
		return nil, err
	}

	// 验证原始私钥长度为256比特，即32字节
	if len(rawPrivKeyBytes) != SM2_PRIVKEY_SIZE {
		return nil, errors.New("This key has wrong format!")
	}

	// 验证密钥数据的算法标识对应SM2签名算法
	if privKey.GetAlgorithm() != S.GetAlgorithm().Code {
		return nil, errors.New("This key is not SM2 private key!")
	}

	key, err := sm2.BytesToPrivKey(rawPrivKeyBytes)
	if err != nil {
		return nil, err
	}
	sign, err := sm2.Sign(key, data)
	if err != nil {
		return nil, err
	}
	// 调用SM2签名算法计算签名结果
	return framework.NewSignatureDigest(S.GetAlgorithm(), sign), nil
}

func (S SM2CryptoFunction) Verify(pubKey *framework.PubKey, data []byte, digest *framework.SignatureDigest) bool {
	rawPubKeyBytes := pubKey.GetRawKeyBytes()
	rawDigestBytes := digest.GetRawDigest()

	// 验证原始公钥长度为520比特，即65字节
	if len(rawPubKeyBytes) != SM2_ECPOINT_SIZE {
		return false
	}

	// 验证密钥数据的算法标识对应SM2签名算法
	if pubKey.GetAlgorithm() != S.GetAlgorithm().Code {
		return false
	}

	// 验证签名数据的算法标识对应SM2签名算法，并且原始签名长度为64字节
	if digest.GetAlgorithm() != S.GetAlgorithm().Code || len(rawDigestBytes) != SM2_SIGNATUREDIGEST_SIZE {
		return false
	}
	key, err := sm2.BytesToPubKey(rawPubKeyBytes)
	if err != nil {
		return false
	}
	// 调用SM2验签算法验证签名结果
	return sm2.Verify(key, data, rawDigestBytes)
}

func (S SM2CryptoFunction) RetrievePubKey(privKey *framework.PrivKey) (*framework.PubKey, error) {
	bytes, err := privKey.GetRawKeyBytes()
	if err != nil {
		return nil, err
	}
	key, err := sm2.BytesToPrivKey(bytes)
	if err != nil {
		return nil, err
	}
	return framework.NewPubKey(S.GetAlgorithm(), sm2.PubKeyToBytes(sm2.RetrievePubKey(key))), nil
}

func (S SM2CryptoFunction) SupportPrivKey(privKeyBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+密钥类型长度+密钥长度，密钥数据的算法标识对应SM2算法，并且密钥类型是私钥
	return len(privKeyBytes) == SM2_PRIVKEY_LENGTH && S.GetAlgorithm().Match(privKeyBytes, 0) && privKeyBytes[framework.ALGORYTHM_CODE_SIZE] == framework.PRIVATE.Code
}

func (S SM2CryptoFunction) ParsePrivKey(privKeyBytes []byte) (*framework.PrivKey, error) {
	if S.SupportPrivKey(privKeyBytes) {
		return framework.ParsePrivKey(privKeyBytes)
	} else {
		return nil, errors.New("invalid privKeyBytes!")
	}
}

func (S SM2CryptoFunction) SupportPubKey(pubKeyBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+密钥类型长度+椭圆曲线点长度，密钥数据的算法标识对应SM2算法，并且密钥类型是公钥
	return len(pubKeyBytes) == SM2_PUBKEY_LENGTH && S.GetAlgorithm().Match(pubKeyBytes, 0) && pubKeyBytes[framework.ALGORYTHM_CODE_SIZE] == framework.PUBLIC.Code
}

func (S SM2CryptoFunction) ParsePubKey(pubKeyBytes []byte) (*framework.PubKey, error) {
	if S.SupportPubKey(pubKeyBytes) {
		return framework.ParsePubKey(pubKeyBytes)
	} else {
		return nil, errors.New("invalid pubKeyBytes!")
	}
}

func (S SM2CryptoFunction) SupportDigest(digestBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+签名长度，字节数组的算法标识对应SM2算法
	return len(digestBytes) == SM2_SIGNATUREDIGEST_LENGTH && S.GetAlgorithm().Match(digestBytes, 0)
}

func (S SM2CryptoFunction) ParseDigest(digestBytes []byte) (*framework.SignatureDigest, error) {
	if S.SupportDigest(digestBytes) {
		return framework.ParseSignatureDigest(digestBytes)
	} else {
		return nil, errors.New("invalid digestBytes!")
	}
}

func (S SM2CryptoFunction) GenerateKeypair() (*framework.AsymmetricKeypair, error) {
	priv, pub, err := sm2.GenerateKeyPair()
	if err != nil {
		return nil, err
	}
	return framework.NewAsymmetricKeypair(framework.NewPubKey(S.GetAlgorithm(), sm2.PubKeyToBytes(pub)), framework.NewPrivKey(S.GetAlgorithm(), sm2.PrivKeyToBytes(priv)))
}

func (S SM2CryptoFunction) GenerateKeypairWithSeed(seed []byte) (keypair *framework.AsymmetricKeypair, err error) {
	return nil, errors.New("not support yet")
}

func (S SM2CryptoFunction) GetAlgorithm() framework.CryptoAlgorithm {
	return SM2_ALGORITHM
}
