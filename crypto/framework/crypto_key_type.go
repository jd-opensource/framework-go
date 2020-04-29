package framework

import "fmt"

/**
 * @Author: imuge
 * @Date: 2020/4/29 9:29 上午
 */

var (
	// 非对称密钥的公钥
	PUBLIC = CryptoKeyType{0x01}
	// 非对称密钥的私钥
	PRIVATE = CryptoKeyType{0x02}
	// 对称密钥
	SYMMETRIC = CryptoKeyType{0x03}
)

type CryptoKeyType struct {
	Code byte
}

func GetCryptoKeyType(code byte) CryptoKeyType {
	switch code {
	case 0x01:
		return PUBLIC
	case 0x02:
		return PRIVATE
	case 0x03:
		return SYMMETRIC
	default:
		panic(fmt.Sprintf("CryptoKeyType doesn't support enum code[%s]!", code))
	}
}
