package ripemd160

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/**
 * @Author: imuge
 * @Date: 2020/5/6 10:10 上午
 */

func TestHash(t *testing.T) {
	bytes := []byte("imuge")
	require.Equal(t, 20, len(Hash(bytes)))
	require.Equal(t, Hash(bytes), Hash(bytes))
}
