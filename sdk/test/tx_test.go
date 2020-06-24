package test

import (
	"fmt"
	"framework-go/crypto/classic"
	"framework-go/ledger_model"
	"framework-go/sdk"
	"framework-go/utils/network"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

/*
 * Author: imuge
 * Date: 2020/5/26 上午11:29
 */

func TestRegisterUser(t *testing.T) {

	// 连接网关，获取节点服务
	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
	service := serviceFactory.GetBlockchainService()

	// 获取账本信息
	ledgerHashs, err := service.GetLedgerHashs()
	require.Nil(t, err)

	// 创建交易
	txTemp := service.NewTransaction(ledgerHashs[0])

	// 生成公私钥对
	user := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)
	// 注册用户
	txTemp.Users().Register(user.GetIdentity())
	// 角色权限配置
	txTemp.Security().Roles().Configure("MANAGER").
		EnableLedgerPermission(ledger_model.REGISTER_USER).
		EnableTransactionPermission(ledger_model.CONTRACT_OPERATION).
		DisableLedgerPermission(ledger_model.WRITE_DATA_ACCOUNT).
		DisableTransactionPermission(ledger_model.DIRECT_OPERATION)
	txTemp.Security().Authorziations().ForUser([][]byte{user.GetAddress()}).Authorize("MANAGER")

	// 注册更多用户
	for i := 0; i < 20; i++ {
		// 生成公私钥对
		user := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)
		// 注册用户
		txTemp.Users().Register(user.GetIdentity())
	}

	// TX 准备就绪；
	prepTx := txTemp.Prepare()

	// 使用网络中已存在用户私钥进行签名；
	prepTx.Sign(NODE_KEY.AsymmetricKeypair)

	// 提交交易；
	resp, err := prepTx.Commit()
	require.Nil(t, err)
	require.True(t, resp.Success)

}

func TestDataAccount(t *testing.T) {

	// 连接网关，获取节点服务
	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
	service := serviceFactory.GetBlockchainService()

	// 获取账本信息
	ledgerHashs, err := service.GetLedgerHashs()
	require.Nil(t, err)

	// 创建交易
	txTemp := service.NewTransaction(ledgerHashs[0])

	for i := 0; i < 20; i++ {
		// 生成公私钥对
		user := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)
		// 注册数据账户
		txTemp.DataAccounts().Register(user.GetIdentity())
		// 插入数据
		for j := 0; j < 20; j++ {
			k := fmt.Sprintf("k%d", j)
			txTemp.DataAccount(user.GetAddress()).SetText(k, "text", -1)
			txTemp.DataAccount(user.GetAddress()).SetInt64(k, int64(64), 0)
			txTemp.DataAccount(user.GetAddress()).SetBytes(k, []byte("bytes"), 1)
			txTemp.DataAccount(user.GetAddress()).SetImage(k, []byte("image"), 2)
			txTemp.DataAccount(user.GetAddress()).SetJSON(k, "json", 3)
			txTemp.DataAccount(user.GetAddress()).SetTimestamp(k, time.Now().Unix(), 4)
		}
		k := "k"
		for j := 0; j < 20; j++ {
			v := fmt.Sprintf("v%d", j)
			txTemp.DataAccount(user.GetAddress()).SetText(k, v, int64(j-1))
		}
	}

	// TX 准备就绪；
	prepTx := txTemp.Prepare()

	// 使用网络中已存在用户私钥进行签名；
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

	// 使用网络中已存在用户私钥进行签名；
	prepTx.Sign(NODE_KEY.AsymmetricKeypair)

	// 提交交易；
	resp, err := prepTx.Commit()
	require.Nil(t, err)
	require.True(t, resp.Success)

}

func TestUserEvent(t *testing.T) {

	// 连接网关，获取节点服务
	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
	service := serviceFactory.GetBlockchainService()

	// 获取账本信息
	ledgerHashs, err := service.GetLedgerHashs()
	require.Nil(t, err)

	// 创建交易
	txTemp := service.NewTransaction(ledgerHashs[0])

	for i := 0; i < 20; i++ {
		// 生成公私钥对
		user := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)
		// 注册事件账户
		txTemp.EventAccounts().Register(user.GetIdentity())
		// 发布事件
		for j := 0; j < 20; j++ {
			e := fmt.Sprintf("e%d", j)
			txTemp.EventAccount(user.GetAddress()).PublishString(e, "text", -1)
			txTemp.EventAccount(user.GetAddress()).PublishInt64(e, int64(64), 0)
			txTemp.EventAccount(user.GetAddress()).PublishBytes(e, []byte("bytes"), 1)
		}
		e := "e"
		for j := 0; j < 20; j++ {
			c := fmt.Sprintf("c%d", j)
			txTemp.EventAccount(user.GetAddress()).PublishString(e, c, int64(j-1))
		}
	}

	// TX 准备就绪；
	prepTx := txTemp.Prepare()

	// 使用网络中已存在用户私钥进行签名；
	prepTx.Sign(NODE_KEY.AsymmetricKeypair)

	// 提交交易；
	resp, err := prepTx.Commit()
	require.Nil(t, err)
	require.True(t, resp.Success)

}
