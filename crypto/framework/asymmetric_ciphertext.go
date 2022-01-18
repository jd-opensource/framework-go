package framework

/**
 * @Author: imuge
 * @Date: 2020/4/29 5:28 下午
 */

var _ Ciphertext = (*AsymmetricCiphertext)(nil)

type AsymmetricCiphertext struct {
	BaseCryptoBytes
}

func NewAsymmetricCiphertext(algorithm CryptoAlgorithm, rawCryptoBytes []byte) *AsymmetricCiphertext {
	return &AsymmetricCiphertext{
		NewBaseCryptoBytes(algorithm, rawCryptoBytes),
	}
}

func ParseAsymmetricCiphertext(cryptoBytes []byte) (*AsymmetricCiphertext, error) {
	bytes, err := ParseBaseCryptoBytes(cryptoBytes, supportAsymmetric)
	if err != nil {
		return nil, err
	}
	return &AsymmetricCiphertext{
		*bytes,
	}, nil
}

func (s AsymmetricCiphertext) GetRawCiphertext() []byte {
	slice, _ := s.GetRawCryptoBytes()
	bytesCopy, _ := slice.GetBytesCopy(0, slice.Size)
	return bytesCopy
}

func supportAsymmetric(algorithm CryptoAlgorithm) bool {
	return algorithm.IsAsymmetricEncryptionAlgorithm()
}
