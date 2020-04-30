package framework

/**
 * @Author: imuge
 * @Date: 2020/4/29 4:39 下午
 */

var (
	KEY_TYPE_BYTES = 1
)

var _ CryptoKey = (*BaseCryptoKey)(nil)

type BaseCryptoKey struct {
	BaseCryptoBytes

	keyType CryptoKeyType
}

func NewBaseCryptoKey(algorithm CryptoAlgorithm, rawKeyBytes []byte, keyType CryptoKeyType) BaseCryptoKey {
	return BaseCryptoKey{
		NewBaseCryptoBytes(algorithm, EncodeKeyBytes(rawKeyBytes, keyType)),
		keyType,
	}
}

func ParseBaseCryptoKey(cryptoBytes []byte) BaseCryptoKey {
	bcbs := ParseBaseCryptoBytes(cryptoBytes, supportAsymmetricOrSymmetric)
	return BaseCryptoKey{
		bcbs,
		DecodeKeyType(bcbs.GetRawCryptoBytes()),
	}
}

func (b BaseCryptoKey) GetKeyType() CryptoKeyType {
	return b.keyType
}

func (b BaseCryptoKey) GetRawKeyBytes() []byte {
	slice := b.GetRawCryptoBytes()
	return slice.GetBytesCopy(1, slice.Size-1)
}

func supportAsymmetricOrSymmetric(algorithm CryptoAlgorithm) bool {
	return algorithm.HasAsymmetricKey() || algorithm.HasSymmetricKey()
}
