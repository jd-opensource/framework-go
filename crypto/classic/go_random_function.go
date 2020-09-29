package classic

import (
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"math/rand"
)

/**
 * @Author: imuge
 * @Date: 2020/4/28 3:46 下午
 */

var _ framework.RandomFunction = (*GoRandomFunction)(nil)

type GoRandomFunction struct {
}

func (g GoRandomFunction) GetAlgorithm() framework.CryptoAlgorithm {
	return GO_RANDOM_ALGORITHM
}

func (g GoRandomFunction) Generate(seed int64) framework.RandomGenerator {
	return GoRandomGenerator{seed}
}

var _ framework.RandomGenerator = (*GoRandomGenerator)(nil)

type GoRandomGenerator struct {
	seed int64
}

func (g GoRandomGenerator) NextBytes(size int) []byte {
	bytes := make([]byte, size)
	rand.Seed(g.seed)
	rand.Read(bytes)

	return bytes
}
