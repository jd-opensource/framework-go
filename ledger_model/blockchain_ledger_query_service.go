package ledger_model

import (
	"framework-go/crypto/framework"
)

/*
 * Author: imuge
 * Date: 2020/5/27 上午11:23
 */

// 区块链账本相关查询器
type BlockChainLedgerQueryService interface {
	// 返回所有的账本的 hash 列表
	GetLedgerHashs() []framework.HashDigest

	// 获取账本信息
	GetLedger(ledgerHash framework.HashDigest) LedgerInfo

	// 获取账本信息
	GetLedgerAdminInfo(ledgerHash framework.HashDigest) LedgerAdminInfo

	// 返回当前账本的参与者信息列表
	GetConsensusParticipants(ledgerHash framework.HashDigest) []ParticipantNode

	// 返回当前账本的元数据
	GetLedgerMetadata(ledgerHash framework.HashDigest) LedgerMetadata

	// 返回指定账本序号的区块
	GetBlockByHeight(ledgerHash framework.HashDigest, height int64) LedgerBlock

	// 返回指定区块hash的区块
	GetBlockByHash(ledgerHash, blockHash framework.HashDigest) LedgerBlock

	// 返回指定高度的区块中记录的交易总数
	GetTransactionCountByHeight(ledgerHash framework.HashDigest, height int64) int64

	// 返回指定高度的区块中记录的交易总数
	GetTransactionCountByHash(ledgerHash, blockHash framework.HashDigest) int64

	// 返回当前账本的交易总数
	GetTransactionTotalCount(ledgerHash framework.HashDigest) int64

	// 返回指定高度的区块中记录的数据账户总数
	GetDataAccountCountByHeight(ledgerHash framework.HashDigest, height int64) int64

	// 返回指定的区块中记录的数据账户总数
	GetDataAccountCountByHash(ledgerHash, blockHash framework.HashDigest) int64

	// 返回当前账本的数据账户总数
	GetDataAccountTotalCount(ledgerHash framework.HashDigest) int64

	// 返回指定高度区块中的用户总数
	GetUserCountByHeight(ledgerHash framework.HashDigest, height int64) int64

	// 返回指定区块中的用户总数
	GetUserCountByHash(ledgerHash, blockHash framework.HashDigest) int64

	// 返回当前账本的用户总数
	GetUserTotalCount(ledgerHash framework.HashDigest) int64

	// 返回指定高度区块中的合约总数
	GetContractCountByHeight(ledgerHash framework.HashDigest, height int64) int64

	// 返回指定区块中的合约总数
	GetContractCountByHash(ledgerHash, blockHash framework.HashDigest) int64

	// 返回当前账本的合约总数
	GetContractTotalCount(ledgerHash framework.HashDigest) int64

	/**
	 * 分页返回指定账本序号的区块中的交易列表；
	 *
	 * @param ledgerHash 账本hash；
	 * @param height     账本高度；
	 * @param fromIndex  开始的记录数；
	 * @param count      本次返回的记录数；
	 */
	GetTransactionsByHeight(ledgerHash framework.HashDigest, height int64, fromIndex, count int32) []LedgerTransaction

	/**
	 * 分页返回指定账本序号的区块中的交易列表；
	 *
	 * @param ledgerHash 账本hash；
	 * @param blockHash  账本高度；
	 * @param fromIndex  开始的记录数；
	 * @param count      本次返回的记录数；
	 * @return
	 */
	GetTransactionsByHash(ledgerHash, blockHash framework.HashDigest, fromIndex, count int32) []LedgerTransaction

	/**
	 * 根据交易内容的哈希获取对应的交易记录；
	 *
	 * @param ledgerHash  账本hash；
	 * @param contentHash 交易内容的hash；
	 * @return
	 */
	GetTransactionByContentHash(ledgerHash, contentHash framework.HashDigest) LedgerTransaction

	/**
	 * 根据交易内容的哈希获取对应的交易状态；
	 *
	 * @param ledgerHash  账本hash；
	 * @param contentHash 交易内容的hash；
	 * @return
	 */
	GetTransactionStateByContentHash(ledgerHash, contentHash framework.HashDigest) TransactionState

	// 返回用户信息
	GetUser(ledgerHash framework.HashDigest, address string) UserInfo

	// 返回数据账户信息
	GetDataAccount(ledgerHash framework.HashDigest, address string) BlockchainIdentity

	/**
	 * 返回数据账户中指定的键的最新值；
	 *
	 * 返回结果的顺序与指定的键的顺序是一致的；
	 *
	 * 如果某个键不存在，则返回版本为 -1 的数据项；
	 *
	 */
	GetLatestDataEntries(ledgerHash framework.HashDigest, address string, keys []string) []TypedKVEntry

	GetDataEntries(ledgerHash framework.HashDigest, address string, kvInfoVO KVInfoVO) []TypedKVEntry

	// 返回指定数据账户中KV数据的总数
	GetDataEntriesTotalCount(ledgerHash framework.HashDigest, address string) int64

	/**
	 * 返回数据账户中指定序号的最新值； 返回结果的顺序与指定的序号的顺序是一致的；
	 *
	 * @param ledgerHash 账本hash；
	 * @param address    数据账户地址；
	 * @param fromIndex  开始的记录数；
	 * @param count      本次返回的记录数；
	 *                   如果参数值为 -1，则返回全部的记录；
	 */
	GetLatestDataEntriesByRange(ledgerHash framework.HashDigest, address string, fromIndex, count int32) []TypedKVEntry

	// 返回合约账户信息
	GetContract(ledgerHash framework.HashDigest, address string) ContractInfo

	// get users by ledgerHash and its range
	GetUsers(ledgerHash framework.HashDigest, fromIndex, count int32) []BlockchainIdentity

	// get data accounts by ledgerHash and its range
	GetDataAccounts(ledgerHash framework.HashDigest, fromIndex, count int) []BlockchainIdentity

	// get contract accounts by ledgerHash and its range
	GetContractAccounts(ledgerHash framework.HashDigest, fromIndex, count int) []BlockchainIdentity

	// return user's roles
	GetUserRoles(ledgerHash framework.HashDigest, userAddress string) RoleSet
}
