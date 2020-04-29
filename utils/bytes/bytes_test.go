package bytes

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/**
 * @Author: imuge
 * @Date: 2020/4/29 8:24 上午
 */

func TestIntToBytes(t *testing.T) {
	i := 10
	bytes := FromInt(i)
	require.Equal(t, i, ToInt(bytes.ToBytes()))
}

func TestInt16ToBytes(t *testing.T) {
	i := int16(10)
	bytes := Int16ToBytes(i)
	require.Equal(t, i, ToInt16(bytes))
}

func TestInt64ToBytes(t *testing.T) {
	i := int64(10)
	bytes := Int64ToBytes(i)
	require.Equal(t, i, ToInt64(bytes))
}

func TestStringToBytes(t *testing.T) {
	p := "prefix"
	i := "imuge"
	bytes := FromString(i)
	require.Equal(t, i, ToString(bytes.ToBytes()))

	prefix := FromString(p)
	bytes = NewBytesWithPrefix(&prefix, bytes.data)
	require.Equal(t, p+i, bytes.ToUTF8String())
}

func TestEquals(t *testing.T) {
	i := "imuge"
	b1 := FromString(i)
	b2 := FromString(i)
	require.True(t, b1.Equals(b2))
}
