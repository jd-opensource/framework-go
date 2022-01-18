package framework

import (
	"errors"
	"fmt"
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
)

/**
 * @Author: imuge
 * @Date: 2020/4/29 3:35 下午
 */

var _ CryptoBytes = (*BaseCryptoBytes)(nil)

type BaseCryptoBytes struct {
	*bytes.Bytes
	algorithm int16
}

func NewBaseCryptoBytes(algorithm CryptoAlgorithm, rawCryptoBytes []byte) BaseCryptoBytes {
	return BaseCryptoBytes{
		bytes.NewBytes(EncodeBytes(algorithm.Code, rawCryptoBytes)),
		algorithm.Code,
	}
}

func ParseBaseCryptoBytes(cryptoBytes []byte, support func(CryptoAlgorithm) bool) (*BaseCryptoBytes, error) {
	algorithm := DecodeAlgorithm(cryptoBytes)
	if !support(CryptoAlgorithm{Code: algorithm}) {
		return nil, errors.New(fmt.Sprintf("Not supported algorithm [code:%d]!", algorithm))
	}
	return &BaseCryptoBytes{
		bytes.NewBytes(cryptoBytes),
		algorithm,
	}, nil
}

func (b BaseCryptoBytes) GetAlgorithm() int16 {
	return b.algorithm
}

func (b BaseCryptoBytes) GetRawCryptoBytes() (*bytes.Slice, error) {
	return bytes.NewSliceWithOffset(b.GetDirectBytes(), CODE_SIZE)
}
