package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:57 下午
 */

// 公钥
type PubKey struct {
	BaseCryptoKey
}

func NewPubKey(algorithm CryptoAlgorithm, rawKeyBytes []byte) *PubKey {
	return &PubKey{NewBaseCryptoKey(algorithm, rawKeyBytes, PUBLIC)}
}

func ParsePubKey(keyBytes []byte) (*PubKey, error) {
	key, err := ParseBaseCryptoKey(keyBytes)
	if err != nil {
		return nil, err
	}
	return &PubKey{*key}, nil
}

func (b PubKey) GetRawKeyBytes() []byte {
	slice, _ := b.GetRawCryptoBytes()
	bytesCopy, _ := slice.GetBytesCopy(1, slice.Size-1)

	return bytesCopy
}
