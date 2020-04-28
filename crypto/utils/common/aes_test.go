package common

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	encrypt, err := AesEncrypt([]byte("abc"), Sha128([]byte("abc")))
	require.Nil(t, err)
	require.NotNil(t, encrypt)
}

func TestAesDecrypt(t *testing.T) {
	encrypt, err := AesEncrypt([]byte("abc"), Sha128([]byte("abc")))
	require.Nil(t, err)
	require.NotNil(t, encrypt)
	decrypt, err := AesDecrypt(encrypt, Sha128([]byte("abc")))
	require.Nil(t, err)
	require.NotNil(t, decrypt)
	require.Equal(t, []byte("abc"), decrypt)

}
