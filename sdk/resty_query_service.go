package sdk

import (
	"errors"
	"fmt"
	binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/ledger_model"
	"github.com/blockchain-jd-com/framework-go/utils/base58"
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
	"github.com/blockchain-jd-com/framework-go/utils/network"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
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

func (r RestyQueryService) query(url string) (data gjson.Result, err error) {
	resp, err := r.client.R().Get(r.baseUrl + url)
	if err != nil {
		return
	}
	if !resp.IsSuccess() {
		return data, errors.New(resp.String())
	}

	wrp := gjson.ParseBytes(resp.Body())
	// 数据不存在
	if wrp.Type == gjson.Null {
		return data, errors.New("empty data")
	}
	if !wrp.Get("success").Bool() {
		return data, errors.New(fmt.Sprintf("error code:%d msg:%s", wrp.Get("error.errorCode").Int(), wrp.Get("error.errorMessage").String()))
	}
	return wrp.Get("data"), nil
}

func (r RestyQueryService) queryWithParams(url string, params map[string]string) (data gjson.Result, err error) {
	resp, err := r.client.R().SetQueryParams(params).Get(r.baseUrl + url)
	if err != nil {
		return data, err
	}
	if !resp.IsSuccess() {
		return data, errors.New(resp.String())
	}

	wrp := gjson.ParseBytes(resp.Body())
	// 数据不存在
	if wrp.Type == gjson.Null {
		return data, errors.New("empty data")
	}
	if !wrp.Get("success").Bool() {
		return data, errors.New(fmt.Sprintf("error code:%d msg:%s", wrp.Get("error.errorCode").Int(), wrp.Get("error.errorMessage").String()))
	}
	return wrp.Get("data"), nil
}

func (r RestyQueryService) queryWithParamsFromValues(url string, params url.Values) (data gjson.Result, err error) {
	resp, err := r.client.R().SetQueryParamsFromValues(params).Get(r.baseUrl + url)
	if err != nil {
		return
	}
	if !resp.IsSuccess() {
		return data, errors.New(resp.String())
	}

	wrp := gjson.ParseBytes(resp.Body())
	// 数据不存在
	if wrp.Type == gjson.Null {
		return data, errors.New("empty data")
	}
	if !wrp.Get("success").Bool() {
		return data, errors.New(fmt.Sprintf("error code:%d msg:%s", wrp.Get("error.errorCode").Int(), wrp.Get("error.errorMessage").String()))
	}
	return wrp.Get("data"), nil
}

func (r RestyQueryService) queryWithBody(url string, params interface{}) (data gjson.Result, err error) {
	resp, err := r.client.R().SetBody(params).Post(r.baseUrl + url)
	if err != nil {
		return
	}
	if !resp.IsSuccess() {
		return data, errors.New(resp.String())
	}

	wrp := gjson.ParseBytes(resp.Body())
	// 数据不存在
	if wrp.Type == gjson.Null {
		return data, errors.New("empty data")
	}
	if !wrp.Get("success").Bool() {
		return data, errors.New(fmt.Sprintf("error code:%d msg:%s", wrp.Get("error.errorCode").Int(), wrp.Get("error.errorMessage").String()))
	}
	return wrp.Get("data"), nil
}

func (r RestyQueryService) GetLedgerHashs() ([]framework.HashDigest, error) {
	wrp, err := r.query("/ledgers")
	if err != nil {
		return nil, err
	}
	ledgers := wrp.Array()
	hashs := make([]framework.HashDigest, len(ledgers))
	for i, m := range ledgers {
		hashs[i] = framework.ParseHashDigest(base58.MustDecode(m.String()))
	}
	return hashs, nil
}

func (r RestyQueryService) GetLedger(ledgerHash framework.HashDigest) (info ledger_model.LedgerInfo, err error) {
	wrp, err := r.query("/ledgers/" + ledgerHash.ToBase58())
	if err != nil {
		return info, err
	}
	info.Hash = framework.ParseHashDigest(base58.MustDecode(wrp.Get("hash").String()))
	info.LatestBlockHash = framework.ParseHashDigest(base58.MustDecode(wrp.Get("latestBlockHash").String()))
	info.LatestBlockHeight = wrp.Get("latestBlockHeight").Int()
	return info, nil
}

func (r RestyQueryService) GetLedgerAdminInfo(ledgerHash framework.HashDigest) (info ledger_model.LedgerAdminInfo, err error) {
	wrp, err := r.query("/ledgers/" + ledgerHash.ToBase58() + "/admininfo")
	if err != nil {
		return info, err
	}
	info.ParticipantCount = wrp.Get("participantCount").Int()
	metadata := wrp.Get("metadata")
	info.Metadata = ledger_model.LedgerMetadata_V2{
		LedgerMetadata: ledger_model.LedgerMetadata{
			Seed:             []byte(metadata.Get("seed").String()),
			ParticipantsHash: base58.MustDecode(metadata.Get("participantsHash").String()),
			SettingsHash:     base58.MustDecode(metadata.Get("settingsHash").String()),
		},
		RolePrivilegesHash: base58.MustDecode(metadata.Get("rolePrivilegesHash").String()),
		UserRolesHash:      base58.MustDecode(metadata.Get("userRolesHash").String()),
	}
	participants := wrp.Get("participants").Array()
	pNodes := make([]ledger_model.ParticipantNode, len(participants))
	for i, node := range participants {
		pNodes[i] = ledger_model.ParticipantNode{
			Id:                   int32(node.Get("id").Int()),
			Name:                 node.Get("name").String(),
			Address:              base58.MustDecode(node.Get("address.value").String()),
			PubKey:               base58.MustDecode(node.Get("pubKey").String()),
			ParticipantNodeState: ledger_model.READY.GetValueByName(node.Get("participantNodeState").String()).(ledger_model.ParticipantNodeState),
		}
	}
	info.Participants = pNodes
	autoVerifyHash := wrp.Get("settings.cryptoSetting.autoVerifyHash").Bool()
	supportedProviders := wrp.Get("settings.cryptoSetting.supportedProviders").Array()
	providers := make([]ledger_model.CryptoProvider, len(supportedProviders))
	for i, node := range supportedProviders {
		algorithmsArray := node.Get("algorithms").Array()
		algorithms := []framework.CryptoAlgorithm{}
		for _, ma := range algorithmsArray {
			if ma.Type != gjson.Null {
				algorithms = append(algorithms, framework.CryptoAlgorithm{
					Code: int16(ma.Get("code").Int()),
					Name: ma.Get("name").String(),
				})
			}
		}
		providers[i] = ledger_model.CryptoProvider{
			Name:       node.Get("name").String(),
			Algorithms: algorithms,
		}
	}
	info.Settings = ledger_model.LedgerSettings{
		ConsensusProvider: wrp.Get("settings.consensusProvider").String(),
		ConsensusSetting:  base58.MustDecode(wrp.Get("settings.consensusSetting.value").String()),
		CryptoSetting: ledger_model.CryptoSetting{
			SupportedProviders: providers,
			HashAlgorithm:      int16(wrp.Get("settings.cryptoSetting.hashAlgorithm").Int()),
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
	participants := wrp.Array()
	info = make([]ledger_model.ParticipantNode, len(participants))
	for i, node := range participants {
		info[i] = ledger_model.ParticipantNode{
			Id:                   int32(node.Get("id").Int()),
			Name:                 node.Get("name").String(),
			Address:              base58.MustDecode(node.Get("address.value").String()),
			PubKey:               base58.MustDecode(node.Get("pubKey").String()),
			ParticipantNodeState: ledger_model.READY.GetValueByName(node.Get("participantNodeState").String()).(ledger_model.ParticipantNodeState),
		}
	}
	return
}

func (r RestyQueryService) GetLedgerMetadata(ledgerHash framework.HashDigest) (info ledger_model.LedgerMetadata, err error) {
	metadata, err := r.query("/ledgers/" + ledgerHash.ToBase58() + "/metadata")
	if err != nil {
		return info, err
	}
	info = ledger_model.LedgerMetadata{
		Seed:             []byte(metadata.Get("seed").String()),
		ParticipantsHash: base58.MustDecode(metadata.Get("participantsHash").String()),
		SettingsHash:     base58.MustDecode(metadata.Get("settingsHash").String()),
	}

	return
}

func (r RestyQueryService) GetBlockByHeight(ledgerHash framework.HashDigest, height int64) (info ledger_model.LedgerBlock, err error) {
	block, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d", ledgerHash.ToBase58(), height))
	if err != nil {
		return info, err
	}

	var PreviousHash []byte
	if block.Get("previousHash").Exists() {
		PreviousHash = base58.MustDecode(block.Get("previousHash").String())
	} else {
		PreviousHash = nil
	}

	info = ledger_model.LedgerBlock{
		BlockBody: ledger_model.BlockBody{
			LedgerDataSnapshot: parseLedgerDataSnapshot(block),
			PreviousHash:       PreviousHash,
			Height:             block.Get("height").Int(),
			TransactionSetHash: base58.MustDecode(block.Get("transactionSetHash").String()),
			Timestamp:          block.Get("timestamp").Int(),
		},
		Hash: base58.MustDecode(block.Get("hash").String()),
	}

	if height > 0 {
		info.BlockBody.LedgerHash = base58.MustDecode(block.Get("ledgerHash").String())
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

	return parseBlock(ledgerHash, wrp)
}

func parseBlock(ledgerHash framework.HashDigest, block gjson.Result) (info ledger_model.LedgerBlock, err error) {
	var PreviousHash []byte
	if block.Get("previousHash").Exists() {
		PreviousHash = base58.MustDecode(block.Get("previousHash").String())
	} else {
		PreviousHash = nil
	}

	info = ledger_model.LedgerBlock{
		BlockBody: ledger_model.BlockBody{
			LedgerDataSnapshot: parseLedgerDataSnapshot(block),
			PreviousHash:       PreviousHash,
			Height:             block.Get("height").Int(),
			TransactionSetHash: base58.MustDecode(block.Get("transactionSetHash").String()),
			Timestamp:          block.Get("timestamp").Int(),
		},
		Hash: base58.MustDecode(block.Get("hash").String()),
	}

	if info.BlockBody.Height > 0 {
		info.BlockBody.LedgerHash = base58.MustDecode(block.Get("ledgerHash").String())
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
	return wrp.Int(), nil
}

func (r RestyQueryService) GetTransactionCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/txs/count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetTransactionTotalCount(ledgerHash framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/txs/count", ledgerHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetDataAccountCountByHeight(ledgerHash framework.HashDigest, height int64) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/accounts/count", ledgerHash.ToBase58(), height))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetDataAccountCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/accounts/count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetDataAccountTotalCount(ledgerHash framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/accounts/count", ledgerHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetUserCountByHeight(ledgerHash framework.HashDigest, height int64) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/users/count", ledgerHash.ToBase58(), height))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetUserCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/users/count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetUserTotalCount(ledgerHash framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/users/count", ledgerHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetContractCountByHeight(ledgerHash framework.HashDigest, height int64) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/contracts/count", ledgerHash.ToBase58(), height))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetContractCountByHash(ledgerHash, blockHash framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/contracts/count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetContractTotalCount(ledgerHash framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/contracts/count", ledgerHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetUser(ledgerHash framework.HashDigest, address string) (info ledger_model.UserInfo, err error) {
	user, err := r.query(fmt.Sprintf("/ledgers/%s/users/address/%s", ledgerHash.ToBase58(), address))
	if err != nil {
		return info, err
	}
	if !user.Exists() {
		return info, errors.New("not exists")
	}
	info = ledger_model.UserInfo{
		UserAccountHeader: ledger_model.UserAccountHeader{
			BlockchainIdentity: parseBlockchainIdentity(user),
		},
	}
	return
}

func (r RestyQueryService) GetDataAccount(ledgerHash framework.HashDigest, address string) (info ledger_model.BlockchainIdentity, err error) {
	id, err := r.query(fmt.Sprintf("/ledgers/%s/accounts/address/%s", ledgerHash.ToBase58(), address))
	if err != nil {
		return info, err
	}
	if !id.Exists() {
		return info, errors.New("not exists")
	}
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
	kvArray := wrp.Array()
	info = make([]ledger_model.TypedKVEntry, len(kvArray))
	for i, id := range kvArray {
		info[i] = ledger_model.TypedKVEntry{
			Key:     id.Get("key").String(),
			Value:   id.Get("value").String(),
			Version: id.Get("version").Int(),
			Type:    ledger_model.NIL.GetValueByName(id.Get("type").String()).(ledger_model.DataType),
		}
	}

	return
}

func (r RestyQueryService) GetDataEntries(ledgerHash framework.HashDigest, address string, kvInfoVO ledger_model.KVInfoVO) (info []ledger_model.TypedKVEntry, err error) {
	wrp, err := r.queryWithBody(fmt.Sprintf("/ledgers/%s/accounts/%s/entries-version", ledgerHash.ToBase58(), address), kvInfoVO)
	if err != nil {
		return info, err
	}
	kvArray := wrp.Array()
	info = make([]ledger_model.TypedKVEntry, len(kvArray))
	for i, id := range kvArray {
		info[i] = ledger_model.TypedKVEntry{
			Key:     id.Get("key").String(),
			Value:   id.Get("value").String(),
			Version: id.Get("version").Int(),
			Type:    ledger_model.NIL.GetValueByName(id.Get("type").String()).(ledger_model.DataType),
		}
	}

	return
}

func (r RestyQueryService) GetDataEntriesTotalCount(ledgerHash framework.HashDigest, address string) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/accounts/address/%s/entries/count", ledgerHash.ToBase58(), address))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
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

	kvArray := wrp.Array()
	info = make([]ledger_model.TypedKVEntry, len(kvArray))
	for i, id := range kvArray {
		info[i] = ledger_model.TypedKVEntry{
			Key:     id.Get("key").String(),
			Value:   id.Get("value").String(),
			Version: id.Get("version").Int(),
			Type:    ledger_model.NIL.GetValueByName(id.Get("type").String()).(ledger_model.DataType),
		}
	}

	return
}

func (r RestyQueryService) GetContract(ledgerHash framework.HashDigest, address string) (info ledger_model.ContractInfo, err error) {
	contract, err := r.query(fmt.Sprintf("/ledgers/%s/contracts/address/%s", ledgerHash.ToBase58(), address))
	if err != nil {
		return
	}
	if !contract.Exists() {
		return info, errors.New("not exists")
	}
	info = ledger_model.ContractInfo{
		BlockchainIdentity: parseBlockchainIdentity(contract),
		MerkleSnapshot: ledger_model.MerkleSnapshot{
			RootHash: base58.MustDecode(contract.Get("rootHash").String()),
		},
		ChainCode: bytes.StringToBytes(contract.Get("chainCode").String()),
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
	idArray := wrp.Array()
	info = make([]ledger_model.BlockchainIdentity, len(idArray))
	for i, id := range idArray {
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
	idArray := wrp.Array()
	info = make([]ledger_model.BlockchainIdentity, len(idArray))
	for i, id := range idArray {
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
	idArray := wrp.Array()
	info = make([]ledger_model.BlockchainIdentity, len(idArray))
	for i, id := range idArray {
		info[i] = parseBlockchainIdentity(id)
	}

	return
}

func (r RestyQueryService) GetUserRoles(ledgerHash framework.HashDigest, userAddress string) (info ledger_model.RoleSet, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/user-role/%s", ledgerHash.ToBase58(), userAddress))
	if err != nil {
		return info, err
	}

	rolesArray := wrp.Get("roleSet").Array()
	roles := make([]string, len(rolesArray))
	for i, role := range rolesArray {
		roles[i] = role.String()
	}
	info = ledger_model.RoleSet{
		Policy: ledger_model.UNION.GetValueByName(wrp.Get("policy").String()).(ledger_model.RolesPolicy),
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
	idArray := wrp.Array()
	info = make([]ledger_model.Event, len(idArray))
	for i, item := range idArray {
		info[i] = parseEvent(item)
	}

	return
}

func parseBytesValue(info gjson.Result) ledger_model.BytesValue {
	return ledger_model.BytesValue{
		Type:  ledger_model.NIL.GetValueByName(info.Get("type").String()).(ledger_model.DataType),
		Bytes: base58.MustDecode(info.Get("bytes.value").String()),
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
	idArray := wrp.Array()
	info = make([]ledger_model.BlockchainIdentity, len(idArray))
	for i, id := range idArray {
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
	idArray := wrp.Array()
	info = make([]ledger_model.Event, len(idArray))
	for i, item := range idArray {
		info[i] = parseEvent(item)
	}

	return
}

func parseEvent(event gjson.Result) ledger_model.Event {
	info := ledger_model.Event{
		Name:        event.Get("name").String(),
		BlockHeight: event.Get("blockHeight").Int(),
		Sequence:    event.Get("sequence").Int(),
	}
	if event.Get("transactionSource").Exists() {
		info.TransactionSource = base58.MustDecode(event.Get("transactionSource").String())
	}
	if event.Get("contractSource").Exists() {
		info.ContractSource = event.Get("contractSource").String()
	}

	if event.Get("eventAccount").Exists() {
		info.EventAccount = base58.MustDecode(event.Get("eventAccount.value").String())
	}
	if !event.Get("content.nil").Bool() {
		info.Content = parseBytesValue(event.Get("content"))
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
	txArray := wrp.Array()
	info = make([]ledger_model.LedgerTransaction, len(txArray))
	for i, tx := range txArray {
		info[i] = ledger_model.LedgerTransaction{
			Request: parseTxRequest(tx.Get("request")),
			Result:  parseTxResult(tx.Get("result")),
		}
	}

	return
}

func parseTxRequest(request gjson.Result) ledger_model.TransactionRequest {
	return ledger_model.TransactionRequest{
		TransactionHash:    base58.MustDecode(request.Get("transactionHash").String()),
		TransactionContent: parseTransactionContent(request.Get("transactionContent")),
		NodeSignatures:     parseSignatures(request.Get("nodeSignatures").Array()),
		EndpointSignatures: parseSignatures(request.Get("endpointSignatures").Array()),
	}
}

func parseTxResult(result gjson.Result) ledger_model.TransactionResult {
	return ledger_model.TransactionResult{
		BlockHeight:     result.Get("blockHeight").Int(),
		ExecutionState:  ledger_model.SUCCESS.GetValueByName(result.Get("executionState").String()).(ledger_model.TransactionState),
		TransactionHash: base58.MustDecode(result.Get("transactionHash").String()),
		DataSnapshot:    parseLedgerDataSnapshot(result.Get("dataSnapshot")),
	}
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
	txArray := wrp.Array()
	info = make([]ledger_model.LedgerTransaction, len(txArray))
	for i, tx := range txArray {
		info[i] = ledger_model.LedgerTransaction{
			Request: parseTxRequest(tx.Get("request")),
			Result:  parseTxResult(tx.Get("result")),
		}
	}

	return
}

func (r RestyQueryService) GetTransactionByContentHash(ledgerHash, contentHash framework.HashDigest) (info ledger_model.LedgerTransaction, err error) {
	tx, err := r.query(fmt.Sprintf("/ledgers/%s/txs/hash/%s", ledgerHash.ToBase58(), contentHash.ToBase58()))
	if err != nil {
		return info, err
	}
	return ledger_model.LedgerTransaction{
		Request: parseTxRequest(tx.Get("request")),
		Result:  parseTxResult(tx.Get("result")),
	}, nil
}

func (r RestyQueryService) GetTransactionStateByContentHash(ledgerHash, contentHash framework.HashDigest) (info ledger_model.TransactionState, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/txs/state/%s", ledgerHash.ToBase58(), contentHash.ToBase58()))
	if err != nil {
		return info, err
	}
	if !wrp.Exists() {
		return info, errors.New("not exists")
	}
	info = ledger_model.SUCCESS.GetValueByName(wrp.String()).(ledger_model.TransactionState)

	return
}

func parseLedgerDataSnapshot(info gjson.Result) ledger_model.LedgerDataSnapshot {
	var adminAccountHash []byte
	if info.Get("adminAccountHash").Exists() {
		adminAccountHash = base58.MustDecode(info.Get("adminAccountHash").String())
	} else {
		adminAccountHash = nil
	}
	var userAccountSetHash []byte
	if info.Get("userAccountSetHash").Exists() {
		userAccountSetHash = base58.MustDecode(info.Get("userAccountSetHash").String())
	} else {
		userAccountSetHash = nil
	}
	var dataAccountSetHash []byte
	if info.Get("dataAccountSetHash").Exists() {
		dataAccountSetHash = base58.MustDecode(info.Get("dataAccountSetHash").String())
	} else {
		dataAccountSetHash = nil
	}
	var contractAccountSetHash []byte
	if info.Get("contractAccountSetHash").Exists() {
		contractAccountSetHash = base58.MustDecode(info.Get("contractAccountSetHash").String())
	} else {
		contractAccountSetHash = nil
	}

	return ledger_model.LedgerDataSnapshot{
		AdminAccountHash:       adminAccountHash,
		UserAccountSetHash:     userAccountSetHash,
		DataAccountSetHash:     dataAccountSetHash,
		ContractAccountSetHash: contractAccountSetHash,
	}
}

func parseTransactionContent(info gjson.Result) ledger_model.TransactionContent {
	return ledger_model.TransactionContent{
		Operations: parseOperations(info.Get("operations").Array()),
		Timestamp:  info.Get("timestamp").Int(),
	}
}

func parseSignatures(info []gjson.Result) []ledger_model.DigitalSignature {
	signatures := make([]ledger_model.DigitalSignature, len(info))
	for i, sign := range info {
		signatures[i] = ledger_model.DigitalSignature{
			DigitalSignatureBody: ledger_model.DigitalSignatureBody{
				PubKey: base58.MustDecode(sign.Get("pubKey").String()),
				Digest: base58.MustDecode(sign.Get("digest").String()),
			},
		}
	}

	return signatures
}

func parseOperations(info []gjson.Result) []binary_proto.DataContract {
	operations := make([]binary_proto.DataContract, len(info))
	for i, operation := range info {
		var dc binary_proto.DataContract
		if operation.Get("userID").Exists() {
			// 注册用户
			dc = parseUserRegisterOperation(operation.Get("userID"))
		} else if operation.Get("accountID").Exists() {
			// 注册数据账户
			dc = parseDataAccountRegisterOperation(operation.Get("accountID"))
		} else if operation.Get("writeSet").Exists() {
			// KV写入
			dc = parseDataAccountKVSetOperation(operation)
		} else if operation.Get("eventAccountID").Exists() {
			// 事件账户注册
			dc = parseEventAccountRegisterOperation(operation.Get("eventAccountID"))
		} else if operation.Get("events").Exists() {
			// 发布事件
			dc = parseEventPublishOperation(operation)
		} else if operation.Get("participantRegisterIdentity").Exists() {
			// 注册参与方
			dc = parseParticipantRegisterOperation(operation.Get("participantRegisterIdentity"))
		} else if operation.Get("stateUpdateIdentity").Exists() {
			// 参与方状态变更
			dc = parseParticipantStateUpdateOperation(operation.Get("stateUpdateIdentity"))
		} else if operation.Get("chainCode").Exists() {
			// 合约部署
			dc = parseContractCodeDeployOperation(operation.Get("chainCode"))
		} else if operation.Get("roles").Exists() {
			// 角色配置
			dc = parseRolesConfigureOperation(operation)
		} else if operation.Get("userRolesAuthorizations").Exists() {
			dc = parseUserAuthorizeOperation(operation.Get("userRolesAuthorizations").Array())
		}
		operations[i] = dc
	}

	return operations
}

func parseUserRegisterOperation(info gjson.Result) binary_proto.DataContract {
	return &ledger_model.UserRegisterOperation{
		UserID: parseBlockchainIdentity(info),
	}
}

func parseDataAccountRegisterOperation(info gjson.Result) binary_proto.DataContract {
	return &ledger_model.DataAccountRegisterOperation{
		AccountID: parseBlockchainIdentity(info),
	}
}

func parseEventAccountRegisterOperation(info gjson.Result) binary_proto.DataContract {
	return &ledger_model.EventAccountRegisterOperation{
		EventAccountID: parseBlockchainIdentity(info),
	}
}

func parseParticipantRegisterOperation(info gjson.Result) binary_proto.DataContract {
	return &ledger_model.ParticipantRegisterOperation{
		ParticipantName:             info.Get("participantName").String(),
		ParticipantRegisterIdentity: parseBlockchainIdentity(info.Get("participantRegisterIdentity")),
	}
}

func parseParticipantStateUpdateOperation(info gjson.Result) binary_proto.DataContract {
	networkAddress := info.Get("networkAddress")
	address := network.NewAddress(networkAddress.Get("host").String(), int32(networkAddress.Get("port").Int()), networkAddress.Get("secure").Bool())
	return &ledger_model.ParticipantStateUpdateOperation{
		State:               ledger_model.READY.GetValueByName(info.Get("state").String()).(ledger_model.ParticipantNodeState),
		StateUpdateIdentity: parseBlockchainIdentity(info.Get("stateUpdateIdentity")),
		NetworkAddress:      address.ToBytes(),
	}
}

func parseContractCodeDeployOperation(info gjson.Result) binary_proto.DataContract {
	return &ledger_model.ContractCodeDeployOperation{
		ContractID: parseBlockchainIdentity(info.Get("contractID")),
		ChainCode:  bytes.StringToBytes(info.Get("chainCode").String()),
	}
}

func parseRolesConfigureOperation(info gjson.Result) binary_proto.DataContract {
	var roles []ledger_model.RolePrivilegeEntry
	if !info.Get("empty").Bool() {
		array := info.Get("roles").Array()
		roles = make([]ledger_model.RolePrivilegeEntry, len(array))
		for i, role := range array {
			roleName := role.Get("roleName").String()
			roles[i] = ledger_model.RolePrivilegeEntry{
				RoleName:                      roleName,
				EnableLedgerPermissions:       parseLedgerPermissions(role.Get("enableLedgerPermissions").Array()),
				DisableLedgerPermissions:      parseLedgerPermissions(role.Get("disableLedgerPermissions").Array()),
				EnableTransactionPermissions:  parseTransactionPermissions(role.Get("enableTransactionPermissions").Array()),
				DisableTransactionPermissions: parseTransactionPermissions(role.Get("disableTransactionPermissions").Array()),
			}
		}
	}
	return &ledger_model.RolesConfigureOperation{
		Roles: roles,
	}
}

func parseUserAuthorizeOperation(array []gjson.Result) binary_proto.DataContract {
	userAuthor := make([]ledger_model.UserRolesEntry, len(array))
	for i, author := range array {
		addressArray := author.Get("userAddresses").Array()
		address := make([][]byte, len(addressArray))
		for j, addr := range addressArray {
			address[j] = base58.MustDecode(addr.Get("value").String())
		}
		userAuthor[i] = ledger_model.UserRolesEntry{
			Policy:            ledger_model.UNION.GetValueByName(author.Get("policy").String()).(ledger_model.RolesPolicy),
			Addresses:         address,
			UnauthorizedRoles: parseStringArray(author.Get("unauthorizedRoles").Array()),
			AuthorizedRoles:   parseStringArray(author.Get("authorizedRoles").Array()),
		}
	}
	return &ledger_model.UserAuthorizeOperation{
		UserRolesAuthorizations: userAuthor,
	}
}

func parseStringArray(info []gjson.Result) []string {
	array := make([]string, len(info))
	for i, item := range info {
		array[i] = item.String()
	}

	return array
}

func parseLedgerPermissions(array []gjson.Result) []ledger_model.LedgerPermission {
	lps := make([]ledger_model.LedgerPermission, len(array))
	for i, item := range array {
		lps[i] = ledger_model.CONFIGURE_ROLES.GetValueByName(item.String()).(ledger_model.LedgerPermission)
	}

	return lps
}

func parseTransactionPermissions(array []gjson.Result) []ledger_model.TransactionPermission {
	tps := make([]ledger_model.TransactionPermission, len(array))
	for i, item := range array {
		tps[i] = ledger_model.DIRECT_OPERATION.GetValueByName(item.String()).(ledger_model.TransactionPermission)
	}

	return tps
}

func parseDataAccountKVSetOperation(info gjson.Result) binary_proto.DataContract {
	kvs := info.Get("writeSet").Array()
	writeSet := make([]ledger_model.KVWriteEntry, len(kvs))
	for i, kv := range kvs {
		kvSet := ledger_model.KVWriteEntry{
			Key:             kv.Get("key").String(),
			Value:           parseBytesValue(kv.Get("value")),
			ExpectedVersion: kv.Get("expectedVersion").Int(),
		}
		writeSet[i] = kvSet
	}
	return &ledger_model.DataAccountKVSetOperation{
		AccountAddress: base58.MustDecode(info.Get("accountAddress.value").String()),
		WriteSet:       writeSet,
	}
}

func parseEventPublishOperation(info gjson.Result) binary_proto.DataContract {
	kvs := info.Get("events").Array()
	writeSet := make([]ledger_model.EventEntry, len(kvs))
	for i, kv := range kvs {
		kvSet := ledger_model.EventEntry{
			Name:     kv.Get("name").String(),
			Content:  parseBytesValue(kv.Get("content")),
			Sequence: kv.Get("sequence").Int(),
		}
		writeSet[i] = kvSet
	}
	return &ledger_model.EventPublishOperation{
		EventAddress: base58.MustDecode(info.Get("eventAddress.value").String()),
		Events:       writeSet,
	}
}

func parseBlockchainIdentity(id gjson.Result) ledger_model.BlockchainIdentity {
	return ledger_model.BlockchainIdentity{
		Address: base58.MustDecode(id.Get("address.value").String()),
		PubKey:  base58.MustDecode(id.Get("pubKey").String()),
	}
}

func (r RestyQueryService) GetSystemEventNameTotalCount(ledgerHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/system/names/count", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
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

	return parseStringArray(wrp.Array()), nil
}

func (r RestyQueryService) GetSystemEventsTotalCount(ledgerHash framework.HashDigest, eventName string) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/system/names/%s/count", ledgerHash.ToBase58(), eventName))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetUserEventAccount(ledgerHash framework.HashDigest, address string) (info ledger_model.BlockchainIdentity, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/user/accounts/%s", ledgerHash.ToBase58(), address))
	if err != nil {
		return info, err
	}

	return parseBlockchainIdentity(wrp), nil
}

func (r RestyQueryService) GetUserEventAccountTotalCount(ledgerHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/user/accounts/count", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
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

	return parseStringArray(wrp.Array()), nil

}

func (r RestyQueryService) GetUserEventNameTotalCount(ledgerHash framework.HashDigest, address string) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/user/accounts/%s/names/count", ledgerHash.ToBase58(), address))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetUserEventsTotalCount(ledgerHash framework.HashDigest, address, eventName string) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/user/accounts/%s/names/%s/count", ledgerHash.ToBase58(), address, eventName))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetLatestSystemEvent(ledgerHash framework.HashDigest, eventName string) (info ledger_model.Event, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/system/names/%s/latest", ledgerHash.ToBase58(), eventName))
	if err != nil {
		return info, err
	}

	return parseEvent(wrp), nil
}

func (r RestyQueryService) GetLatestUserEvent(ledgerHash framework.HashDigest, address string, eventName string) (info ledger_model.Event, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/user/accounts/%s/names/%s/latest", ledgerHash.ToBase58(), address, eventName))
	if err != nil {
		return info, err
	}

	return parseEvent(wrp), nil
}

func (r RestyQueryService) GetLatestBlock(ledgerHash framework.HashDigest) (info ledger_model.LedgerBlock, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/latest", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return parseBlock(ledgerHash, wrp)
}

func (r RestyQueryService) GetAdditionalTransactionCountByHeight(ledgerHash framework.HashDigest, blockHeight int64) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/txs/additional-count", ledgerHash.ToBase58(), blockHeight))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalTransactionCountByHash(ledgerHash, blockHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/txs/additional-count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalTransactionCount(ledgerHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/txs/additional-count", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalDataAccountCountByHeight(ledgerHash framework.HashDigest, blockHeight int64) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/accounts/additional-count", ledgerHash.ToBase58(), blockHeight))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalDataAccountCountByHash(ledgerHash, blockHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/accounts/additional-count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalDataAccountCount(ledgerHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/accounts/additional-count", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalUserCountByHeight(ledgerHash framework.HashDigest, blockHeight int64) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/users/additional-count", ledgerHash.ToBase58(), blockHeight))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalUserCountByHash(ledgerHash, blockHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/users/additional-count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalUserCount(ledgerHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/users/additional-count", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalContractCountByHeight(ledgerHash framework.HashDigest, blockHeight int64) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/contracts/additional-count", ledgerHash.ToBase58(), blockHeight))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalContractCountByHash(ledgerHash, blockHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/contracts/additional-count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalContractCount(ledgerHash framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/contracts/additional-count", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetLedgersCount() (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/count"))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetRolePrivileges(ledgerHash framework.HashDigest, roleName string) (info ledger_model.RolePrivileges, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/authorization/role/%s", ledgerHash.ToBase58(), roleName))
	if err != nil {
		return info, err
	}

	return parsePrivilegeSet(wrp)
}

func parsePrivilegeSet(rolePrivileges gjson.Result) (info ledger_model.RolePrivileges, err error) {
	info = ledger_model.RolePrivileges{
		RoleName: rolePrivileges.Get("roleName").String(),
		Version:  rolePrivileges.Get("version").Int(),
	}

	if rolePrivileges.Get("transactionPrivilege").Exists() {
		transactionPrivilege := ledger_model.TransactionPrivilegeBitset{
			PermissionCount: int32(rolePrivileges.Get("transactionPrivilege.permissionCount").Int()),
		}
		transactionPermissions := make([]ledger_model.TransactionPermission, transactionPrivilege.PermissionCount)
		permissions := rolePrivileges.Get("transactionPrivilege.privilege").Array()
		for i := int32(0); i < transactionPrivilege.PermissionCount; i++ {
			transactionPermissions[i] = ledger_model.DIRECT_OPERATION.GetValueByName(permissions[i].String()).(ledger_model.TransactionPermission)
		}
		transactionPrivilege.Privilege = transactionPermissions

		info.TransactionPrivilege = transactionPrivilege
	}
	if rolePrivileges.Get("ledgerPrivilege").Exists() {
		ledgerPrivilege := ledger_model.LedgerPrivilegeBitset{
			PermissionCount: int32(rolePrivileges.Get("ledgerPrivilege.permissionCount").Int()),
		}
		ledgerPermissions := make([]ledger_model.LedgerPermission, ledgerPrivilege.PermissionCount)
		permissions := rolePrivileges.Get("ledgerPrivilege.privilege").Array()
		for i := int32(0); i < ledgerPrivilege.PermissionCount; i++ {
			ledgerPermissions[i] = ledger_model.CONFIGURE_ROLES.GetValueByName(permissions[i].String()).(ledger_model.LedgerPermission)
		}
		ledgerPrivilege.Privilege = ledgerPermissions

		info.LedgerPrivilege = ledgerPrivilege
	}

	return
}

func (r RestyQueryService) GetUserPrivileges(ledgerHash framework.HashDigest, userAddress string) (info ledger_model.UserRolesPrivileges, err error) {
	userPrivilegeMap, err := r.query(fmt.Sprintf("/ledgers/%s/authorization/user/%s", ledgerHash.ToBase58(), userAddress))
	if err != nil {
		return info, err
	}

	userRoles := parseStringArray(userPrivilegeMap.Get("userRole").Array())
	info = ledger_model.UserRolesPrivileges{
		UserAddress: base58.MustDecode(userAddress),
		UserRoles:   userRoles,
	}
	if userPrivilegeMap.Get("ledgerPrivilegesBitset").Exists() {
		ledgerPrivilegesBitset := ledger_model.LedgerPrivilegeBitset{
			PermissionCount: int32(userPrivilegeMap.Get("ledgerPrivilegesBitset.permissionCount").Int()),
			Privilege:       parseLedgerPermissions(userPrivilegeMap.Get("ledgerPrivilegesBitset.privilege").Array()),
		}
		info.LedgerPrivilegesBitset = ledgerPrivilegesBitset
	}
	if userPrivilegeMap.Get("transactionPrivilegesBitset").Exists() {
		transactionPrivilegeBitset := ledger_model.TransactionPrivilegeBitset{
			PermissionCount: int32(userPrivilegeMap.Get("transactionPrivilegesBitset.permissionCount").Int()),
			Privilege:       parseTransactionPermissions(userPrivilegeMap.Get("transactionPrivilegesBitset.privilege").Array()),
		}
		info.TransactionPrivilegesBitset = transactionPrivilegeBitset
	}

	return
}

func (r RestyQueryService) GetAdditionalTransactionsByHeight(ledgerHash framework.HashDigest, height int64, fromIndex, count int64) (info []ledger_model.LedgerTransaction, err error) {
	params := map[string]string{
		"fromIndex": strconv.FormatInt(fromIndex, 10),
		"count":     strconv.FormatInt(count, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/blocks/height/%d/txs/additional-txs", ledgerHash.ToBase58(), height), params)
	if err != nil {
		return info, err
	}
	txArray := wrp.Array()
	info = make([]ledger_model.LedgerTransaction, len(txArray))
	for i, tx := range txArray {
		info[i] = ledger_model.LedgerTransaction{
			Request: parseTxRequest(tx.Get("request")),
			Result:  parseTxResult(tx.Get("result")),
		}
	}

	return
}

func (r RestyQueryService) GetAdditionalTransactionsByHash(ledgerHash, blockHash framework.HashDigest, fromIndex, count int64) (info []ledger_model.LedgerTransaction, err error) {
	params := map[string]string{
		"fromIndex": strconv.FormatInt(fromIndex, 10),
		"count":     strconv.FormatInt(count, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/txs/additional-txs", ledgerHash.ToBase58(), blockHash.ToBase58()), params)
	if err != nil {
		return info, err
	}
	txArray := wrp.Array()
	info = make([]ledger_model.LedgerTransaction, len(txArray))
	for i, tx := range txArray {
		info[i] = ledger_model.LedgerTransaction{
			Request: parseTxRequest(tx.Get("request")),
			Result:  parseTxResult(tx.Get("result")),
		}
	}

	return
}