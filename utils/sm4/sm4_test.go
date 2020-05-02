package sm4

import (
	"framework-go/utils/bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

/**
 * @Author: imuge
 * @Date: 2020/5/2 12:05 下午
 */

func TestAesEncrypt(t *testing.T) {
	key := GenerateSymmetricKey()
	encrypt := Encrypt(key, []byte("abc"))
	require.NotNil(t, encrypt)
}

func TestAesDecrypt(t *testing.T) {
	key := GenerateSymmetricKey()
	encrypt := Encrypt(key, []byte("abc"))
	require.NotNil(t, encrypt)
	decrypt := Decrypt(key, encrypt)
	require.NotNil(t, decrypt)
	require.True(t, bytes.Equals([]byte("abc"), decrypt))

}
