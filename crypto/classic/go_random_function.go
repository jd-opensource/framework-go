package classic

import "framework-go/crypto/framework"

/**
 * @Author: imuge
 * @Date: 2020/4/28 3:46 下午
 */

var _ framework.RandomFunction = (GoRandomFunction)(nil)

// TODO
type GoRandomFunction struct {

}

func (g GoRandomFunction) GetAlgorithm() framework.CryptoAlgorithm {
	panic("implement me")
}

func (g GoRandomFunction) Generate(seed []byte) framework.RandomGenerator {
	panic("implement me")
}
