package classic

import (
	rsa2 "crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/blockchain-jd-com/framework-go/crypto/ca"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/utils/base64"
	bytes2 "github.com/blockchain-jd-com/framework-go/utils/bytes"
	"github.com/blockchain-jd-com/framework-go/utils/rsa"
	"math/big"
	"strings"
)

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:52 下午
 */

var (
	// modulus.length = 256, publicExponent.length = 3
	RSA_PUBKEY_SIZE = 259
	// modulus.length = 256, publicExponent.length = 3, privateExponent.length = 256, p.length = 128, q.length =128,
	// dP.length = 128, dQ.length = 128, qInv.length = 128
	RSA_PRIVKEY_SIZE = 1155

	RSA_SIGNATUREDIGEST_SIZE = 256
	RSA_CIPHERTEXTBLOCK_SIZE = 256

	RSA_PUBKEY_LENGTH          = framework.ALGORYTHM_CODE_SIZE + framework.KEY_TYPE_BYTES + RSA_PUBKEY_SIZE
	RSA_PRIVKEY_LENGTH         = framework.ALGORYTHM_CODE_SIZE + framework.KEY_TYPE_BYTES + RSA_PRIVKEY_SIZE
	RSA_SIGNATUREDIGEST_LENGTH = framework.ALGORYTHM_CODE_SIZE + RSA_SIGNATUREDIGEST_SIZE
)

var _ framework.AsymmetricEncryptionFunction = (*RSACryptoFunction)(nil)
var _ framework.SignatureFunction = (*RSACryptoFunction)(nil)

type RSACryptoFunction struct {
}

func (R RSACryptoFunction) RetrievePrivKey(privateKey string) (framework.PrivKey, error) {
	privateKey = strings.ReplaceAll(privateKey, "-----BEGIN PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "-----BEGIN RSA PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "-----END PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "-----END RSA PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "\n", "")
	privateKey = strings.Trim(privateKey, "")
	encoded := base64.MustDecode(privateKey)
	key, err := x509.ParsePKCS8PrivateKey(encoded)
	if err != nil {
		key, err = x509.ParsePKCS1PrivateKey(encoded)
	}
	rsaKey, ok := key.(*rsa2.PrivateKey)
	if ok {
		return framework.NewPrivKey(R.GetAlgorithm(),
				bytes2.Concat(rsaKey.PublicKey.N.Bytes(),
					big.NewInt(int64(rsaKey.PublicKey.E)).Bytes(),
					rsaKey.D.Bytes(),
					rsaKey.Primes[0].Bytes(),
					rsaKey.Primes[1].Bytes(),
					rsaKey.Precomputed.Dp.Bytes(),
					rsaKey.Precomputed.Dq.Bytes(),
					rsaKey.Precomputed.Qinv.Bytes())),
			nil
	}

	return framework.PrivKey{}, errors.New("not rsa private key")
}

func (R RSACryptoFunction) RetrieveEncrypedPrivKey(privateKey string, pwd []byte) (framework.PrivKey, error) {
	privateKey = strings.ReplaceAll(privateKey, "-----BEGIN PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "-----BEGIN RSA PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "-----END PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "-----END RSA PRIVATE KEY-----", "")
	privateKey = strings.ReplaceAll(privateKey, "\n", "")
	privateKey = strings.Trim(privateKey, "")
	encoded := base64.MustDecode(privateKey)
	block, rest := pem.Decode(encoded)
	if len(rest) > 0 {
		return framework.PrivKey{}, errors.New("extra data included in key")
	}
	encoded, err := x509.DecryptPEMBlock(block, pwd)
	if err != nil {
		return framework.PrivKey{}, errors.New("decrypt error")
	}
	key, err := x509.ParsePKCS8PrivateKey(encoded)
	if err != nil {
		key, err = x509.ParsePKCS1PrivateKey(encoded)
	}
	rsaKey, ok := key.(*rsa2.PrivateKey)
	if ok {
		return framework.NewPrivKey(R.GetAlgorithm(),
				bytes2.Concat(rsaKey.PublicKey.N.Bytes(),
					big.NewInt(int64(rsaKey.PublicKey.E)).Bytes(),
					rsaKey.D.Bytes(),
					rsaKey.Primes[0].Bytes(),
					rsaKey.Primes[1].Bytes(),
					rsaKey.Precomputed.Dp.Bytes(),
					rsaKey.Precomputed.Dq.Bytes(),
					rsaKey.Precomputed.Qinv.Bytes())),
			nil
	}

	return framework.PrivKey{}, errors.New("not rsa private key")
}

func (R RSACryptoFunction) RetrievePubKeyFromCA(cert *ca.Certificate) framework.PubKey {
	key := cert.ClassicCert.PublicKey.(*rsa2.PublicKey)
	return framework.NewPubKey(R.GetAlgorithm(), bytes2.Concat(key.N.Bytes(), big.NewInt(int64(key.E)).Bytes()))
}

func (R RSACryptoFunction) Encrypt(pubKey framework.PubKey, data []byte) framework.AsymmetricCiphertext {
	rawPubBytes := pubKey.GetRawKeyBytes()

	// 验证原始公钥长度为257字节
	if len(rawPubBytes) != RSA_PUBKEY_SIZE {
		panic("This key has wrong format!")
	}

	// 验证密钥数据的算法标识对应RSA算法
	if pubKey.GetAlgorithm() != R.GetAlgorithm().Code {
		panic("The is not RSA public key!")
	}

	return framework.NewAsymmetricCiphertext(R.GetAlgorithm(), rsa.Encrypt(rsa.BytesToPubKey(rawPubBytes), data))
}

func (R RSACryptoFunction) Decrypt(privKey framework.PrivKey, ciphertext framework.AsymmetricCiphertext) []byte {
	rawPrivBytes := privKey.GetRawKeyBytes()
	rawCiphertextBytes := ciphertext.GetRawCiphertext()

	// 验证原始私钥长度为1153字节
	if len(rawPrivBytes) != RSA_PRIVKEY_SIZE {
		panic("This key has wrong format!")
	}

	// 验证密钥数据的算法标识对应RSA算法
	if privKey.GetAlgorithm() != R.GetAlgorithm().Code {
		panic("This key is not RSA private key!")
	}

	// 验证密文数据的算法标识对应RSA算法，并且密文是分组长度的整数倍
	if ciphertext.GetAlgorithm() != R.GetAlgorithm().Code || len(rawCiphertextBytes)%RSA_CIPHERTEXTBLOCK_SIZE != 0 {
		panic("This is not RSA ciphertext!")
	}

	// 调用RSA解密算法得到明文结果
	return rsa.Decrypt(rsa.BytesToPrivKey(rawPrivBytes), rawCiphertextBytes)
}

func (R RSACryptoFunction) SupportCiphertext(ciphertextBytes []byte) bool {
	// 验证输入字节数组长度=密文分组的整数倍，字节数组的算法标识对应RSA算法
	return (len(ciphertextBytes)%RSA_CIPHERTEXTBLOCK_SIZE == framework.ALGORYTHM_CODE_SIZE) && R.GetAlgorithm().Match(ciphertextBytes, 0)
}

func (R RSACryptoFunction) ParseCiphertext(ciphertextBytes []byte) framework.AsymmetricCiphertext {
	return framework.ParseAsymmetricCiphertext(ciphertextBytes)
}

func (R RSACryptoFunction) Sign(privKey framework.PrivKey, data []byte) framework.SignatureDigest {
	rawPrivKeyBytes := privKey.GetRawKeyBytes()

	// 验证原始私钥长度为1153字节
	if len(rawPrivKeyBytes) != RSA_PRIVKEY_SIZE {
		panic("This key has wrong format!")
	}

	// 验证密钥数据的算法标识对应RSA签名算法
	if privKey.GetAlgorithm() != R.GetAlgorithm().Code {
		panic("This key is not RSA private key!")
	}

	// 调用RSA签名算法计算签名结果
	return framework.NewSignatureDigest(R.GetAlgorithm(), rsa.Sign(rsa.BytesToPrivKey(rawPrivKeyBytes), data))
}

func (R RSACryptoFunction) Verify(pubKey framework.PubKey, data []byte, digest framework.SignatureDigest) bool {
	rawPubKeyBytes := pubKey.GetRawKeyBytes()
	rawDigestBytes := digest.GetRawDigest()

	// 验证原始公钥长度为257字节
	if len(rawPubKeyBytes) != RSA_PUBKEY_SIZE {
		panic("This key has wrong format!")
	}

	// 验证密钥数据的算法标识对应RSA签名算法
	if pubKey.GetAlgorithm() != R.GetAlgorithm().Code {
		panic("This key is not RSA public key!")
	}

	// 验证签名数据的算法标识对应RSA签名算法，并且原始签名长度为256字节
	if digest.GetAlgorithm() != R.GetAlgorithm().Code || len(rawDigestBytes) != RSA_SIGNATUREDIGEST_SIZE {
		panic("This is not RSA signature digest!")
	}

	// 调用RSA验签算法验证签名结果
	return rsa.Verify(rsa.BytesToPubKey(rawPubKeyBytes), data, rawDigestBytes)
}

func (R RSACryptoFunction) RetrievePubKey(privKey framework.PrivKey) framework.PubKey {
	priv := rsa.BytesToPrivKey(privKey.GetRawKeyBytes())
	return framework.NewPubKey(R.GetAlgorithm(), rsa.PubKeyToBytes(&priv.PublicKey))
}

func (R RSACryptoFunction) SupportPrivKey(privKeyBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+密钥类型长度+密钥长度，密钥数据的算法标识对应RSA算法，并且密钥类型是私钥
	return len(privKeyBytes) == RSA_PRIVKEY_LENGTH && R.GetAlgorithm().Match(privKeyBytes, 0) && privKeyBytes[framework.ALGORYTHM_CODE_SIZE] == framework.PRIVATE.Code
}

func (R RSACryptoFunction) ParsePrivKey(privKeyBytes []byte) framework.PrivKey {
	return framework.ParsePrivKey(privKeyBytes)
}

func (R RSACryptoFunction) SupportPubKey(pubKeyBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+密钥类型长度+椭圆曲线点长度，密钥数据的算法标识对应RSA算法，并且密钥类型是公钥
	return len(pubKeyBytes) == RSA_PUBKEY_LENGTH && R.GetAlgorithm().Match(pubKeyBytes, 0) && pubKeyBytes[framework.ALGORYTHM_CODE_SIZE] == framework.PUBLIC.Code
}

func (R RSACryptoFunction) ParsePubKey(pubKeyBytes []byte) framework.PubKey {
	return framework.ParsePubKey(pubKeyBytes)
}

func (R RSACryptoFunction) SupportDigest(digestBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+签名长度，字节数组的算法标识对应RSA算法
	return len(digestBytes) == RSA_SIGNATUREDIGEST_LENGTH && RSA_ALGORITHM.Match(digestBytes, 0)
}

func (R RSACryptoFunction) ParseDigest(digestBytes []byte) framework.SignatureDigest {
	return framework.ParseSignatureDigest(digestBytes)
}

func (R RSACryptoFunction) GenerateKeypair() framework.AsymmetricKeypair {
	priv := rsa.GenerateKeyPair()
	return framework.NewAsymmetricKeypair(framework.NewPubKey(R.GetAlgorithm(), rsa.PubKeyToBytes(&priv.PublicKey)), framework.NewPrivKey(R.GetAlgorithm(), rsa.PrivKeyToBytes(priv)))
}

func (R RSACryptoFunction) GenerateKeypairWithSeed(seed []byte) (keypair framework.AsymmetricKeypair, err error) {
	panic("not support yet")
}

func (R RSACryptoFunction) GetAlgorithm() framework.CryptoAlgorithm {
	return RSA_ALGORITHM
}
