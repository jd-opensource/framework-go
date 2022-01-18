package sm2

import (
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

/**
 * @Author: imuge
 * @Date: 2020/5/2 10:24 上午
 */

func TestEncrypt(t *testing.T) {
	priv, pub, err := GenerateKeyPair()
	require.Nil(t, err)
	plainBytes := []byte("imuge")
	encrypt, err := Encrypt(pub, plainBytes)
	require.Nil(t, err)
	decrypt, err := Decrypt(priv, encrypt)
	require.Nil(t, err)
	require.True(t, bytes.Equals(plainBytes, decrypt))
}

func TestSign(t *testing.T) {
	priv, pub, err := GenerateKeyPair()
	require.Nil(t, err)
	plainBytes := []byte("imuge")
	sign, err := Sign(priv, plainBytes)
	require.Nil(t, err)
	require.True(t, Verify(pub, plainBytes, sign))

	pubBytes := PubKeyToBytes(pub)
	pubDec, err := BytesToPubKey(pubBytes)
	require.Nil(t, err)
	privBytes := PrivKeyToBytes(priv)
	privDec, err := BytesToPrivKey(privBytes)
	require.Nil(t, err)
	sign, err = Sign(privDec, plainBytes)
	require.Nil(t, err)
	require.True(t, Verify(pubDec, plainBytes, sign))
}

func TestPubKeyToBytes(t *testing.T) {
	_, pub, err := GenerateKeyPair()
	require.Nil(t, err)
	pubBytes := PubKeyToBytes(pub)
	pubDec, err := BytesToPubKey(pubBytes)
	require.Nil(t, err)
	require.Equal(t, pub, pubDec)
}

func TestPrivKeyToBytes(t *testing.T) {
	priv, _, err := GenerateKeyPair()
	require.Nil(t, err)
	privBytes := PrivKeyToBytes(priv)
	privDec, err := BytesToPrivKey(privBytes)
	require.Nil(t, err)
	privDecBytes := PrivKeyToBytes(privDec)
	require.True(t, bytes.Equals(privBytes, privDecBytes))
}
