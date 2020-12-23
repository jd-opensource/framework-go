package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
	"github.com/blockchain-jd-com/framework-go/utils/sha"
	"math/big"
)

/**
 * @Author: imuge
 * @Date: 2020/4/30 9:52 上午
 */

func GenerateKeyPair() *rsa.PrivateKey {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	return priv
}

func GenerateKeyPairWithSeed(seed []byte) *rsa.PrivateKey {
	panic("not support yet")
}

func PubKeyToBytes(pub *rsa.PublicKey) []byte {
	rawBytes := append(pub.N.Bytes(), bytes.IntToBytes(pub.E)[1:]...)
	return rawBytes
}

func BytesToPubKey(b []byte) (pub *rsa.PublicKey) {
	n := &big.Int{}
	n.SetBytes(b[:256])
	return &rsa.PublicKey{
		N: n,
		E: bytes.ToInt(b[256:]),
	}
}

func PrivKeyToBytes(priv *rsa.PrivateKey) []byte {
	N := priv.N
	E := priv.PublicKey.E
	D := priv.D
	P := priv.Primes[0]
	Q := priv.Primes[1]
	Dp := priv.Precomputed.Dp
	Dq := priv.Precomputed.Dq
	Qinv := priv.Precomputed.Qinv
	rawBytes := append(N.Bytes(), bytes.IntToBytes(E)[1:]...)
	rawBytes = append(rawBytes, D.Bytes()...)
	rawBytes = append(rawBytes, P.Bytes()...)
	rawBytes = append(rawBytes, Q.Bytes()...)
	rawBytes = append(rawBytes, Dp.Bytes()...)
	rawBytes = append(rawBytes, Dq.Bytes()...)
	rawBytes = append(rawBytes, Qinv.Bytes()...)
	return rawBytes
}

func BytesToPrivKey(b []byte) *rsa.PrivateKey {
	n := &big.Int{}
	n.SetBytes(b[:256])
	e := bytes.ToInt(b[256 : 256+3])
	d := &big.Int{}
	d.SetBytes(b[256+3 : 256+3+256])
	p := &big.Int{}
	p.SetBytes(b[256+3+256 : 256+3+256+128])
	q := &big.Int{}
	q.SetBytes(b[256+3+256+128 : 256+3+256+128+128])
	dp := &big.Int{}
	dp.SetBytes(b[256+3+256+128+128 : 256+3+256+128+128+128])
	dq := &big.Int{}
	dq.SetBytes(b[256+3+256+128+128+128 : 256+3+256+128+128+128+128])
	qInv := &big.Int{}
	qInv.SetBytes(b[256+3+256+128+128+128+128 : 256+3+256+128+128+128+128+128])

	priv := &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: n,
			E: e,
		},
		D: d,
		Primes: []*big.Int{
			p,
			q,
		},
		Precomputed: rsa.PrecomputedValues{
			Dp:   dp,
			Dq:   dq,
			Qinv: qInv,
		},
	}
	return priv
}

func Encrypt(pub *rsa.PublicKey, plainBytes []byte) []byte {
	encrypt, err := rsa.EncryptPKCS1v15(rand.Reader, pub, plainBytes)
	if err != nil {
		panic(err)
	}

	return encrypt
}

func Decrypt(priv *rsa.PrivateKey, cipherBytes []byte) []byte {
	decrypt, err := rsa.DecryptPKCS1v15(rand.Reader, priv, cipherBytes)
	if err != nil {
		panic(err)
	}

	return decrypt
}

func Sign(priv *rsa.PrivateKey, plainBytes []byte) []byte {
	hashed := sha.Sha256(plainBytes)
	signed, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed[:])
	if err != nil {
		panic(err)
	}

	return signed
}

func Verify(pub *rsa.PublicKey, plainBytes, cipherBytes []byte) bool {
	hashed := sha256.Sum256(plainBytes)
	err := rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], cipherBytes)
	if err != nil {
		return false
	}

	return true
}
