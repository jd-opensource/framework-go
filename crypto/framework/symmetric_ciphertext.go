package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:26 下午
 */

var _ Ciphertext = (*SymmetricCiphertext)(nil)

type SymmetricCiphertext struct {
	BaseCryptoBytes
}

func NewSymmetricCiphertext(algorithm CryptoAlgorithm, rawCryptoBytes []byte) SymmetricCiphertext {
	return SymmetricCiphertext{
		NewBaseCryptoBytes(algorithm, rawCryptoBytes),
	}
}

func ParseSymmetricCiphertext(cryptoBytes []byte) SymmetricCiphertext {
	return SymmetricCiphertext{
		ParseBaseCryptoBytes(cryptoBytes, supportSymmetric),
	}
}

func (s SymmetricCiphertext) GetRawCiphertext() []byte {
	slice := s.GetRawCryptoBytes()
	return slice.GetBytesCopy(0, slice.Size)
}

func supportSymmetric(algorithm CryptoAlgorithm) bool {
	return algorithm.HasSymmetricKey()
}
