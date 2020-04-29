package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:40 下午
 */

var _ CryptoDigest = (*HashDigest)(nil)

type HashDigest struct {
	BaseCryptoBytes
}

func NewHashDigest(algorithm CryptoAlgorithm, rawCryptoBytes []byte) HashDigest {
	return HashDigest{
		NewBaseCryptoBytes(algorithm, rawCryptoBytes),
	}
}

func ParseHashDigest(cryptoBytes []byte) HashDigest {
	return HashDigest{
		ParseBaseCryptoBytes(cryptoBytes, supportHash),
	}
}

func (s HashDigest) GetRawDigest() []byte {
	slice := s.GetRawCryptoBytes()
	return slice.GetBytesCopy(0, slice.Size)
}

func supportHash(algorithm CryptoAlgorithm) bool {
	return algorithm.IsHashAlgorithm()
}
