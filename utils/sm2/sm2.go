package sm2

import (
	"crypto/rand"
	"github.com/ZZMarquis/gm/sm2"
	"github.com/ZZMarquis/gm/sm3"
	"math/big"
)

/**
 * @Author: imuge
 * @Date: 2020/5/2 8:47 上午
 */

var DEFAULT_UID = []byte("1234567812345678")

func GenerateKeyPair() (*sm2.PrivateKey, *sm2.PublicKey, error) {
	priv, pub, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	return priv, pub, nil
}

func GenerateKeyPairWithSeed(seed []byte) (*sm2.PrivateKey, *sm2.PublicKey, error) {
	panic("not support yet")
}

func sm3Hash(d []byte) []byte {
	s := sm3.Sum(d)
	return s[:]
}

func PubKeyToBytes(pub *sm2.PublicKey) []byte {
	return append([]byte{0x04}, append(pub.X.Bytes(), pub.Y.Bytes()...)...)
}

func BytesToPubKey(b []byte) (*sm2.PublicKey, error) {
	pub, err := sm2.RawBytesToPublicKey(b[1:])
	if err != nil {
		return nil, err
	}

	return pub, nil
}

func RawBytesToPubKey(b []byte) (*sm2.PublicKey, error) {
	pub, err := sm2.RawBytesToPublicKey(b)
	if err != nil {
		return nil, err
	}

	return pub, nil
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

func BytesToPrivKey(b []byte) (*sm2.PrivateKey, error) {
	priv, err := sm2.RawBytesToPrivateKey(b)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func Encrypt(pub *sm2.PublicKey, plainBytes []byte) ([]byte, error) {
	encrypt, err := sm2.Encrypt(pub, plainBytes, sm2.C1C3C2)
	if err != nil {
		return nil, err
	}

	return encrypt, nil
}

func Decrypt(priv *sm2.PrivateKey, cipherBytes []byte) ([]byte, error) {
	decrypt, err := sm2.Decrypt(priv, cipherBytes, sm2.C1C3C2)
	if err != nil {
		return nil, err
	}

	return decrypt, nil
}

func Sign(priv *sm2.PrivateKey, plainBytes []byte) ([]byte, error) {
	r, s, err := sm2.SignToRS(priv, DEFAULT_UID, plainBytes)
	if err != nil {
		return nil, err
	}

	return append(r.Bytes(), s.Bytes()...), nil
}

func Verify(pub *sm2.PublicKey, plainBytes, cipherBytes []byte) bool {
	r := new(big.Int).SetBytes(cipherBytes[0:32])
	s := new(big.Int).SetBytes(cipherBytes[32:])
	return sm2.VerifyByRS(pub, DEFAULT_UID, plainBytes, r, s)
}
