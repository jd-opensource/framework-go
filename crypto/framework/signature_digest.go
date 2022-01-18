package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:57 下午
 */

var _ CryptoDigest = (*SignatureDigest)(nil)

type SignatureDigest struct {
	BaseCryptoBytes
}

func NewSignatureDigest(algorithm CryptoAlgorithm, rawCryptoBytes []byte) *SignatureDigest {
	return &SignatureDigest{
		NewBaseCryptoBytes(algorithm, rawCryptoBytes),
	}
}

func ParseSignatureDigest(cryptoBytes []byte) (*SignatureDigest, error) {
	bytes, err := ParseBaseCryptoBytes(cryptoBytes, supportSignature)
	if err != nil {
		return nil, err
	}
	return &SignatureDigest{
		*bytes,
	}, nil
}

func (s SignatureDigest) GetRawDigest() []byte {
	slice, _ := s.GetRawCryptoBytes()
	bytesCopy, _ := slice.GetBytesCopy(0, slice.Size)
	return bytesCopy
}

func supportSignature(algorithm CryptoAlgorithm) bool {
	return algorithm.IsSignatureAlgorithm()
}
