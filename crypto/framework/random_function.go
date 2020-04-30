package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:32 下午
 */

type RandomFunction interface {
	CryptoFunction

	Generate(seed int64) RandomGenerator
}
