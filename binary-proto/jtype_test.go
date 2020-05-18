package binary_proto

import (
	"fmt"
	"framework-go/utils/base58"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEncode(t *testing.T) {
	bytes, err := Cdc.Encode(NewJType())
	require.Nil(t, err)
	fmt.Println(base58.Encode(bytes))
}
