package sdk

import (
	"fmt"
	"framework-go/crypto"
	"framework-go/ledger_model"
	"framework-go/utils/base58"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
 * Author: imuge
 * Date: 2020/5/25 下午6:16
 */

func TestQuery(t *testing.T) {
	nodePrivKey := crypto.DecodePrivKey("177gjzfT217HTByHAe2FEhirUj8hVYyNL4HfJFvdE5KQ52aDPa75xbuBNior2ia2sv3EXqG", base58.MustDecode("8EjkXVSTxMFjCvNNsTo8RBMDEVQmk7gYkW4SCDuvdsBG"))
	nodePubKey := crypto.DecodePubKey("3snPdw7i7PYQzfmYsrrmqWsM9RefGSobRoLa8vEyFbVazfdkvPuF1J")
	factory := Connect("localhost", 8081, false, ledger_model.NewBlockchainKeypair(nodePubKey, nodePrivKey))

	// /ledgers
	ledgers, err := factory.GetBlockchainService().GetLedgerHashs()
	require.Nil(t, err)
	for _, ledger := range ledgers {
		fmt.Println(ledger.ToBase58())
	}

	// /ledgers/j5uhqzPUtc3DSadTNPUG4sXxkjXC56oWBmqAdJbtq7MNNj
	ledger, err := factory.GetBlockchainService().GetLedger(ledgers[0])
	require.Nil(t, err)
	require.Equal(t, ledgers[0], ledger.Hash)

	// /ledgers/j5uhqzPUtc3DSadTNPUG4sXxkjXC56oWBmqAdJbtq7MNNj/
	_, err = factory.GetBlockchainService().GetLedgerAdminInfo(ledgers[0])
	require.Nil(t, err)
}
