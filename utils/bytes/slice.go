package bytes

import "errors"

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

func NewSliceWithOffset(bytes []byte, offset int) (*Slice, error) {
	if offset >= len(bytes) {
		return nil, errors.New("offset out of bounds")
	}
	return NewSliceWithOffsetAndSize(bytes, offset, len(bytes)-offset)
}

func NewSliceWithOffsetAndSize(bytes []byte, offset, size int) (*Slice, error) {
	if offset+size > len(bytes) {
		return nil, errors.New("index out of bounds")
	}
	return &Slice{
		bytes,
		offset,
		size,
	}, nil
}

func (s Slice) IsEmpty() bool {
	return s.Size == 0
}

func (s Slice) GetByte(offset int) (byte, error) {
	off := s.Offset + offset
	if !s.checkBoundary(off, 1) {
		return 0, errors.New("accessing index is out of BytesSlice's bounds!")
	}
	return s.Bytes[off], nil
}

func (s Slice) GetInt16(offset int) (int16, error) {
	off := s.Offset + offset
	if !s.checkBoundary(off, 2) {
		return 0, errors.New("accessing index is out of BytesSlice's bounds!")
	}
	return ToInt16(s.Bytes[off : off+2]), nil
}

func (s Slice) GetInt32(offset int) (int32, error) {
	off := s.Offset + offset
	if !s.checkBoundary(off, 4) {
		return 0, errors.New("accessing index is out of BytesSlice's bounds!")
	}
	return ToInt32(s.Bytes[off : off+4]), nil
}

func (s Slice) GetInt64(offset int) (int64, error) {
	off := s.Offset + offset
	if !s.checkBoundary(off, 8) {
		return 0, errors.New("accessing index is out of BytesSlice's bounds!")
	}
	return ToInt64(s.Bytes[off : off+8]), nil
}

func (s Slice) GetString() string {
	return ToString(s.Bytes[s.Offset : s.Offset+s.Size])
}

func (s Slice) GetBytesCopy(offset, size int) ([]byte, error) {
	newOffset := s.Offset + offset
	if !s.checkBoundary(newOffset, size) {
		return nil, errors.New("accessing index is out of BytesSlice's bounds!")
	}
	if size == 0 {
		return []byte{}, nil
	}
	dst := make([]byte, size)
	copy(dst, s.Bytes[newOffset:newOffset+size])
	return dst, nil
}

func (s Slice) GetSlice(offset, size int) (*Slice, error) {
	newOffset := s.Offset + offset
	if !s.checkBoundary(newOffset, size) {
		return nil, errors.New("accessing index is out of BytesSlice's bounds!")
	}
	return NewSliceWithOffsetAndSize(s.Bytes, newOffset, size)
}

func (s Slice) ToBytes() []byte {
	bs, _ := s.GetBytesCopy(0, s.Size)
	return bs
}

func (s Slice) checkBoundary(offset, len int) bool {
	if offset < s.Offset || offset+len > s.Offset+s.Size {
		return false
	}

	return true
}
