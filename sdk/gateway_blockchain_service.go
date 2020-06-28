package sdk

import (
	"framework-go/crypto/framework"
	"framework-go/ledger_model"
)

/*
 * Author: imuge
 * Date: 2020/5/27 下午4:18
 */

var _ BlockchainService = (*GatewayBlockchainService)(nil)

type GatewayBlockchainService struct {
	QueryService ledger_model.BlockchainQueryService
	TxService    ledger_model.TransactionService
}

func (b *GatewayBlockchainService) NewTransaction(ledgerHash framework.HashDigest) ledger_model.TransactionTemplate {
	return ledger_model.NewTxTemplate(ledgerHash, b.TxService)
}

func (b *GatewayBlockchainService) PrepareTransaction(content ledger_model.TransactionContent) ledger_model.PreparedTransaction {
	txReqBuilder := ledger_model.NewTxRequestBuilder(content)
	return ledger_model.NewPreparedTx(txReqBuilder, b.TxService)
}

func NewGatewayBlockchainService(txService ledger_model.TransactionService, queryService ledger_model.BlockchainQueryService) *GatewayBlockchainService {
	return &GatewayBlockchainService{
		QueryService: queryService,
		TxService:    txService,
	}
}

func (b *GatewayBlockchainService) GetLedgerHashs() ([]framework.HashDigest, error) {
	return b.QueryService.GetLedgerHashs()
}

func (b *GatewayBlockchainService) GetLedger(ledgerHash framework.HashDigest) (ledger_model.LedgerInfo, error) {
	return b.QueryService.GetLedger(ledgerHash)
}

func (b *GatewayBlockchainService) GetLedgerAdminInfo(ledgerHash framework.HashDigest) (ledger_model.LedgerAdminInfo, error) {
	return b.QueryService.GetLedgerAdminInfo(ledgerHash)
}

func (b *GatewayBlockchainService) GetConsensusParticipants(ledgerHash framework.HashDigest) ([]ledger_model.ParticipantNode, error) {
	return b.QueryService.GetConsensusParticipants(ledgerHash)
}

func (b *GatewayBlockchainService) GetLedgerMetadata(ledgerHash framework.HashDigest) (ledger_model.LedgerMetadata, error) {
	return b.QueryService.GetLedgerMetadata(ledgerHash)
}

func (b *GatewayBlockchainService) GetBlockByHeight(ledgerHash framework.HashDigest, height int64) (ledger_model.LedgerBlock, error) {
	return b.QueryService.GetBlockByHeight(ledgerHash, height)
}

func (b *GatewayBlockchainService) GetBlockByHash(ledgerHash, blockHash framework.HashDigest) (ledger_model.LedgerBlock, error) {
	return b.QueryService.GetBlockByHash(ledgerHash, blockHash)
}

func (b *GatewayBlockchainService) GetTransactionCountByHeight(ledgerHash framework.HashDigest, height int64) (int64, error) {
	return b.QueryService.GetTransactionCountByHeight(ledgerHash, height)
}

func (b *GatewayBlockchainService) GetTransactionCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	return b.QueryService.GetTransactionCountByHash(ledgerHash, blockHash)
}

func (b *GatewayBlockchainService) GetTransactionTotalCount(ledgerHash framework.HashDigest) (int64, error) {
	return b.QueryService.GetTransactionTotalCount(ledgerHash)
}

func (b *GatewayBlockchainService) GetDataAccountCountByHeight(ledgerHash framework.HashDigest, height int64) (int64, error) {
	return b.QueryService.GetDataAccountCountByHeight(ledgerHash, height)
}

func (b *GatewayBlockchainService) GetDataAccountCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	return b.QueryService.GetDataAccountCountByHash(ledgerHash, blockHash)
}

func (b *GatewayBlockchainService) GetDataAccountTotalCount(ledgerHash framework.HashDigest) (int64, error) {
	return b.QueryService.GetDataAccountTotalCount(ledgerHash)
}

func (b *GatewayBlockchainService) GetUserCountByHeight(ledgerHash framework.HashDigest, height int64) (int64, error) {
	return b.QueryService.GetUserCountByHeight(ledgerHash, height)
}

func (b *GatewayBlockchainService) GetUserCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	return b.QueryService.GetUserCountByHash(ledgerHash, blockHash)
}

func (b *GatewayBlockchainService) GetUserTotalCount(ledgerHash framework.HashDigest) (int64, error) {
	return b.QueryService.GetUserTotalCount(ledgerHash)
}

func (b *GatewayBlockchainService) GetContractCountByHeight(ledgerHash framework.HashDigest, height int64) (int64, error) {
	return b.QueryService.GetContractCountByHeight(ledgerHash, height)
}

func (b *GatewayBlockchainService) GetContractCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	return b.QueryService.GetContractCountByHash(ledgerHash, blockHash)
}

func (b *GatewayBlockchainService) GetContractTotalCount(ledgerHash framework.HashDigest) (int64, error) {
	return b.QueryService.GetContractTotalCount(ledgerHash)
}

func (b *GatewayBlockchainService) GetTransactionsByHeight(ledgerHash framework.HashDigest, height int64, fromIndex, count int64) ([]ledger_model.LedgerTransaction, error) {
	return b.QueryService.GetTransactionsByHeight(ledgerHash, height, fromIndex, count)
}

func (b *GatewayBlockchainService) GetTransactionsByHash(ledgerHash, blockHash framework.HashDigest, fromIndex, count int64) ([]ledger_model.LedgerTransaction, error) {
	return b.QueryService.GetTransactionsByHash(ledgerHash, blockHash, fromIndex, count)
}

func (b *GatewayBlockchainService) GetTransactionByContentHash(ledgerHash, contentHash framework.HashDigest) (ledger_model.LedgerTransaction, error) {
	return b.QueryService.GetTransactionByContentHash(ledgerHash, contentHash)
}

func (b *GatewayBlockchainService) GetTransactionStateByContentHash(ledgerHash, contentHash framework.HashDigest) (ledger_model.TransactionState, error) {
	return b.QueryService.GetTransactionStateByContentHash(ledgerHash, contentHash)
}

func (b *GatewayBlockchainService) GetUser(ledgerHash framework.HashDigest, address string) (ledger_model.UserInfo, error) {
	return b.QueryService.GetUser(ledgerHash, address)
}

func (b *GatewayBlockchainService) GetDataAccount(ledgerHash framework.HashDigest, address string) (ledger_model.BlockchainIdentity, error) {
	return b.QueryService.GetDataAccount(ledgerHash, address)
}

func (b *GatewayBlockchainService) GetLatestDataEntries(ledgerHash framework.HashDigest, address string, keys []string) ([]ledger_model.TypedKVEntry, error) {
	return b.QueryService.GetLatestDataEntries(ledgerHash, address, keys)
}

func (b *GatewayBlockchainService) GetDataEntries(ledgerHash framework.HashDigest, address string, kvInfoVO ledger_model.KVInfoVO) ([]ledger_model.TypedKVEntry, error) {
	return b.QueryService.GetDataEntries(ledgerHash, address, kvInfoVO)
}

func (b *GatewayBlockchainService) GetDataEntriesTotalCount(ledgerHash framework.HashDigest, address string) (int64, error) {
	return b.QueryService.GetDataEntriesTotalCount(ledgerHash, address)
}

func (b *GatewayBlockchainService) GetLatestDataEntriesByRange(ledgerHash framework.HashDigest, address string, fromIndex, count int64) ([]ledger_model.TypedKVEntry, error) {
	return b.QueryService.GetLatestDataEntriesByRange(ledgerHash, address, fromIndex, count)
}

func (b *GatewayBlockchainService) GetContract(ledgerHash framework.HashDigest, address string) (ledger_model.ContractInfo, error) {
	return b.QueryService.GetContract(ledgerHash, address)
}

func (b *GatewayBlockchainService) GetUsers(ledgerHash framework.HashDigest, fromIndex, count int64) ([]ledger_model.BlockchainIdentity, error) {
	return b.QueryService.GetUsers(ledgerHash, fromIndex, count)
}

func (b *GatewayBlockchainService) GetDataAccounts(ledgerHash framework.HashDigest, fromIndex, count int64) ([]ledger_model.BlockchainIdentity, error) {
	return b.QueryService.GetDataAccounts(ledgerHash, fromIndex, count)
}

func (b *GatewayBlockchainService) GetContractAccounts(ledgerHash framework.HashDigest, fromIndex, count int64) ([]ledger_model.BlockchainIdentity, error) {
	return b.QueryService.GetContractAccounts(ledgerHash, fromIndex, count)
}

func (b *GatewayBlockchainService) GetUserRoles(ledgerHash framework.HashDigest, userAddress string) (ledger_model.RoleSet, error) {
	return b.QueryService.GetUserRoles(ledgerHash, userAddress)
}

func (b *GatewayBlockchainService) GetSystemEvents(ledgerHash framework.HashDigest, eventName string, fromSequence int64, maxCount int64) ([]ledger_model.Event, error) {
	return b.QueryService.GetSystemEvents(ledgerHash, eventName, fromSequence, maxCount)
}

func (b *GatewayBlockchainService) GetUserEventAccounts(ledgerHash framework.HashDigest, fromIndex int64, maxCount int64) ([]ledger_model.BlockchainIdentity, error) {
	return b.QueryService.GetUserEventAccounts(ledgerHash, fromIndex, maxCount)
}

func (b *GatewayBlockchainService) GetUserEvents(ledgerHash framework.HashDigest, address string, eventName string, fromSequence int64, maxCount int64) ([]ledger_model.Event, error) {
	return b.QueryService.GetUserEvents(ledgerHash, address, eventName, fromSequence, maxCount)
}

func (b *GatewayBlockchainService) GetSystemEventNameTotalCount(digest framework.HashDigest) (int64, error) {
	return b.QueryService.GetSystemEventNameTotalCount(digest)
}

func (b *GatewayBlockchainService) GetSystemEventNames(digest framework.HashDigest, fromIndex, count int64) ([]string, error) {
	return b.QueryService.GetSystemEventNames(digest, fromIndex, count)
}

func (b *GatewayBlockchainService) GetSystemEventsTotalCount(digest framework.HashDigest, eventName string) (int64, error) {
	return b.QueryService.GetSystemEventsTotalCount(digest, eventName)
}

func (b *GatewayBlockchainService) GetUserEventAccount(digist framework.HashDigest, address string) (ledger_model.BlockchainIdentity, error) {
	return b.QueryService.GetUserEventAccount(digist, address)
}

func (b *GatewayBlockchainService) GetUserEventAccountTotalCount(digest framework.HashDigest) (int64, error) {
	return b.QueryService.GetUserEventAccountTotalCount(digest)
}

func (b *GatewayBlockchainService) GetUserEventNames(ledgerHash framework.HashDigest, address string, fromIndex, count int64) ([]string, error) {
	return b.QueryService.GetUserEventNames(ledgerHash, address, fromIndex, count)
}

func (b *GatewayBlockchainService) GetUserEventNameTotalCount(digest framework.HashDigest, address string) (int64, error) {
	return b.QueryService.GetUserEventNameTotalCount(digest, address)
}

func (b *GatewayBlockchainService) GetUserEventsTotalCount(digest framework.HashDigest, address, eventName string) (int64, error) {
	return b.QueryService.GetUserEventsTotalCount(digest, address, eventName)
}

func (b *GatewayBlockchainService) GetLatestSystemEvent(ledgerHash framework.HashDigest, eventName string) (ledger_model.Event, error) {
	return b.QueryService.GetLatestSystemEvent(ledgerHash, eventName)
}

func (b *GatewayBlockchainService) GetLatestUserEvent(ledgerHash framework.HashDigest, address string, eventName string) (ledger_model.Event, error) {
	return b.QueryService.GetLatestUserEvent(ledgerHash, address, eventName)
}

func (b *GatewayBlockchainService) GetLatestBlock(ledgerHash framework.HashDigest) (ledger_model.LedgerBlock, error) {
	return b.QueryService.GetLatestBlock(ledgerHash)
}

func (b *GatewayBlockchainService) GetAdditionalTransactionCountByHeight(ledgerHash framework.HashDigest, blockHeight int64) (int64, error) {
	return b.QueryService.GetAdditionalTransactionCountByHeight(ledgerHash, blockHeight)
}

func (b *GatewayBlockchainService) GetAdditionalTransactionCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	return b.QueryService.GetAdditionalTransactionCountByHash(ledgerHash, blockHash)
}

func (b *GatewayBlockchainService) GetAdditionalTransactionCount(ledgerHash framework.HashDigest) (int64, error) {
	return b.QueryService.GetAdditionalTransactionCount(ledgerHash)
}

func (b *GatewayBlockchainService) GetAdditionalDataAccountCountByHeight(ledgerHash framework.HashDigest, blockHeight int64) (int64, error) {
	return b.QueryService.GetAdditionalDataAccountCountByHeight(ledgerHash, blockHeight)
}

func (b *GatewayBlockchainService) GetAdditionalDataAccountCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	return b.QueryService.GetAdditionalDataAccountCountByHash(ledgerHash, blockHash)
}

func (b *GatewayBlockchainService) GetAdditionalDataAccountCount(ledgerHash framework.HashDigest) (int64, error) {
	return b.QueryService.GetAdditionalDataAccountCount(ledgerHash)
}

func (b *GatewayBlockchainService) GetAdditionalUserCountByHeight(ledgerHash framework.HashDigest, blockHeight int64) (int64, error) {
	return b.QueryService.GetAdditionalUserCountByHeight(ledgerHash, blockHeight)
}

func (b *GatewayBlockchainService) GetAdditionalUserCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	return b.QueryService.GetAdditionalUserCountByHash(ledgerHash, blockHash)
}

func (b *GatewayBlockchainService) GetAdditionalUserCount(ledgerHash framework.HashDigest) (int64, error) {
	return b.QueryService.GetAdditionalUserCount(ledgerHash)
}

func (b *GatewayBlockchainService) GetAdditionalContractCountByHeight(ledgerHash framework.HashDigest, blockHeight int64) (int64, error) {
	return b.QueryService.GetAdditionalContractCountByHeight(ledgerHash, blockHeight)
}

func (b *GatewayBlockchainService) GetAdditionalContractCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	return b.QueryService.GetAdditionalContractCountByHash(ledgerHash, blockHash)
}

func (b *GatewayBlockchainService) GetAdditionalContractCount(ledgerHash framework.HashDigest) (int64, error) {
	return b.QueryService.GetAdditionalContractCount(ledgerHash)
}

func (b *GatewayBlockchainService) GetLedgersCount() (int64, error) {
	return b.QueryService.GetLedgersCount()
}

func (b *GatewayBlockchainService) GetRolePrivileges(ledgerHash framework.HashDigest, roleName string) (ledger_model.PrivilegeSetVO, error) {
	return b.QueryService.GetRolePrivileges(ledgerHash, roleName)
}

func (b *GatewayBlockchainService) GetUserPrivileges(ledgerHash framework.HashDigest, userAddress string) (ledger_model.UserPrivilege, error) {
	return b.QueryService.GetUserPrivileges(ledgerHash, userAddress)
}
