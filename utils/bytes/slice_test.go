package bytes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author: imuge
 * @Date: 2020/5/6 10:29 上午
 */

func TestSlice_IsEmpty(t *testing.T) {
	cases := []struct {
		slice  Slice
		expect bool
	}{
		{NewSlice(nil), true},
		{NewSlice([]byte("imuge")), false},
	}
	for _, c := range cases {
		assert.Equal(t, c.expect, c.slice.IsEmpty())
	}
}

func TestSlice_GetByte(t *testing.T) {
	cases := []struct {
		slice  Slice
		offset int
		expect byte
	}{
		{NewSlice([]byte{0x01}), 0, byte(0x01)},
		{NewSlice([]byte{0x00}), 0, byte(0x00)},
		{NewSlice([]byte{0x10, 0x11}), 1, byte(0x11)},
	}
	for _, c := range cases {
		assert.Equal(t, c.expect, c.slice.GetByte(c.offset))
	}
}

func TestSlice_GetInt16(t *testing.T) {
	cases := []struct {
		slice  Slice
		offset int
		expect int16
	}{
		{NewSlice(Int16ToBytes(1)), 0, int16(1)},
		{NewSlice(Int16ToBytes(2)), 0, int16(2)},
		{NewSlice(append(Int16ToBytes(1), Int16ToBytes(3)...)), 2, int16(3)},
	}
	for _, c := range cases {
		assert.Equal(t, c.expect, c.slice.GetInt16(c.offset))
	}
}

func TestSlice_GetInt32(t *testing.T) {
	cases := []struct {
		slice  Slice
		offset int
		expect int32
	}{
		{NewSlice(Int32ToBytes(1)), 0, 1},
		{NewSlice(Int32ToBytes(2)), 0, 2},
		{NewSlice(append(Int32ToBytes(1), Int32ToBytes(3)...)), 4, 3},
	}
	for _, c := range cases {
		assert.Equal(t, c.expect, c.slice.GetInt32(c.offset))
	}
}

func TestSlice_GetInt64(t *testing.T) {
	cases := []struct {
		slice  Slice
		offset int
		expect int64
	}{
		{NewSlice(Int64ToBytes(1)), 0, 1},
		{NewSlice(Int64ToBytes(2)), 0, 2},
		{NewSlice(append(Int64ToBytes(1), Int64ToBytes(3)...)), 8, 3},
	}
	for _, c := range cases {
		assert.Equal(t, c.expect, c.slice.GetInt64(c.offset))
	}
}

func TestSlice_GetSlice(t *testing.T) {
	cases := []struct {
		slice  Slice
		offset int
		size   int
	}{
		{NewSlice(Int64ToBytes(1)), 0, 1},
		{NewSlice(Int64ToBytes(2)), 0, 2},
		{NewSlice(append(Int64ToBytes(1), Int64ToBytes(3)...)), 8, 3},
	}
	for _, c := range cases {
		assert.Equal(t, c.slice.Bytes, c.slice.GetSlice(c.offset, c.size).Bytes)
		assert.Equal(t, c.offset, c.slice.GetSlice(c.offset, c.size).Offset)
		assert.Equal(t, c.size, c.slice.GetSlice(c.offset, c.size).Size)
	}
}

func TestSlice_GetBytesCopy(t *testing.T) {
	cases := []struct {
		bytes  []byte
		offset int
		size   int
	}{
		{Int64ToBytes(1), 0, 8},
		{Int64ToBytes(2), 0, 4},
		{append(Int64ToBytes(1), Int64ToBytes(3)...), 4, 4},
	}
	for _, c := range cases {
		assert.Equal(t, c.bytes[c.offset*2:c.offset*2+c.size], NewSliceWithOffset(c.bytes, c.offset).GetBytesCopy(c.offset, c.size))
	}
}

func TestSlice_ToBytes(t *testing.T) {
	cases := []struct {
		bytes  []byte
		offset int
	}{
		{Int64ToBytes(1), 0},
		{Int64ToBytes(2), 0},
		{append(Int64ToBytes(1), Int64ToBytes(3)...), 8},
	}
	for _, c := range cases {
		assert.Equal(t, c.bytes[c.offset:len(c.bytes)], NewSliceWithOffset(c.bytes, c.offset).ToBytes())
	}
}
