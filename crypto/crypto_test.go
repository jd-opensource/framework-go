package crypto

import (
	"fmt"
	"framework-go/crypto/framework"
	"framework-go/utils/base58"
	"github.com/stretchr/testify/require"
	"testing"
)

/**
 * @Author: imuge
 * @Date: 2020/4/29 7:06 下午
 */

func TestSHA256(t *testing.T) {
	data := []byte("imuge")
	function := GetHashFunctionByName("SHA256")
	hash := function.Hash(data)
	fmt.Println("hash: " + hash.ToBase58())
	require.True(t, function.Verify(hash, data))

	// hash from JD Cahin
	jdHash, _ := base58.Decode("j5vkSRxmUjJzo9KBX79cTMRwD8Aw3J7Ke2JnPzS1eq4fH1")
	require.True(t, function.Verify(framework.ParseHashDigest(jdHash), data))
}
