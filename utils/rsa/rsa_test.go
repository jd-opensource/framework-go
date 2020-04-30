package rsa

import (
	"framework-go/utils/bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

/**
 * @Author: imuge
 * @Date: 2020/4/30 11:07 上午
 */

func TestEncrypt(t *testing.T) {
	priv := GenerateKeyPair()
	plainBytes := []byte("imuge")
	encrypt := Encrypt(&priv.PublicKey, plainBytes)
	decrypt := Decrypt(priv, encrypt)
	require.True(t, bytes.Equals(plainBytes, decrypt))
}

func TestSign(t *testing.T) {
	priv := GenerateKeyPair()
	plainBytes := []byte("imuge")
	sign := Sign(priv, plainBytes)
	require.True(t, Verify(&priv.PublicKey, plainBytes, sign))

	pub := priv.PublicKey
	pubBytes := PubKeyToBytes(&pub)
	pubDec := BytesToPubKey(pubBytes)
	privBytes := PrivKeyToBytes(priv)
	privDec := BytesToPrivKey(privBytes)
	sign = Sign(privDec, plainBytes)
	require.True(t, Verify(pubDec, plainBytes, sign))
}

func TestPubKeyToBytes(t *testing.T) {
	pub := GenerateKeyPair().PublicKey
	pubBytes := PubKeyToBytes(&pub)
	pubDec := BytesToPubKey(pubBytes)
	require.Equal(t, pub.E, pubDec.E)
	require.Equal(t, pub.N.Bytes(), pubDec.N.Bytes())
}

func TestPrivKeyToBytes(t *testing.T) {
	priv := GenerateKeyPair()
	privBytes := PrivKeyToBytes(priv)
	privDec := BytesToPrivKey(privBytes)
	privDecBytes := PrivKeyToBytes(privDec)
	require.True(t, bytes.Equals(privBytes, privDecBytes))
}
