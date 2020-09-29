package classic

import (
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
	"github.com/blockchain-jd-com/framework-go/utils/ripemd160"
)

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:42 下午
 */

var (
	RIPEMD160_DIGEST_BYTES  = 160 / 8
	RIPEMD160_DIGEST_LENGTH = framework.ALGORYTHM_CODE_SIZE + RIPEMD160_DIGEST_BYTES
)

var _ framework.HashFunction = (*RIPEMD160HashFunction)(nil)

type RIPEMD160HashFunction struct {
}

func (R RIPEMD160HashFunction) GetAlgorithm() framework.CryptoAlgorithm {
	return RIPEMD160_ALGORITHM
}

func (R RIPEMD160HashFunction) Hash(data []byte) framework.HashDigest {
	if data == nil {
		panic("data is null!")
	}

	return framework.NewHashDigest(R.GetAlgorithm(), ripemd160.Hash(data))
}

func (R RIPEMD160HashFunction) Verify(digest framework.HashDigest, data []byte) bool {
	hashDigest := R.Hash(data)
	return bytes.Equals(hashDigest.ToBytes(), digest.ToBytes())
}

func (R RIPEMD160HashFunction) SupportHashDigest(digestBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+摘要长度，以及算法标识；
	return RIPEMD160_DIGEST_LENGTH == len(digestBytes) && R.GetAlgorithm().Match(digestBytes, 0)
}

func (R RIPEMD160HashFunction) ParseHashDigest(digestBytes []byte) framework.HashDigest {
	return framework.ParseHashDigest(digestBytes)
}
