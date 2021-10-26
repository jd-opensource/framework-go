package test

import (
	"fmt"
	"github.com/blockchain-jd-com/framework-go/crypto/classic"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/ledger_model"
	"github.com/blockchain-jd-com/framework-go/sdk"
	"github.com/blockchain-jd-com/framework-go/utils/base58"
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
	"github.com/blockchain-jd-com/framework-go/utils/ca"
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
	//
	// 创建交易
	txTemp := service.NewTransaction(ledgerHashs[0])

	// 生成公私钥对
	userCert, _ := ca.RetrieveCertificate("-----BEGIN CERTIFICATE-----\nMIIB3zCCAYagAwIBAgIEFgKY7jAKBggqhkjOPQQDAjBwMQwwCgYDVQQKDANKRFQxDTALBgNVBAsM\nBFJPT1QxCzAJBgNVBAYTAkNOMQswCQYDVQQIDAJCSjELMAkGA1UEBwwCQkoxDTALBgNVBAMMBHJv\nb3QxGzAZBgkqhkiG9w0BCQEWDGltdWdlQGpkLmNvbTAeFw0yMTEwMTUxMTQxMjVaFw0zMTEwMTMx\nMTQxMjVaMHExDDAKBgNVBAoMA0pEVDENMAsGA1UECwwEVVNFUjELMAkGA1UEBhMCQ04xCzAJBgNV\nBAgMAkJKMQswCQYDVQQHDAJCSjEOMAwGA1UEAwwFdXNlcjExGzAZBgkqhkiG9w0BCQEWDGltdWdl\nQGpkLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABAI1OVeItk5prWS/+Bc23ExVz8420VGh\n0oa/NmzIf/aJewN0KpT1j8+wybmEyEWwdqV2rUbuJktjepkTPtpdjcyjDTALMAkGA1UdEwQCMAAw\nCgYIKoZIzj0EAwIDRwAwRAIgAtgfbwZS3yJdtYfnkoCZKM29jtBIvJLj5qXcDOHWW/YCIF0XKgwh\ng5RDHhdI3a7lh6CE5vGNZJH781MFHVCO6Ma5\n-----END CERTIFICATE-----")
	pubkey := ca.RetrievePubKey(userCert)
	address := framework.GenerateAddress(pubkey)
	// 注册用户
	txTemp.Users().RegisterWithCA(userCert)
	// 角色权限配置
	txTemp.Security().Roles().Configure("MANAGER").
		EnableLedgerPermission(ledger_model.REGISTER_USER).
		EnableTransactionPermission(ledger_model.CONTRACT_OPERATION).
		DisableLedgerPermission(ledger_model.WRITE_DATA_ACCOUNT).
		DisableTransactionPermission(ledger_model.DIRECT_OPERATION)
	txTemp.Security().Authorziations().ForUser([][]byte{address}).Authorize("MANAGER")

	// TX 准备就绪；
	prepTx := txTemp.Prepare()

	// 使用网络中已存在用户私钥进行签名；
	prepTx.Sign(NODE_KEY.AsymmetricKeypair)

	// 提交交易；
	resp, err := prepTx.Commit()
	require.Nil(t, err)
	require.True(t, resp.Success)

}

func TestUserState(t *testing.T) {
	// 连接网关，获取节点服务
	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
	service := serviceFactory.GetBlockchainService()

	// 获取账本信息
	ledgerHashs, err := service.GetLedgerHashs()
	require.Nil(t, err)

	// 创建交易
	txTemp := service.NewTransaction(ledgerHashs[0])

	// 更新用户状态
	txTemp.User(base58.MustDecode("LdeNr6RxwmsXVMgwBBCcFFvpYwEJwkmrcgd7w")).State(ledger_model.NORMAL)

	// TX 准备就绪；
	prepTx := txTemp.Prepare()

	// 使用私钥进行签名；
	prepTx.Sign(NODE_KEY.AsymmetricKeypair)

	// 提交交易；
	resp, err := prepTx.Commit()
	require.Nil(t, err)
	require.True(t, resp.Success)
}

func TestDataAccountRegister(t *testing.T) {

	// 连接网关，获取节点服务
	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
	service := serviceFactory.GetBlockchainService()

	// 获取账本信息
	ledgerHashs, err := service.GetLedgerHashs()
	require.Nil(t, err)

	// 创建交易
	txTemp := service.NewTransaction(ledgerHashs[0])

	// 生成公私钥对
	dataAccount := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)
	// 注册数据账户
	txTemp.DataAccounts().Register(dataAccount.GetIdentity())

	// TX 准备就绪；
	prepTx := txTemp.Prepare()

	// 使用网络中已存在用户私钥进行签名；
	prepTx.Sign(NODE_KEY.AsymmetricKeypair)

	// 提交交易；
	resp, err := prepTx.Commit()
	require.Nil(t, err)
	require.True(t, resp.Success)

}

func TestDataAccountSetKV(t *testing.T) {

	// 连接网关，获取节点服务
	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
	service := serviceFactory.GetBlockchainService()

	// 获取账本信息
	ledgerHashs, err := service.GetLedgerHashs()
	require.Nil(t, err)

	// 创建交易
	txTemp := service.NewTransaction(ledgerHashs[0])

	dataAccountAddress := base58.MustDecode("LdeNr6RxwmsXVMgwBBCcFFvpYwEJwkmrcgd7w")

	// 插入数据
	txTemp.DataAccount(dataAccountAddress).SetText("key", "text", -1)
	txTemp.DataAccount(dataAccountAddress).SetInt64("key", int64(64), 0)
	txTemp.DataAccount(dataAccountAddress).SetBytes("key", []byte("bytes"), 1)
	txTemp.DataAccount(dataAccountAddress).SetImage("key", []byte("image"), 2)
	txTemp.DataAccount(dataAccountAddress).SetJSON("key", "json", 3)
	txTemp.DataAccount(dataAccountAddress).SetTimestamp("key", time.Now().Unix(), 4)

	// TX 准备就绪；
	prepTx := txTemp.Prepare()

	// 使用网络中已存在用户私钥进行签名；
	prepTx.Sign(NODE_KEY.AsymmetricKeypair)

	// 提交交易；
	resp, err := prepTx.Commit()
	require.Nil(t, err)
	require.True(t, resp.Success)

}

func TestDataAccountPermission(t *testing.T) {

	// 连接网关，获取节点服务
	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
	service := serviceFactory.GetBlockchainService()

	// 获取账本信息
	ledgerHashs, err := service.GetLedgerHashs()
	require.Nil(t, err)

	// 创建交易
	txTemp := service.NewTransaction(ledgerHashs[0])

	dataAccountAddress := base58.MustDecode("LdeNvKC8tVkED4nRyhjY1t9hdNQugSC7XrhRd")

	// 更新数据账户权限
	txTemp.DataAccount(dataAccountAddress).Permission().Role("IMUGE")

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
	file, err := os.Open("contract-samples-1.4.0.RELEASE.car")
	defer file.Close()
	require.Nil(t, err)
	contract, err := ioutil.ReadAll(file)
	require.Nil(t, err)
	txTemp.Contracts().Deploy(user.GetIdentity(), contract, -1)

	// TX 准备就绪；
	prepTx := txTemp.Prepare()

	// 使用私钥进行签名；
	prepTx.Sign(NODE_KEY.AsymmetricKeypair)

	// 提交交易；
	resp, err := prepTx.Commit()
	require.Nil(t, err)
	require.True(t, resp.Success)

	// 创建合约调用交易，请修改数据账户地址为链上已经存在的数据账户地址
	txTemp = service.NewTransaction(ledgerHashs[0])
	err = txTemp.ContractEvents().Send(user.GetAddress(), 0, "register-user", "至少32位字节数-----------------------------")
	require.Nil(t, err)
	// TX 准备就绪；
	prepTx = txTemp.Prepare()

	// 使用私钥进行签名；
	prepTx.Sign(NODE_KEY.AsymmetricKeypair)

	// 提交交易；
	resp, err = prepTx.Commit()
	require.Nil(t, err)
	require.True(t, resp.Success)
	res := resp.OperationResults
	require.Equal(t, 1, len(res))
	fmt.Println(bytes.ToString(res[0].Result.Bytes))
}

func TestContractState(t *testing.T) {
	// 连接网关，获取节点服务
	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
	service := serviceFactory.GetBlockchainService()

	// 获取账本信息
	ledgerHashs, err := service.GetLedgerHashs()
	require.Nil(t, err)

	// 创建交易
	txTemp := service.NewTransaction(ledgerHashs[0])

	// 更新合约状态
	txTemp.Contract(base58.MustDecode("LdeNxyo5qifskkKQW3PRKBjEHuHeXUFLC1GXL")).State(ledger_model.FREEZE)

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

	// 连接网关，获取节点服务
	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
	service := serviceFactory.GetBlockchainService()

	// 获取账本信息
	ledgerHashs, err := service.GetLedgerHashs()
	require.Nil(t, err)

	// 创建交易
	txTemp := service.NewTransaction(ledgerHashs[0])

	// 生成公私钥对
	eventAccount := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)
	// 注册事件账户
	txTemp.EventAccounts().Register(eventAccount.GetIdentity())
	// 发布事件
	txTemp.EventAccount(eventAccount.GetAddress()).PublishString("topic", "text", -1)
	txTemp.EventAccount(eventAccount.GetAddress()).PublishInt64("topic", int64(64), 0)
	txTemp.EventAccount(eventAccount.GetAddress()).PublishBytes("topic", []byte("bytes"), 1)

	// TX 准备就绪；
	prepTx := txTemp.Prepare()

	// 使用网络中已存在用户私钥进行签名；
	prepTx.Sign(NODE_KEY.AsymmetricKeypair)

	// 提交交易；
	resp, err := prepTx.Commit()
	require.Nil(t, err)
	require.True(t, resp.Success)

}

func TestUserEventListener(t *testing.T) {
	// 连接网关，获取节点服务
	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
	service := serviceFactory.GetBlockchainService()

	// 获取账本信息
	ledgerHashs, err := service.GetLedgerHashs()
	require.Nil(t, err)

	user := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)
	handler := service.MonitorUserEvent(ledgerHashs[0], base58.Encode(user.GetAddress()), "e", 0, EUserEventListener{})

	// 创建交易
	txTemp := service.NewTransaction(ledgerHashs[0])

	// 注册事件账户
	txTemp.EventAccounts().Register(user.GetIdentity())
	e := "e"
	for j := 0; j < 20; j++ {
		c := fmt.Sprintf("c%d", j)
		// 发布事件
		txTemp.EventAccount(user.GetAddress()).PublishString(e, c, int64(j-1))
	}

	// TX 准备就绪；
	prepTx := txTemp.Prepare()

	// 使用网络中已存在用户私钥进行签名；
	prepTx.Sign(NODE_KEY.AsymmetricKeypair)

	// 提交交易；
	resp, err := prepTx.Commit()
	require.Nil(t, err)
	require.True(t, resp.Success)

	time.Sleep(time.Minute)

	handler.Cancel()

}

var _ sdk.UserEventListener = (*EUserEventListener)(nil)

type EUserEventListener struct {
}

func (E EUserEventListener) OnEvent(event ledger_model.Event, context sdk.UserEventContext) {
	fmt.Printf("event topic : %s \r\n", event.Name)
	fmt.Printf("event sequence : %d \n", event.Sequence)
	switch event.Content.Type {
	case ledger_model.INT64:
		fmt.Printf("event content : %d \n", bytes.ToInt64(event.Content.Bytes))
		break
	case ledger_model.TIMESTAMP:
		fmt.Printf("event content : %d \n", bytes.ToInt64(event.Content.Bytes))
		break
	case ledger_model.TEXT:
		fmt.Printf("event content : %s \n", bytes.ToString(event.Content.Bytes))
		break
	case ledger_model.JSON:
		fmt.Printf("event content : %s \n", bytes.ToString(event.Content.Bytes))
		break
	case ledger_model.XML:
		fmt.Printf("event content : %s \n", bytes.ToString(event.Content.Bytes))
		break
	case ledger_model.BYTES:
		fmt.Printf("event content : %v \n", event.Content.Bytes)
		break
	case ledger_model.IMG:
		fmt.Printf("event content : %v \n", event.Content.Bytes)
		break
	default:
		fmt.Println("not support content value")
	}

	// 可通过下面的代码停止监听
	//context.EventHandler.Cancel()
}

func TestSystemEventListener(t *testing.T) {
	// 连接网关，获取节点服务
	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
	service := serviceFactory.GetBlockchainService()

	// 获取账本信息
	ledgerHashs, err := service.GetLedgerHashs()
	require.Nil(t, err)

	// 监听新区快产生事件，目前只有这一种系统事件
	handler := service.MonitorSystemEvent(ledgerHashs[0], sdk.NewSystemEventPoint("new_block", 0), ESystemEventListener{})

	// 提交交易
	for i := 0; i < 20; i++ {
		txTemp := service.NewTransaction(ledgerHashs[0])
		user := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)
		txTemp.EventAccounts().Register(user.GetIdentity())
		prepTx := txTemp.Prepare()
		prepTx.Sign(NODE_KEY.AsymmetricKeypair)
		resp, err := prepTx.Commit()
		require.Nil(t, err)
		require.True(t, resp.Success)
	}

	time.Sleep(time.Minute)

	handler.Cancel()

}

var _ sdk.SystemEventListener = (*ESystemEventListener)(nil)

type ESystemEventListener struct {
}

func (E ESystemEventListener) OnEvents(events []ledger_model.Event, context sdk.SystemEventContext) {
	for _, event := range events {
		fmt.Printf("event topic : %s \n", event.Name)
		fmt.Printf("event sequence : %d \n", event.Sequence)
		fmt.Printf("event content : %d \n", bytes.ToInt64(event.Content.Bytes))
	}

	// 可通过下面的代码停止监听
	//context.EventHandler.Cancel()
}
