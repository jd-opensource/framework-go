package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:15 下午
 */

// 对称秘钥
type SymmetricKey struct {
	BaseCryptoKey
}

func NewSymmetricKey(algorithm CryptoAlgorithm, rawKeyBytes []byte) *SymmetricKey {
	return &SymmetricKey{NewBaseCryptoKey(algorithm, rawKeyBytes, SYMMETRIC)}
}

func ParseSymmetricKey(keyBytes []byte) (*SymmetricKey, error) {
	key, err := ParseBaseCryptoKey(keyBytes)
	if err != nil {
		return nil, err
	}
	return &SymmetricKey{*key}, nil
}

func (b SymmetricKey) GetRawKeyBytes() []byte {
	slice, _ := b.GetRawCryptoBytes()
	bytesCopy, _ := slice.GetBytesCopy(1, slice.Size-1)

	return bytesCopy
}
