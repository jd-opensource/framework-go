package common

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBase58Encode(t *testing.T) {
	encode := Base58Encode([]byte("abc"))
	require.NotNil(t, encode)
}

func TestBase58Decode(t *testing.T) {
	encode := Base58Encode([]byte("abc"))
	require.NotNil(t, encode)
	decode, err := Base58Decode(encode)
	require.Nil(t, err)
	require.NotNil(t, decode)
	require.Equal(t, []byte("abc"), decode)
}