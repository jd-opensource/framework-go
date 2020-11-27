### SDK

JD Chain[交易](#交易)提交和[查询](#查询)

运行测试用例前请正确配置[constants](../sdk/test/constants.go)中相关配置信息。

#### 交易

测试用例参照[tx_test](../sdk/test/tx_test.go)

- [用户](#用户)
- [角色权限](#角色权限)
- [数据账户](#数据账户)
- [合约](#合约)
- [事件](#事件)
- [节点添加/移除](#节点添加/移除)

##### 用户

```go
// 生成公私钥对
user := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)

// 连接网关，获取节点服务
serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
service := serviceFactory.GetBlockchainService()

// 获取账本信息
ledgerHashs, err := service.GetLedgerHashs()

// 选择一个账本，创建交易
txTemp := service.NewTransaction(ledgerHashs[0])

// 注册用户
txTemp.Users().Register(user.GetIdentity())

// TX 准备就绪；
prepTx := txTemp.Prepare()

// 使用网络中已存在用户私钥进行签名；
prepTx.Sign(NODE_KEY.AsymmetricKeypair)

// 提交交易；
resp, err := prepTx.Commit()
```

##### 角色权限

```go
// 生成公私钥对
user := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)

// 连接网关，获取节点服务
serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
service := serviceFactory.GetBlockchainService()

// 获取账本信息
ledgerHashs, err := service.GetLedgerHashs()

// 选择一个账本，创建交易
txTemp := service.NewTransaction(ledgerHashs[0])

// 配置角色权限
txTemp.Security().Roles().Configure("MANAGER").
    EnableLedgerPermission(ledger_model.REGISTER_USER).
    EnableTransactionPermission(ledger_model.CONTRACT_OPERATION).
    DisableLedgerPermission(ledger_model.WRITE_DATA_ACCOUNT).
    DisableTransactionPermission(ledger_model.DIRECT_OPERATION)
txTemp.Security().Authorziations().ForUser([][]byte{user1.GetAddress(), user2.GetAddress()}).Authorize("MANAGER")

// TX 准备就绪；
prepTx := txTemp.Prepare()

// 使用网络中已存在用户私钥进行签名；
prepTx.Sign(NODE_KEY.AsymmetricKeypair)

// 提交交易；
resp, err := prepTx.Commit()
```

##### 数据账户

```go
// 生成公私钥对
user := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)

// 连接网关，获取节点服务
serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
service := serviceFactory.GetBlockchainService()

// 获取账本信息
ledgerHashs, err := service.GetLedgerHashs()

// 选择一个账本，创建交易
txTemp := service.NewTransaction(ledgerHashs[0])

// 注册数据账户
txTemp.DataAccounts().Register(user.GetIdentity())
// 插入数据
txTemp.DataAccount(user.GetAddress()).SetText("k1", "v1", -1).SetText("k2", "v2", -1)
txTemp.DataAccount(user.GetAddress()).SetText("k3", "v3", -1)

// TX 准备就绪；
prepTx := txTemp.Prepare()

// 使用网络中已存在用户私钥进行签名；
prepTx.Sign(NODE_KEY.AsymmetricKeypair)

// 提交交易；
resp, err := prepTx.Commit()
```

##### 合约

```go
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
txTemp.Contracts().Deploy(user.GetIdentity(), contract, -1)

// TX 准备就绪；
prepTx := txTemp.Prepare()

// 使用私钥进行签名；
prepTx.Sign(NODE_KEY.AsymmetricKeypair)

// 提交交易；
resp, err := prepTx.Commit()
require.Nil(t, err)
require.True(t, resp.Success)

// 创建合约调用交易
txTemp = service.NewTransaction(ledgerHashs[0])
txTemp.ContractEvents().Send(user.GetAddress(), 合约版本, "方法名", 参数列表...)
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
require.EqualValues(t, "success", bytes.ToString(res[0].Result.Bytes))
```

##### 事件

```go
// 生成公私钥对
eventAccount := sdk.NewBlockchainKeyGenerator().Generate(classic.ED25519_ALGORITHM)

// 连接网关，获取节点服务
serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
service := serviceFactory.GetBlockchainService()

// 获取账本信息
ledgerHashs, err := service.GetLedgerHashs()

// 选择一个账本，创建交易
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

// 使用网络中已存在用户私钥进行签名；
prepTx.Sign(NODE_KEY.AsymmetricKeypair)

// 提交交易；
resp, err := prepTx.Commit()
```

##### 节点添加/移除

1. 注册新节点
```go
// 生成公私钥对
participantPriviteKey := crypto.DecodePrivKey(string(MustLoadFile("nodes/peer4/config/keys/jd.priv")), base58.MustDecode(string(MustLoadFile("nodes/peer4/config/keys/jd.pwd"))))
participantPublicKey := crypto.DecodePubKey(string(MustLoadFile("nodes/peer4/config/keys/jd.pub")))
participant := ledger_model.NewBlockchainKeypair(participantPublicKey, participantPriviteKey)

// 连接网关，获取节点服务
serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
service := serviceFactory.GetBlockchainService()

// 获取账本信息
ledgerHashs, err := service.GetLedgerHashs()
require.Nil(t, err)

// 创建交易
txTemp := service.NewTransaction(ledgerHashs[0])

name := "peer4"
identity := participant.GetIdentity()

// 注册
txTemp.Participants().Register(name, identity)

// TX 准备就绪；
prepTx := txTemp.Prepare()

// 使用网络中已存在用户私钥进行签名；
prepTx.Sign(NODE_KEY.AsymmetricKeypair)

// 提交交易；
resp, err := prepTx.Commit()
require.Nil(t, err)
require.True(t, resp.Success)
```

2. 激活新节点
```go
// 连接网关，获取节点服务
serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
service := serviceFactory.GetBlockchainService()

// 获取账本信息
ledgerHashs, err := service.GetLedgerHashs()

// 激活，向新节点发送激活请求
consensusAService := sdk.NewRestyConsensusService("127.0.0.1", 7084, false)
resp, err := consensusAService.ActivateParticipant(ledgerHashs[0].ToBase58(), "127.0.0.1", 20000, "127.0.0.1", 7080)
require.Nil(t, err)
require.True(t, resp)
```

3. 移除节点
```go
// 连接网关，获取节点服务
serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)
service := serviceFactory.GetBlockchainService()

// 获取账本信息
ledgerHashs, err := service.GetLedgerHashs()

// 移除，向待移除节点发送移除请求
consensusAService := sdk.NewRestyConsensusService("127.0.0.1", 7084, false)
resp, err := consensusAService.InactivateParticipant(ledgerHashs[0].ToBase58(), "LdeNj9UCKucz5QmVnRYn9cB3G7EE5mabpn3Pq", "127.0.0.1", 7080)
require.Nil(t, err)
require.True(t, resp)
```

#### 查询

测试用例参照[query_test](../sdk/test/query_test.go)

```go
// 返回所有的账本的 hash 列表
ledgers, err := serviceFactory.GetBlockchainService().GetLedgerHashs()
ledger := ledgers[0]

// 获取账本信息
_, err = serviceFactory.GetBlockchainService().GetLedger(ledger)

// 获取账本信息
_, err = serviceFactory.GetBlockchainService().GetLedgerAdminInfo(ledger)

// 返回当前账本的参与者信息列表
_, err = serviceFactory.GetBlockchainService().GetConsensusParticipants(ledger)

// 返回当前账本的元数据
_, err = serviceFactory.GetBlockchainService().GetLedgerMetadata(ledger)

// 返回指定账本序号的区块
block, err := serviceFactory.GetBlockchainService().GetBlockByHeight(ledger, 0)
blockHash := framework.ParseHashDigest(block.Hash)

// 返回指定区块hash的区块
_, err = serviceFactory.GetBlockchainService().GetBlockByHash(ledger, blockHash)

// 返回指定高度的区块中记录的交易总数
_, err = serviceFactory.GetBlockchainService().GetTransactionCountByHeight(ledger, 0)

// 返回指定高度的区块中记录的交易总数
_, err = serviceFactory.GetBlockchainService().GetTransactionCountByHash(ledger, blockHash)

// 返回当前账本的交易总数
_, err = serviceFactory.GetBlockchainService().GetTransactionTotalCount(ledger)

// 返回指定高度的区块中记录的数据账户总数
_, err = serviceFactory.GetBlockchainService().GetDataAccountCountByHeight(ledger, 0)

// 返回指定的区块中记录的数据账户总数
_, err = serviceFactory.GetBlockchainService().GetDataAccountCountByHash(ledger, blockHash)

// 返回当前账本的数据账户总数
_, err = serviceFactory.GetBlockchainService().GetDataAccountTotalCount(ledger)

// 返回指定高度区块中的用户总数
_, err = serviceFactory.GetBlockchainService().GetUserCountByHeight(ledger, 0)

// 返回指定区块中的用户总数
_, err = serviceFactory.GetBlockchainService().GetUserCountByHash(ledger, blockHash)

// 返回当前账本的用户总数
_, err = serviceFactory.GetBlockchainService().GetUserTotalCount(ledger)

// 返回指定高度区块中的合约总数
_, err = serviceFactory.GetBlockchainService().GetContractCountByHeight(ledger, 0)

// 返回指定区块中的合约总数
_, err = serviceFactory.GetBlockchainService().GetContractCountByHash(ledger, blockHash)

// 返回当前账本的合约总数
_, err = serviceFactory.GetBlockchainService().GetContractTotalCount(ledger)

// get users by ledgerHash and its range
users, err := serviceFactory.GetBlockchainService().GetUsers(ledger, 0, 10)
user := base58.Encode(users[0].Address)

// 返回用户信息
_, err = serviceFactory.GetBlockchainService().GetUser(ledger, user)

// get data accounts by ledgerHash and its range
dataAccounts, err := serviceFactory.GetBlockchainService().GetDataAccounts(ledger, 0, 10)
dataAccount := base58.Encode(dataAccounts[0].Address)

// 返回数据账户信息
_, err = serviceFactory.GetBlockchainService().GetDataAccount(ledger, dataAccount)
require.Nil(t, err)

// 数据账户中指定的键的最新值
_, err = serviceFactory.GetBlockchainService().GetLatestDataEntries(ledger, dataAccount, []string{"test"})
require.Nil(t, err)

_, err = serviceFactory.GetBlockchainService().GetDataEntries(ledger, dataAccount, ledger_model.KVInfoVO{
    []ledger_model.KVDataVO{
        {
            Key:     "test",
            Version: []int64{0},
        },
    },
})

// 返回指定数据账户中KV数据的总数
_, err = serviceFactory.GetBlockchainService().GetDataEntriesTotalCount(ledger, dataAccount)

// 数据账户中指定序号的最新值； 返回结果的顺序与指定的序号的顺序是一致的
_, err = serviceFactory.GetBlockchainService().GetLatestDataEntriesByRange(ledger, dataAccount, 0, 10)

// return user's roles
_, err = serviceFactory.GetBlockchainService().GetUserRoles(ledger, user)

// get contract accounts by ledgerHash and its range
contractAccounts, err := serviceFactory.GetBlockchainService().GetContractAccounts(ledger, 0, 10)
contractAccount := base58.Encode(contractAccounts[0].Address)

// 返回合约账户信息
_, err = serviceFactory.GetBlockchainService().GetContract(ledger, contractAccount)

// 事件账户列表
eventAccounts, err := serviceFactory.GetBlockchainService().GetUserEventAccounts(ledger, 0, 10)
eventAccount := base58.Encode(eventAccounts[0].Address)

// 用户事件列表
_, err = serviceFactory.GetBlockchainService().GetUserEvents(ledger, eventAccount, "e1", 0, 10)

//分页返回指定账本序号的区块中的交易列表
txs, err := serviceFactory.GetBlockchainService().GetTransactionsByHeight(ledger, 0, 0, 100)
txHash := framework.ParseHashDigest(txs[0].TransactionContent.Hash)

// 分页返回指定账本序号的区块中的交易列表
_, err = serviceFactory.GetBlockchainService().GetTransactionsByHash(ledger, blockHash, 0, 10)

// 根据交易内容的哈希获取对应的交易记录
_, err = serviceFactory.GetBlockchainService().GetTransactionByContentHash(ledger, txHash)

// 根据交易内容的哈希获取对应的交易状态
_, err = serviceFactory.GetBlockchainService().GetTransactionStateByContentHash(ledger, txHash)

// 用户事件账户总数
_, err = serviceFactory.GetBlockchainService().GetUserEventAccountTotalCount(ledger)

// 用户事件账户列表
accounts, err := serviceFactory.GetBlockchainService().GetUserEventAccounts(ledger, 0, 10)

userEventAccountAddress := base58.Encode(accounts[0].Address)

// 用户事件账户
_, err = serviceFactory.GetBlockchainService().GetUserEventAccount(ledger, userEventAccountAddress)

// 用户事件名总数
namesCount, err := serviceFactory.GetBlockchainService().GetUserEventNameTotalCount(ledger, userEventAccountAddress)

// 用户事件名列表
names, err := serviceFactory.GetBlockchainService().GetUserEventNames(ledger, userEventAccountAddress, 0, int(namesCount))
eventName := names[0]

// 最新用户事件
event, err := serviceFactory.GetBlockchainService().GetLatestUserEvent(ledger, userEventAccountAddress, eventName)
require.Equal(t, eventName, event.Name)

// 用户事件总数
userEventsCount, err := serviceFactory.GetBlockchainService().GetUserEventsTotalCount(ledger, userEventAccountAddress, eventName)

// 用户事件列表
_, err = serviceFactory.GetBlockchainService().GetUserEvents(ledger, userEventAccountAddress, eventName, 0, int32(userEventsCount))
```