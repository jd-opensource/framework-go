package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:57 下午
 */

// 公钥
type PubKey struct {
	BaseCryptoKey
}

func NewPubKey(algorithm CryptoAlgorithm, rawKeyBytes []byte) PubKey {
	return PubKey{NewBaseCryptoKey(algorithm, rawKeyBytes, PUBLIC)}
}

func ParsePubKey(keyBytes []byte) PubKey {
	return PubKey{ParseBaseCryptoKey(keyBytes)}
}
