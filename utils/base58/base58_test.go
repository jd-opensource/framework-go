package base58

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBase58Encode(t *testing.T) {
	encode := Encode([]byte("abc"))
	require.NotNil(t, encode)
}

func TestBase58Decode(t *testing.T) {
	encode := Encode([]byte("abc"))
	require.NotNil(t, encode)
	decode, err := Decode(encode)
	require.Nil(t, err)
	require.NotNil(t, decode)
	require.Equal(t, []byte("abc"), decode)
}
