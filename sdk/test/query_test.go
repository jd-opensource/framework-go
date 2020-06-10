package test

import (
	"framework-go/crypto/framework"
	"framework-go/ledger_model"
	"framework-go/sdk"
	"framework-go/utils/base58"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
 * Author: imuge
 * Date: 2020/5/25 下午6:16
 */

func TestQuery(t *testing.T) {
	serviceFactory := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY)

	// 返回所有的账本的 hash 列表
	ledgers, err := serviceFactory.GetBlockchainService().GetLedgerHashs()
	require.Nil(t, err)

	ledger := ledgers[0]

	// 获取账本信息
	_, err = serviceFactory.GetBlockchainService().GetLedger(ledger)
	require.Nil(t, err)

	// 获取账本信息
	_, err = serviceFactory.GetBlockchainService().GetLedgerAdminInfo(ledger)
	require.Nil(t, err)

	// 返回当前账本的参与者信息列表
	_, err = serviceFactory.GetBlockchainService().GetConsensusParticipants(ledger)
	require.Nil(t, err)

	// 返回当前账本的元数据
	_, err = serviceFactory.GetBlockchainService().GetLedgerMetadata(ledger)
	require.Nil(t, err)

	// 返回指定账本序号的区块
	block, err := serviceFactory.GetBlockchainService().GetBlockByHeight(ledger, 0)
	require.Nil(t, err)
	blockHash := framework.ParseHashDigest(block.Hash)

	// 返回指定区块hash的区块
	_, err = serviceFactory.GetBlockchainService().GetBlockByHash(ledger, blockHash)
	require.Nil(t, err)

	// 返回指定高度的区块中记录的交易总数
	_, err = serviceFactory.GetBlockchainService().GetTransactionCountByHeight(ledger, 0)
	require.Nil(t, err)

	// 返回指定高度的区块中记录的交易总数
	_, err = serviceFactory.GetBlockchainService().GetTransactionCountByHash(ledger, blockHash)
	require.Nil(t, err)

	// 返回当前账本的交易总数
	_, err = serviceFactory.GetBlockchainService().GetTransactionTotalCount(ledger)
	require.Nil(t, err)

	// 返回指定高度的区块中记录的数据账户总数
	_, err = serviceFactory.GetBlockchainService().GetDataAccountCountByHeight(ledger, 0)
	require.Nil(t, err)

	// 返回指定的区块中记录的数据账户总数
	_, err = serviceFactory.GetBlockchainService().GetDataAccountCountByHash(ledger, blockHash)
	require.Nil(t, err)

	// 返回当前账本的数据账户总数
	_, err = serviceFactory.GetBlockchainService().GetDataAccountTotalCount(ledger)
	require.Nil(t, err)

	// 返回指定高度区块中的用户总数
	_, err = serviceFactory.GetBlockchainService().GetUserCountByHeight(ledger, 0)
	require.Nil(t, err)

	// 返回指定区块中的用户总数
	_, err = serviceFactory.GetBlockchainService().GetUserCountByHash(ledger, blockHash)
	require.Nil(t, err)

	// 返回当前账本的用户总数
	_, err = serviceFactory.GetBlockchainService().GetUserTotalCount(ledger)
	require.Nil(t, err)

	// 返回指定高度区块中的合约总数
	_, err = serviceFactory.GetBlockchainService().GetContractCountByHeight(ledger, 0)
	require.Nil(t, err)

	// 返回指定区块中的合约总数
	_, err = serviceFactory.GetBlockchainService().GetContractCountByHash(ledger, blockHash)
	require.Nil(t, err)

	// 返回当前账本的合约总数
	_, err = serviceFactory.GetBlockchainService().GetContractTotalCount(ledger)
	require.Nil(t, err)

	// get users by ledgerHash and its range
	users, err := serviceFactory.GetBlockchainService().GetUsers(ledger, 0, 10)
	require.Nil(t, err)
	user := base58.Encode(users[0].Address)

	// 返回用户信息
	_, err = serviceFactory.GetBlockchainService().GetUser(ledger, user)
	require.Nil(t, err)

	// get data accounts by ledgerHash and its range
	dataAccounts, err := serviceFactory.GetBlockchainService().GetDataAccounts(ledger, 0, 10)
	require.Nil(t, err)
	dataAccount := base58.Encode(dataAccounts[0].Address)

	// 返回数据账户信息
	_, err = serviceFactory.GetBlockchainService().GetDataAccount(ledger, dataAccount)
	require.Nil(t, err)

	// 数据账户中指定的键的最新值
	_, err = serviceFactory.GetBlockchainService().GetLatestDataEntries(ledger, dataAccount, []string{"imuge"})
	require.Nil(t, err)

	_, err = serviceFactory.GetBlockchainService().GetDataEntries(ledger, dataAccount, ledger_model.KVInfoVO{
		[]ledger_model.KVDataVO{
			{
				Key:     "imuge",
				Version: []int64{0},
			},
		},
	})
	require.Nil(t, err)

	// 返回指定数据账户中KV数据的总数
	_, err = serviceFactory.GetBlockchainService().GetDataEntriesTotalCount(ledger, dataAccount)
	require.Nil(t, err)

	// 数据账户中指定序号的最新值； 返回结果的顺序与指定的序号的顺序是一致的
	_, err = serviceFactory.GetBlockchainService().GetLatestDataEntriesByRange(ledger, dataAccount, 0, 10)
	require.Nil(t, err)

	// return user's roles
	_, err = serviceFactory.GetBlockchainService().GetUserRoles(ledger, user)
	require.Nil(t, err)

	// get contract accounts by ledgerHash and its range
	contractAccounts, err := serviceFactory.GetBlockchainService().GetContractAccounts(ledger, 0, 10)
	require.Nil(t, err)
	contractAccount := base58.Encode(contractAccounts[0].Address)

	// 返回合约账户信息
	_, err = serviceFactory.GetBlockchainService().GetContract(ledger, contractAccount)
	require.Nil(t, err)

	// 事件账户列表
	eventAccounts, err := serviceFactory.GetBlockchainService().GetUserEventAccounts(ledger, 0, 10)
	require.Nil(t, err)
	eventAccount := base58.Encode(eventAccounts[0].Address)

	// 用户事件列表
	_, err = serviceFactory.GetBlockchainService().GetUserEvents(ledger, eventAccount, "e1", 0, 10)
	require.Nil(t, err)

	//分页返回指定账本序号的区块中的交易列表
	txs, err := serviceFactory.GetBlockchainService().GetTransactionsByHeight(ledger, 0, 0, 100)
	require.Nil(t, err)
	txHash := framework.ParseHashDigest(txs[0].TransactionContent.Hash)

	// 分页返回指定账本序号的区块中的交易列表
	_, err = serviceFactory.GetBlockchainService().GetTransactionsByHash(ledger, blockHash, 0, 10)
	require.Nil(t, err)

	// 根据交易内容的哈希获取对应的交易记录
	_, err = serviceFactory.GetBlockchainService().GetTransactionByContentHash(ledger, txHash)
	require.Nil(t, err)

	// 根据交易内容的哈希获取对应的交易状态
	_, err = serviceFactory.GetBlockchainService().GetTransactionStateByContentHash(ledger, txHash)
	require.NotNil(t, err)

	// 用户事件账户总数
	_, err = serviceFactory.GetBlockchainService().GetUserEventAccountTotalCount(ledger)
	require.Nil(t, err)

	// 用户事件账户列表
	accounts, err := serviceFactory.GetBlockchainService().GetUserEventAccounts(ledger, 0, 10)
	require.Nil(t, err)

	userEventAccountAddress := base58.Encode(accounts[0].Address)

	// 用户事件账户
	_, err = serviceFactory.GetBlockchainService().GetUserEventAccount(ledger, userEventAccountAddress)
	require.Nil(t, err)

	// 用户事件名总数
	namesCount, err := serviceFactory.GetBlockchainService().GetUserEventNameTotalCount(ledger, userEventAccountAddress)
	require.Nil(t, err)

	// 用户事件名列表
	names, err := serviceFactory.GetBlockchainService().GetUserEventNames(ledger, userEventAccountAddress, 0, int(namesCount))
	require.Nil(t, err)
	eventName := names[0]

	// 最新用户事件
	event, err := serviceFactory.GetBlockchainService().GetLatestUserEvent(ledger, userEventAccountAddress, eventName)
	require.Nil(t, err)
	require.Equal(t, eventName, event.Name)

	// 用户事件总数
	userEventsCount, err := serviceFactory.GetBlockchainService().GetUserEventsTotalCount(ledger, userEventAccountAddress, eventName)
	require.Nil(t, err)

	// 用户事件列表
	_, err = serviceFactory.GetBlockchainService().GetUserEvents(ledger, userEventAccountAddress, eventName, 0, int32(userEventsCount))
	require.Nil(t, err)
}