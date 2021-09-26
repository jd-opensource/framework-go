package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/blockchain-jd-com/framework-go/utils/random"
	"github.com/blockchain-jd-com/framework-go/utils/sha"
	"math/big"
)

/**
 * @Author: imuge
 * @Date: 2020/4/30 9:25 下午
 */

func GenerateKeyPair() *ecdsa.PrivateKey {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return priv
}

func GenerateKeyPairWithSeed(seed []byte) *ecdsa.PrivateKey {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), random.NewHashSecureRandom(seed, 32, random.Sha256))
	return priv
}

func PubKeyToBytes(pub *ecdsa.PublicKey) []byte {
	x := pub.X.Bytes()
	y := pub.Y.Bytes()
	return append(append([]byte{0x04}, x...), y...)
}

func BytesToPubKey(b []byte) (pub *ecdsa.PublicKey) {
	x := new(big.Int).SetBytes(b[1:33])
	y := new(big.Int).SetBytes(b[33:])
	return &ecdsa.PublicKey{elliptic.P256(), x, y}
}

func PrivKeyToBytes(priv *ecdsa.PrivateKey) []byte {
	return priv.D.Bytes()
}

func BytesToPrivKey(b []byte) *ecdsa.PrivateKey {
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = elliptic.P256()
	priv.D = new(big.Int).SetBytes(b)
	priv.PublicKey.X, priv.PublicKey.Y = elliptic.P256().ScalarBaseMult(b)

	return priv
}

func Sign(priv *ecdsa.PrivateKey, plainBytes []byte) []byte {
	r, s, _ := ecdsa.Sign(rand.Reader, priv, sha.Sha256(plainBytes))

	return append(r.Bytes(), s.Bytes()...)
}

func Verify(pub *ecdsa.PublicKey, plainBytes, cipherBytes []byte) bool {
	return ecdsa.Verify(pub, sha.Sha256(plainBytes), new(big.Int).SetBytes(cipherBytes[:32]), new(big.Int).SetBytes(cipherBytes[32:]))
}
