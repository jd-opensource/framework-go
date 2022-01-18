package crypto

import (
	"fmt"
	"github.com/blockchain-jd-com/framework-go/crypto/classic"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/utils/base58"
	"github.com/blockchain-jd-com/framework-go/utils/sha"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEncodeDecodePubKey(t *testing.T) {
	for _, am := range algorithms {
		fmt.Println(am.Name + ": ")
		function := GetCryptoFunctionByName(am.Name)
		f1, ok := function.(framework.AsymmetricKeypairGenerator)
		if !ok {
			continue
		}
		keypair, err := f1.GenerateKeypair()
		require.Nil(t, err)
		base58PubKey := EncodePubKey(keypair.PubKey)
		decPubKey, err := DecodePubKey(base58PubKey)
		require.Nil(t, err)
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
		keypair, err := f1.GenerateKeypair()
		require.Nil(t, err)
		pwd := []byte("abc")

		base58PrivKey, err := EncodePrivKey(keypair.PrivKey, sha.Sha256(pwd))
		require.Nil(t, err)
		fmt.Println(base58PrivKey)
		decPrivKey, err := DecodePrivKey(base58PrivKey, sha.Sha256(pwd))
		require.Nil(t, err)
		require.NotNil(t, decPrivKey)
		keyEncoded, err := EncodePrivKey(decPrivKey, sha.Sha256(pwd))
		require.Nil(t, err)
		require.Equal(t, base58PrivKey, keyEncoded)
	}
}

func TestEncodeDecodePrivKeyWithRawPwd(t *testing.T) {
	for _, am := range algorithms {
		function := GetCryptoFunctionByName(am.Name)
		f1, ok := function.(framework.AsymmetricKeypairGenerator)
		if !ok {
			continue
		}
		keypair, err := f1.GenerateKeypair()
		require.Nil(t, err)
		pwd := "abc"
		base58PrivKey, err := EncodePrivKeyWithRawPwd(keypair.PrivKey, pwd)
		require.Nil(t, err)
		decPrivKey, err := DecodePrivKeyWithRawPwd(base58PrivKey, pwd)
		require.Nil(t, err)
		require.NotNil(t, decPrivKey)
		encodedKey, err := EncodePrivKeyWithRawPwd(decPrivKey, pwd)
		require.Nil(t, err)
		require.Equal(t, base58PrivKey, encodedKey)
	}
}

// 从公钥生成地址
func TestGenerateAddress(t *testing.T) {
	function := GetCryptoFunctionByName(classic.ED25519_ALGORITHM.Name).(framework.AsymmetricKeypairGenerator)
	keyPair, err := function.GenerateKeypair()
	require.Nil(t, err)
	address := framework.GenerateAddress(keyPair.PubKey)
	fmt.Println(base58.Encode(address))
}
