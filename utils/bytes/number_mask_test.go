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

func TestLongNumberMask(t *testing.T) {
	require.True(t, NUMBERMASK_TINY.MAX_HEADER_LENGTH == 1)
	require.EqualValues(t, 256, NUMBERMASK_TINY.GetBoundarySize(1))

	require.True(t, NUMBERMASK_SHORT.MAX_HEADER_LENGTH == 2)
	require.EqualValues(t, 128, NUMBERMASK_SHORT.GetBoundarySize(1))
	require.EqualValues(t, 32768, NUMBERMASK_SHORT.GetBoundarySize(2))

	require.True(t, NUMBERMASK_NORMAL.MAX_HEADER_LENGTH == 4)
	require.EqualValues(t, int32(64), NUMBERMASK_NORMAL.GetBoundarySize(1))
	require.EqualValues(t, int32(16384), NUMBERMASK_NORMAL.GetBoundarySize(2))
	require.EqualValues(t, int32(4194304), NUMBERMASK_NORMAL.GetBoundarySize(3))
	require.EqualValues(t, int32(1073741824), NUMBERMASK_NORMAL.GetBoundarySize(4))

	require.True(t, NUMBERMASK_LONG.MAX_HEADER_LENGTH == 8)
	require.EqualValues(t, int64(32), NUMBERMASK_LONG.GetBoundarySize(1))
	require.EqualValues(t, int64(8192), NUMBERMASK_LONG.GetBoundarySize(2))
	require.EqualValues(t, int64(2097152), NUMBERMASK_LONG.GetBoundarySize(3))
	require.EqualValues(t, int64(536870912), NUMBERMASK_LONG.GetBoundarySize(4))
	require.EqualValues(t, int64(137438953472), NUMBERMASK_LONG.GetBoundarySize(5))
	require.EqualValues(t, int64(35184372088832), NUMBERMASK_LONG.GetBoundarySize(6))
	require.EqualValues(t, int64(9007199254740992), NUMBERMASK_LONG.GetBoundarySize(7))
	require.EqualValues(t, int64(2305843009213693952), NUMBERMASK_LONG.GetBoundarySize(8))

	testLong(t, 0, 1)
	testLong(t, 17, 1)
	testLong(t, 31, 1)

	testLong(t, 32, 2)
	testLong(t, 57, 2)
	testLong(t, 8191, 2)

	testLong(t, 8192, 3)
	testLong(t, 103200, 3)
	testLong(t, 2000320, 3)
	testLong(t, 2097151, 3)

	testLong(t, 2097152, 4)
	testLong(t, 403332200, 4)
	testLong(t, 536870911, 4)

	testLong(t, 536870912, 5)
	testLong(t, 103388332200, 5)
	testLong(t, 137438953471, 5)

	testLong(t, 137438953472, 6)
	testLong(t, 25243388332201, 6)
	testLong(t, 35184372088831, 6)

	testLong(t, 35184372088832, 7)
	testLong(t, 7985243388332201, 7)
	testLong(t, 9007199254740991, 7)

	testLong(t, 9007199254740992, 8)
	testLong(t, 1985932243388332201, 8)
	testLong(t, 2305843009213693951, 8)
}

func testLong(t *testing.T, number int64, expectedLength int32) {
	bytes := NUMBERMASK_LONG.GenerateMask(number)

	length := NUMBERMASK_LONG.GetMaskLength(number)

	resolvedLen, err := NUMBERMASK_LONG.ResolveMaskLength(bytes[0])
	require.Nil(t, err)
	resolvedNumber, err := NUMBERMASK_LONG.ResolveMaskedNumber(bytes)
	require.Nil(t, err)
	require.Equal(t, number, resolvedNumber)
	require.Equal(t, int(expectedLength), len(bytes))
	require.Equal(t, expectedLength, length)
	require.Equal(t, expectedLength, resolvedLen)
}

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
