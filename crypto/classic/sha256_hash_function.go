package classic

import "framework-go/crypto/framework"

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:41 下午
 */

var _ framework.HashFunction = (*SHA256HashFunction)(nil)

// TODO

type SHA256HashFunction struct {
	
}

func (S SHA256HashFunction) GetAlgorithm() framework.CryptoAlgorithm {
	panic("implement me")
}

func (S SHA256HashFunction) Hash(data []byte) framework.HashDigest {
	panic("implement me")
}

func (S SHA256HashFunction) Verify(digest framework.HashDigest, data []byte) bool {
	panic("implement me")
}

func (S SHA256HashFunction) SupportHashDigest(digestBytes []byte) bool {
	panic("implement me")
}

func (S SHA256HashFunction) ResolveHashDigest(digestBytes []byte) framework.HashDigest {
	panic("implement me")
}
