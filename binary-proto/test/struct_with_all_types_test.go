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
		"111d1Vm4bW1fmUDZyB1kprs8R5rSkVCJaPwa1CvsNuRStdwzzhKggkM8pJgpJw3Mivfkgddv5sT2wjCqLLgM7nRonjqK5amk91Yj14QThV7X2RzDFZQgG9PLd5yv1tYqqnbHrWXqrBuudD7zBQfCaTWiZgRepvgBnbG7oTPcQKTcznRzQ2Gpydey6EXBpqpfS",
		base58.Encode(bytes))
}
