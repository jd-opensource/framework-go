package classic

import "framework-go/crypto/framework"

/**
 * @Author: imuge
 * @Date: 2020/4/28 3:06 下午
 */

var (
	SM2 = SM2CryptoFunction{}
	SM3 = SM3HashFunction{}
	SM4 = SM4EncryptionFunction{}
)

var _ framework.CryptoService = (*SMCryptoService)(nil)

type SMCryptoService struct {
	functions []framework.CryptoFunction
}

func NewClassicCryptoService() SMCryptoService {
	return SMCryptoService{
		[]framework.CryptoFunction{SM2, SM3, SM4},
	}
}

func (c SMCryptoService) GetFunctions() []framework.CryptoFunction {
	return c.functions
}
