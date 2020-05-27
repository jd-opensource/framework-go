package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:33 下午
 */

type RandomGenerator interface {
	NextBytes(size int) []byte
}
