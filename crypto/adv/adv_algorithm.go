package adv

import "github.com/blockchain-jd-com/framework-go/crypto/framework"

/**
 * @Author: imuge
 * @Date: 2022/5/10 20:11 下午
 */

var (
	ELGAMAL_ALGORITHM  = framework.DefineSignature("ELGAMAL", true, 31)
	PAILLIER_ALGORITHM = framework.DefineSignature("PAILLIER", true, 32)
)
