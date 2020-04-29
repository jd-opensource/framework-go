package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:58 下午
 */

// 私钥
type PrivKey struct {
	BaseCryptoKey
}

func NewPrivKey(algorithm CryptoAlgorithm, rawKeyBytes []byte) PrivKey {
	return PrivKey{NewBaseCryptoKey(algorithm, rawKeyBytes, PRIVATE)}
}

func ParsePrivKey(keyBytes []byte) PrivKey {
	return PrivKey{ParseBaseCryptoKey(keyBytes)}
}