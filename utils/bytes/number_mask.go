package bytes

import "fmt"

/**
 * @Author: imuge
 * @Date: 2020/5/21 6:00 下午
 */

var (
	NUMBERMASK_NONE   = NumberMask{}
	NUMBERMASK_TINY   = newNumberMask("TINY", 0)   // TINY编码
	NUMBERMASK_SHORT  = newNumberMask("SHORT", 1)  // SHORT编码
	NUMBERMASK_NORMAL = newNumberMask("NORMAL", 2) // NORMAL编码
	NUMBERMASK_LONG   = newNumberMask("LONG", 3)   // LONG编码
)

func GetNumberMask(name string) NumberMask {
	switch name {
	case NUMBERMASK_TINY.Name:
		return NUMBERMASK_TINY
	case NUMBERMASK_SHORT.Name:
		return NUMBERMASK_SHORT
	case NUMBERMASK_NORMAL.Name:
		return NUMBERMASK_NORMAL
	case NUMBERMASK_LONG.Name:
		return NUMBERMASK_LONG
	default:
		panic(fmt.Sprintf("unknown number mask name:%s", name))
	}
}

// 数值掩码；用于以更少的字节空间输出整数的字节数组
type NumberMask struct {
	Name string

	// 掩码位的个数
	BIT_COUNT byte
	// 头部长度的最大值
	MAX_HEADER_LENGTH int32

	// 最大边界值
	MAX_BOUNDARY_SIZE int64
	// 此常量对于 TINY、SHORT、NORMAL 有效
	BOUNDARY_SIZE_0 int64
	BOUNDARY_SIZE_1 int64
	BOUNDARY_SIZE_2 int64
	BOUNDARY_SIZE_3 int64
	BOUNDARY_SIZE_4 int64
	BOUNDARY_SIZE_5 int64
	BOUNDARY_SIZE_6 int64
	BOUNDARY_SIZE_7 int64

	boundarySizes []int64
}

func newNumberMask(name string, bitCount byte) NumberMask {
	mask := NumberMask{}
	mask.Name = name
	mask.BIT_COUNT = bitCount
	mask.MAX_HEADER_LENGTH = 1 << bitCount
	mask.boundarySizes = make([]int64, mask.MAX_HEADER_LENGTH)
	for i := 0; i < int(mask.MAX_HEADER_LENGTH); i++ {
		mask.boundarySizes[i] = mask.computeBoundarySize(int32(i + 1))
	}

	mask.MAX_BOUNDARY_SIZE = mask.boundarySizes[mask.MAX_HEADER_LENGTH-1]
	if bitCount == 0 {
		// TINY;
		mask.BOUNDARY_SIZE_0 = mask.boundarySizes[0]
		mask.BOUNDARY_SIZE_1 = -1
		mask.BOUNDARY_SIZE_2 = -1
		mask.BOUNDARY_SIZE_3 = -1
		mask.BOUNDARY_SIZE_4 = -1
		mask.BOUNDARY_SIZE_5 = -1
		mask.BOUNDARY_SIZE_6 = -1
		mask.BOUNDARY_SIZE_7 = -1
	} else if bitCount == 1 {
		// SHORT;
		mask.BOUNDARY_SIZE_0 = mask.boundarySizes[0]
		mask.BOUNDARY_SIZE_1 = mask.boundarySizes[1]
		mask.BOUNDARY_SIZE_2 = -1
		mask.BOUNDARY_SIZE_3 = -1
		mask.BOUNDARY_SIZE_4 = -1
		mask.BOUNDARY_SIZE_5 = -1
		mask.BOUNDARY_SIZE_6 = -1
		mask.BOUNDARY_SIZE_7 = -1
	} else if bitCount == 2 {
		// NORMAL;
		mask.BOUNDARY_SIZE_0 = mask.boundarySizes[0]
		mask.BOUNDARY_SIZE_1 = mask.boundarySizes[1]
		mask.BOUNDARY_SIZE_2 = mask.boundarySizes[2]
		mask.BOUNDARY_SIZE_3 = mask.boundarySizes[3]
		mask.BOUNDARY_SIZE_4 = -1
		mask.BOUNDARY_SIZE_5 = -1
		mask.BOUNDARY_SIZE_6 = -1
		mask.BOUNDARY_SIZE_7 = -1
	} else if bitCount == 3 {
		// LONG;
		mask.BOUNDARY_SIZE_0 = mask.boundarySizes[0]
		mask.BOUNDARY_SIZE_1 = mask.boundarySizes[1]
		mask.BOUNDARY_SIZE_2 = mask.boundarySizes[2]
		mask.BOUNDARY_SIZE_3 = mask.boundarySizes[3]
		mask.BOUNDARY_SIZE_4 = mask.boundarySizes[4]
		mask.BOUNDARY_SIZE_5 = mask.boundarySizes[5]
		mask.BOUNDARY_SIZE_6 = mask.boundarySizes[6]
		mask.BOUNDARY_SIZE_7 = mask.boundarySizes[7]
	}

	return mask
}

func (mask *NumberMask) Equals(mask2 NumberMask) bool {
	return mask.Name == mask2.Name && mask.BIT_COUNT == mask2.BIT_COUNT
}

// 在指定的头部长度下能够表示的数据大小的临界值（不含）
// headerLength 值范围必须大于 0 ，且小于等于 MAX_HEADER_LENGTH
func (mask *NumberMask) GetBoundarySize(headerLength int32) int64 {
	return mask.boundarySizes[headerLength-1]
}

func (mask *NumberMask) computeBoundarySize(headerLength int32) int64 {
	boundarySize := 1 << (int64(headerLength)*8 - int64(mask.BIT_COUNT))
	return int64(boundarySize)
}

// 获取能够表示指定的数值的掩码长度，即掩码所需的字节数
// number 要表示的数值；如果值范围超出掩码的有效范围，将引起恐慌
func (mask *NumberMask) GetMaskLength(number int64) int32 {
	if number > -1 {
		if number < mask.BOUNDARY_SIZE_0 {
			return 1
		}
		if number < mask.BOUNDARY_SIZE_1 {
			return 2
		}
		if number < mask.BOUNDARY_SIZE_2 {
			return 3
		}
		if number < mask.BOUNDARY_SIZE_3 {
			return 4
		}
		if number < mask.BOUNDARY_SIZE_4 {
			return 5
		}
		if number < mask.BOUNDARY_SIZE_5 {
			return 6
		}
		if number < mask.BOUNDARY_SIZE_6 {
			return 7
		}
		if number < mask.BOUNDARY_SIZE_7 {
			return 8
		}
	}
	panic(fmt.Sprintf("Number is out of the illegal range! --[number=%d]", number))
}

// 生成指定数值的掩码
// number 要表示的数值；如果值范围超出掩码的有效范围，将引起恐慌
func (mask *NumberMask) GenerateMask(number int64) []byte {
	// 计算掩码占用的字节长度；
	maskLen := mask.GetMaskLength(number)
	maskBytes := make([]byte, maskLen)
	mask.writeMask(number, maskLen, maskBytes, 0)
	return maskBytes
}

func (mask *NumberMask) writeMask(number int64, maskLen int32, buffer []byte, offset int32) int32 {
	// 计算掩码占用的字节长度；
	for i := maskLen; i > 0; i-- {
		buffer[offset+i-1] = (byte)((number >> (8 * (maskLen - i))) & 0xFF)
	}

	// 计算头字节的标识位；
	indicatorByte := (byte)((maskLen - 1) << (8 - mask.BIT_COUNT))
	// 设置标识位；
	buffer[offset] = indicatorByte | buffer[offset]
	return maskLen
}

func (mask *NumberMask) WriteMask(number int64) []byte {
	return mask.GenerateMask(number)
}

// 解析掩码的头字节获得该掩码实例的完整长度
// headByte 掩码的头字节；即掩码的字节序列的首个字节
// 返回掩码实例的完整长度
// 注：在字节流中，对首字节解析获取该值后减 1，可以得到该掩码后续要读取的字节长度
func (mask *NumberMask) ResolveMaskLength(headByte byte) int32 {
	len := int32(((headByte & 0xFF) >> (8 - mask.BIT_COUNT)) + 1)
	if len < 1 {
		panic(
			fmt.Sprintf("Illegal length [%d] was resolved from the head byte of NumberMask!", len))
	}
	if len > mask.MAX_HEADER_LENGTH {
		panic(fmt.Sprintf(
			"Illegal length [%d] was resolved from the head byte of NumberMask!", len))
	}
	return len
}

// 从字节中解析掩码表示的数值
func (mask *NumberMask) ResolveMaskedNumber(markBytes []byte) int64 {
	maskLen := mask.ResolveMaskLength(markBytes[0])

	// 清除首字节的标识位；
	numberHead := markBytes[0] & (0xFF >> mask.BIT_COUNT)

	// 转换字节大小；
	number := numberHead & 0xFF
	for i := int32(1); i < maskLen; i++ {
		number = (number << 8) | (markBytes[i] & 0xFF)
	}

	return int64(number)
}
