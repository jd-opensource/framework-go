package sm

import (
	"framework-go/crypto/framework"
	"framework-go/utils/sm4"
)

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:30 下午
 */

var (
	SM4_KEY_SIZE   = 128 / 8
	SM4_BLOCK_SIZE = 128 / 8
	// SM4-CBC
	SM4_PLAINTEXT_BUFFER_LENGTH  = 256
	SM4_CIPHERTEXT_BUFFER_LENGTH = 256 + 32 + 2
	SM4_SYMMETRICKEY_LENGTH      = framework.ALGORYTHM_CODE_SIZE + framework.KEY_TYPE_BYTES + SM4_KEY_SIZE
)

var _ framework.SymmetricEncryptionFunction = (*SM4EncryptionFunction)(nil)

type SM4EncryptionFunction struct {
}

func (S SM4EncryptionFunction) GenerateSymmetricKey() framework.SymmetricKey {
	return framework.NewSymmetricKey(S.GetAlgorithm(), sm4.GenerateSymmetricKey())
}

func (S SM4EncryptionFunction) GetAlgorithm() framework.CryptoAlgorithm {
	return SM4_ALGORITHM
}

func (S SM4EncryptionFunction) Encrypt(key framework.SymmetricKey, data []byte) framework.SymmetricCiphertext {
	return framework.NewSymmetricCiphertext(S.GetAlgorithm(), sm4.Encrypt(key.GetRawKeyBytes(), data))
}

func (S SM4EncryptionFunction) Decrypt(key framework.SymmetricKey, ciphertext framework.SymmetricCiphertext) []byte {
	return sm4.Decrypt(key.GetRawKeyBytes(), ciphertext.GetRawCiphertext())
}

func (S SM4EncryptionFunction) SupportSymmetricKey(symmetricKeyBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+密钥类型长度+密钥长度，字节数组的算法标识对应SM4算法且密钥密钥类型是对称密钥
	return len(symmetricKeyBytes) == SM4_SYMMETRICKEY_LENGTH && S.GetAlgorithm().Match(symmetricKeyBytes, 0) && symmetricKeyBytes[framework.ALGORYTHM_CODE_SIZE] == framework.SYMMETRIC.Code
}

func (S SM4EncryptionFunction) ParseSymmetricKey(symmetricKeyBytes []byte) framework.SymmetricKey {
	if S.SupportSymmetricKey(symmetricKeyBytes) {
		return framework.ParseSymmetricKey(symmetricKeyBytes)
	} else {
		panic("symmetricKeyBytes is invalid!")
	}
}

func (S SM4EncryptionFunction) SupportCiphertext(ciphertextBytes []byte) bool {
	// 验证(输入字节数组长度-算法标识长度)是分组长度的整数倍，字节数组的算法标识对应SM4算法
	return (len(ciphertextBytes)-framework.ALGORYTHM_CODE_SIZE)%SM4_BLOCK_SIZE == 0 && S.GetAlgorithm().Match(ciphertextBytes, 0)
}

func (S SM4EncryptionFunction) ParseCiphertext(ciphertextBytes []byte) framework.SymmetricCiphertext {
	if S.SupportCiphertext(ciphertextBytes) {
		return framework.ParseSymmetricCiphertext(ciphertextBytes)
	} else {
		panic("ciphertextBytes is invalid!")
	}
}
