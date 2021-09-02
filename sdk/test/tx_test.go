package test

import (
	"fmt"
	"github.com/blockchain-jd-com/framework-go/crypto/classic"
	"github.com/blockchain-jd-com/framework-go/ledger_model"
	"github.com/blockchain-jd-com/framework-go/sdk"
	"github.com/blockchain-jd-com/framework-go/utils/base58"
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
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

	// 生成公私钥对
	dataAccount := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)
	// 注册数据账户
	txTemp.DataAccounts().Register(dataAccount.GetIdentity())
	// 插入数据
	//txTemp.DataAccount(dataAccount.GetAddress()).SetText("key", "text", -1)
	//txTemp.DataAccount(dataAccount.GetAddress()).SetInt64("key", int64(64), 0)
	//txTemp.DataAccount(dataAccount.GetAddress()).SetBytes("key", []byte("bytes"), 1)
	//txTemp.DataAccount(dataAccount.GetAddress()).SetImage("key", []byte("image"), 2)
	//txTemp.DataAccount(dataAccount.GetAddress()).SetJSON("key", "json", 3)
	//txTemp.DataAccount(dataAccount.GetAddress()).SetTimestamp("key", time.Now().Unix(), 4)
//[0, 0, 2, 0, -54, -118, -15, 111, -117, -17, -30, -6, 34, 32, 24, -125, -47, 10, -22, 57, 57, 0, 64, 74, 114, 57, -40, 55, 22, 124, 3, 4, 122, 6, -88, -67, 39, -55, 42, -74, -14, 109, -59, 1, -112, -78, 63, 1, 64, 105, 0, 0, 3, 32, 13, 122, 17, -30, -73, 64, 95, 0, 64, 76, 0, 0, 0, -112, -4, 112, 68, 86, -99, -4, -52, -66, 27, -111, 65, 21, 95, -28, -56, -96, -4, -79, -123, -31, 90, 95, -29, 74, 54, -22, 107, -52, 120, 47, -77, 48, 21, 28, 123, -34, 35, 65, 21, 1, -61, -119, -86, -112, 23, -113, -54, -106, 107, -87, -41, 94, -46, -93, 32, -93, -105, 77, -61, 91, 127, -101, -28, -51, 26, 121, -43, -94, 123, 72, 96, -6, 14, 0, 0, 11, 48, -4, 112, 68, 86, -99, -4, -52, -66, 0, 0, 0, 0, 1, 123, -90, 59, 26, 13]
//[0, 0, 2, 0, -54, -118, -15, 111, -117, -17, -30, -6, 34, 32, 24, -125, -47, 10, -22, 57, 57, 0, 64, 74, 114, 57, -40, 55, 22, 124, 3, 4, 122, 6, -88, -67, 39, -55, 42, -74, -14, 109, -59, 1, -112, -78, 63, 1, 64, 103, 0, 0, 3, 32, 13, 122, 17, -30, -73, 64, 95, 0, 64, 76, 0, 0, 0, -112, -4, 112, 68, 86, -99, -4, -52, -66, 27, -111, 65, 21, 95, -28, -56, -96, -4, -79, -123, -31, 90, 95, -29, 74, 54, -22, 107, -52, 120, 47, -77, 48, 21, 28, 123, -34, 35, 65, 21, 1, -61, -119, -86, -112, 23, -113, -54, -106, 107, -87, -41, 94, -46, -93, 32, -93, -105, 77, -61, 91, 127, -101, -28, -51, 26, 121, -43, -94, 123, 72, 96, -6, 12, 0, 0, 11, 48, -4, 112, 68, 86, -99, -4, -52, -66, 0, 0, 1, 123, -90, 59, 26, 13]
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
