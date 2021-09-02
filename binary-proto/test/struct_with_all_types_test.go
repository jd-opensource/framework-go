package test

import (
	"fmt"
	"github.com/blockchain-jd-com/framework-go/binary-proto"
	"github.com/blockchain-jd-com/framework-go/ledger_model"
	"github.com/blockchain-jd-com/framework-go/utils/base58"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEncode(t *testing.T) {
	bytes, err := binary_proto.NewCodec().Encode(NewStructWithAllTypes())
	require.Nil(t, err)
	require.Equal(t,
		"1112RQsc83DvDVJomSUvEPnkBCq1qVQWMKfWYyx2ERvu9oeypyDLzivbM3Rg5icMKzFf2cEBm7rrcYtDomYAnWF4G8DcLq1gtd186eWaSNVWdQzTZkErc4ky53VxMMizsYB2msM8Qsq4qBeNeVbhDa6fSzynbzuM8PXWrCqX6cg9SDJG9J8nEyMWDMGGW41kx49uT9x2A9HNyUA1NrsAx",
		base58.Encode(bytes))
}

func TestDecode(t *testing.T) {
	ll := ledger_model.DigitalSignatureBody{}
	fmt.Println(ll.PubKey)
	decode := base58.MustDecode("11Gb7H97yLwfptkD5iqwj6DG8UrSE8LGKeZKf894hXiWKwZYLS9ro7RwKG2sfo5ZVCY1UqMUsGD18AQBXWzVoENyfdMy6h27QRZKiTgDLUKWAF4EGbqriQM2sdAR9sFN3mBANJ3xniJPRkpctxyE1ihsaMPtf34Jj6PCaAd5FuMaf6aRec29ey6Wb8JyHnPHQ9wXrvZoYTnrjdoTRQnEtnHAD2H4Q")
	obj, err := binary_proto.NewCodec().Decode(decode)
	require.Nil(t, err)
	contract := obj.(ledger_model.TransactionContent)
	fmt.Println(contract.Timestamp)
	operation := contract.Operations[0].(ledger_model.DataAccountRegisterOperation)
	fmt.Println(operation.AddressSignature)

	encode, err := binary_proto.NewCodec().Encode(contract)
	require.Nil(t, err)
	fmt.Println(len(encode))
}

func TestVersion(t *testing.T) {
	cdc := binary_proto.NewCodec()
	contract1 := RefContract{}
	binary_proto.RegisterContract(contract1)
	cdc.CalculateVersion(contract1)
	require.Equal(t, int64(-4451409565821993051), cdc.VersionMap[contract1.ContractCode()])

	contract2 := RefGeneric{}
	binary_proto.RegisterContract(contract2)
	cdc.CalculateVersion(contract2)
	require.Equal(t, int64(-2039914840885289964), cdc.VersionMap[contract2.ContractCode()])

	binary_proto.RegisterEnum(ONE)

	contract3 := StructWithAllTypes{}
	binary_proto.RegisterContract(contract3)
	cdc.CalculateVersion(contract3)
	require.Equal(t, int64(-4218456988248628983), cdc.VersionMap[contract3.ContractCode()])
}
