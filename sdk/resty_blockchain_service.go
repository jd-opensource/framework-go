package sdk

import (
	"errors"
	"fmt"
	"framework-go/crypto/framework"
	"framework-go/ledger_model"
	"framework-go/utils/base58"
	"github.com/go-resty/resty/v2"
)

/*
 * Author: imuge
 * Date: 2020/5/27 下午5:55
 */

var _ ledger_model.BlockChainLedgerQueryService = (*RestyBlockchainService)(nil)

type RestyBlockchainService struct {
	host    string
	port    int32
	secure  bool
	client  *resty.Client
	baseUrl string
}

func NewRestyBlockchainService(host string, port int32, secure bool) *RestyBlockchainService {
	var baseUrl string
	if secure {
		baseUrl = fmt.Sprintf("https://%s:%d", host, port)
	} else {
		baseUrl = fmt.Sprintf("http://%s:%d", host, port)
	}
	return &RestyBlockchainService{
		host:    host,
		port:    port,
		secure:  secure,
		client:  resty.New(),
		baseUrl: baseUrl,
	}
}

func (r RestyBlockchainService) query(url string) (*WebResponse, error) {
	resp, err := r.client.R().SetResult(WebResponse{}).Get(r.baseUrl + url)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	wrp, ok := resp.Result().(*WebResponse)
	if !ok {
		return nil, errors.New("unparseable response")
	}
	if !wrp.Success {
		return nil, errors.New(fmt.Sprintf("error code:%d msg:%s", wrp.Error.ErrorCode, wrp.Error.ErrorMessage))
	}
	return wrp, nil
}

func (r RestyBlockchainService) GetLedgerHashs() ([]framework.HashDigest, error) {
	wrp, err := r.query("/ledgers")
	if err != nil {
		return nil, err
	}
	ledgers := wrp.Data.([]interface{})
	hashs := make([]framework.HashDigest, len(ledgers))
	for i, m := range ledgers {
		im := m.(map[string]interface{})
		hashs[i] = framework.ParseHashDigest(base58.MustDecode(im["value"].(string)))
	}
	return hashs, nil
}

func (r RestyBlockchainService) GetLedger(ledgerHash framework.HashDigest) (info ledger_model.LedgerInfo, err error) {
	wrp, err := r.query("/ledgers/" + ledgerHash.ToBase58())
	if err != nil {
		return info, err
	}
	infoMap := wrp.Data.(map[string]interface{})
	info.Hash = framework.ParseHashDigest(base58.MustDecode(infoMap["hash"].(map[string]interface{})["value"].(string)))
	info.LatestBlockHash = framework.ParseHashDigest(base58.MustDecode(infoMap["latestBlockHash"].(map[string]interface{})["value"].(string)))
	info.LatestBlockHeight = int64(infoMap["latestBlockHeight"].(float64))
	return info, nil
}

func (r RestyBlockchainService) GetLedgerAdminInfo(ledgerHash framework.HashDigest) (info ledger_model.LedgerAdminInfo, err error) {
	_, err = r.query("/ledgers/" + ledgerHash.ToBase58() + "/admininfo")
	if err != nil {
		return info, err
	}
	panic("implement me")
}

func (r RestyBlockchainService) GetConsensusParticipants(ledgerHash framework.HashDigest) ([]ledger_model.ParticipantNode, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetLedgerMetadata(ledgerHash framework.HashDigest) (ledger_model.LedgerMetadata, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetBlockByHeight(ledgerHash framework.HashDigest, height int64) (ledger_model.LedgerBlock, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetBlockByHash(ledgerHash, blockHash framework.HashDigest) (ledger_model.LedgerBlock, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetTransactionCountByHeight(ledgerHash framework.HashDigest, height int64) (int64, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetTransactionCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetTransactionTotalCount(ledgerHash framework.HashDigest) (int64, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetDataAccountCountByHeight(ledgerHash framework.HashDigest, height int64) (int64, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetDataAccountCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetDataAccountTotalCount(ledgerHash framework.HashDigest) (int64, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetUserCountByHeight(ledgerHash framework.HashDigest, height int64) (int64, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetUserCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetUserTotalCount(ledgerHash framework.HashDigest) (int64, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetContractCountByHeight(ledgerHash framework.HashDigest, height int64) (int64, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetContractCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetContractTotalCount(ledgerHash framework.HashDigest) (int64, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetTransactionsByHeight(ledgerHash framework.HashDigest, height int64, fromIndex, count int32) ([]ledger_model.LedgerTransaction, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetTransactionsByHash(ledgerHash, blockHash framework.HashDigest, fromIndex, count int32) ([]ledger_model.LedgerTransaction, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetTransactionByContentHash(ledgerHash, contentHash framework.HashDigest) (ledger_model.LedgerTransaction, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetTransactionStateByContentHash(ledgerHash, contentHash framework.HashDigest) (ledger_model.TransactionState, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetUser(ledgerHash framework.HashDigest, address string) (ledger_model.UserInfo, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetDataAccount(ledgerHash framework.HashDigest, address string) (ledger_model.BlockchainIdentity, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetLatestDataEntries(ledgerHash framework.HashDigest, address string, keys []string) ([]ledger_model.TypedKVEntry, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetDataEntries(ledgerHash framework.HashDigest, address string, kvInfoVO ledger_model.KVInfoVO) ([]ledger_model.TypedKVEntry, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetDataEntriesTotalCount(ledgerHash framework.HashDigest, address string) (int64, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetLatestDataEntriesByRange(ledgerHash framework.HashDigest, address string, fromIndex, count int32) ([]ledger_model.TypedKVEntry, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetContract(ledgerHash framework.HashDigest, address string) (ledger_model.ContractInfo, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetUsers(ledgerHash framework.HashDigest, fromIndex, count int32) ([]ledger_model.BlockchainIdentity, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetDataAccounts(ledgerHash framework.HashDigest, fromIndex, count int) ([]ledger_model.BlockchainIdentity, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetContractAccounts(ledgerHash framework.HashDigest, fromIndex, count int) ([]ledger_model.BlockchainIdentity, error) {
	panic("implement me")
}

func (r RestyBlockchainService) GetUserRoles(ledgerHash framework.HashDigest, userAddress string) (ledger_model.RoleSet, error) {
	panic("implement me")
}
