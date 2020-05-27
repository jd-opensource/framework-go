package test

import (
	"framework-go/crypto/classic"
	"framework-go/ledger_model"
	"framework-go/sdk"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
 * Author: imuge
 * Date: 2020/5/26 上午11:29
 */

func TestRegisterUser(t *testing.T) {
	// 交易内容
	user1 := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)
	user2 := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)

	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
	service := serviceFactory.GetBlockchainService()

	ledgerHashs, err := service.GetLedgerHashs()
	require.Nil(t, err)
	// 在本地定义注册账号的 TX；
	txTemp := service.NewTransaction(ledgerHashs[0])

	// 注册
	txTemp.Users().Register(user1.GetIdentity())
	txTemp.Users().Register(user2.GetIdentity())
	txTemp.Security().Roles().Configure("MANAGER").EnableLedgerPermission(ledger_model.REGISTER_USER).EnableTransactionPermission(ledger_model.CONTRACT_OPERATION)
	txTemp.Security().Authorziations().ForUser([][]byte{user1.GetAddress(), user2.GetAddress()}).Authorize("MANAGER")

	// TX 准备就绪；
	prepTx := txTemp.Prepare()

	// 使用私钥进行签名；
	prepTx.Sign(NODE_KEY.AsymmetricKeypair)

	// 提交交易；
	resp, err := prepTx.Commit()
	require.Nil(t, err)
	require.True(t, resp.Success)

}

func TestDataAccount(t *testing.T) {
	// 交易内容
	user := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)

	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
	service := serviceFactory.GetBlockchainService()

	ledgerHashs, err := service.GetLedgerHashs()
	require.Nil(t, err)
	// 在本地定义注册账号的 TX；
	txTemp := service.NewTransaction(ledgerHashs[0])

	// 注册
	txTemp.Users().Register(user.GetIdentity())
	// 注册
	txTemp.DataAccounts().Register(user.GetIdentity())
	txTemp.DataAccount(user.GetAddress()).SetText("imuge", "nice", -1).SetText("xiuxiu", "nice", -1)
	txTemp.DataAccount(user.GetAddress()).SetText("wang", "nice", -1)

	// TX 准备就绪；
	prepTx := txTemp.Prepare()

	// 使用私钥进行签名；
	prepTx.Sign(NODE_KEY.AsymmetricKeypair)

	// 提交交易；
	resp, err := prepTx.Commit()
	require.Nil(t, err)
	require.True(t, resp.Success)

}
