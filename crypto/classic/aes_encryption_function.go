package classic

import (
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/utils/aes"
)

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:28 下午
 */

var (
	AES_KEY_SIZE   = 128 / 8
	AES_BLOCK_SIZE = 128 / 8

	// AES-ECB
	AES_PLAINTEXT_BUFFER_LENGTH  = 256
	AES_CIPHERTEXT_BUFFER_LENGTH = 256 + 16 + 2
	AES_SYMMETRICKEY_LENGTH      = framework.ALGORYTHM_CODE_SIZE + framework.KEY_TYPE_BYTES + AES_KEY_SIZE
)

var _ framework.SymmetricEncryptionFunction = (*AESEncryptionFunction)(nil)

type AESEncryptionFunction struct {
}

func (A AESEncryptionFunction) GenerateSymmetricKey() framework.SymmetricKey {
	return framework.NewSymmetricKey(A.GetAlgorithm(), aes.GenerateRandomBytes(16))
}

func (A AESEncryptionFunction) GetAlgorithm() framework.CryptoAlgorithm {
	return AES_ALGORITHM
}

func (A AESEncryptionFunction) Encrypt(key framework.SymmetricKey, data []byte) framework.SymmetricCiphertext {
	rawKeyBytes := key.GetRawKeyBytes()

	// 验证原始密钥长度为128比特，即16字节
	if len(rawKeyBytes) != AES_KEY_SIZE {
		panic("This key has wrong format!")
	}

	// 验证密钥数据的算法标识对应AES算法
	if key.GetAlgorithm() != A.GetAlgorithm().Code {
		panic("The is not AES symmetric key!")
	}

	// 调用底层AES128算法并计算密文数据
	encrypt, err := aes.Encrypt(rawKeyBytes, data)
	if err != nil {
		panic(err)
	}
	return framework.NewSymmetricCiphertext(A.GetAlgorithm(), encrypt)
}

func (A AESEncryptionFunction) Decrypt(key framework.SymmetricKey, ciphertext framework.SymmetricCiphertext) []byte {
	rawKeyBytes := key.GetRawKeyBytes()
	rawCiphertextBytes := ciphertext.GetRawCiphertext()

	// 验证原始密钥长度为128比特，即16字节
	if len(rawKeyBytes) != AES_KEY_SIZE {
		panic("This key has wrong format!")
	}

	// 验证密钥数据的算法标识对应AES算法
	if key.GetAlgorithm() != A.GetAlgorithm().Code {
		panic("The is not AES symmetric key!")
	}

	// 验证原始密文长度为分组长度的整数倍
	if len(rawCiphertextBytes)%AES_BLOCK_SIZE != 0 {
		panic("This ciphertext has wrong format!")
	}

	// 验证密文数据算法标识对应AES算法
	if ciphertext.GetAlgorithm() != A.GetAlgorithm().Code {
		panic("This is not AES ciphertext!")
	}

	// 调用底层AES128算法解密，得到明文
	descrypt, err := aes.Decrypt(rawKeyBytes, rawCiphertextBytes)
	if err != nil {
		panic(err)
	}
	return descrypt
}

func (A AESEncryptionFunction) SupportSymmetricKey(symmetricKeyBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+密钥类型长度+密钥长度，字节数组的算法标识对应AES算法且密钥密钥类型是对称密钥
	return len(symmetricKeyBytes) == AES_SYMMETRICKEY_LENGTH && A.GetAlgorithm().Match(symmetricKeyBytes, 0) && symmetricKeyBytes[framework.ALGORYTHM_CODE_SIZE] == framework.SYMMETRIC.Code
}

func (A AESEncryptionFunction) ParseSymmetricKey(symmetricKeyBytes []byte) framework.SymmetricKey {
	if A.SupportSymmetricKey(symmetricKeyBytes) {
		return framework.ParseSymmetricKey(symmetricKeyBytes)
	} else {
		panic("symmetricKeyBytes is invalid!")
	}
}

func (A AESEncryptionFunction) SupportCiphertext(ciphertextBytes []byte) bool {
	// 验证(输入字节数组长度-算法标识长度)是分组长度的整数倍，字节数组的算法标识对应AES算法
	return (len(ciphertextBytes)-framework.ALGORYTHM_CODE_SIZE)%AES_BLOCK_SIZE == 0 && A.GetAlgorithm().Match(ciphertextBytes, 0)
}

func (A AESEncryptionFunction) ParseCiphertext(ciphertextBytes []byte) framework.SymmetricCiphertext {
	if A.SupportCiphertext(ciphertextBytes) {
		return framework.ParseSymmetricCiphertext(ciphertextBytes)
	} else {
		panic("ciphertextBytes is invalid!")
	}
}
