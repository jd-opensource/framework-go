package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:26 下午
 */

var _ Ciphertext = (*SymmetricCiphertext)(nil)

type SymmetricCiphertext struct {
	BaseCryptoBytes
}

func NewSymmetricCiphertext(algorithm CryptoAlgorithm, rawCryptoBytes []byte) *SymmetricCiphertext {
	return &SymmetricCiphertext{
		NewBaseCryptoBytes(algorithm, rawCryptoBytes),
	}
}

func ParseSymmetricCiphertext(cryptoBytes []byte) (*SymmetricCiphertext, error) {
	bytes, err := ParseBaseCryptoBytes(cryptoBytes, supportSymmetric)
	if err != nil {
		return nil, err
	}
	return &SymmetricCiphertext{
		*bytes,
	}, nil
}

func (s SymmetricCiphertext) GetRawCiphertext() []byte {
	slice, _ := s.GetRawCryptoBytes()
	bytesCopy, _ := slice.GetBytesCopy(0, slice.Size)
	return bytesCopy
}

func supportSymmetric(algorithm CryptoAlgorithm) bool {
	return algorithm.HasSymmetricKey()
}
