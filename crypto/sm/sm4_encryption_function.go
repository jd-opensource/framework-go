package sm

import "framework-go/crypto/framework"

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:30 下午
 */

var _ framework.SymmetricEncryptionFunction = (*SM4EncryptionFunction)(nil)

// TODO
type SM4EncryptionFunction struct {
}

func (S SM4EncryptionFunction) GenerateSymmetricKey() framework.SymmetricKey {
	panic("implement me")
}

func (S SM4EncryptionFunction) GetAlgorithm() framework.CryptoAlgorithm {
	panic("implement me")
}

func (S SM4EncryptionFunction) Encrypt(key framework.SymmetricKey, data []byte) framework.SymmetricCiphertext {
	panic("implement me")
}

func (S SM4EncryptionFunction) Decrypt(key framework.SymmetricKey, ciphertext framework.SymmetricCiphertext) []byte {
	panic("implement me")
}

func (S SM4EncryptionFunction) SupportSymmetricKey(symmetricKeyBytes []byte) bool {
	panic("implement me")
}

func (S SM4EncryptionFunction) ParseSymmetricKey(symmetricKeyBytes []byte) framework.SymmetricKey {
	panic("implement me")
}

func (S SM4EncryptionFunction) SupportCiphertext(ciphertextBytes []byte) bool {
	panic("implement me")
}

func (S SM4EncryptionFunction) ParseCiphertext(ciphertextBytes []byte) framework.SymmetricCiphertext {
	panic("implement me")
}
