package test

import (
	"framework-go/sdk"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
 * Author: imuge
 * Date: 2020/5/25 下午6:16
 */

func TestQuery(t *testing.T) {
	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)

	ledgers, err := serviceFactory.GetBlockchainService().GetLedgerHashs()
	require.Nil(t, err)

	_, err = serviceFactory.GetBlockchainService().GetBlockByHeight(ledgers[0], 1)
	require.Nil(t, err)
}
