package test

import (
	"framework-go/binary-proto"
	"framework-go/utils/base58"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEncode(t *testing.T) {
	bytes, err := binary_proto.Cdc.Encode(NewStructWithAllTypes())
	require.Nil(t, err)
	require.Equal(t,
		"111Dt763NCtQ6JQaq8fzWFC1upKUTWhiCBcosnccHR88f5uzcvF6V4QckNiQTUzNCB2qVNWooUtAmapa8mE4XvQksV57AuNyZDao6FqvsSosMfDm5KQJ4e5qd7jvuUZGkAmgVARdkPbrEnYdiDsioyQLLgC6dync5oEvRbsrFPvwUUFxyrP8RSC9ccsSavfcSvT51Eo5QtqwjaUgpK2droKYHL1h19VxNg2QY7wWqquGrrZaXaPoC2duZPQt",
		base58.Encode(bytes))
}

func TestDecode(t *testing.T) {
	bytes, err := binary_proto.Cdc.Encode(NewStructWithAllTypes())
	require.Nil(t, err)
	obj, err := binary_proto.Cdc.Decode(bytes)
	require.Nil(t, err)
	contract := obj.(StructWithAllTypes)
	require.Equal(t, "bytes", string(contract.Bytes))
	require.Equal(t, int8(8), contract.I8)
	require.Equal(t, int16(16), contract.I16)
	require.Equal(t, int32(32), contract.I32)
	require.Equal(t, int64(64), contract.I64)
	require.Equal(t, "text", contract.Text)
	require.Equal(t, int8(8), contract.I8m)
	require.Equal(t, int16(16), contract.I16m)
	require.Equal(t, int32(32), contract.I32m)
	require.Equal(t, int64(64), contract.I64m)
	require.True(t, contract.Bool)
}
