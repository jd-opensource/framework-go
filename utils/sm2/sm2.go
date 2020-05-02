package sm2

import (
	"crypto/rand"
	"github.com/ZZMarquis/gm/sm2"
	"math/big"
)

/**
 * @Author: imuge
 * @Date: 2020/5/2 8:47 上午
 */

var DEFAULT_UID = []byte("1234567812345678")

func GenerateKeyPair() (*sm2.PrivateKey, *sm2.PublicKey) {
	priv, pub, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	return priv, pub
}

func PubKeyToBytes(pub *sm2.PublicKey) []byte {
	return append([]byte{0x04}, append(pub.X.Bytes(), pub.Y.Bytes()...)...)
}

func BytesToPubKey(b []byte) (pub *sm2.PublicKey) {
	pub, err := sm2.RawBytesToPublicKey(b[1:])
	if err != nil {
		panic(err)
	}

	return pub
}

func RetrievePubKey(priv *sm2.PrivateKey) *sm2.PublicKey {
	pub := new(sm2.PublicKey)
	pub.Curve = priv.Curve
	pub.X, pub.Y = priv.Curve.ScalarBaseMult(priv.D.Bytes())
	return pub
}

func PrivKeyToBytes(priv *sm2.PrivateKey) []byte {
	return priv.D.Bytes()
}

func BytesToPrivKey(b []byte) *sm2.PrivateKey {
	priv, err := sm2.RawBytesToPrivateKey(b)
	if err != nil {
		panic(err)
	}

	return priv
}

func Encrypt(pub *sm2.PublicKey, plainBytes []byte) []byte {
	encrypt, err := sm2.Encrypt(pub, plainBytes, sm2.C1C3C2)
	if err != nil {
		panic(err)
	}

	return encrypt
}

func Decrypt(priv *sm2.PrivateKey, cipherBytes []byte) []byte {
	decrypt, err := sm2.Decrypt(priv, cipherBytes, sm2.C1C3C2)
	if err != nil {
		panic(err)
	}

	return decrypt
}

func Sign(priv *sm2.PrivateKey, plainBytes []byte) []byte {
	r, s, err := sm2.SignToRS(priv, DEFAULT_UID, plainBytes)
	if err != nil {
		panic(err)
	}

	return append(r.Bytes(), s.Bytes()...)
}

func Verify(pub *sm2.PublicKey, plainBytes, cipherBytes []byte) bool {
	r := new(big.Int).SetBytes(cipherBytes[0:32])
	s := new(big.Int).SetBytes(cipherBytes[32:])
	return sm2.VerifyByRS(pub, DEFAULT_UID, plainBytes, r, s)
}
