package ecdsa

import (
	"crypto/rand"
	"github.com/blockchain-jd-com/framework-go/utils/random"
	"github.com/blockchain-jd-com/framework-go/utils/sha"
	"math/big"
)

/**
 * @Author: imuge
 * @Date: 2020/4/30 9:25 下午
 */

func GenerateKeyPair() *PrivateKey {
	priv, _ := GenerateKey(S256(), rand.Reader)
	return priv
}

func GenerateKeyPairWithSeed(seed []byte) *PrivateKey {
	priv, _ := GenerateKey(S256(), random.NewHashSecureRandom(seed, 32, random.Sha256))
	return priv
}

func PubKeyToBytes(pub *PublicKey) []byte {
	x := pub.X.Bytes()
	y := pub.Y.Bytes()
	return append(append([]byte{0x04}, x...), y...)
}

func BytesToPubKey(b []byte) (pub *PublicKey) {
	x := new(big.Int).SetBytes(b[1:33])
	y := new(big.Int).SetBytes(b[33:])
	return &PublicKey{S256(), x, y}
}

func PrivKeyToBytes(priv *PrivateKey) []byte {
	return priv.D.Bytes()
}

func BytesToPrivKey(b []byte) *PrivateKey {
	priv := new(PrivateKey)
	priv.PublicKey.BitCurve = S256()
	priv.D = new(big.Int).SetBytes(b)
	priv.PublicKey.X, priv.PublicKey.Y = S256().ScalarBaseMult(b)

	return priv
}

func Sign(priv *PrivateKey, plainBytes []byte) []byte {
	hashed := sha.Sha256(plainBytes)
	r, s, _ := sign(rand.Reader, priv, hashed)

	return append(r.Bytes(), s.Bytes()...)
}

func Verify(pub *PublicKey, plainBytes, cipherBytes []byte) bool {
	return verify(pub, sha.Sha256(plainBytes), new(big.Int).SetBytes(cipherBytes[:32]), new(big.Int).SetBytes(cipherBytes[32:]))
}
