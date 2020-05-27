package sdk

import (
	"errors"
	"fmt"
	"framework-go/crypto/framework"
	"framework-go/ledger_model"
	"framework-go/utils/base58"
	"github.com/go-resty/resty/v2"
	"net/url"
)

/*
 * Author: imuge
 * Date: 2020/5/27 下午5:55
 */

var _ ledger_model.BlockChainLedgerQueryService = (*RestyQueryService)(nil)

type RestyQueryService struct {
	host    string
	port    int
	secure  bool
	client  *resty.Client
	baseUrl string
}

func NewRestyQueryService(host string, port int, secure bool) *RestyQueryService {
	var baseUrl string
	if secure {
		baseUrl = fmt.Sprintf("https://%s:%d", host, port)
	} else {
		baseUrl = fmt.Sprintf("http://%s:%d", host, port)
	}
	return &RestyQueryService{
		host:    host,
		port:    port,
		secure:  secure,
		client:  resty.New(),
		baseUrl: baseUrl,
	}
}

func (r RestyQueryService) query(url string) (interface{}, error) {
	resp, err := r.client.R().SetResult(WebResponse{}).Get(r.baseUrl + url)
	if err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("%s \n %v \n", url, resp))
	if !resp.IsSuccess() {
		return nil, errors.New(resp.String())
	}
	wrp, ok := resp.Result().(*WebResponse)
	if !ok {
		return nil, errors.New("unparseable response")
	}
	if !wrp.Success {
		return nil, errors.New(fmt.Sprintf("error code:%d msg:%s", wrp.Error.ErrorCode, wrp.Error.ErrorMessage))
	}
	return wrp.Data, nil
}

func (r RestyQueryService) queryWithFormData(url string, params map[string]string) (interface{}, error) {
	resp, err := r.client.R().SetFormData(params).SetResult(WebResponse{}).Get(r.baseUrl + url)
	if err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("%s \n %v \n", url, resp))
	if !resp.IsSuccess() {
		return nil, errors.New(resp.String())
	}
	wrp, ok := resp.Result().(*WebResponse)
	if !ok {
		return nil, errors.New("unparseable response")
	}
	if !wrp.Success {
		return nil, errors.New(fmt.Sprintf("error code:%d msg:%s", wrp.Error.ErrorCode, wrp.Error.ErrorMessage))
	}
	return wrp.Data, nil
}

func (r RestyQueryService) queryWithFormDataFromValues(url string, params url.Values) (interface{}, error) {
	resp, err := r.client.R().SetQueryParamsFromValues(params).SetResult(WebResponse{}).Get(r.baseUrl + url)
	if err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("%s \n %v \n", url, resp))
	if !resp.IsSuccess() {
		return nil, errors.New(resp.String())
	}
	wrp, ok := resp.Result().(*WebResponse)
	if !ok {
		return nil, errors.New("unparseable response")
	}
	if !wrp.Success {
		return nil, errors.New(fmt.Sprintf("error code:%d msg:%s", wrp.Error.ErrorCode, wrp.Error.ErrorMessage))
	}
	return wrp.Data, nil
}

func (r RestyQueryService) queryWithBody(url string, params interface{}) (interface{}, error) {
	resp, err := r.client.R().SetBody(params).SetResult(WebResponse{}).Post(r.baseUrl + url)
	if err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("%s \n %v \n", url, resp))
	if !resp.IsSuccess() {
		return nil, errors.New(resp.String())
	}
	wrp, ok := resp.Result().(*WebResponse)
	if !ok {
		return nil, errors.New("unparseable response")
	}
	if !wrp.Success {
		return nil, errors.New(fmt.Sprintf("error code:%d msg:%s", wrp.Error.ErrorCode, wrp.Error.ErrorMessage))
	}
	return wrp.Data, nil
}

func (r RestyQueryService) GetLedgerHashs() ([]framework.HashDigest, error) {
	wrp, err := r.query("/ledgers")
	if err != nil {
		return nil, err
	}
	ledgers := wrp.([]interface{})
	hashs := make([]framework.HashDigest, len(ledgers))
	for i, m := range ledgers {
		im := m.(map[string]interface{})
		hashs[i] = framework.ParseHashDigest(base58.MustDecode(im["value"].(string)))
	}
	return hashs, nil
}

func (r RestyQueryService) GetLedger(ledgerHash framework.HashDigest) (info ledger_model.LedgerInfo, err error) {
	wrp, err := r.query("/ledgers/" + ledgerHash.ToBase58())
	if err != nil {
		return info, err
	}
	infoMap := wrp.(map[string]interface{})
	info.Hash = framework.ParseHashDigest(base58.MustDecode(infoMap["hash"].(map[string]interface{})["value"].(string)))
	info.LatestBlockHash = framework.ParseHashDigest(base58.MustDecode(infoMap["latestBlockHash"].(map[string]interface{})["value"].(string)))
	info.LatestBlockHeight = int64(infoMap["latestBlockHeight"].(float64))
	return info, nil
}

func (r RestyQueryService) GetLedgerAdminInfo(ledgerHash framework.HashDigest) (info ledger_model.LedgerAdminInfo, err error) {
	wrp, err := r.query("/ledgers/" + ledgerHash.ToBase58() + "/admininfo")
	if err != nil {
		return info, err
	}
	infoMap := wrp.(map[string]interface{})
	info.ParticipantCount = int64(infoMap["participantCount"].(float64))
	metadata := infoMap["metadata"].(map[string]interface{})
	info.Metadata = ledger_model.LedgerMetadata_V2{
		LedgerMetadata: ledger_model.LedgerMetadata{
			Seed:             []byte(metadata["seed"].(string)),
			ParticipantsHash: base58.MustDecode(metadata["participantsHash"].(map[string]interface{})["value"].(string)),
			SettingsHash:     base58.MustDecode(metadata["settingsHash"].(map[string]interface{})["value"].(string)),
		},
		RolePrivilegesHash: base58.MustDecode(metadata["rolePrivilegesHash"].(map[string]interface{})["value"].(string)),
		UserRolesHash:      base58.MustDecode(metadata["userRolesHash"].(map[string]interface{})["value"].(string)),
	}
	participants := infoMap["participants"].([]interface{})
	pNodes := make([]ledger_model.ParticipantNode, len(participants))
	for i, p := range participants {
		node := p.(map[string]interface{})
		pNodes[i] = ledger_model.ParticipantNode{
			Id:                   int32(node["id"].(float64)),
			Name:                 node["name"].(string),
			Address:              base58.MustDecode(node["address"].(map[string]interface{})["value"].(string)),
			PubKey:               base58.MustDecode(node["pubKey"].(map[string]interface{})["value"].(string)),
			ParticipantNodeState: ledger_model.REGISTERED.GetValueByName(node["participantNodeState"].(string)).(ledger_model.ParticipantNodeState),
		}
	}
	info.Participants = pNodes
	settings := infoMap["settings"].(map[string]interface{})
	cryptoSetting := settings["cryptoSetting"].(map[string]interface{})
	autoVerifyHash, _ := cryptoSetting["autoVerifyHash"].(bool)
	supportedProviders := cryptoSetting["supportedProviders"].([]interface{})
	providers := make([]ledger_model.CryptoProvider, len(participants))
	for i, p := range supportedProviders {
		node := p.(map[string]interface{})
		algorithmsArray := node["algorithms"].([]interface{})
		algorithms := make([]framework.CryptoAlgorithm, len(algorithmsArray))
		for j, a := range algorithmsArray {
			if a == nil {
				continue
			}
			ma := a.(map[string]interface{})
			algorithms[j] = framework.CryptoAlgorithm{
				Code: int16(ma["code"].(float64)),
				Name: ma["name"].(string),
			}
		}
		providers[i] = ledger_model.CryptoProvider{
			Name:       node["name"].(string),
			Algorithms: algorithms,
		}
	}
	info.Settings = ledger_model.LedgerSettings{
		ConsensusProvider: settings["consensusProvider"].(string),
		ConsensusSetting:  base58.MustDecode(settings["consensusSetting"].(map[string]interface{})["value"].(string)),
		CryptoSetting: ledger_model.CryptoSetting{
			SupportedProviders: providers,
			HashAlgorithm:      int16(cryptoSetting["hashAlgorithm"].(float64)),
			AutoVerifyHash:     autoVerifyHash,
		},
	}
	return
}

func (r RestyQueryService) GetConsensusParticipants(ledgerHash framework.HashDigest) (info []ledger_model.ParticipantNode, err error) {
	wrp, err := r.query("/ledgers/" + ledgerHash.ToBase58() + "/participants")
	if err != nil {
		return info, err
	}
	participants := wrp.([]interface{})
	info = make([]ledger_model.ParticipantNode, len(participants))
	for i, p := range participants {
		node := p.(map[string]interface{})
		info[i] = ledger_model.ParticipantNode{
			Id:                   int32(node["id"].(float64)),
			Name:                 node["name"].(string),
			Address:              base58.MustDecode(node["address"].(map[string]interface{})["value"].(string)),
			PubKey:               base58.MustDecode(node["pubKey"].(map[string]interface{})["value"].(string)),
			ParticipantNodeState: ledger_model.REGISTERED.GetValueByName(node["participantNodeState"].(string)).(ledger_model.ParticipantNodeState),
		}
	}
	return
}

func (r RestyQueryService) GetLedgerMetadata(ledgerHash framework.HashDigest) (info ledger_model.LedgerMetadata, err error) {
	wrp, err := r.query("/ledgers/" + ledgerHash.ToBase58() + "/metadata")
	if err != nil {
		return info, err
	}
	metadata := wrp.(map[string]interface{})
	info = ledger_model.LedgerMetadata{
		Seed:             []byte(metadata["seed"].(string)),
		ParticipantsHash: base58.MustDecode(metadata["participantsHash"].(map[string]interface{})["value"].(string)),
		SettingsHash:     base58.MustDecode(metadata["settingsHash"].(map[string]interface{})["value"].(string)),
	}

	return
}

func (r RestyQueryService) GetBlockByHeight(ledgerHash framework.HashDigest, height int64) (info ledger_model.LedgerBlock, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d", ledgerHash.ToBase58(), height))
	if err != nil {
		return info, err
	}

	block := wrp.(map[string]interface{})

	var adminAccountHash []byte
	if adminAccountHashI, ok := block["adminAccountHash"]; ok {
		adminAccountHash = base58.MustDecode(adminAccountHashI.(map[string]interface{})["value"].(string))
	} else {
		adminAccountHash = nil
	}
	var userAccountSetHash []byte
	if adminAccountHashI, ok := block["adminAccountHash"]; ok {
		adminAccountHash = base58.MustDecode(adminAccountHashI.(map[string]interface{})["value"].(string))
	} else {
		adminAccountHash = nil
	}
	var dataAccountSetHash []byte
	if adminAccountHashI, ok := block["adminAccountHash"]; ok {
		adminAccountHash = base58.MustDecode(adminAccountHashI.(map[string]interface{})["value"].(string))
	} else {
		adminAccountHash = nil
	}
	var contractAccountSetHash []byte
	if adminAccountHashI, ok := block["adminAccountHash"]; ok {
		adminAccountHash = base58.MustDecode(adminAccountHashI.(map[string]interface{})["value"].(string))
	} else {
		adminAccountHash = nil
	}
	var PreviousHash []byte
	if PreviousHashI, ok := block["previousHash"]; ok {
		PreviousHash = base58.MustDecode(PreviousHashI.(map[string]interface{})["value"].(string))
	} else {
		PreviousHash = nil
	}

	info = ledger_model.LedgerBlock{
		BlockBody: ledger_model.BlockBody{
			LedgerDataSnapshot: ledger_model.LedgerDataSnapshot{
				AdminAccountHash:       adminAccountHash,
				UserAccountSetHash:     userAccountSetHash,
				DataAccountSetHash:     dataAccountSetHash,
				ContractAccountSetHash: contractAccountSetHash,
			},
			PreviousHash:       PreviousHash,
			LedgerHash:         base58.MustDecode(block["ledgerHash"].(map[string]interface{})["value"].(string)),
			Height:             nil,
			TransactionSetHash: base58.MustDecode(block["transactionSetHash"].(map[string]interface{})["value"].(string)),
			Timestamp:          nil,
		},
		Hash: base58.MustDecode(block["hash"].(map[string]interface{})["value"].(string)),
	}

	return
}

func (r RestyQueryService) GetBlockByHash(ledgerHash, blockHash framework.HashDigest) (info ledger_model.LedgerBlock, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return info, err
	}

	block := wrp.(map[string]interface{})

	var adminAccountHash []byte
	if adminAccountHashI, ok := block["adminAccountHash"]; ok {
		adminAccountHash = base58.MustDecode(adminAccountHashI.(map[string]interface{})["value"].(string))
	} else {
		adminAccountHash = nil
	}
	var userAccountSetHash []byte
	if adminAccountHashI, ok := block["adminAccountHash"]; ok {
		adminAccountHash = base58.MustDecode(adminAccountHashI.(map[string]interface{})["value"].(string))
	} else {
		adminAccountHash = nil
	}
	var dataAccountSetHash []byte
	if adminAccountHashI, ok := block["adminAccountHash"]; ok {
		adminAccountHash = base58.MustDecode(adminAccountHashI.(map[string]interface{})["value"].(string))
	} else {
		adminAccountHash = nil
	}
	var contractAccountSetHash []byte
	if adminAccountHashI, ok := block["adminAccountHash"]; ok {
		adminAccountHash = base58.MustDecode(adminAccountHashI.(map[string]interface{})["value"].(string))
	} else {
		adminAccountHash = nil
	}
	var PreviousHash []byte
	if PreviousHashI, ok := block["previousHash"]; ok {
		PreviousHash = base58.MustDecode(PreviousHashI.(map[string]interface{})["value"].(string))
	} else {
		PreviousHash = nil
	}

	info = ledger_model.LedgerBlock{
		BlockBody: ledger_model.BlockBody{
			LedgerDataSnapshot: ledger_model.LedgerDataSnapshot{
				AdminAccountHash:       adminAccountHash,
				UserAccountSetHash:     userAccountSetHash,
				DataAccountSetHash:     dataAccountSetHash,
				ContractAccountSetHash: contractAccountSetHash,
			},
			PreviousHash:       PreviousHash,
			LedgerHash:         base58.MustDecode(block["ledgerHash"].(map[string]interface{})["value"].(string)),
			Height:             nil,
			TransactionSetHash: base58.MustDecode(block["transactionSetHash"].(map[string]interface{})["value"].(string)),
			Timestamp:          nil,
		},
		Hash: base58.MustDecode(block["hash"].(map[string]interface{})["value"].(string)),
	}

	return
}

func (r RestyQueryService) GetTransactionCountByHeight(ledgerHash framework.HashDigest, height int64) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/txs/count", ledgerHash.ToBase58(), height))
	if err != nil {
		return 0, err
	}
	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetTransactionCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/txs/count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetTransactionTotalCount(ledgerHash framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/txs/count", ledgerHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetDataAccountCountByHeight(ledgerHash framework.HashDigest, height int64) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/accounts/count", ledgerHash.ToBase58(), height))
	if err != nil {
		return 0, err
	}
	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetDataAccountCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/accounts/count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetDataAccountTotalCount(ledgerHash framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/accounts/count", ledgerHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetUserCountByHeight(ledgerHash framework.HashDigest, height int64) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/users/count", ledgerHash.ToBase58(), height))
	if err != nil {
		return 0, err
	}
	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetUserCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/users/count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetUserTotalCount(ledgerHash framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/users/count", ledgerHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetContractCountByHeight(ledgerHash framework.HashDigest, height int64) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/contracts/count", ledgerHash.ToBase58(), height))
	if err != nil {
		return 0, err
	}
	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetContractCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/contracts/count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetContractTotalCount(ledgerHash framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/contracts/count", ledgerHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetTransactionsByHeight(ledgerHash framework.HashDigest, height int64, fromIndex, count int32) ([]ledger_model.LedgerTransaction, error) {
	panic("implement me")
}

func (r RestyQueryService) GetTransactionsByHash(ledgerHash, blockHash framework.HashDigest, fromIndex, count int32) ([]ledger_model.LedgerTransaction, error) {
	panic("implement me")
}

func (r RestyQueryService) GetTransactionByContentHash(ledgerHash, contentHash framework.HashDigest) (ledger_model.LedgerTransaction, error) {
	panic("implement me")
}

func (r RestyQueryService) GetTransactionStateByContentHash(ledgerHash, contentHash framework.HashDigest) (info ledger_model.TransactionState, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/txs/state/%s", ledgerHash.ToBase58(), contentHash.ToBase58()))
	if err != nil {
		return info, err
	}
	if wrp == nil {
		return info, errors.New("not exists")
	}
	info = ledger_model.SUCCESS.GetValueByName(wrp.(string)).(ledger_model.TransactionState)

	return
}

func (r RestyQueryService) GetUser(ledgerHash framework.HashDigest, address string) (info ledger_model.UserInfo, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/users/address/%s", ledgerHash.ToBase58(), address))
	if err != nil {
		return info, err
	}
	if wrp == nil {
		return info, errors.New("not exists")
	}
	user := wrp.(map[string]interface{})
	info = ledger_model.UserInfo{
		UserAccountHeader: ledger_model.UserAccountHeader{
			BlockchainIdentity: ledger_model.BlockchainIdentity{
				Address: base58.MustDecode(user["address"].(map[string]interface{})["value"].(string)),
				PubKey:  base58.MustDecode(user["pubKey"].(map[string]interface{})["value"].(string)),
			},
		},
	}
	return
}

func (r RestyQueryService) GetDataAccount(ledgerHash framework.HashDigest, address string) (info ledger_model.BlockchainIdentity, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/accounts/address/%s", ledgerHash.ToBase58(), address))
	if err != nil {
		return info, err
	}
	if wrp == nil {
		return info, errors.New("not exists")
	}
	id := wrp.(map[string]interface{})
	info = ledger_model.BlockchainIdentity{
		Address: base58.MustDecode(id["address"].(map[string]interface{})["value"].(string)),
		PubKey:  base58.MustDecode(id["pubKey"].(map[string]interface{})["value"].(string)),
	}
	return
}

func (r RestyQueryService) GetLatestDataEntries(ledgerHash framework.HashDigest, address string, keys []string) (info []ledger_model.TypedKVEntry, err error) {
	params := url.Values{
		"keys": keys,
	}
	wrp, err := r.queryWithFormDataFromValues(fmt.Sprintf("/ledgers/%s/accounts/%s/entries", ledgerHash.ToBase58(), address), params)
	if err != nil {
		return info, err
	}
	kvArray := wrp.([]interface{})
	info = make([]ledger_model.TypedKVEntry, len(kvArray))
	for i, item := range kvArray {
		id := item.(map[string]interface{})
		info[i] = ledger_model.TypedKVEntry{
			Key:     id["key"].(string),
			Value:   id["value"],
			Version: int64(id["version"].(float64)),
			Type:    ledger_model.NIL.GetValueByName(id["type"].(string)).(ledger_model.DataType),
		}
	}

	return
}

func (r RestyQueryService) GetDataEntries(ledgerHash framework.HashDigest, address string, kvInfoVO ledger_model.KVInfoVO) (info []ledger_model.TypedKVEntry, err error) {
	wrp, err := r.queryWithBody(fmt.Sprintf("/ledgers/%s/accounts/%s/entries-version", ledgerHash.ToBase58(), address), kvInfoVO)
	if err != nil {
		return info, err
	}
	kvArray := wrp.([]interface{})
	info = make([]ledger_model.TypedKVEntry, len(kvArray))
	for i, item := range kvArray {
		id := item.(map[string]interface{})
		info[i] = ledger_model.TypedKVEntry{
			Key:     id["key"].(string),
			Value:   id["value"],
			Version: int64(id["version"].(float64)),
			Type:    ledger_model.NIL.GetValueByName(id["type"].(string)).(ledger_model.DataType),
		}
	}

	return
}

func (r RestyQueryService) GetDataEntriesTotalCount(ledgerHash framework.HashDigest, address string) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/accounts/address/%s/entries/count", ledgerHash.ToBase58(), address))
	if err != nil {
		return 0, err
	}
	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetLatestDataEntriesByRange(ledgerHash framework.HashDigest, address string, fromIndex, count int32) (info []ledger_model.TypedKVEntry, err error) {
	params := map[string]string{
		"fromIndex": string(fromIndex),
		"count":     string(count),
	}
	wrp, err := r.queryWithFormData(fmt.Sprintf("/ledgers/%s/accounts/address/%s/entries", ledgerHash.ToBase58(), address), params)
	if err != nil {
		return info, err
	}
	kvArray := wrp.([]interface{})
	info = make([]ledger_model.TypedKVEntry, len(kvArray))
	for i, item := range kvArray {
		id := item.(map[string]interface{})
		info[i] = ledger_model.TypedKVEntry{
			Key:     id["key"].(string),
			Value:   id["value"],
			Version: int64(id["version"].(float64)),
			Type:    ledger_model.NIL.GetValueByName(id["type"].(string)).(ledger_model.DataType),
		}
	}

	return
}

func (r RestyQueryService) GetContract(ledgerHash framework.HashDigest, address string) (info ledger_model.ContractInfo, err error) {
	wrp, err := r.query(fmt.Sprintf("ledgers/%s/contracts/address/%s", ledgerHash.ToBase58(), address))
	if err != nil {
		return
	}
	if wrp == nil {
		return info, errors.New("not exists")
	}
	//_ := wrp.(map[string]interface{})
	info = ledger_model.ContractInfo{
		BlockchainIdentity: ledger_model.BlockchainIdentity{},
		MerkleSnapshot:     ledger_model.MerkleSnapshot{},
		ChainCode:          nil,
	}
	//TODO
	return
}

func (r RestyQueryService) GetUsers(ledgerHash framework.HashDigest, fromIndex, count int32) (info []ledger_model.BlockchainIdentity, err error) {
	params := map[string]string{
		"fromIndex": string(fromIndex),
		"count":     string(count),
	}
	wrp, err := r.queryWithFormData(fmt.Sprintf("/ledgers/%s/users", ledgerHash.ToBase58()), params)
	if err != nil {
		return info, err
	}
	idArray := wrp.([]interface{})
	info = make([]ledger_model.BlockchainIdentity, len(idArray))
	for i, item := range idArray {
		id := item.(map[string]interface{})
		info[i] = ledger_model.BlockchainIdentity{
			Address: base58.MustDecode(id["address"].(map[string]interface{})["value"].(string)),
			PubKey:  base58.MustDecode(id["pubKey"].(map[string]interface{})["value"].(string)),
		}
	}

	return
}

func (r RestyQueryService) GetDataAccounts(ledgerHash framework.HashDigest, fromIndex, count int) (info []ledger_model.BlockchainIdentity, err error) {
	params := map[string]string{
		"fromIndex": string(fromIndex),
		"count":     string(count),
	}
	wrp, err := r.queryWithFormData(fmt.Sprintf("/ledgers/%s/accounts", ledgerHash.ToBase58()), params)
	if err != nil {
		return info, err
	}
	idArray := wrp.([]interface{})
	info = make([]ledger_model.BlockchainIdentity, len(idArray))
	for i, item := range idArray {
		id := item.(map[string]interface{})
		info[i] = ledger_model.BlockchainIdentity{
			Address: base58.MustDecode(id["address"].(map[string]interface{})["value"].(string)),
			PubKey:  base58.MustDecode(id["pubKey"].(map[string]interface{})["value"].(string)),
		}
	}

	return
}

func (r RestyQueryService) GetContractAccounts(ledgerHash framework.HashDigest, fromIndex, count int) (info []ledger_model.BlockchainIdentity, err error) {
	params := map[string]string{
		"fromIndex": string(fromIndex),
		"count":     string(count),
	}
	wrp, err := r.queryWithFormData(fmt.Sprintf("/ledgers/%s/contracts", ledgerHash.ToBase58()), params)
	if err != nil {
		return info, err
	}
	idArray := wrp.([]interface{})
	info = make([]ledger_model.BlockchainIdentity, len(idArray))
	for i, item := range idArray {
		id := item.(map[string]interface{})
		info[i] = ledger_model.BlockchainIdentity{
			Address: base58.MustDecode(id["address"].(map[string]interface{})["value"].(string)),
			PubKey:  base58.MustDecode(id["pubKey"].(map[string]interface{})["value"].(string)),
		}
	}

	return
}

func (r RestyQueryService) GetUserRoles(ledgerHash framework.HashDigest, userAddress string) (info ledger_model.RoleSet, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/userrole/%s", ledgerHash.ToBase58(), userAddress))
	if err != nil {
		return info, err
	}
	roleSet := wrp.(map[string]interface{})
	rolesArray := roleSet["roleSet"].([]interface{})
	roles := make([]string, len(rolesArray))
	for i, role := range rolesArray {
		roles[i] = role.(string)
	}
	info = ledger_model.RoleSet{
		Policy: ledger_model.UNION.GetValueByName(roleSet["policy"].(string)).(ledger_model.RolesPolicy),
		Roles:  roles,
	}

	return
}
