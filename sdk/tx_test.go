package sdk

import (
	"fmt"
	binary_proto "framework-go/binary-proto"
	"framework-go/crypto"
	"framework-go/crypto/classic"
	"framework-go/crypto/framework"
	"framework-go/ledger_model"
	"framework-go/utils/base58"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

/*
 * Author: imuge
 * Date: 2020/5/26 上午11:29
 */

func TestRegisterUser(t *testing.T) {
	// 账本哈兮
	ledger, _ := base58.Decode("j5uhqzPUtc3DSadTNPUG4sXxkjXC56oWBmqAdJbtq7MNNj")

	nodePrivKey := crypto.DecodePrivKey("177gjzfT217HTByHAe2FEhirUj8hVYyNL4HfJFvdE5KQ52aDPa75xbuBNior2ia2sv3EXqG", base58.MustDecode("8EjkXVSTxMFjCvNNsTo8RBMDEVQmk7gYkW4SCDuvdsBG"))
	nodePubKey := crypto.DecodePubKey("3snPdw7i7PYQzfmYsrrmqWsM9RefGSobRoLa8vEyFbVazfdkvPuF1J")

	// 交易内容
	keyPair := crypto.GetCryptoFunctionByName(classic.ED25519_ALGORITHM.Name).(framework.AsymmetricKeypairGenerator).GenerateKeypair()
	contentBody := ledger_model.TransactionContentBody{
		LedgerHash: ledger,
		Operations: []binary_proto.DataContract{
			ledger_model.UserRegisterOperation{
				UserID: ledger_model.BlockchainIdentity{
					Address: framework.GenerateAddress(keyPair.PubKey),
					PubKey:  keyPair.PubKey.ToBytes(),
				},
			},
		},
		Timestamp: time.Now().Unix(),
	}
	content := ledger_model.TransactionContent{
		TransactionContentBody: contentBody,
		Hash:       nil,
	}
	contentBytes, err := binary_proto.Cdc.Encode(contentBody)
	if err != nil {
		panic(err)
	}
	hasher := crypto.GetHashFunctionByName(classic.SHA256_ALGORITHM.Name)
	hashDigest := hasher.Hash(contentBytes)
	content.Hash = hashDigest.ToBytes()

	signer := crypto.GetSignatureFunctionByCode(nodePrivKey.GetAlgorithm())
	// 网关签名
	endSign := signer.Sign(nodePrivKey, content.Hash)
	txMessage := ledger_model.TransactionRequest{
		TransactionContent: content,
		EndpointSignatures: []ledger_model.DigitalSignature{
			{
				nodePubKey.ToBytes(),
				endSign.ToBytes(),
			},
		},
	}

	require.True(t, signer.Verify(nodePubKey, content.Hash, endSign))

	reqBytes, err := binary_proto.Cdc.Encode(txMessage)
	if err != nil {
		panic(err)
	}
	reqHash := hasher.Hash(reqBytes)
	txMessage.Hash = reqHash.ToBytes()

	msg, _ := binary_proto.Cdc.Encode(txMessage)

	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/bin-obj").
		SetBody(msg).
		Post("http://localhost:8081/rpc/tx")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("Error      :", err)
	fmt.Println("Status Code:", resp.StatusCode())
	fmt.Println("Status     :", resp.Status())
	fmt.Println("Proto      :", resp.Proto())
	fmt.Println("Time       :", resp.Time())
	fmt.Println("Received At:", resp.ReceivedAt())
	fmt.Println("Body       :\n", resp)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("DNSLookup    :", ti.DNSLookup)
	fmt.Println("ConnTime     :", ti.ConnTime)
	fmt.Println("TCPConnTime  :", ti.TCPConnTime)
	fmt.Println("TLSHandshake :", ti.TLSHandshake)
	fmt.Println("ServerTime   :", ti.ServerTime)
	fmt.Println("ResponseTime :", ti.ResponseTime)
	fmt.Println("TotalTime    :", ti.TotalTime)
	fmt.Println("IsConnReused :", ti.IsConnReused)
	fmt.Println("IsConnWasIdle:", ti.IsConnWasIdle)
	fmt.Println("ConnIdleTime :", ti.ConnIdleTime)
}
