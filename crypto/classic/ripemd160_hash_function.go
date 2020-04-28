package classic

import "framework-go/crypto/framework"

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:42 下午
 */

var _ framework.HashFunction = (*RIPEMD160HashFunction)(nil)

// TODO

type RIPEMD160HashFunction struct {
	
}

func (R RIPEMD160HashFunction) GetAlgorithm() framework.CryptoAlgorithm {
	panic("implement me")
}

func (R RIPEMD160HashFunction) Hash(data []byte) framework.HashDigest {
	panic("implement me")
}

func (R RIPEMD160HashFunction) Verify(digest framework.HashDigest, data []byte) bool {
	panic("implement me")
}

func (R RIPEMD160HashFunction) SupportHashDigest(digestBytes []byte) bool {
	panic("implement me")
}

func (R RIPEMD160HashFunction) ResolveHashDigest(digestBytes []byte) framework.HashDigest {
	panic("implement me")
}
