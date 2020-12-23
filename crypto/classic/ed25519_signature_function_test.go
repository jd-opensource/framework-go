package classic

import (
	"github.com/blockchain-jd-com/framework-go/utils/random"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestED25519SignatureFunction_GenerateKeypairWithSeed(t *testing.T) {
	// 小于32长度seed
	for i := 0; i < 32; i++ {
		_, err := ED25519.GenerateKeypairWithSeed([]byte(random.RandString(i)))
		require.NotNil(t, err)
	}

	// 正常生成
	seed := []byte("abcdefghijklmnopqrstuvwxyz123456")
	keypair, err := ED25519.GenerateKeypairWithSeed(seed)
	require.Nil(t, err)

	// generated from jd chain
	require.Equal(t, "7VeRF4s9UCtANKRFJ9f9DB3iWTfg14KBrecNiBgWa6zF7t5F", keypair.PubKey.ToBase58())
}
