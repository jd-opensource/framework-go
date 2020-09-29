package ecdsa

import (
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

/**
 * @Author: imuge
 * @Date: 2020/4/30 10:16 下午
 */

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
	require.Equal(t, pub.X, pubDec.X)
	require.Equal(t, pub.Y, pubDec.Y)
}

func TestPrivKeyToBytes(t *testing.T) {
	priv := GenerateKeyPair()
	privBytes := PrivKeyToBytes(priv)
	privDec := BytesToPrivKey(privBytes)
	privDecBytes := PrivKeyToBytes(privDec)
	require.True(t, bytes.Equals(privBytes, privDecBytes))
}
