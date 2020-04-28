package common

import (
	"crypto/sha256"
)

// 对指定的字节数组进行 SHA128 哈希, 返回长度为 16 的字节数组
func Sha128(bytes []byte) []byte {
	return Sha256(bytes)[:16]
}

// 对指定的字节数组进行 SHA256 哈希, 返回长度为 32 的字节数组
func Sha256(bytes []byte) []byte {
	hasher := sha256.New()
	hasher.Write(bytes)
	sha := hasher.Sum(nil)
	return sha[:]
}
