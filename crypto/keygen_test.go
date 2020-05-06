package crypto

import (
	"fmt"
	"framework-go/crypto/framework"
	"framework-go/utils/sha"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEncodeDecodePubKey(t *testing.T) {
	for _, am := range algorithms {
		function := GetCryptoFunctionByName(am.Name)
		f1, ok := function.(framework.AsymmetricKeypairGenerator)
		if !ok {
			continue
		}
		keypair := f1.GenerateKeypair()

		base58PubKey := EncodePubKey(keypair.PubKey)
		decPubKey := DecodePubKey(base58PubKey)
		fmt.Println(base58PubKey)
		require.NotNil(t, decPubKey)
		require.Equal(t, base58PubKey, EncodePubKey(decPubKey))
	}
}

func TestEncodeDecodePrivKey(t *testing.T) {
	for _, am := range algorithms {
		function := GetCryptoFunctionByName(am.Name)
		f1, ok := function.(framework.AsymmetricKeypairGenerator)
		if !ok {
			continue
		}
		keypair := f1.GenerateKeypair()
		pwd := []byte("abc")

		base58PrivKey := EncodePrivKey(keypair.PrivKey, sha.Sha256(pwd))
		fmt.Println(base58PrivKey)
		decPrivKey := DecodePrivKey(base58PrivKey, sha.Sha256(pwd))
		require.NotNil(t, decPrivKey)
		require.Equal(t, base58PrivKey, EncodePrivKey(decPrivKey, sha.Sha256(pwd)))
	}
}

func TestEncodeDecodePrivKeyWithRawPwd(t *testing.T) {
	for _, am := range algorithms {
		function := GetCryptoFunctionByName(am.Name)
		f1, ok := function.(framework.AsymmetricKeypairGenerator)
		if !ok {
			continue
		}
		keypair := f1.GenerateKeypair()

		pwd := "abc"
		base58PrivKey := EncodePrivKeyWithRawPwd(keypair.PrivKey, pwd)
		decPrivKey := DecodePrivKeyWithRawPwd(base58PrivKey, pwd)
		require.NotNil(t, decPrivKey)
		require.Equal(t, base58PrivKey, EncodePrivKeyWithRawPwd(decPrivKey, pwd))
	}
}
