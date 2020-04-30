package ripemd160

import (
	"golang.org/x/crypto/ripemd160"
)

/**
 * @Author: imuge
 * @Date: 2020/4/30 5:25 下午
 */

func Hash(data []byte) []byte {
	hasher := ripemd160.New()
	hasher.Write(data)
	hashBytes := hasher.Sum(nil)
	return hashBytes[:20]
}
