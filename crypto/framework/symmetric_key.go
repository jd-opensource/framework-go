package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:15 下午
 */

// 对称秘钥
type SymmetricKey struct {
	BaseCryptoKey
}

func NewSymmetricKey(algorithm CryptoAlgorithm, rawKeyBytes []byte) SymmetricKey {
	return SymmetricKey{NewBaseCryptoKey(algorithm, rawKeyBytes, SYMMETRIC)}
}

func ParseSymmetricKey(keyBytes []byte) SymmetricKey {
	return SymmetricKey{ParseBaseCryptoKey(keyBytes)}
}
