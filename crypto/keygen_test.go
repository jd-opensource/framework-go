package crypto

import (
	"framework-go/crypto/classic"
	"framework-go/crypto/framework"
	"framework-go/utils/sha"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEncodeDecodePubKey(t *testing.T) {
	function := GetCryptoFunctionByName(classic.ED25519_ALGORITHM.Name)
	f1 := function.(framework.AsymmetricKeypairGenerator)
	keypair := f1.GenerateKeypair()

	base58PubKey := EncodePubKey(keypair.PubKey)
	decPubKey := DecodePubKey(base58PubKey)
	require.NotNil(t, decPubKey)
	require.Equal(t, base58PubKey, EncodePubKey(decPubKey))
}

func TestEncodeDecodePrivKey(t *testing.T) {
	function := GetCryptoFunctionByName(classic.ED25519_ALGORITHM.Name)
	f1 := function.(framework.AsymmetricKeypairGenerator)
	keypair := f1.GenerateKeypair()
	pwd := []byte("abc")

	base58PrivKey := EncodePrivKey(keypair.PrivKey, sha.Sha256(pwd))
	decPrivKey := DecodePrivKey(base58PrivKey, sha.Sha256(pwd))
	require.NotNil(t, decPrivKey)
	require.Equal(t, base58PrivKey, EncodePrivKey(decPrivKey, sha.Sha256(pwd)))
}

func TestEncodeDecodePrivKeyWithRawPwd(t *testing.T) {
	function := GetCryptoFunctionByName(classic.ED25519_ALGORITHM.Name)
	f1 := function.(framework.AsymmetricKeypairGenerator)
	keypair := f1.GenerateKeypair()

	pwd := "abc"
	base58PrivKey := EncodePrivKeyWithRawPwd(keypair.PrivKey, pwd)
	decPrivKey := DecodePrivKeyWithRawPwd(base58PrivKey, pwd)
	require.NotNil(t, decPrivKey)
	require.Equal(t, base58PrivKey, EncodePrivKeyWithRawPwd(decPrivKey, pwd))
}
