package bytes

import (
	"framework-go/utils/base58"
	"io"
)

/**
 * @Author: imuge
 * @Date: 2020/4/28 5:38 下午
 */

var (
	EMPTY_BYTES = NewBytes([]byte{})
	MAX_CACHE   = 256
	INT_BYTES   [256]Bytes
	LONG_BYTES  [256]Bytes
)

var _ BytesSerializable = (*Bytes)(nil)

type Bytes struct {
	prefix *Bytes
	data   []byte
}

func NewBytes(b []byte) *Bytes {
	if b == nil {
		panic("data is null!")
	}
	return &Bytes{
		prefix: nil,
		data:   b,
	}
}

func NewBytesWithPrefix(prefix *Bytes, b []byte) *Bytes {
	if b == nil {
		panic("data is null!")
	}
	return &Bytes{
		prefix: prefix,
		data:   b,
	}
}

func (b *Bytes) Size() int {
	if b.prefix == nil {
		return len(b.data)
	} else {
		return b.prefix.Size() + len(b.data)
	}
}

// 返回当前的字节数组（不包含前缀对象）
func (b *Bytes) GetDirectBytes() []byte {
	return b.data
}

func (b *Bytes) ConcatBytes(key Bytes) *Bytes {
	return NewBytesWithPrefix(b, key.data)
}

func (b *Bytes) Concat(key []byte) *Bytes {
	return NewBytesWithPrefix(b, key)
}

func (b *Bytes) WriteTo(out io.Writer) int {
	size := 0
	if b.prefix != nil {
		size = b.prefix.WriteTo(out)
	}
	n, err := out.Write(b.data)
	if err != nil {
		panic(err)
	}
	size += n
	return size
}

func (b *Bytes) Equals(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if b == obj {
		return true
	}
	ob, ok := obj.(Bytes)
	if !ok {
		return false
	}
	prefixIsEqual := false
	if b.prefix == nil && ob.prefix == nil {
		prefixIsEqual = true
	} else if b.prefix == nil {
		prefixIsEqual = false
	} else {
		prefixIsEqual = b.prefix.Equals(ob.prefix)
	}
	if !prefixIsEqual {
		return false
	}
	return Equals(b.data, ob.data)
}

func (b *Bytes) CopyTo(buffer []byte, offset int, size int) int {
	if size < 0 {
		panic("Argument len is negative!")
	}
	if size == 0 {
		return 0
	}
	s := 0
	if b.prefix != nil {
		s = b.prefix.CopyTo(buffer, offset, size)
	}
	if s < size {
		l := size - s
		if l >= len(b.data) {
			l = len(b.data)
		}
		copy(buffer[offset+s:], b.data[:])
		s += l
	}
	return s
}

func (b *Bytes) ToBytes() []byte {
	if b.prefix == nil || b.prefix.Size() == 0 {
		return b.data
	}
	size := b.Size()
	buffer := make([]byte, size)
	b.CopyTo(buffer, 0, size)
	return buffer
}

func init() {
	for i := 0; i < MAX_CACHE; i++ {
		INT_BYTES[i] = *NewBytes(IntToBytes(i))
		LONG_BYTES[i] = *NewBytes(Int64ToBytes(int64(i)))
	}
}

func (b *Bytes) ToBase58() string {
	return base58.Encode(b.ToBytes())
}

/**
 * 返回 Base58 编码的字符；
 */
func (b *Bytes) ToString() string {
	return b.ToBase58()
}

func (b *Bytes) ToUTF8String() string {
	return ToString(b.ToBytes())
}

func FromInt(value int) *Bytes {
	if value > -1 && value < MAX_CACHE {
		return &INT_BYTES[value]
	}
	return NewBytes(IntToBytes(value))
}

func FromString(str string) *Bytes {
	return NewBytes(StringToBytes(str))
}

func FromBase58(str string) *Bytes {
	bytes, err := base58.Decode(str)
	if err != nil {
		panic(err)
	}
	return NewBytes(bytes)
}
