package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:14 下午
 */
type SymmetricKeyGenerator interface {
	GenerateSymmetricKey() SymmetricKey
}
