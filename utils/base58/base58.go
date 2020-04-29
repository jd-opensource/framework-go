package base58

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

var (
	defaultAlphabet = newAlphabet("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

	ErrInvalidBase58String = errors.New("invalid base58 string")
)

type base58Alphabet struct {
	encodeTable        [58]rune
	decodeTable        [256]int
	unicodeDecodeTable []rune
}

func (alphabet base58Alphabet) String() string {
	return string(alphabet.encodeTable[:])
}

func newAlphabet(alphabet string) *base58Alphabet {
	if utf8.RuneCountInString(alphabet) != 58 {
		panic(fmt.Sprintf("Base58 base58Alphabet length must 58, but %d", utf8.RuneCountInString(alphabet)))
	}

	ret := new(base58Alphabet)
	for i := range ret.decodeTable {
		ret.decodeTable[i] = -1
	}
	ret.unicodeDecodeTable = make([]rune, 0, 58*2)
	var idx int
	var ch rune
	for _, ch = range alphabet {
		ret.encodeTable[idx] = ch
		if ch >= 0 && ch < 256 {
			ret.decodeTable[byte(ch)] = idx
		} else {
			ret.unicodeDecodeTable = append(ret.unicodeDecodeTable, ch)
			ret.unicodeDecodeTable = append(ret.unicodeDecodeTable, rune(idx))
		}
		idx++
	}
	return ret
}

func Encode(input []byte) string {
	inputLength := len(input)
	prefixZeroes := 0
	for prefixZeroes < inputLength && input[prefixZeroes] == 0 {
		prefixZeroes++
	}

	capacity := (inputLength-prefixZeroes)*138/100 + 1
	output := make([]byte, capacity)
	outputReverseEnd := capacity - 1

	var carry uint32
	var outputIdx int
	for _, inputByte := range input[prefixZeroes:] {
		carry = uint32(inputByte)

		outputIdx = capacity - 1
		for ; outputIdx > outputReverseEnd || carry != 0; outputIdx-- {
			carry += (uint32(output[outputIdx]) << 8)
			output[outputIdx] = byte(carry % 58)
			carry /= 58
		}
		outputReverseEnd = outputIdx
	}

	encodeTable := defaultAlphabet.encodeTable
	if len(defaultAlphabet.unicodeDecodeTable) == 0 {
		retStrBytes := make([]byte, prefixZeroes+(capacity-1-outputReverseEnd))
		for i := 0; i < prefixZeroes; i++ {
			retStrBytes[i] = byte(encodeTable[0])
		}
		for i, n := range output[outputReverseEnd+1:] {
			retStrBytes[prefixZeroes+i] = byte(encodeTable[n])
		}
		return string(retStrBytes)
	}
	retStrRunes := make([]rune, prefixZeroes+(capacity-1-outputReverseEnd))
	for i := 0; i < prefixZeroes; i++ {
		retStrRunes[i] = encodeTable[0]
	}
	for i, n := range output[outputReverseEnd+1:] {
		retStrRunes[prefixZeroes+i] = encodeTable[n]
	}
	return string(retStrRunes)
}

func Decode(input string) ([]byte, error) {
	capacity := utf8.RuneCountInString(input)*733/1000 + 1
	output := make([]byte, capacity)
	outputReverseEnd := capacity - 1
	var carry, outputIdx, i int
	var target rune

	zero58Byte := defaultAlphabet.encodeTable[0]
	prefixZeroes := 0
	skipZeros := false

	for _, target = range input {
		if !skipZeros {
			if target == zero58Byte {
				prefixZeroes++
				continue
			} else {
				skipZeros = true
			}
		}

		carry = -1
		if target >= 0 && target < 256 {
			carry = defaultAlphabet.decodeTable[target]
		} else { // unicode
			for i = 0; i < len(defaultAlphabet.unicodeDecodeTable); i += 2 {
				if defaultAlphabet.unicodeDecodeTable[i] == target {
					carry = int(defaultAlphabet.unicodeDecodeTable[i+1])
					break
				}
			}
		}
		if carry == -1 {
			return nil, ErrInvalidBase58String
		}

		outputIdx = capacity - 1
		for ; outputIdx > outputReverseEnd || carry != 0; outputIdx-- {
			carry += 58 * int(output[outputIdx])
			output[outputIdx] = byte(uint32(carry) & 0xff)
			carry >>= 8
		}
		outputReverseEnd = outputIdx
	}

	retBytes := make([]byte, prefixZeroes+(capacity-1-outputReverseEnd))
	copy(retBytes[prefixZeroes:], output[outputReverseEnd+1:])
	return retBytes, nil
}
