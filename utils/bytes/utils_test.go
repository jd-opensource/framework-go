package bytes

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/**
 * @Author: imuge
 * @Date: 2020/5/6 10:13 上午
 */

func TestIntToBytes(t *testing.T) {
	i := 10
	bytes := FromInt(i)
	require.Equal(t, i, ToInt(bytes.ToBytes()))
}

func TestInt32ToBytes(t *testing.T) {
	i := int32(10)
	bytes := FromInt32(i)
	require.Equal(t, i, ToInt32(bytes.ToBytes()))
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
	i := "imuge"
	bytes := StringToBytes(i)
	require.Equal(t, i, ToString(bytes))
}

func TestBoolToBytes(t *testing.T) {
	i := true
	byte := BoolToBytes(i)
	require.Equal(t, i, ToBoolean(byte))

	i = false
	byte = BoolToBytes(i)
	require.Equal(t, i, ToBoolean(byte))
}