package framework

/**
 * @Author: imuge
 * @Date: 2020/4/29 9:27 上午
 */

// 秘钥
type CryptoKey interface {
	CryptoBytes

	// 密钥的类型
	GetKeyType() CryptoKeyType

	// 原始的密钥数据
	GetRawKeyBytes() ([]byte, error)
}
