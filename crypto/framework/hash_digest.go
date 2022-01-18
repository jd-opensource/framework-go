package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:40 下午
 */

var _ CryptoDigest = (*HashDigest)(nil)

type HashDigest struct {
	BaseCryptoBytes
}

func NewHashDigest(algorithm CryptoAlgorithm, rawCryptoBytes []byte) *HashDigest {
	return &HashDigest{
		NewBaseCryptoBytes(algorithm, rawCryptoBytes),
	}
}

func ParseHashDigest(cryptoBytes []byte) (*HashDigest, error) {
	bytes, err := ParseBaseCryptoBytes(cryptoBytes, supportHash)
	if err != nil {
		return nil, err
	}
	return &HashDigest{
		*bytes,
	}, nil
}

func (s HashDigest) GetRawDigest() []byte {
	slice, _ := s.GetRawCryptoBytes()
	bytesCopy, _ := slice.GetBytesCopy(0, slice.Size)
	return bytesCopy
}

func supportHash(algorithm CryptoAlgorithm) bool {
	return algorithm.IsHashAlgorithm()
}
