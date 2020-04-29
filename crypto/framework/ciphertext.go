package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:25 下午
 */

// 密文
type Ciphertext interface {
	CryptoBytes

	// 原始的密文数据
	GetRawCiphertext() []byte
}
