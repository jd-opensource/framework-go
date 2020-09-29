package aes

import (
	"github.com/blockchain-jd-com/framework-go/utils/sha"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	encrypt, err := Encrypt(sha.Sha128([]byte("abc")), []byte("abc"))
	require.Nil(t, err)
	require.NotNil(t, encrypt)
}

func TestAesDecrypt(t *testing.T) {
	encrypt, err := Encrypt(sha.Sha128([]byte("abc")), []byte("abc"))
	require.Nil(t, err)
	require.NotNil(t, encrypt)
	decrypt, err := Decrypt(sha.Sha128([]byte("abc")), encrypt)
	require.Nil(t, err)
	require.NotNil(t, decrypt)
	require.Equal(t, []byte("abc"), decrypt)

}
