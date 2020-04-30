package classic

import "framework-go/crypto/framework"

/**
 * @Author: imuge
 * @Date: 2020/4/28 3:06 下午
 */

var (
	AES       = AESEncryptionFunction{}
	ED25519   = ED25519SignatureFunction{}
	RIPEMD160 = RIPEMD160HashFunction{}
	SHA256    = SHA256HashFunction{}
	GO_RANDOM = GoRandomFunction{}
	ECDSA     = ECDSASignatureFunction{}
	RSA       = RSACryptoFunction{}
)

var _ framework.CryptoService = (*ClassicCryptoService)(nil)

type ClassicCryptoService struct {
	functions []framework.CryptoFunction
}

func NewClassicCryptoService() ClassicCryptoService {
	// TODO
	return ClassicCryptoService{
		[]framework.CryptoFunction{SHA256, RSA, RIPEMD160, GO_RANDOM, ECDSA, AES},
	}
}

func (c ClassicCryptoService) GetFunctions() []framework.CryptoFunction {
	return c.functions
}
