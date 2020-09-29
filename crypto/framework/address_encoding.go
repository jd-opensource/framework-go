package framework

import (
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
	"github.com/blockchain-jd-com/framework-go/utils/ripemd160"
	"github.com/blockchain-jd-com/framework-go/utils/sha"
)

/*
 * Author: imuge
 * Date: 2020/5/26 上午9:29
 */

// 从公钥生成地址
func GenerateAddress(pubKey PubKey) []byte {
	h1Bytes := sha.Sha256(pubKey.GetRawKeyBytes())
	h2Bytes := ripemd160.Hash(h1Bytes)
	xBytes := bytes.Concat([]byte{ADDRESSVERSION_V1}, bytes.Int16ToBytes(pubKey.GetAlgorithm()), h2Bytes)
	checksum := sha.Sha256(sha.Sha256(xBytes))[:4]
	addressBytes := append(xBytes, checksum...)

	return addressBytes
}
