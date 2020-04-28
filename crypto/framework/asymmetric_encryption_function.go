package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:49 下午
 */

type AsymmetricEncryptionFunction interface {
	AsymmetricKeypairGenerator
	CryptoFunction
}
