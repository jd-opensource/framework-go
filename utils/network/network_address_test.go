package network

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/*
 * Author: imuge
 * Date: 2020/6/11 上午9:20
 */

func TestFromBytes(t *testing.T) {
	a1 := NewAddress("localhost", 1000, false)
	a2 := FromBytes(a1.ToBytes())
	require.Equal(t, a1.Host, a2.Host)
	require.Equal(t, a1.Port, a2.Port)
	require.Equal(t, a1.Secure, a2.Secure)
}