package bytes

import (
	"github.com/blockchain-jd-com/framework-go/utils/base58"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
 * Author: imuge
 * Date: 2020/5/21 下午7:43
 */

func TestMaskTiny(t *testing.T) {
	buf := NUMBERMASK_TINY.WriteMask(8)
	require.Equal(t, "9", base58.Encode(buf))
}

func TestMaskShort(t *testing.T) {
	buf := NUMBERMASK_TINY.WriteMask(16)
	require.Equal(t, "H", base58.Encode(buf))
}

func TestMaskNormal(t *testing.T) {
	buf := NUMBERMASK_NORMAL.WriteMask(32)
	require.Equal(t, "Z", base58.Encode(buf))
}

func TestMaskLong(t *testing.T) {
	buf := NUMBERMASK_LONG.WriteMask(64)
	require.Equal(t, "3TM", base58.Encode(buf))
}
