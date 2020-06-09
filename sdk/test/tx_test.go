package test

import (
	"framework-go/crypto/classic"
	"framework-go/ledger_model"
	"framework-go/sdk"
	"framework-go/utils/bytes"
	"framework-go/utils/network"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

/*
 * Author: imuge
 * Date: 2020/5/26 上午11:29
 */

func TestRegisterUser(t *testing.T) {
	// 生成公私钥对
	user1 := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)
	user2 := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)

	// 连接网关，获取节点服务
	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
	service := serviceFactory.GetBlockchainService()

	// 获取账本信息
	ledgerHashs, err := service.GetLedgerHashs()
	require.Nil(t, err)

	// 创建交易
	txTemp := service.NewTransaction(ledgerHashs[0])

	// 注册用户
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
	// 生成公私钥对
	user := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)

	// 连接网关，获取节点服务
	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
	service := serviceFactory.GetBlockchainService()

	// 获取账本信息
	ledgerHashs, err := service.GetLedgerHashs()
	require.Nil(t, err)

	// 创建交易
	txTemp := service.NewTransaction(ledgerHashs[0])

	// 注册数据账户
	txTemp.DataAccounts().Register(user.GetIdentity())
	// 插入数据
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

func TestContract(t *testing.T) {
	// 生成公私钥对
	user := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)

	// 连接网关，获取节点服务
	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
	service := serviceFactory.GetBlockchainService()

	// 获取账本信息
	ledgerHashs, err := service.GetLedgerHashs()
	require.Nil(t, err)

	// 创建交易
	txTemp := service.NewTransaction(ledgerHashs[0])

	// 部署合约
	file, err := os.Open("contract.car")
	defer file.Close()
	require.Nil(t, err)
	contract, err := ioutil.ReadAll(file)
	require.Nil(t, err)
	txTemp.Contracts().Deploy(user.GetIdentity(), contract)

	// TX 准备就绪；
	prepTx := txTemp.Prepare()

	// 使用私钥进行签名；
	prepTx.Sign(NODE_KEY.AsymmetricKeypair)

	// 提交交易；
	resp, err := prepTx.Commit()
	require.Nil(t, err)
	require.True(t, resp.Success)

}

func TestParticipant(t *testing.T) {
	// 生成公私钥对
	participant := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)

	// 连接网关，获取节点服务
	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
	service := serviceFactory.GetBlockchainService()

	// 获取账本信息
	ledgerHashs, err := service.GetLedgerHashs()
	require.Nil(t, err)

	// 创建交易
	txTemp := service.NewTransaction(ledgerHashs[0])

	name := "PARTICIPANT"
	identity := participant.GetIdentity()
	networkAddress := network.NewAddress("127.0.0.1", 20000, false).ToBytes()

	// 注册
	txTemp.Participants().Register(name, identity, networkAddress)
	// 激活
	txTemp.States().Update(identity, networkAddress, ledger_model.ACTIVED)

	// TX 准备就绪；
	prepTx := txTemp.Prepare()

	// 使用私钥进行签名；
	prepTx.Sign(NODE_KEY.AsymmetricKeypair)

	// 提交交易；
	resp, err := prepTx.Commit()
	require.Nil(t, err)
	require.True(t, resp.Success)

}

func TestUserEvent(t *testing.T) {
	// 生成公私钥对
	eventAccount := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)

	// 连接网关，获取节点服务
	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
	service := serviceFactory.GetBlockchainService()

	// 获取账本信息
	ledgerHashs, err := service.GetLedgerHashs()
	require.Nil(t, err)

	// 创建交易
	txTemp := service.NewTransaction(ledgerHashs[0])

	// 注册事件账户
	txTemp.EventAccounts().Register(eventAccount.GetIdentity())

	// 发布事件
	txTemp.EventAccount(eventAccount.GetAddress()).
		PublishBytes("e1", bytes.StringToBytes("bytes"), -1).
		PublishString("e2", "string", -1).
		PublishInt64("e3", 64, -1)

	// TX 准备就绪；
	prepTx := txTemp.Prepare()

	// 使用私钥进行签名；
	prepTx.Sign(NODE_KEY.AsymmetricKeypair)

	// 提交交易；
	resp, err := prepTx.Commit()
	require.Nil(t, err)
	require.True(t, resp.Success)

}
