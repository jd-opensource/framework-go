package adv

import "github.com/blockchain-jd-com/framework-go/crypto/framework"

/**
 * @Author: imuge
 * @Date: 2022/5/10 20:11 下午
 */

var (
	ELGAMAL  = ElgamalCryptoFunction{}
	PAILLIER = PaillierCryptoFunction{}
)

var _ framework.CryptoService = (*AdvCryptoService)(nil)

type AdvCryptoService struct {
	functions []framework.CryptoFunction
}

func NewAdvCryptoService() AdvCryptoService {
	return AdvCryptoService{
		[]framework.CryptoFunction{ELGAMAL, PAILLIER},
	}
}

func (c AdvCryptoService) GetFunctions() []framework.CryptoFunction {
	return c.functions
}
