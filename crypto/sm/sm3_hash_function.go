package classic

import "framework-go/crypto/framework"

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:46 下午
 */

var _ framework.HashFunction = (*SM3HashFunction)(nil)

// TODO

type SM3HashFunction struct {

}

func (S SM3HashFunction) GetAlgorithm() framework.CryptoAlgorithm {
	panic("implement me")
}

func (S SM3HashFunction) Hash(data []byte) framework.HashDigest {
	panic("implement me")
}

func (S SM3HashFunction) Verify(digest framework.HashDigest, data []byte) bool {
	panic("implement me")
}

func (S SM3HashFunction) SupportHashDigest(digestBytes []byte) bool {
	panic("implement me")
}

func (S SM3HashFunction) ResolveHashDigest(digestBytes []byte) framework.HashDigest {
	panic("implement me")
}
