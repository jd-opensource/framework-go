package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:16 下午
 */
type AsymmetricKeypairGenerator interface {
	GenerateKeypair() AsymmetricKeypair
	GenerateKeypairWithSeed(seed []byte) (AsymmetricKeypair, error)
}
