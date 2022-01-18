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

func ParseBaseCryptoKey(cryptoBytes []byte) (*BaseCryptoKey, error) {
	bcbs, err := ParseBaseCryptoBytes(cryptoBytes, supportAsymmetricOrSymmetric)
	if err != nil {
		return nil, err
	}
	bytes, err := bcbs.GetRawCryptoBytes()
	if err != nil {
		return nil, err
	}
	keyType, err := DecodeKeyType(*bytes)
	if err != nil {
		return nil, err
	}
	return &BaseCryptoKey{
		*bcbs,
		keyType,
	}, nil
}

func (b BaseCryptoKey) GetKeyType() CryptoKeyType {
	return b.keyType
}

func (b BaseCryptoKey) GetRawKeyBytes() ([]byte, error) {
	slice, err := b.GetRawCryptoBytes()
	if err != nil {
		return nil, err
	}
	return slice.GetBytesCopy(1, slice.Size-1)
}

func supportAsymmetricOrSymmetric(algorithm CryptoAlgorithm) bool {
	return algorithm.HasAsymmetricKey() || algorithm.HasSymmetricKey()
}
