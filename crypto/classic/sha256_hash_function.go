package classic

import (
	"framework-go/crypto/framework"
	"framework-go/utils/bytes"
	"framework-go/utils/sha"
)

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:41 下午
 */

var _ framework.HashFunction = (*SHA256HashFunction)(nil)

var SHA256_DIGEST_BYTES = 256 / 8

type SHA256HashFunction struct {
}

func (S SHA256HashFunction) GetAlgorithm() framework.CryptoAlgorithm {
	return SHA256_ALGORITHM
}

func (S SHA256HashFunction) Hash(data []byte) framework.HashDigest {
	if data == nil {
		panic("data is null!")
	}

	return framework.NewHashDigest(SHA256_ALGORITHM, sha.Sha256(data))
}

func (S SHA256HashFunction) Verify(digest framework.HashDigest, data []byte) bool {
	hashDigest := S.Hash(data)
	return bytes.Equals(hashDigest.ToBytes(), digest.ToBytes())
}

func (S SHA256HashFunction) SupportHashDigest(digestBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+摘要长度，以及算法标识；
	return SHA256_DIGEST_BYTES == len(digestBytes) && SHA256_ALGORITHM.Match(digestBytes, 0)
}

func (S SHA256HashFunction) ResolveHashDigest(digestBytes []byte) framework.HashDigest {
	if S.SupportHashDigest(digestBytes) {
		return framework.NewHashDigest(SHA256_ALGORITHM, digestBytes)
	} else {
		panic("digestBytes is invalid!")
	}
}
