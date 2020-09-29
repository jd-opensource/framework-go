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
	priv, pub := GenerateKeyPair()
	plainBytes := []byte("imuge")
	encrypt := Encrypt(pub, plainBytes)
	decrypt := Decrypt(priv, encrypt)
	require.True(t, bytes.Equals(plainBytes, decrypt))
}

func TestSign(t *testing.T) {
	priv, pub := GenerateKeyPair()
	plainBytes := []byte("imuge")
	sign := Sign(priv, plainBytes)
	require.True(t, Verify(pub, plainBytes, sign))

	pubBytes := PubKeyToBytes(pub)
	pubDec := BytesToPubKey(pubBytes)
	privBytes := PrivKeyToBytes(priv)
	privDec := BytesToPrivKey(privBytes)
	sign = Sign(privDec, plainBytes)
	require.True(t, Verify(pubDec, plainBytes, sign))
}

func TestPubKeyToBytes(t *testing.T) {
	_, pub := GenerateKeyPair()
	pubBytes := PubKeyToBytes(pub)
	pubDec := BytesToPubKey(pubBytes)
	require.Equal(t, pub, pubDec)
}

func TestPrivKeyToBytes(t *testing.T) {
	priv, _ := GenerateKeyPair()
	privBytes := PrivKeyToBytes(priv)
	privDec := BytesToPrivKey(privBytes)
	privDecBytes := PrivKeyToBytes(privDec)
	require.True(t, bytes.Equals(privBytes, privDecBytes))
}
