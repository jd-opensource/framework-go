package framework

import (
	"errors"
	"fmt"
)

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

func GetCryptoKeyType(code byte) (CryptoKeyType, error) {
	switch code {
	case 0x01:
		return PUBLIC, nil
	case 0x02:
		return PRIVATE, nil
	case 0x03:
		return SYMMETRIC, nil
	default:
		return PUBLIC, errors.New(fmt.Sprintf("CryptoKeyType doesn't support enum code[%s]!", code))
	}
}
