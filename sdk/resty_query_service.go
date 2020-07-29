package sdk

import (
	"errors"
	"fmt"
	binary_proto "framework-go/binary-proto"
	"framework-go/crypto/framework"
	"framework-go/ledger_model"
	"framework-go/utils/base58"
	"framework-go/utils/bytes"
	"framework-go/utils/network"
	"github.com/go-resty/resty/v2"
	"net/url"
	"strconv"
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

func (r RestyQueryService) queryWithParams(url string, params map[string]string) (interface{}, error) {
	resp, err := r.client.R().SetQueryParams(params).SetResult(WebResponse{}).Get(r.baseUrl + url)
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

func (r RestyQueryService) queryWithParamsFromValues(url string, params url.Values) (interface{}, error) {
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
			ParticipantNodeState: ledger_model.READY.GetValueByName(node["participantNodeState"].(string)).(ledger_model.ParticipantNodeState),
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
			ParticipantNodeState: ledger_model.READY.GetValueByName(node["participantNodeState"].(string)).(ledger_model.ParticipantNodeState),
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

	var PreviousHash []byte
	if PreviousHashI, ok := block["previousHash"]; ok {
		PreviousHash = base58.MustDecode(PreviousHashI.(map[string]interface{})["value"].(string))
	} else {
		PreviousHash = nil
	}

	info = ledger_model.LedgerBlock{
		BlockBody: ledger_model.BlockBody{
			LedgerDataSnapshot: parseLedgerDataSnapshot(block),
			PreviousHash:       PreviousHash,
			Height:             int64(block["height"].(float64)),
			TransactionSetHash: base58.MustDecode(block["transactionSetHash"].(map[string]interface{})["value"].(string)),
		},
		Hash: base58.MustDecode(block["hash"].(map[string]interface{})["value"].(string)),
	}

	if height > 0 {
		info.BlockBody.LedgerHash = base58.MustDecode(block["ledgerHash"].(map[string]interface{})["value"].(string))
	} else {
		info.BlockBody.LedgerHash = ledgerHash.ToBytes()
	}

	return
}

func (r RestyQueryService) GetBlockByHash(ledgerHash, blockHash framework.HashDigest) (info ledger_model.LedgerBlock, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return parseBlock(ledgerHash, wrp.(map[string]interface{}))
}

func parseBlock(ledgerHash framework.HashDigest, block map[string]interface{}) (info ledger_model.LedgerBlock, err error) {
	var PreviousHash []byte
	if PreviousHashI, ok := block["previousHash"]; ok {
		PreviousHash = base58.MustDecode(PreviousHashI.(map[string]interface{})["value"].(string))
	} else {
		PreviousHash = nil
	}

	info = ledger_model.LedgerBlock{
		BlockBody: ledger_model.BlockBody{
			LedgerDataSnapshot: parseLedgerDataSnapshot(block),
			PreviousHash:       PreviousHash,
			Height:             int64(block["height"].(float64)),
			TransactionSetHash: base58.MustDecode(block["transactionSetHash"].(map[string]interface{})["value"].(string)),
		},
		Hash: base58.MustDecode(block["hash"].(map[string]interface{})["value"].(string)),
	}

	if info.BlockBody.Height > 0 {
		info.BlockBody.LedgerHash = base58.MustDecode(block["ledgerHash"].(map[string]interface{})["value"].(string))
	} else {
		info.BlockBody.LedgerHash = ledgerHash.ToBytes()
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
			BlockchainIdentity: parseBlockchainIdentity(user),
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
	info = parseBlockchainIdentity(id)
	return
}

func (r RestyQueryService) GetLatestDataEntries(ledgerHash framework.HashDigest, address string, keys []string) (info []ledger_model.TypedKVEntry, err error) {
	params := url.Values{
		"keys": keys,
	}
	wrp, err := r.queryWithParamsFromValues(fmt.Sprintf("/ledgers/%s/accounts/%s/entries", ledgerHash.ToBase58(), address), params)
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

func (r RestyQueryService) GetLatestDataEntriesByRange(ledgerHash framework.HashDigest, address string, fromIndex, count int64) (info []ledger_model.TypedKVEntry, err error) {
	params := map[string]string{
		"fromIndex": strconv.FormatInt(fromIndex, 10),
		"count":     strconv.FormatInt(count, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/accounts/address/%s/entries", ledgerHash.ToBase58(), address), params)
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
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/contracts/address/%s", ledgerHash.ToBase58(), address))
	if err != nil {
		return
	}
	if wrp == nil {
		return info, errors.New("not exists")
	}
	contract := wrp.(map[string]interface{})
	info = ledger_model.ContractInfo{
		BlockchainIdentity: parseBlockchainIdentity(contract),
		MerkleSnapshot: ledger_model.MerkleSnapshot{
			RootHash: base58.MustDecode(contract["rootHash"].(map[string]interface{})["value"].(string)),
		},
		ChainCode: bytes.StringToBytes(contract["chainCode"].(string)),
	}

	return
}

func (r RestyQueryService) GetUsers(ledgerHash framework.HashDigest, fromIndex, count int64) (info []ledger_model.BlockchainIdentity, err error) {
	params := map[string]string{
		"fromIndex": strconv.FormatInt(fromIndex, 10),
		"count":     strconv.FormatInt(count, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/users", ledgerHash.ToBase58()), params)
	if err != nil {
		return info, err
	}
	idArray := wrp.([]interface{})
	info = make([]ledger_model.BlockchainIdentity, len(idArray))
	for i, item := range idArray {
		id := item.(map[string]interface{})
		info[i] = parseBlockchainIdentity(id)
	}

	return
}

func (r RestyQueryService) GetDataAccounts(ledgerHash framework.HashDigest, fromIndex, count int64) (info []ledger_model.BlockchainIdentity, err error) {
	params := map[string]string{
		"fromIndex": strconv.FormatInt(fromIndex, 10),
		"count":     strconv.FormatInt(count, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/accounts", ledgerHash.ToBase58()), params)
	if err != nil {
		return info, err
	}
	idArray := wrp.([]interface{})
	info = make([]ledger_model.BlockchainIdentity, len(idArray))
	for i, item := range idArray {
		id := item.(map[string]interface{})
		info[i] = parseBlockchainIdentity(id)
	}

	return
}

func (r RestyQueryService) GetContractAccounts(ledgerHash framework.HashDigest, fromIndex, count int64) (info []ledger_model.BlockchainIdentity, err error) {
	params := map[string]string{
		"fromIndex": strconv.FormatInt(fromIndex, 10),
		"count":     strconv.FormatInt(count, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/contracts", ledgerHash.ToBase58()), params)
	if err != nil {
		return info, err
	}
	idArray := wrp.([]interface{})
	info = make([]ledger_model.BlockchainIdentity, len(idArray))
	for i, item := range idArray {
		id := item.(map[string]interface{})
		info[i] = parseBlockchainIdentity(id)
	}

	return
}

func (r RestyQueryService) GetUserRoles(ledgerHash framework.HashDigest, userAddress string) (info ledger_model.RoleSet, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/user-role/%s", ledgerHash.ToBase58(), userAddress))
	if err != nil {
		return info, err
	}

	return parseRoleSet(wrp.(map[string]interface{}))
}

func parseRoleSet(roleSet map[string]interface{}) (info ledger_model.RoleSet, err error) {
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

func (r RestyQueryService) GetSystemEvents(ledgerHash framework.HashDigest, eventName string, fromSequence int64, maxCount int64) (info []ledger_model.Event, err error) {
	params := map[string]string{
		"fromSequence": strconv.FormatInt(fromSequence, 10),
		"count":        strconv.FormatInt(maxCount, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/events/system/names/%s", ledgerHash.ToBase58(), eventName), params)
	if err != nil {
		return info, err
	}
	idArray := wrp.([]interface{})
	info = make([]ledger_model.Event, len(idArray))
	for i, item := range idArray {
		info[i] = parseEvent(item.(map[string]interface{}))
	}

	return
}

func parseBytesValue(info map[string]interface{}) ledger_model.BytesValue {
	return ledger_model.BytesValue{
		Type:  ledger_model.NIL.GetValueByName(info["type"].(string)).(ledger_model.DataType),
		Bytes: base58.MustDecode(info["bytes"].(map[string]interface{})["value"].(string)),
	}
}

func (r RestyQueryService) GetUserEventAccounts(ledgerHash framework.HashDigest, fromIndex int64, maxCount int64) (info []ledger_model.BlockchainIdentity, err error) {
	params := map[string]string{
		"fromIndex": strconv.FormatInt(fromIndex, 10),
		"count":     strconv.FormatInt(maxCount, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/events/user/accounts", ledgerHash.ToBase58()), params)
	if err != nil {
		return info, err
	}
	idArray := wrp.([]interface{})
	info = make([]ledger_model.BlockchainIdentity, len(idArray))
	for i, item := range idArray {
		id := item.(map[string]interface{})
		info[i] = parseBlockchainIdentity(id)
	}

	return
}

func (r RestyQueryService) GetUserEvents(ledgerHash framework.HashDigest, address string, eventName string, fromSequence int64, maxCount int64) (info []ledger_model.Event, err error) {
	params := map[string]string{
		"fromSequence": strconv.FormatInt(fromSequence, 10),
		"count":        strconv.FormatInt(maxCount, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/events/user/accounts/%s/names/%s", ledgerHash.ToBase58(), address, eventName), params)
	if err != nil {
		return info, err
	}
	idArray := wrp.([]interface{})
	info = make([]ledger_model.Event, len(idArray))
	for i, item := range idArray {
		info[i] = parseEvent(item.(map[string]interface{}))
	}

	return
}

func parseEvent(event map[string]interface{}) ledger_model.Event {
	info := ledger_model.Event{
		Name:        event["name"].(string),
		BlockHeight: int64(event["blockHeight"].(float64)),
		Sequence:    int64(event["sequence"].(float64)),
	}
	transactionSource, ok := event["transactionSource"]
	if ok {
		info.TransactionSource = base58.MustDecode(transactionSource.(map[string]interface{})["value"].(string))
	}
	contractSource, ok := event["contractSource"]
	if ok {
		info.ContractSource = contractSource.(string)
	}
	eventAccount, ok := event["eventAccount"]
	if ok {
		info.EventAccount = base58.MustDecode(eventAccount.(map[string]interface{})["value"].(string))
	}
	content := event["content"].(map[string]interface{})
	if !content["nil"].(bool) {
		info.Content = parseBytesValue(content)
	}

	return info
}

func (r RestyQueryService) GetTransactionsByHeight(ledgerHash framework.HashDigest, height int64, fromIndex, count int64) (info []ledger_model.LedgerTransaction, err error) {
	params := map[string]string{
		"fromIndex": strconv.FormatInt(fromIndex, 10),
		"count":     strconv.FormatInt(count, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/blocks/height/%d/txs", ledgerHash.ToBase58(), height), params)
	if err != nil {
		return info, err
	}
	txArray := wrp.([]interface{})
	info = make([]ledger_model.LedgerTransaction, len(txArray))
	for i, item := range txArray {
		tx := item.(map[string]interface{})

		info[i] = ledger_model.LedgerTransaction{
			LedgerDataSnapshot: parseLedgerDataSnapshot(tx),
			Transaction:        parseTransaction(tx),
		}
	}

	return
}

func (r RestyQueryService) GetTransactionsByHash(ledgerHash, blockHash framework.HashDigest, fromIndex, count int64) (info []ledger_model.LedgerTransaction, err error) {
	params := map[string]string{
		"fromIndex": strconv.FormatInt(fromIndex, 10),
		"count":     strconv.FormatInt(count, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/txs", ledgerHash.ToBase58(), blockHash.ToBase58()), params)
	if err != nil {
		return info, err
	}
	txArray := wrp.([]interface{})
	info = make([]ledger_model.LedgerTransaction, len(txArray))
	for i, item := range txArray {
		tx := item.(map[string]interface{})

		info[i] = ledger_model.LedgerTransaction{
			LedgerDataSnapshot: parseLedgerDataSnapshot(tx),
			Transaction:        parseTransaction(tx),
		}
	}

	return
}

func (r RestyQueryService) GetTransactionByContentHash(ledgerHash, contentHash framework.HashDigest) (info ledger_model.LedgerTransaction, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/txs/hash/%s", ledgerHash.ToBase58(), contentHash.ToBase58()))
	if err != nil {
		return info, err
	}
	tx := wrp.(map[string]interface{})
	return ledger_model.LedgerTransaction{
		LedgerDataSnapshot: parseLedgerDataSnapshot(tx),
		Transaction:        parseTransaction(tx),
	}, nil
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

func parseLedgerDataSnapshot(info map[string]interface{}) ledger_model.LedgerDataSnapshot {
	var adminAccountHash []byte
	if adminAccountHashI, ok := info["adminAccountHash"]; ok {
		adminAccountHash = base58.MustDecode(adminAccountHashI.(map[string]interface{})["value"].(string))
	} else {
		adminAccountHash = nil
	}
	var userAccountSetHash []byte
	if adminAccountHashI, ok := info["adminAccountHash"]; ok {
		adminAccountHash = base58.MustDecode(adminAccountHashI.(map[string]interface{})["value"].(string))
	} else {
		adminAccountHash = nil
	}
	var dataAccountSetHash []byte
	if adminAccountHashI, ok := info["adminAccountHash"]; ok {
		adminAccountHash = base58.MustDecode(adminAccountHashI.(map[string]interface{})["value"].(string))
	} else {
		adminAccountHash = nil
	}
	var contractAccountSetHash []byte
	if adminAccountHashI, ok := info["adminAccountHash"]; ok {
		adminAccountHash = base58.MustDecode(adminAccountHashI.(map[string]interface{})["value"].(string))
	} else {
		adminAccountHash = nil
	}

	return ledger_model.LedgerDataSnapshot{
		AdminAccountHash:       adminAccountHash,
		UserAccountSetHash:     userAccountSetHash,
		DataAccountSetHash:     dataAccountSetHash,
		ContractAccountSetHash: contractAccountSetHash,
	}
}

func parseTransaction(info map[string]interface{}) ledger_model.Transaction {
	// TODO OperationResults

	return ledger_model.Transaction{
		NodeRequest: ledger_model.NodeRequest{
			NodeSignatures: parseNodeSignatures(info["nodeSignatures"].([]interface{})),
			EndpointRequest: ledger_model.EndpointRequest{
				TransactionContent: parseTransactionContent(info["transactionContent"].(map[string]interface{})),
			},
		},
		BlockHeight:    int64(info["blockHeight"].(float64)),
		ExecutionState: ledger_model.SUCCESS.GetValueByName(info["executionState"].(string)).(ledger_model.TransactionState),
	}
}

func parseTransactionContent(info map[string]interface{}) ledger_model.TransactionContent {
	return ledger_model.TransactionContent{
		TransactionContentBody: ledger_model.TransactionContentBody{
			Operations: parseOperations(info["operations"].([]interface{})),
		},
		Hash: base58.MustDecode(info["hash"].(map[string]interface{})["value"].(string)),
	}
}

func parseNodeSignatures(info []interface{}) []ledger_model.DigitalSignature {
	signatures := make([]ledger_model.DigitalSignature, len(info))
	for i, item := range info {
		sign := item.(map[string]interface{})
		signatures[i] = ledger_model.DigitalSignature{
			DigitalSignatureBody: ledger_model.DigitalSignatureBody{
				PubKey: base58.MustDecode(sign["pubKey"].(map[string]interface{})["value"].(string)),
				Digest: base58.MustDecode(sign["digest"].(map[string]interface{})["value"].(string)),
			},
		}
	}

	return signatures
}

func parseOperations(info []interface{}) []binary_proto.DataContract {
	operations := make([]binary_proto.DataContract, len(info))
	for i, item := range info {
		var dc binary_proto.DataContract
		operation := item.(map[string]interface{})
		if _, ok := operation["userID"]; ok {
			// 注册用户
			dc = parseUserRegisterOperation(operation["userID"].(map[string]interface{}))
		} else if _, ok := operation["accountID"]; ok {
			// 注册数据账户
			dc = parseDataAccountRegisterOperation(operation["accountID"].(map[string]interface{}))
		} else if _, ok := operation["writeSet"]; ok {
			// KV写入
			dc = parseDataAccountKVSetOperation(operation)
		} else if _, ok := operation["eventAccountID"]; ok {
			// 事件账户注册
			dc = parseEventAccountRegisterOperation(operation["eventAccountID"].(map[string]interface{}))
		} else if _, ok := operation["events"]; ok {
			// 发布事件
			dc = parseEventPublishOperation(operation)
		} else if _, ok := operation["participantRegisterIdentity"]; ok {
			// 注册参与方
			dc = parseParticipantRegisterOperation(operation)
		} else if _, ok := operation["stateUpdateIdentity"]; ok {
			// 参与方状态变更
			dc = parseParticipantStateUpdateOperation(operation)
		} else if _, ok := operation["chainCode"]; ok {
			// 合约部署
			dc = parseContractCodeDeployOperation(operation)
		} else if _, ok := operation["roles"]; ok {
			// 角色配置
			dc = parseRolesConfigureOperation(operation)
		} else if _, ok := operation["userRolesAuthorizations"]; ok {
			dc = parseUserAuthorizeOperation(operation["userRolesAuthorizations"].([]interface{}))
		}
		operations[i] = dc
	}

	return operations
}

func parseUserRegisterOperation(info map[string]interface{}) binary_proto.DataContract {
	return &ledger_model.UserRegisterOperation{
		UserID: parseBlockchainIdentity(info),
	}
}

func parseDataAccountRegisterOperation(info map[string]interface{}) binary_proto.DataContract {
	return &ledger_model.DataAccountRegisterOperation{
		AccountID: parseBlockchainIdentity(info),
	}
}

func parseEventAccountRegisterOperation(info map[string]interface{}) binary_proto.DataContract {
	return &ledger_model.EventAccountRegisterOperation{
		EventAccountID: parseBlockchainIdentity(info),
	}
}

func parseParticipantRegisterOperation(info map[string]interface{}) binary_proto.DataContract {
	return &ledger_model.ParticipantRegisterOperation{
		ParticipantName:             info["participantName"].(string),
		ParticipantRegisterIdentity: parseBlockchainIdentity(info["participantRegisterIdentity"].(map[string]interface{})),
	}
}

func parseParticipantStateUpdateOperation(info map[string]interface{}) binary_proto.DataContract {
	networkAddress := info["networkAddress"].(map[string]interface{})
	address := network.NewAddress(networkAddress["host"].(string), int32(networkAddress["port"].(float64)), networkAddress["secure"].(bool))
	return &ledger_model.ParticipantStateUpdateOperation{
		State:               ledger_model.READY.GetValueByName(info["state"].(string)).(ledger_model.ParticipantNodeState),
		StateUpdateIdentity: parseBlockchainIdentity(info["stateUpdateIdentity"].(map[string]interface{})),
		NetworkAddress:      address.ToBytes(),
	}
}

func parseContractCodeDeployOperation(info map[string]interface{}) binary_proto.DataContract {
	return &ledger_model.ContractCodeDeployOperation{
		ContractID: parseBlockchainIdentity(info["contractID"].(map[string]interface{})),
		ChainCode:  bytes.StringToBytes(info["chainCode"].(string)),
	}
}

func parseRolesConfigureOperation(info map[string]interface{}) binary_proto.DataContract {
	empty := info["empty"].(bool)
	var roles []ledger_model.RolePrivilegeEntry
	if !empty {
		array := info["roles"].([]interface{})
		roles = make([]ledger_model.RolePrivilegeEntry, len(array))
		for i, item := range array {
			role := item.(map[string]interface{})
			roleName := role["roleName"].(string)
			roles[i] = ledger_model.RolePrivilegeEntry{
				RoleName:                      roleName,
				EnableLedgerPermissions:       parseLedgerPermissions(role["enableLedgerPermissions"].([]interface{})),
				DisableLedgerPermissions:      parseLedgerPermissions(role["disableLedgerPermissions"].([]interface{})),
				EnableTransactionPermissions:  parseTransactionPermissions(role["enableTransactionPermissions"].([]interface{})),
				DisableTransactionPermissions: parseTransactionPermissions(role["disableTransactionPermissions"].([]interface{})),
			}
		}
	}
	return &ledger_model.RolesConfigureOperation{
		Roles: roles,
	}
}

func parseUserAuthorizeOperation(array []interface{}) binary_proto.DataContract {
	userAuthor := make([]ledger_model.UserRolesEntry, len(array))
	for i, item := range array {
		author := item.(map[string]interface{})
		addressArray := author["userAddresses"].([]interface{})
		address := make([][]byte, len(addressArray))
		for j, addr := range addressArray {
			address[j] = base58.MustDecode(addr.(map[string]interface{})["value"].(string))
		}
		userAuthor[i] = ledger_model.UserRolesEntry{
			Policy:            ledger_model.UNION.GetValueByName(author["policy"].(string)).(ledger_model.RolesPolicy),
			Addresses:         address,
			UnauthorizedRoles: parseStringArray(author["unauthorizedRoles"].([]interface{})),
			AuthorizedRoles:   parseStringArray(author["authorizedRoles"].([]interface{})),
		}
	}
	return &ledger_model.UserAuthorizeOperation{
		UserRolesAuthorizations: userAuthor,
	}
}

func parseStringArray(info []interface{}) []string {
	array := make([]string, len(info))
	for i, item := range info {
		array[i] = item.(string)
	}

	return array
}

func parseLedgerPermissions(array []interface{}) []ledger_model.LedgerPermission {
	lps := make([]ledger_model.LedgerPermission, len(array))
	for i, item := range array {
		lps[i] = ledger_model.CONFIGURE_ROLES.GetValueByName(item.(string)).(ledger_model.LedgerPermission)
	}

	return lps
}

func parseTransactionPermissions(array []interface{}) []ledger_model.TransactionPermission {
	tps := make([]ledger_model.TransactionPermission, len(array))
	for i, item := range array {
		tps[i] = ledger_model.DIRECT_OPERATION.GetValueByName(item.(string)).(ledger_model.TransactionPermission)
	}

	return tps
}

func parseDataAccountKVSetOperation(info map[string]interface{}) binary_proto.DataContract {
	kvs := info["writeSet"].([]interface{})
	writeSet := make([]ledger_model.KVWriteEntry, len(kvs))
	for i, item := range kvs {
		kv := item.(map[string]interface{})
		kvSet := ledger_model.KVWriteEntry{
			Key:             kv["key"].(string),
			Value:           parseBytesValue(kv["value"].(map[string]interface{})),
			ExpectedVersion: int64(kv["expectedVersion"].(float64)),
		}
		writeSet[i] = kvSet
	}
	return &ledger_model.DataAccountKVSetOperation{
		AccountAddress: base58.MustDecode(info["accountAddress"].(map[string]interface{})["value"].(string)),
		WriteSet:       writeSet,
	}
}

func parseEventPublishOperation(info map[string]interface{}) binary_proto.DataContract {
	kvs := info["events"].([]interface{})
	writeSet := make([]ledger_model.EventEntry, len(kvs))
	for i, item := range kvs {
		kv := item.(map[string]interface{})
		kvSet := ledger_model.EventEntry{
			Name:     kv["name"].(string),
			Content:  parseBytesValue(kv["content"].(map[string]interface{})),
			Sequence: int64(kv["sequence"].(float64)),
		}
		writeSet[i] = kvSet
	}
	return &ledger_model.EventPublishOperation{
		EventAddress: base58.MustDecode(info["eventAddress"].(map[string]interface{})["value"].(string)),
		Events:       writeSet,
	}
}

func parseBlockchainIdentity(id map[string]interface{}) ledger_model.BlockchainIdentity {
	return ledger_model.BlockchainIdentity{
		Address: base58.MustDecode(id["address"].(map[string]interface{})["value"].(string)),
		PubKey:  base58.MustDecode(id["pubKey"].(map[string]interface{})["value"].(string)),
	}
}

func (r RestyQueryService) GetSystemEventNameTotalCount(ledgerHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/system/names/count", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetSystemEventNames(ledgerHash framework.HashDigest, fromIndex, count int64) (info []string, err error) {
	params := map[string]string{
		"fromIndex": strconv.FormatInt(fromIndex, 10),
		"count":     strconv.FormatInt(count, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/events/system/names", ledgerHash.ToBase58()), params)
	if err != nil {
		return info, err
	}

	return parseStringArray(wrp.([]interface{})), nil
}

func (r RestyQueryService) GetSystemEventsTotalCount(ledgerHash framework.HashDigest, eventName string) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/system/names/%s/count", ledgerHash.ToBase58(), eventName))
	if err != nil {
		return info, err
	}

	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetUserEventAccount(ledgerHash framework.HashDigest, address string) (info ledger_model.BlockchainIdentity, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/user/accounts/%s", ledgerHash.ToBase58(), address))
	if err != nil {
		return info, err
	}

	return parseBlockchainIdentity(wrp.(map[string]interface{})), nil
}

func (r RestyQueryService) GetUserEventAccountTotalCount(ledgerHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/user/accounts/count", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetUserEventNames(ledgerHash framework.HashDigest, address string, fromIndex, count int64) (info []string, err error) {
	params := map[string]string{
		"fromIndex": strconv.FormatInt(fromIndex, 10),
		"count":     strconv.FormatInt(count, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/events/user/accounts/%s/names", ledgerHash.ToBase58(), address), params)
	if err != nil {
		return info, err
	}

	return parseStringArray(wrp.([]interface{})), nil

}

func (r RestyQueryService) GetUserEventNameTotalCount(ledgerHash framework.HashDigest, address string) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/user/accounts/%s/names/count", ledgerHash.ToBase58(), address))
	if err != nil {
		return info, err
	}

	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetUserEventsTotalCount(ledgerHash framework.HashDigest, address, eventName string) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/user/accounts/%s/names/%s/count", ledgerHash.ToBase58(), address, eventName))
	if err != nil {
		return info, err
	}

	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetLatestSystemEvent(ledgerHash framework.HashDigest, eventName string) (info ledger_model.Event, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/system/names/%s/latest", ledgerHash.ToBase58(), eventName))
	if err != nil {
		return info, err
	}

	return parseEvent(wrp.(map[string]interface{})), nil
}

func (r RestyQueryService) GetLatestUserEvent(ledgerHash framework.HashDigest, address string, eventName string) (info ledger_model.Event, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/user/accounts/%s/names/%s/latest", ledgerHash.ToBase58(), address, eventName))
	if err != nil {
		return info, err
	}

	return parseEvent(wrp.(map[string]interface{})), nil
}

func (r RestyQueryService) GetLatestBlock(ledgerHash framework.HashDigest) (info ledger_model.LedgerBlock, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/latest", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return parseBlock(ledgerHash, wrp.(map[string]interface{}))
}

func (r RestyQueryService) GetAdditionalTransactionCountByHeight(ledgerHash framework.HashDigest, blockHeight int64) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/txs/additional-count", ledgerHash.ToBase58(), blockHeight))
	if err != nil {
		return info, err
	}

	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetAdditionalTransactionCountByHash(ledgerHash, blockHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/txs/additional-count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetAdditionalTransactionCount(ledgerHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/txs/additional-count", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetAdditionalDataAccountCountByHeight(ledgerHash framework.HashDigest, blockHeight int64) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/accounts/additional-count", ledgerHash.ToBase58(), blockHeight))
	if err != nil {
		return info, err
	}

	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetAdditionalDataAccountCountByHash(ledgerHash, blockHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/accounts/additional-count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetAdditionalDataAccountCount(ledgerHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/accounts/additional-count", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetAdditionalUserCountByHeight(ledgerHash framework.HashDigest, blockHeight int64) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/users/additional-count", ledgerHash.ToBase58(), blockHeight))
	if err != nil {
		return info, err
	}

	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetAdditionalUserCountByHash(ledgerHash, blockHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/users/additional-count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetAdditionalUserCount(ledgerHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/users/additional-count", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetAdditionalContractCountByHeight(ledgerHash framework.HashDigest, blockHeight int64) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/contracts/additional-count", ledgerHash.ToBase58(), blockHeight))
	if err != nil {
		return info, err
	}

	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetAdditionalContractCountByHash(ledgerHash, blockHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/contracts/additional-count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetAdditionalContractCount(ledgerHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/contracts/additional-count", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetLedgersCount() (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/count"))
	if err != nil {
		return info, err
	}

	return int64(wrp.(float64)), nil
}

func (r RestyQueryService) GetRolePrivileges(ledgerHash framework.HashDigest, roleName string) (info ledger_model.RolePrivileges, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/authorization/role/%s", ledgerHash.ToBase58(), roleName))
	if err != nil {
		return info, err
	}

	return parsePrivilegeSet(wrp.(map[string]interface{}))
}

func parsePrivilegeSet(rolePrivileges map[string]interface{}) (info ledger_model.RolePrivileges, err error) {
	info = ledger_model.RolePrivileges{
		RoleName: rolePrivileges["roleName"].(string),
		Version:  int64(rolePrivileges["version"].(float64)),
	}

	transactionPrivilegeMap, ok := rolePrivileges["transactionPrivilege"].(map[string]interface{})
	if ok {
		transactionPrivilege := ledger_model.TransactionPrivilegeBitset{
			PermissionCount: int32(transactionPrivilegeMap["permissionCount"].(float64)),
		}
		transactionPermissions := make([]ledger_model.TransactionPermission, transactionPrivilege.PermissionCount)
		permissions := transactionPrivilegeMap["privilege"].([]interface{})
		for i := int32(0); i < transactionPrivilege.PermissionCount; i++ {
			transactionPermissions[i] = ledger_model.DIRECT_OPERATION.GetValueByName(permissions[i].(string)).(ledger_model.TransactionPermission)
		}
		transactionPrivilege.Privilege = transactionPermissions

		info.TransactionPrivilege = transactionPrivilege
	}
	ledgerPrivilegeMap, ok := rolePrivileges["ledgerPrivilege"].(map[string]interface{})
	if ok {
		ledgerPrivilege := ledger_model.LedgerPrivilegeBitset{
			PermissionCount: int32(ledgerPrivilegeMap["permissionCount"].(float64)),
		}
		ledgerPermissions := make([]ledger_model.LedgerPermission, ledgerPrivilege.PermissionCount)
		permissions := ledgerPrivilegeMap["privilege"].([]interface{})
		for i := int32(0); i < ledgerPrivilege.PermissionCount; i++ {
			ledgerPermissions[i] = ledger_model.CONFIGURE_ROLES.GetValueByName(permissions[i].(string)).(ledger_model.LedgerPermission)
		}
		ledgerPrivilege.Privilege = ledgerPermissions

		info.LedgerPrivilege = ledgerPrivilege
	}

	return
}

func (r RestyQueryService) GetUserPrivileges(ledgerHash framework.HashDigest, userAddress string) (info ledger_model.UserRolesPrivileges, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/authorization/user/%s", ledgerHash.ToBase58(), userAddress))
	if err != nil {
		return info, err
	}

	userPrivilegeMap := wrp.(map[string]interface{})
	userRoles := parseStringArray(userPrivilegeMap["userRole"].([]interface{}))
	info = ledger_model.UserRolesPrivileges{
		UserAddress: base58.MustDecode(userAddress),
		UserRoles:   userRoles,
	}
	ledgerPrivilegesBitsetMap, ok := userPrivilegeMap["ledgerPrivilegesBitset"]
	if ok {
		ledgerPrivilegesInfo := ledgerPrivilegesBitsetMap.(map[string]interface{})
		ledgerPrivilegesBitset := ledger_model.LedgerPrivilegeBitset{
			PermissionCount: int32(ledgerPrivilegesInfo["permissionCount"].(float64)),
			Privilege:       parseLedgerPermissions(ledgerPrivilegesInfo["privilege"].([]interface{})),
		}
		info.LedgerPrivilegesBitset = ledgerPrivilegesBitset
	}
	transactionPrivilegesBitsetMap, ok := userPrivilegeMap["transactionPrivilegesBitset"]
	if ok {
		transactionPrivilegesBitsetInfo := transactionPrivilegesBitsetMap.(map[string]interface{})
		transactionPrivilegeBitset := ledger_model.TransactionPrivilegeBitset{
			PermissionCount: int32(transactionPrivilegesBitsetInfo["permissionCount"].(float64)),
			Privilege:       parseTransactionPermissions(transactionPrivilegesBitsetInfo["privilege"].([]interface{})),
		}
		info.TransactionPrivilegesBitset = transactionPrivilegeBitset
	}

	return
}
