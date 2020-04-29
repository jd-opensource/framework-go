package aes

import (
	"framework-go/utils/sha"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	encrypt, err := Encrypt([]byte("abc"), sha.Sha128([]byte("abc")))
	require.Nil(t, err)
	require.NotNil(t, encrypt)
}

func TestAesDecrypt(t *testing.T) {
	encrypt, err := Encrypt([]byte("abc"), sha.Sha128([]byte("abc")))
	require.Nil(t, err)
	require.NotNil(t, encrypt)
	decrypt, err := Decrypt(encrypt, sha.Sha128([]byte("abc")))
	require.Nil(t, err)
	require.NotNil(t, decrypt)
	require.Equal(t, []byte("abc"), decrypt)

}
