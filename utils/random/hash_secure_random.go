package random

import (
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
	"github.com/blockchain-jd-com/framework-go/utils/sha"
	"io"
)

/*
 * Author: imuge
 * Date: 2020/12/1 下午5:30
 */

var _ io.Reader = (*HashSecureRandom)(nil)

type HashSecureRandom struct {
	state  []byte
	output []byte
	i      int64
	hf     HashFunc

	hashSize      int
	outputOffset  int
	counterOffset int
	availableSize int
}

func NewHashSecureRandom(seed []byte, hashSize int, hf HashFunc) *HashSecureRandom {
	hsr := &HashSecureRandom{}
	hsr.hf = hf
	hsr.hashSize = hashSize
	hsr.outputOffset = hashSize
	hsr.counterOffset = hashSize * 2
	// 定义状态数据的空间：种子 + 上一次输出 + 轮次；
	initState := make([]byte, hsr.counterOffset+8)
	// 将原始的种子数据计算哈希摘要，作为后续随机数计算的种子；
	hf(seed, initState, 0)
	// 复制哈希种子，初始化“上一次输出”状态；
	copy(initState[hsr.outputOffset:(hsr.outputOffset+hashSize)], initState[0:hashSize])

	hsr.output = make([]byte, hashSize)
	hsr.availableSize = 0
	hsr.state = initState

	return hsr
}

func (sr *HashSecureRandom) Read(p []byte) (n int, err error) {
	// 用随机数填充数组；
	left := len(p)
	offset := 0
	for left > 0 && offset < len(p) {
		if sr.availableSize == 0 {
			sr.i = sr.i + 1
			iBytes := bytes.Int64ToBytes(sr.i)
			copy(sr.state[sr.counterOffset:sr.counterOffset+8], iBytes[0:8])
			sr.hf(sr.state, sr.output, 0)
			copy(sr.state[sr.outputOffset:sr.outputOffset+sr.hashSize], sr.output[0:sr.hashSize])
			sr.availableSize = sr.hashSize
		}
		copySize := Min(left, sr.availableSize)
		copy(p[offset:copySize+offset], sr.output[sr.hashSize-sr.availableSize:sr.hashSize-sr.availableSize+copySize])
		offset += copySize + offset
		left -= copySize
		sr.availableSize -= copySize
	}

	return len(p), nil
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type HashFunc = func(bytes []byte, output []byte, offset int)

var Sha256 = func(bytes, output []byte, offset int) {
	ob := sha.Sha256(bytes)
	copy(output[:], ob[offset:])
}
