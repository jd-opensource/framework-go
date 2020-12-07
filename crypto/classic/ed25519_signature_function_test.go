package classic

import (
	"github.com/blockchain-jd-com/framework-go/utils/random"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestED25519SignatureFunction_GenerateKeypairWithSeed(t *testing.T) {
	// 小于32长度seed panic
	for i := 0; i < 32; i++ {
		b := false
		defer func() {
			if err := recover(); err != nil {
				b = true
			}
		}()
		require.True(t, b)
	}
	for i := 0; i < 32; i++ {
		b := false
		defer func() {
			if err := recover(); err != nil {
				b = true
			}
		}()
		ED25519.GenerateKeypairWithSeed([]byte(random.RandString(i)))
		require.True(t, b)
	}

}
