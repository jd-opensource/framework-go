package classic

import "github.com/blockchain-jd-com/framework-go/crypto/framework"

/**
 * @Author: imuge
 * @Date: 2020/4/29 6:08 下午
 */

var (
	ED25519_ALGORITHM   = framework.DefineSignature("ED25519", false, 21)
	ECDSA_ALGORITHM     = framework.DefineSignature("ECDSA", false, 22)
	RSA_ALGORITHM       = framework.DefineSignature("RSA", true, 23)
	SHA256_ALGORITHM    = framework.DefineHash("SHA256", 24)
	RIPEMD160_ALGORITHM = framework.DefineHash("RIPEMD160", 25)
	AES_ALGORITHM       = framework.DefineSymmetricEncryption("AES", 26)
	GO_RANDOM_ALGORITHM = framework.DefineRandom("GO-RANDOM", 27)
)
