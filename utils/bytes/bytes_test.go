package bytes

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

/**
 * @Author: imuge
 * @Date: 2020/4/29 8:24 上午
 */

func TestEquals(t *testing.T) {
	i := "imuge"
	b1 := FromString(i)
	b2 := FromString(i)
	require.True(t, b1.Equals(b2))
}

func TestFromBase58(t *testing.T) {
	b := FromString("imuge")
	b58 := b.ToBase58()
	require.Equal(t, b, FromBase58(b58))
}

func TestBytes_Size(t *testing.T) {
	cases := []struct {
		prefix *Bytes
		data   []byte
		size   int
	}{
		{nil, []byte("imuge"), 5},
		{NewBytes([]byte("liu")), []byte("imuge"), 8},
	}
	for _, c := range cases {
		assert.Equal(t, NewBytesWithPrefix(c.prefix, c.data).Size(), c.size)
	}

}

func TestBytes_GetDirectBytes(t *testing.T) {
	cases := []struct {
		prefix *Bytes
		data   []byte
		direct []byte
	}{
		{nil, []byte("imuge"), []byte("imuge")},
		{NewBytes([]byte("liu")), []byte("imuge"), []byte("imuge")},
	}
	for _, c := range cases {
		assert.Equal(t, NewBytesWithPrefix(c.prefix, c.data).GetDirectBytes(), c.direct)
	}
}
