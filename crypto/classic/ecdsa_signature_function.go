package classic

import (
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/utils/ecdsa"
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

func (E ECDSASignatureFunction) GenerateKeypair() framework.AsymmetricKeypair {
	priv := ecdsa.GenerateKeyPair()
	return framework.NewAsymmetricKeypair(framework.NewPubKey(E.GetAlgorithm(), ecdsa.PubKeyToBytes(&priv.PublicKey)), framework.NewPrivKey(E.GetAlgorithm(), ecdsa.PrivKeyToBytes(priv)))
}

func (E ECDSASignatureFunction) GetAlgorithm() framework.CryptoAlgorithm {
	return ECDSA_ALGORITHM
}

func (E ECDSASignatureFunction) Sign(privKey framework.PrivKey, data []byte) framework.SignatureDigest {

	rawPrivKeyBytes := privKey.GetRawKeyBytes()

	// 验证原始私钥长度为256比特，即32字节
	if len(rawPrivKeyBytes) != ECDSA_PRIVKEY_SIZE {
		panic("This key has wrong format!")
	}

	// 验证密钥数据的算法标识对应ECDSA签名算法
	if privKey.GetAlgorithm() != E.GetAlgorithm().Code {
		panic("This key is not ECDSA private key!")
	}

	// 调用ECDSA签名算法计算签名结果
	return framework.NewSignatureDigest(E.GetAlgorithm(), ecdsa.Sign(ecdsa.BytesToPrivKey(rawPrivKeyBytes), data))
}

func (E ECDSASignatureFunction) Verify(pubKey framework.PubKey, data []byte, digest framework.SignatureDigest) bool {
	rawPubKeyBytes := pubKey.GetRawKeyBytes()
	rawDigestBytes := digest.GetRawDigest()

	// 验证原始公钥长度为256比特，即32字节
	if len(rawPubKeyBytes) != ECDSA_PUBKEY_SIZE {
		panic("This key has wrong format!")
	}

	// 验证密钥数据的算法标识对应ECDSA签名算法
	if pubKey.GetAlgorithm() != E.GetAlgorithm().Code {
		panic("This key is not ECDSA public key!")
	}

	// 验证签名数据的算法标识对应ECDSA签名算法，并且原始摘要长度为64字节
	if digest.GetAlgorithm() != E.GetAlgorithm().Code || len(rawDigestBytes) != ECDSA_SIGNATUREDIGEST_SIZE {
		panic("This is not ECDSA signature digest!")
	}

	// 调用ECDSA验签算法验证签名结果
	return ecdsa.Verify(ecdsa.BytesToPubKey(rawPubKeyBytes), data, rawDigestBytes)
}

func (E ECDSASignatureFunction) RetrievePubKey(privKey framework.PrivKey) framework.PubKey {
	return framework.NewPubKey(E.GetAlgorithm(), ecdsa.PubKeyToBytes(&ecdsa.BytesToPrivKey(privKey.GetRawKeyBytes()).PublicKey))
}

func (E ECDSASignatureFunction) SupportPrivKey(privKeyBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+密钥类型长度+密钥长度，密钥数据的算法标识对应ECDSA签名算法，并且密钥类型是私钥
	return len(privKeyBytes) == ECDSA_PRIVKEY_LENGTH && E.GetAlgorithm().Match(privKeyBytes, 0) && privKeyBytes[framework.ALGORYTHM_CODE_SIZE] == framework.PRIVATE.Code
}

func (E ECDSASignatureFunction) ParsePrivKey(privKeyBytes []byte) framework.PrivKey {
	if E.SupportPrivKey(privKeyBytes) {
		return framework.ParsePrivKey(privKeyBytes)
	} else {
		panic("privKeyBytes are invalid!")
	}
}

func (E ECDSASignatureFunction) SupportPubKey(pubKeyBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+密钥类型长度+密钥长度，密钥数据的算法标识对应ECDSA签名算法，并且密钥类型是公钥
	return len(pubKeyBytes) == ECDSA_PUBKEY_LENGTH && E.GetAlgorithm().Match(pubKeyBytes, 0) && pubKeyBytes[framework.ALGORYTHM_CODE_SIZE] == framework.PUBLIC.Code
}

func (E ECDSASignatureFunction) ParsePubKey(pubKeyBytes []byte) framework.PubKey {
	if E.SupportPubKey(pubKeyBytes) {
		return framework.ParsePubKey(pubKeyBytes)
	} else {
		panic("pubKeyBytes are invalid!")
	}
}

func (E ECDSASignatureFunction) SupportDigest(digestBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+摘要长度，字节数组的算法标识对应ECDSA算法
	return len(digestBytes) == ECDSA_SIGNATUREDIGEST_LENGTH && E.GetAlgorithm().Match(digestBytes, 0)
}

func (E ECDSASignatureFunction) ParseDigest(digestBytes []byte) framework.SignatureDigest {
	if E.SupportDigest(digestBytes) {
		return framework.ParseSignatureDigest(digestBytes)
	} else {
		panic("digestBytes are invalid!")
	}
}
