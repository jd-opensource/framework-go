package sha

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSha128(t *testing.T) {
	require.Equal(t, 16, len(Sha128([]byte("abc"))))
}

func TestSha256(t *testing.T) {
	require.Equal(t, 32, len(Sha256([]byte("abc"))))
}