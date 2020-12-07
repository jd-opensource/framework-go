package random

import (
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
	"io"
)

/*
 * Author: imuge
 * Date: 2020/12/1 下午5:30
 */

var _ io.Reader = (*HashSecureRandom)(nil)

type HashSecureRandom struct {
	state []byte
	i     int64
	hf    HashFunc
}

func NewHashSecureRandom(seed []byte, hf HashFunc) *HashSecureRandom {
	return &HashSecureRandom{
		state: append(bytes.Int64ToBytes(0), hf(seed)...),
		i:     int64(0),
		hf:    hf,
	}
}

func (sr *HashSecureRandom) Read(p []byte) (n int, err error) {
	// 更新状态；
	sr.i += 1
	copy(sr.state[0:8], bytes.Int64ToBytes(sr.i))

	// 计算哈希值作为随机数输出；
	randomOutput := sr.hf(sr.state)

	// 用随机数填充数组；
	left := len(p)
	offset := 0
	for left > 0 {
		copySize := Min(left, len(randomOutput))
		copy(p[offset:offset+copySize], randomOutput[0:copySize])
		offset += copySize
		left -= copySize
	}

	return len(p), err
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type HashFunc = func([]byte) []byte
