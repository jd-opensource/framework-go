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
