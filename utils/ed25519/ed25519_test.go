package ed25519

import (
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

/**
 * @Author: imuge
 * @Date: 2020/5/1 11:37 下午
 */

func TestSign(t *testing.T) {
	pub, seed := GenerateKeyPair()
	plainBytes := []byte("imuge")
	sign := Sign(seed, plainBytes)
	require.True(t, Verify(pub, plainBytes, sign))

	require.True(t, bytes.Equals(pub, RetrievePubKey(seed)))
}
