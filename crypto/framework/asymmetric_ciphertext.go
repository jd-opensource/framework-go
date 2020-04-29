package framework

/**
 * @Author: imuge
 * @Date: 2020/4/29 5:28 下午
 */

var _ Ciphertext = (*AsymmetricCiphertext)(nil)

type AsymmetricCiphertext struct {
	BaseCryptoBytes
}

func NewAsymmetricCiphertext(algorithm CryptoAlgorithm, rawCryptoBytes []byte) AsymmetricCiphertext {
	return AsymmetricCiphertext{
		NewBaseCryptoBytes(algorithm, rawCryptoBytes),
	}
}

func ParseAsymmetricCiphertext(cryptoBytes []byte) AsymmetricCiphertext {
	return AsymmetricCiphertext{
		ParseBaseCryptoBytes(cryptoBytes, supportAsymmetric),
	}
}

func (s AsymmetricCiphertext) GetRawCiphertext() []byte {
	slice := s.GetRawCryptoBytes()
	return slice.GetBytesCopy(0, slice.Size)
}

func supportAsymmetric(algorithm CryptoAlgorithm) bool {
	return algorithm.IsAsymmetricEncryptionAlgorithm()
}
