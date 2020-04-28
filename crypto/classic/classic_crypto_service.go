package classic

import "framework-go/crypto/framework"

/**
 * @Author: imuge
 * @Date: 2020/4/28 3:06 下午
 */

var _ framework.CryptoService = (*ClassicCryptoService)(nil)

// TODO

type ClassicCryptoService struct {

}

func (c ClassicCryptoService) GetFunctions() []framework.CryptoFunction {
	panic("implement me")
}
