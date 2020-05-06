package bytes

/**
 * @Author: imuge
 * @Date: 2020/4/29 9:21 上午
 */

var EMPTY_SLICE = NewSlice([]byte{})

var _ BytesSerializable = (*Slice)(nil)

type Slice struct {
	Bytes  []byte
	Offset int
	Size   int
}

func NewSlice(bytes []byte) Slice {
	return Slice{
		bytes,
		0,
		len(bytes),
	}
}

func NewSliceWithOffset(bytes []byte, offset int) Slice {
	return NewSliceWithOffsetAndSize(bytes, offset, len(bytes)-offset)
}

func NewSliceWithOffsetAndSize(bytes []byte, offset, size int) Slice {
	if offset+size > len(bytes) {
		panic("index out of bounds")
	}
	return Slice{
		bytes,
		offset,
		size,
	}
}

func (s Slice) IsEmpty() bool {
	return s.Size == 0
}

func (s Slice) GetByte(offset int) byte {
	off := s.Offset + offset
	s.checkBoundary(off, 1)
	return s.Bytes[off]
}

func (s Slice) GetInt16(offset int) int16 {
	off := s.Offset + offset
	s.checkBoundary(off, 2)
	return ToInt16(s.Bytes[off : off+2])
}

func (s Slice) GetInt(offset int) int {
	off := s.Offset + offset
	s.checkBoundary(off, 4)
	return ToInt(s.Bytes[off : off+4])
}

func (s Slice) GetInt64(offset int) int64 {
	off := s.Offset + offset
	s.checkBoundary(off, 8)
	return ToInt64(s.Bytes[off : off+8])
}

func (s Slice) GetString() string {
	return ToString(s.Bytes[s.Offset : s.Offset+s.Size])
}

func (s Slice) GetBytesCopy(offset, size int) []byte {
	newOffset := s.Offset + offset
	s.checkBoundary(newOffset, size)

	if size == 0 {
		return []byte{}
	}
	dst := make([]byte, size)
	copy(dst, s.Bytes[newOffset:newOffset+size])
	return dst
}

func (s Slice) GetSlice(offset, size int) Slice {
	newOffset := s.Offset + offset
	s.checkBoundary(newOffset, size)
	return NewSliceWithOffsetAndSize(s.Bytes, newOffset, size)
}

func (s Slice) ToBytes() []byte {
	return s.GetBytesCopy(0, s.Size)
}

func (s Slice) checkBoundary(offset, len int) {
	if offset < s.Offset || offset+len > s.Offset+s.Size {
		panic("The accessing index is out of BytesSlice's bounds!")
	}
}
