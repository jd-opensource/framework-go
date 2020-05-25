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
		"111HWkwFfXMBaPBx2XAv4LKoGTXngcj8dG7boBohWaFGvMiX9MSVuk876e4ZTVxzfeRmovv1g232FzyVj4szMpDHSVPCrgm6HnqiW1CGC4gRvvcwU5JBz9YM1ohr6SY8evP4ghC8mZa2eRoo5vR8V63FSps23vxD3mHWpGp3hGazd6Pgix6y1xYUcECVT9RD8dfixtL7UZaJzR6hjbecqqVoWSH4ZSEGqsU7BJPDzsMdpdv1PLHdQGFyToXn",
		base58.Encode(bytes))
}

func TestDecode(t *testing.T) {
	origin := NewStructWithAllTypes()
	bytes, err := binary_proto.Cdc.Encode(origin)
	require.Nil(t, err)
	obj, err := binary_proto.Cdc.Decode(bytes)
	require.Nil(t, err)
	contract := obj.(StructWithAllTypes)
	require.True(t, origin.Equals(contract))
}

func TestVersion(t *testing.T) {
	cdc := binary_proto.NewCodec()
	contract1 := RefContract{}
	cdc.RegisterContract(contract1)
	require.Equal(t, int64(-4451409565821993051), cdc.VersionMap[contract1.Code()])

	contract2 := RefGeneric{}
	cdc.RegisterContract(contract2)
	require.Equal(t, int64(-2039914840885289964), cdc.VersionMap[contract2.Code()])

	cdc.RegisterEnum(ONE)

	contract3 := StructWithAllTypes{}
	cdc.RegisterContract(contract3)
	require.Equal(t, int64(-4218456988248628983), cdc.VersionMap[contract3.Code()])
}
