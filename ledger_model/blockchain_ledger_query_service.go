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
	GetLedgerHashs() ([]framework.HashDigest, error)

	// 获取账本信息
	GetLedger(ledgerHash framework.HashDigest) (LedgerInfo, error)

	// 获取账本信息
	GetLedgerAdminInfo(ledgerHash framework.HashDigest) (LedgerAdminInfo, error)

	// 返回当前账本的参与者信息列表
	GetConsensusParticipants(ledgerHash framework.HashDigest) ([]ParticipantNode, error)

	// 返回当前账本的元数据
	GetLedgerMetadata(ledgerHash framework.HashDigest) (LedgerMetadata, error)

	// 返回指定账本序号的区块
	GetBlockByHeight(ledgerHash framework.HashDigest, height int64) (LedgerBlock, error)

	// 返回指定区块hash的区块
	GetBlockByHash(ledgerHash, blockHash framework.HashDigest) (LedgerBlock, error)

	// 返回指定高度的区块中记录的交易总数
	GetTransactionCountByHeight(ledgerHash framework.HashDigest, height int64) (int64, error)

	// 返回指定高度的区块中记录的交易总数
	GetTransactionCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error)

	// 返回当前账本的交易总数
	GetTransactionTotalCount(ledgerHash framework.HashDigest) (int64, error)

	// 返回指定高度的区块中记录的数据账户总数
	GetDataAccountCountByHeight(ledgerHash framework.HashDigest, height int64) (int64, error)

	// 返回指定的区块中记录的数据账户总数
	GetDataAccountCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error)

	// 返回当前账本的数据账户总数
	GetDataAccountTotalCount(ledgerHash framework.HashDigest) (int64, error)

	// 返回指定高度区块中的用户总数
	GetUserCountByHeight(ledgerHash framework.HashDigest, height int64) (int64, error)

	// 返回指定区块中的用户总数
	GetUserCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error)

	// 返回当前账本的用户总数
	GetUserTotalCount(ledgerHash framework.HashDigest) (int64, error)

	// 返回指定高度区块中的合约总数
	GetContractCountByHeight(ledgerHash framework.HashDigest, height int64) (int64, error)

	// 返回指定区块中的合约总数
	GetContractCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error)

	// 返回当前账本的合约总数
	GetContractTotalCount(ledgerHash framework.HashDigest) (int64, error)

	/**
	 * 分页返回指定账本序号的区块中的交易列表；
	 *
	 * @param ledgerHash 账本hash；
	 * @param height     账本高度；
	 * @param fromIndex  开始的记录数；
	 * @param count      本次返回的记录数；
	 */
	GetTransactionsByHeight(ledgerHash framework.HashDigest, height int64, fromIndex, count int64) ([]LedgerTransaction, error)

	/**
	 * 分页返回指定账本序号的区块中的交易列表；
	 *
	 * @param ledgerHash 账本hash；
	 * @param blockHash  账本高度；
	 * @param fromIndex  开始的记录数；
	 * @param count      本次返回的记录数；
	 * @return
	 */
	GetTransactionsByHash(ledgerHash, blockHash framework.HashDigest, fromIndex, count int64) ([]LedgerTransaction, error)

	/**
	 * 根据交易内容的哈希获取对应的交易记录；
	 *
	 * @param ledgerHash  账本hash；
	 * @param contentHash 交易内容的hash；
	 * @return
	 */
	GetTransactionByContentHash(ledgerHash, contentHash framework.HashDigest) (LedgerTransaction, error)

	/**
	 * 根据交易内容的哈希获取对应的交易状态；
	 *
	 * @param ledgerHash  账本hash；
	 * @param contentHash 交易内容的hash；
	 * @return
	 */
	GetTransactionStateByContentHash(ledgerHash, contentHash framework.HashDigest) (TransactionState, error)

	// 返回用户信息
	GetUser(ledgerHash framework.HashDigest, address string) (UserInfo, error)

	// 返回数据账户信息
	GetDataAccount(ledgerHash framework.HashDigest, address string) (BlockchainIdentity, error)

	/**
	 * 返回数据账户中指定的键的最新值；
	 *
	 * 返回结果的顺序与指定的键的顺序是一致的；
	 *
	 * 如果某个键不存在，则返回版本为 -1 的数据项；
	 *
	 */
	GetLatestDataEntries(ledgerHash framework.HashDigest, address string, keys []string) ([]TypedKVEntry, error)

	GetDataEntries(ledgerHash framework.HashDigest, address string, kvInfoVO KVInfoVO) ([]TypedKVEntry, error)

	// 返回指定数据账户中KV数据的总数
	GetDataEntriesTotalCount(ledgerHash framework.HashDigest, address string) (int64, error)

	/**
	 * 返回数据账户中指定序号的最新值； 返回结果的顺序与指定的序号的顺序是一致的；
	 *
	 * @param ledgerHash 账本hash；
	 * @param address    数据账户地址；
	 * @param fromIndex  开始的记录数；
	 * @param count      本次返回的记录数；
	 *                   如果参数值为 -1，则返回全部的记录；
	 */
	GetLatestDataEntriesByRange(ledgerHash framework.HashDigest, address string, fromIndex, count int64) ([]TypedKVEntry, error)

	// 返回合约账户信息
	GetContract(ledgerHash framework.HashDigest, address string) (ContractInfo, error)

	// get users by ledgerHash and its range
	GetUsers(ledgerHash framework.HashDigest, fromIndex, count int64) ([]BlockchainIdentity, error)

	// get data accounts by ledgerHash and its range
	GetDataAccounts(ledgerHash framework.HashDigest, fromIndex, count int64) ([]BlockchainIdentity, error)

	// get contract accounts by ledgerHash and its range
	GetContractAccounts(ledgerHash framework.HashDigest, fromIndex, count int64) ([]BlockchainIdentity, error)

	// return user's roles
	GetUserRoles(ledgerHash framework.HashDigest, userAddress string) (RoleSet, error)

	/**
	 * 返回系统事件；
	 *
	 * @param ledgerHash   账本哈希；
	 * @param eventName    事件名；
	 * @param fromSequence 开始的事件序列号；
	 * @param maxCount     最大数量；
	 * @return
	 */
	GetSystemEvents(ledgerHash framework.HashDigest, eventName string, fromSequence int64, maxCount int64) ([]Event, error)

	/**
	 * 返回自定义事件账户；
	 * @param ledgerHash
	 * @param fromIndex
	 * @param count
	 * @return
	 */
	GetUserEventAccounts(ledgerHash framework.HashDigest, fromSequence int64, maxCount int64) ([]BlockchainIdentity, error)

	/**
	 * 返回自定义事件；
	 *
	 * @param ledgerHash   账本哈希；
	 * @param address      事件账户地址；
	 * @param eventName    事件名；
	 * @param fromSequence 开始的事件序列号；
	 * @param maxCount     最大数量；
	 * @return
	 */
	GetUserEvents(ledgerHash framework.HashDigest, address string, eventName string, fromSequence int64, maxCount int64) ([]Event, error)

	/**
	 * 返回系统事件名称总数； <br>
	 *
	 * @param ledgerHash
	 * @return
	 */
	GetSystemEventNameTotalCount(digest framework.HashDigest) (int64, error)

	/**
	 * 返回系统事件名称列表； <br>
	 *
	 * @param ledgerHash
	 * @param fromIndex
	 * @param count
	 * @return
	 */
	GetSystemEventNames(digest framework.HashDigest, fromIndex, count int64) ([]string, error)

	/**
	 * 返回指定系统事件名称下事件总数； <br>
	 *
	 * @param ledgerHash
	 * @param eventName
	 * @return
	 */
	GetSystemEventsTotalCount(digest framework.HashDigest, eventName string) (int64, error)

	/**
	 * 返回事件账户信息；
	 *
	 * @param ledgerHash
	 * @param address
	 * @return
	 */
	GetUserEventAccount(digist framework.HashDigest, address string) (BlockchainIdentity, error)

	/**
	 * 返回事件账户总数； <br>
	 *
	 * @param ledgerHash
	 * @return
	 */
	GetUserEventAccountTotalCount(digest framework.HashDigest) (int64, error)

	/**
	 * 返回指定事件账户事件名称列表； <br>
	 *
	 * @param ledgerHash
	 * @param address
	 * @return
	 */
	GetUserEventNames(ledgerHash framework.HashDigest, address string, fromIndex, count int64) ([]string, error)

	/**
	 * 返回指定事件账户事件名称总数； <br>
	 *
	 * @param ledgerHash
	 * @param address
	 * @return
	 */
	GetUserEventNameTotalCount(digest framework.HashDigest, address string) (int64, error)

	/**
	 * 返回指定事件账户，指定事件名称下事件总数； <br>
	 *
	 * @param ledgerHash
	 * @param address
	 * @param eventName
	 * @return
	 */
	GetUserEventsTotalCount(digest framework.HashDigest, address, eventName string) (int64, error)

	/**
	 * 返回最新系统事件； <br>
	 *
	 * @param ledgerHash
	 * @param eventName
	 * @return
	 */
	GetLatestSystemEvent(ledgerHash framework.HashDigest, eventName string) (Event, error)

	/**
	 * 返回最新用户自定义事件； <br>
	 *
	 * @param ledgerHash
	 * @param address
	 * @param eventName
	 * @return
	 */
	GetLatestUserEvent(ledgerHash framework.HashDigest, address string, eventName string) (Event, error)

	/**
	 * 获取最新区块
	 *
	 * @param ledgerHash
	 *         账本Hash
	 * @return
	 */
	GetLatestBlock(ledgerHash framework.HashDigest) (LedgerBlock, error)

	/**
	 * 获取指定区块高度中新增的交易总数（即该区块中交易集合的数量）
	 *
	 * @param ledgerHash
	 *         账本Hash
	 * @param blockHeight
	 *         区块高度
	 * @return
	 */
	GetAdditionalTransactionCountByHeight(ledgerHash framework.HashDigest, blockHeight int64) (int64, error)

	/**
	 * 获取指定区块Hash中新增的交易总数（即该区块中交易集合的数量）
	 *
	 * @param ledgerHash
	 *         账本Hash
	 * @param blockHash
	 *         区块Hash
	 * @return
	 */
	GetAdditionalTransactionCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error)

	/**
	 * 获取指定账本最新区块附加的交易数量
	 *
	 * @param ledgerHash
	 *         账本Hash
	 * @return
	 */
	GetAdditionalTransactionCount(ledgerHash framework.HashDigest) (int64, error)

	/**
	 * 获取指定区块高度中新增的数据账户总数（即该区块中数据账户集合的数量）
	 *
	 * @param ledgerHash
	 *         账本Hash
	 * @param blockHeight
	 *         区块高度
	 * @return
	 */
	GetAdditionalDataAccountCountByHeight(ledgerHash framework.HashDigest, blockHeight int64) (int64, error)

	/**
	 * 获取指定区块Hash中新增的数据账户总数（即该区块中数据账户集合的数量）
	 *
	 * @param ledgerHash
	 *         账本Hash
	 * @param blockHash
	 *         区块Hash
	 * @return
	 */
	GetAdditionalDataAccountCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error)

	/**
	 * 获取指定账本中附加的数据账户数量
	 *
	 * @param ledgerHash
	 *         账本Hash
	 * @return
	 */
	GetAdditionalDataAccountCount(ledgerHash framework.HashDigest) (int64, error)
	/**
	 * 获取指定区块高度中新增的用户总数（即该区块中用户集合的数量）
	 *
	 * @param ledgerHash
	 *         账本Hash
	 * @param blockHeight
	 *         区块高度
	 * @return
	 */
	GetAdditionalUserCountByHeight(ledgerHash framework.HashDigest, blockHeight int64) (int64, error)

	/**
	 * 获取指定区块Hash中新增的用户总数（即该区块中用户集合的数量）
	 *
	 * @param ledgerHash
	 *         账本Hash
	 * @param blockHash
	 *         区块Hash
	 * @return
	 */
	GetAdditionalUserCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error)

	/**
	 * 获取指定账本中新增的用户数量
	 *
	 * @param ledgerHash
	 *         账本Hash
	 * @return
	 */
	GetAdditionalUserCount(ledgerHash framework.HashDigest) (int64, error)

	/**
	 * 获取指定区块高度中新增的合约总数（即该区块中合约集合的数量）
	 *
	 * @param ledgerHash
	 *         账本Hash
	 * @param blockHeight
	 *         区块高度
	 * @return
	 */
	GetAdditionalContractCountByHeight(ledgerHash framework.HashDigest, blockHeight int64) (int64, error)

	/**
	 * 获取指定区块Hash中新增的合约总数（即该区块中合约集合的数量）
	 *
	 * @param ledgerHash
	 *         账本Hash
	 * @param blockHash
	 *         区块Hash
	 * @return
	 */
	GetAdditionalContractCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error)

	/**
	 * 获取指定账本中新增的合约数量
	 *
	 * @param ledgerHash
	 *         账本Hash
	 * @return
	 */
	GetAdditionalContractCount(ledgerHash framework.HashDigest) (int64, error)

	/**
	 *  get all ledgers count;
	 */
	GetLedgersCount() (int64, error)

	/**
	 * return role's privileges;
	 * @param ledgerHash
	 * @param roleName
	 * @return
	 */
	GetRolePrivileges(ledgerHash framework.HashDigest, roleName string) (PrivilegeSetVO, error)

	/**
	 * 返回user's priveleges;
	 *
	 * @param userAddress
	 * @return
	 */
	GetUserPrivileges(ledgerHash framework.HashDigest, userAddress string) (UserPrivilege, error)
}
