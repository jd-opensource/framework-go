package test

import (
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/ledger_model"
	"github.com/blockchain-jd-com/framework-go/sdk"
	"github.com/blockchain-jd-com/framework-go/utils/base58"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
 * Author: imuge
 * Date: 2020/5/25 下午6:16
 */

func TestQuery(t *testing.T) {
	blockchainService := sdk.Connect(GATEWAY_HOST, GATEWAY_PORT, SECURE, NODE_KEY).GetBlockchainService()

	// 返回所有的账本的 hash 列表
	ledgers, err := blockchainService.GetLedgerHashs()
	require.Nil(t, err)

	// 账本数量
	ledgersCount, err := blockchainService.GetLedgersCount()
	require.Nil(t, err)
	require.Equal(t, int64(len(ledgers)), ledgersCount)

	// 遍历所有账本
	for _, ledger := range ledgers {

		// 获取账本信息
		ledgerInfo, err := blockchainService.GetLedger(ledger)
		require.Nil(t, err)
		require.Equal(t, ledger, ledgerInfo.Hash)

		// 获取账本信息
		adminInfo, err := blockchainService.GetLedgerAdminInfo(ledger)
		require.Nil(t, err)

		// 返回当前账本的参与者信息列表
		participantNodes, err := blockchainService.GetConsensusParticipants(ledger)
		require.Nil(t, err)
		require.Equal(t, adminInfo.ParticipantCount, int64(len(participantNodes)))

		// 返回当前账本的元数据
		metadata, err := blockchainService.GetLedgerMetadata(ledger)
		require.Nil(t, err)
		require.Equal(t, metadata, adminInfo.Metadata.LedgerMetadata)

		// 遍历所有区块
		for height := int64(0); height <= ledgerInfo.LatestBlockHeight; height++ {

			// 返回指定账本序号的区块
			block, err := blockchainService.GetBlockByHeight(ledger, height)
			require.Nil(t, err)
			blockHash := framework.ParseHashDigest(block.Hash)

			// 返回指定区块hash的区块
			ledgerBlock, err := blockchainService.GetBlockByHash(ledger, blockHash)
			require.Nil(t, err)
			require.Equal(t, block, ledgerBlock)

			// 返回指定高度的区块中记录的交易总数
			txCountByHeight, err := blockchainService.GetTransactionCountByHeight(ledger, height)
			require.Nil(t, err)

			// 返回指定高度的区块中记录的交易总数
			txCountByHash, err := blockchainService.GetTransactionCountByHash(ledger, blockHash)
			require.Nil(t, err)
			require.Equal(t, txCountByHeight, txCountByHash)

			// 返回指定高度的区块中记录的数据账户总数
			daCountByHeight, err := blockchainService.GetDataAccountCountByHeight(ledger, height)
			require.Nil(t, err)

			// 返回指定的区块中记录的数据账户总数
			daCountByHash, err := blockchainService.GetDataAccountCountByHash(ledger, blockHash)
			require.Nil(t, err)
			require.Equal(t, daCountByHeight, daCountByHash)

			// 返回指定高度区块中的用户总数
			uCountByHeight, err := blockchainService.GetUserCountByHeight(ledger, height)
			require.Nil(t, err)

			// 返回指定区块中的用户总数
			uCountByHash, err := blockchainService.GetUserCountByHash(ledger, blockHash)
			require.Nil(t, err)
			require.Equal(t, uCountByHeight, uCountByHash)

			// 返回指定高度区块中的合约总数
			cCountByHeight, err := blockchainService.GetContractCountByHeight(ledger, height)
			require.Nil(t, err)

			// 返回指定区块中的合约总数
			cCountByHash, err := blockchainService.GetContractCountByHash(ledger, blockHash)
			require.Nil(t, err)
			require.Equal(t, cCountByHeight, cCountByHash)

			//分页返回指定账本序号的区块中的交易列表
			txsByHeight, err := blockchainService.GetTransactionsByHeight(ledger, height, 0, txCountByHeight)
			require.Nil(t, err)

			// 分页返回指定账本序号的区块中的交易列表
			txsByHash, err := blockchainService.GetTransactionsByHash(ledger, blockHash, 0, txCountByHeight)
			require.Nil(t, err)
			require.Equal(t, txsByHeight, txsByHash)

			// 遍历所有交易
			for _, tx := range txsByHeight {

				txHash := framework.ParseHashDigest(tx.TransactionContent.Hash)

				// 根据交易内容的哈希获取对应的交易记录
				txByHash, err := blockchainService.GetTransactionByContentHash(ledger, txHash)
				require.Nil(t, err)
				require.Equal(t, tx, txByHash)

				// 根据交易内容的哈希获取对应的交易状态
				txState, err := blockchainService.GetTransactionStateByContentHash(ledger, txHash)
				require.Nil(t, err)
				require.Equal(t, tx.ExecutionState, txState)
			}

			// 获取指定区块高度中新增的交易总数（即该区块中交易集合的数量）
			txacByHeight, err := blockchainService.GetAdditionalTransactionCountByHeight(ledger, height)
			require.Nil(t, err)
			require.True(t, txacByHeight > 0)

			// 获取指定区块Hash中新增的交易总数（即该区块中交易集合的数量）
			txacByHash, err := blockchainService.GetAdditionalTransactionCountByHash(ledger, blockHash)
			require.Nil(t, err)
			require.True(t, txacByHash > 0)
			require.Equal(t, txacByHeight, txacByHash)

			// 获取指定区块高度中新增的数据账户总数（即该区块中数据账户集合的数量）
			daacByHeight, err := blockchainService.GetAdditionalDataAccountCountByHeight(ledger, height)
			require.Nil(t, err)

			// 获取指定区块Hash中新增的数据账户总数（即该区块中数据账户集合的数量）
			daacByHash, err := blockchainService.GetAdditionalDataAccountCountByHash(ledger, blockHash)
			require.Nil(t, err)
			require.Equal(t, daacByHeight, daacByHash)

			// 获取指定区块高度中新增的用户总数（即该区块中用户集合的数量）
			uacByHeight, err := blockchainService.GetAdditionalUserCountByHeight(ledger, height)
			require.Nil(t, err)

			// 获取指定区块Hash中新增的用户总数（即该区块中用户集合的数量）
			uacByHash, err := blockchainService.GetAdditionalUserCountByHash(ledger, blockHash)
			require.Nil(t, err)
			require.Equal(t, uacByHeight, uacByHash)

			// 获取指定区块高度中新增的合约总数（即该区块中合约集合的数量）
			ccByHeight, err := blockchainService.GetAdditionalContractCountByHeight(ledger, height)
			require.Nil(t, err)

			// 获取指定区块Hash中新增的合约总数（即该区块中合约集合的数量）
			ccByHash, err := blockchainService.GetAdditionalContractCountByHash(ledger, blockHash)
			require.Nil(t, err)
			require.Equal(t, ccByHeight, ccByHash)

			if height == ledgerInfo.LatestBlockHeight {
				// 获取指定账本最新区块附加的交易数量
				ltxac, err := blockchainService.GetAdditionalTransactionCount(ledger)
				require.Nil(t, err)
				require.Equal(t, txacByHash, ltxac)

				// 获取指定账本中附加的数据账户数量
				ldaac, err := blockchainService.GetAdditionalDataAccountCount(ledger)
				require.Nil(t, err)
				require.Equal(t, daacByHash, ldaac)

				// 获取指定账本中新增的用户数量
				luac, err := blockchainService.GetAdditionalUserCount(ledger)
				require.Nil(t, err)
				require.Equal(t, uacByHash, luac)

				// 获取指定账本中新增的合约数量
				lcc, err := blockchainService.GetAdditionalContractCount(ledger)
				require.Nil(t, err)
				require.Equal(t, ccByHash, lcc)
			}
		}

		// 返回当前账本的交易总数
		_, err = blockchainService.GetTransactionTotalCount(ledger)
		require.Nil(t, err)

		// 返回当前账本的数据账户总数
		totalDAccount, err := blockchainService.GetDataAccountTotalCount(ledger)
		require.Nil(t, err)

		// 返回当前账本的用户总数
		totalUser, err := blockchainService.GetUserTotalCount(ledger)
		require.Nil(t, err)

		// 返回当前账本的合约总数
		totalContract, err := blockchainService.GetContractTotalCount(ledger)
		require.Nil(t, err)

		// get users by ledgerHash and its range
		users, err := blockchainService.GetUsers(ledger, 0, totalUser)
		require.Nil(t, err)

		for _, user := range users {

			u := base58.Encode(user.Address)

			// 返回用户信息
			_, err = blockchainService.GetUser(ledger, u)
			require.Nil(t, err)

			// 返回user's priveleges;
			userPrivilege, err := blockchainService.GetUserPrivileges(ledger, u)
			require.Nil(t, err)
			// return role's privileges;
			for _, role := range userPrivilege.UserRoles {
				rolePrivilegeSet, err := blockchainService.GetRolePrivileges(ledger, role)
				require.Nil(t, err)
				require.Equal(t, role, rolePrivilegeSet.RoleName)
			}
		}

		// get data accounts by ledgerHash and its range
		dataAccounts, err := blockchainService.GetDataAccounts(ledger, 0, totalDAccount)
		require.Nil(t, err)

		for _, da := range dataAccounts {
			dataAccount := base58.Encode(da.Address)

			// 返回数据账户信息
			daid, err := blockchainService.GetDataAccount(ledger, dataAccount)
			require.Nil(t, err)
			require.Equal(t, da, daid)

			// 返回指定数据账户中KV数据的总数
			dEntriesCount, err := blockchainService.GetDataEntriesTotalCount(ledger, dataAccount)
			require.Nil(t, err)

			// 数据账户中指定序号的最新值； 返回结果的顺序与指定的序号的顺序是一致的
			dEntries, err := blockchainService.GetLatestDataEntriesByRange(ledger, dataAccount, 0, dEntriesCount)
			require.Nil(t, err)
			require.Equal(t, dEntriesCount, int64(len(dEntries)))

			for _, entry := range dEntries {
				// 数据账户中指定的键的最新值
				latestEntry, err := blockchainService.GetLatestDataEntries(ledger, dataAccount, []string{entry.Key})
				require.Nil(t, err)
				require.Equal(t, entry, latestEntry[0])

				vers := make([]int64, entry.Version+1)
				for v := int64(0); v <= entry.Version; v++ {
					vers[v] = v
				}
				oldEntres, err := blockchainService.GetDataEntries(ledger, dataAccount, ledger_model.KVInfoVO{
					[]ledger_model.KVDataVO{
						{
							Key:     entry.Key,
							Version: vers,
						},
					},
				})
				require.Nil(t, err)
				require.Equal(t, len(vers), len(oldEntres))
				require.Equal(t, entry, oldEntres[len(vers)-1])
			}
		}

		// get contract accounts by ledgerHash and its range
		contractAccounts, err := blockchainService.GetContractAccounts(ledger, 0, totalContract)
		require.Nil(t, err)
		require.Equal(t, totalContract, int64(len(contractAccounts)))

		for _, contract := range contractAccounts {
			contractAccount := base58.Encode(contract.Address)

			// 返回合约账户信息
			contractInfo, err := blockchainService.GetContract(ledger, contractAccount)
			require.Nil(t, err)
			require.Equal(t, contract, contractInfo.BlockchainIdentity)

		}

		// 用户事件账户总数
		totalEventAccount, err := blockchainService.GetUserEventAccountTotalCount(ledger)
		require.Nil(t, err)

		// 事件账户列表
		eventAccounts, err := blockchainService.GetUserEventAccounts(ledger, 0, totalEventAccount)
		require.Nil(t, err)
		require.Equal(t, totalEventAccount, int64(len(eventAccounts)))

		for _, ea := range eventAccounts {
			eventAddress := base58.Encode(ea.Address)

			// 用户事件账户
			eventAccount, err := blockchainService.GetUserEventAccount(ledger, eventAddress)
			require.Nil(t, err)
			require.Equal(t, ea, eventAccount)

			// 用户事件名总数
			namesCount, err := blockchainService.GetUserEventNameTotalCount(ledger, eventAddress)
			require.Nil(t, err)

			// 用户事件名列表
			names, err := blockchainService.GetUserEventNames(ledger, eventAddress, 0, namesCount)
			require.Nil(t, err)
			require.Equal(t, namesCount, int64(len(names)))

			for _, eventName := range names {

				// 用户事件总数
				userEventsCount, err := blockchainService.GetUserEventsTotalCount(ledger, eventAddress, eventName)
				require.Nil(t, err)

				// 用户事件列表
				events, err := blockchainService.GetUserEvents(ledger, eventAddress, eventName, 0, userEventsCount)
				require.Nil(t, err)
				require.Equal(t, userEventsCount, int64(len(events)))

				// 最新用户事件
				event, err := blockchainService.GetLatestUserEvent(ledger, eventAddress, eventName)
				require.Nil(t, err)
				require.Equal(t, event, events[len(events)-1])
			}
		}

		// 获取最新区块
		latestBlock, err := blockchainService.GetLatestBlock(ledger)
		require.Nil(t, err)
		require.Equal(t, ledgerInfo.Hash.ToBytes(), latestBlock.LedgerHash)
		require.Equal(t, ledgerInfo.LatestBlockHeight, latestBlock.Height)
		require.Equal(t, ledgerInfo.LatestBlockHash.ToBytes(), latestBlock.Hash)
	}
}
