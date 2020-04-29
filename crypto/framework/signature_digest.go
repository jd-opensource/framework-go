package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:57 下午
 */

var _ CryptoDigest = (*SignatureDigest)(nil)

type SignatureDigest struct {
	BaseCryptoBytes
}

func NewSignatureDigest(algorithm CryptoAlgorithm, rawCryptoBytes []byte) SignatureDigest {
	return SignatureDigest{
		NewBaseCryptoBytes(algorithm, rawCryptoBytes),
	}
}

func ParseSignatureDigest(cryptoBytes []byte) SignatureDigest {
	return SignatureDigest{
		ParseBaseCryptoBytes(cryptoBytes, supportSignature),
	}
}

func (s SignatureDigest) GetRawDigest() []byte {
	slice := s.GetRawCryptoBytes()
	return slice.GetBytesCopy(0, slice.Size)
}

func supportSignature(algorithm CryptoAlgorithm) bool {
	return algorithm.IsSignatureAlgorithm()
}
