package framework

/**
 * @Author: imuge
 * @Date: 2020/4/29 9:26 上午
 */

// 摘要
type CryptoDigest interface {
	CryptoBytes

	// 原始的摘要数据
	GetRawDigest() []byte
}
