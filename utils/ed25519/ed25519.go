package ed25519

import (
	"crypto/rand"
	"github.com/blockchain-jd-com/framework-go/utils/random"
	"golang.org/x/crypto/ed25519"
)

/**
 * @Author: imuge
 * @Date: 2020/5/1 10:52 下午
 */

func GenerateKeyPair() (ed25519.PublicKey, []byte) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	return pub, priv.Seed()
}

func GenerateKeyPairWithSeed(seed []byte) (ed25519.PublicKey, []byte) {
	pub, priv, err := ed25519.GenerateKey(random.NewHashSecureRandom(seed, 32, random.Sha256))
	if err != nil {
		panic(err)
	}

	return pub, priv.Seed()
}

func RetrievePubKey(seed []byte) ed25519.PublicKey {
	return ed25519.PublicKey(ed25519.NewKeyFromSeed(seed)[32:])
}

func Sign(seed []byte, plainBytes []byte) []byte {
	priv := ed25519.NewKeyFromSeed(seed)
	return ed25519.Sign(priv, plainBytes)
}

func Verify(pub ed25519.PublicKey, plainBytes, cipherBytes []byte) bool {
	return ed25519.Verify(pub, plainBytes, cipherBytes)
}
