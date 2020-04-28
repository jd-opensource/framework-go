package classic

import "framework-go/crypto/framework"

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:28 下午
 */

var _ framework.SymmetricEncryptionFunction = (*AESEncryptionFunction)(nil)

// TODO
type AESEncryptionFunction struct {

}

func (A AESEncryptionFunction) GenerateSymmetricKey() framework.SymmetricKey {
	panic("implement me")
}

func (A AESEncryptionFunction) GetAlgorithm() framework.CryptoAlgorithm {
	panic("implement me")
}

func (A AESEncryptionFunction) Encrypt(key framework.SymmetricKey, data []byte) framework.Ciphertext {
	panic("implement me")
}

func (A AESEncryptionFunction) Decrypt(key framework.SymmetricKey, ciphertext framework.Ciphertext) []byte {
	panic("implement me")
}

func (A AESEncryptionFunction) SupportSymmetricKey(symmetricKeyBytes []byte) bool {
	panic("implement me")
}

func (A AESEncryptionFunction) ResolveSymmetricKey(symmetricKeyBytes []byte) framework.SymmetricKey {
	panic("implement me")
}

func (A AESEncryptionFunction) SupportCiphertext(ciphertextBytes []byte) bool {
	panic("implement me")
}

func (A AESEncryptionFunction) ResolveCiphertext(ciphertextBytes []byte) framework.SymmetricCiphertext {
	panic("implement me")
}
