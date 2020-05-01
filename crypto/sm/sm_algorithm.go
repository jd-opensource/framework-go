package sm

import "framework-go/crypto/framework"

/**
 * @Author: imuge
 * @Date: 2020/5/1 10:11 下午
 */

var (
	SM2_ALGORITHM = framework.DefineSignature("SM2", true, 2)
	SM3_ALGORITHM = framework.DefineHash("SM3", 3)
	SM4_ALGORITHM = framework.DefineSymmetricEncryption("SM4", 4)
)
