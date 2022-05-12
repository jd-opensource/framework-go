package sdk

import (
	gbytes "bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/gmsm/gmtls"
	gmx509 "github.com/blockchain-jd-com/framework-go/gmsm/x509"
	"github.com/blockchain-jd-com/framework-go/ledger_model"
	"github.com/blockchain-jd-com/framework-go/utils/base58"
	"github.com/blockchain-jd-com/framework-go/utils/base64"
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

/*
 * Author: imuge
 * Date: 2020/5/27 下午5:55
 */

var _ ledger_model.BlockChainLedgerQueryService = (*RestyQueryService)(nil)

type RestyQueryService struct {
	host       string
	port       int
	secure     bool
	gmSecure   bool
	client     *resty.Client
	gmClient   *http.Client // 国密TLS客户端
	baseUrl    string
	security   *SSLSecurity
	gmSecurity *GMSSLSecurity
}

func NewRestyQueryService(host string, port int) *RestyQueryService {
	baseUrl := fmt.Sprintf("http://%s:%d", host, port)
	return &RestyQueryService{
		host:    host,
		port:    port,
		secure:  false,
		client:  resty.New(),
		baseUrl: baseUrl,
	}
}

func NewSecureRestyQueryService(host string, port int, security *SSLSecurity) *RestyQueryService {
	baseUrl := fmt.Sprintf("https://%s:%d", host, port)
	r := &RestyQueryService{
		host:     host,
		port:     port,
		secure:   true,
		client:   resty.New(),
		baseUrl:  baseUrl,
		security: security,
	}
	if r.security != nil {
		r.client.SetTLSClientConfig(&tls.Config{
			RootCAs:      r.security.RootCerts,
			Certificates: r.security.ClientCerts,
		})
	} else {
		r.client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return r
}

func NewGMSecureRestyQueryService(host string, port int, security *GMSSLSecurity) *RestyQueryService {
	baseUrl := fmt.Sprintf("https://%s:%d", host, port)

	var certPool *gmx509.CertPool
	if security != nil {
		certPool = security.RootCerts
	}
	var certificates []gmtls.Certificate
	if nil != security && nil != security.EncCert && nil != security.SigCert {
		certificates = []gmtls.Certificate{*security.SigCert, *security.EncCert}
	} else {
		certificates = []gmtls.Certificate{}
	}
	config := &gmtls.Config{
		GMSupport:          &gmtls.GMSupport{},
		RootCAs:            certPool,
		ClientAuth:         gmtls.NoClientCert,
		Certificates:       certificates,
		InsecureSkipVerify: security == nil || security.RootCerts == nil,
	}

	r := &RestyQueryService{
		host:       host,
		port:       port,
		secure:     true,
		gmSecure:   true,
		gmClient:   gmtls.NewCustomHTTPSClient(config),
		baseUrl:    baseUrl,
		gmSecurity: security,
	}

	return r
}

func (r RestyQueryService) query(url string) (data gjson.Result, err error) {
	var wrp gjson.Result
	if !r.gmSecure {
		resp, err := r.client.R().Get(r.baseUrl + url)
		if err != nil {
			return data, err
		}
		if !resp.IsSuccess() {
			return data, errors.New(resp.String())
		}
		wrp = gjson.ParseBytes(resp.Body())
	} else {
		response, err := r.gmClient.Get(r.baseUrl + url)
		if err != nil {
			return data, err
		}
		defer response.Body.Close()
		raw, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return data, err
		}
		wrp = gjson.ParseBytes(raw)
	}

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
	var wrp gjson.Result
	if !r.gmSecure {
		resp, err := r.client.R().SetQueryParams(params).Get(r.baseUrl + url)
		if err != nil {
			return data, err
		}
		if !resp.IsSuccess() {
			return data, errors.New(resp.String())
		}
		wrp = gjson.ParseBytes(resp.Body())
	} else {
		req, err := http.NewRequest("GET", r.baseUrl+url, nil)
		if err != nil {
			return data, err
		}
		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()

		response, err := r.gmClient.Do(req)
		if err != nil {
			return data, err
		}
		defer response.Body.Close()
		raw, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return data, err
		}
		wrp = gjson.ParseBytes(raw)
	}

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
	var wrp gjson.Result
	if !r.gmSecure {
		resp, err := r.client.R().SetQueryParamsFromValues(params).Get(r.baseUrl + url)
		if err != nil {
			return data, err
		}
		if !resp.IsSuccess() {
			return data, errors.New(resp.String())
		}
		wrp = gjson.ParseBytes(resp.Body())
	} else {
		req, err := http.NewRequest("GET", r.baseUrl+url, strings.NewReader(params.Encode()))
		if err != nil {
			return data, err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		response, err := r.gmClient.Do(req)
		if err != nil {
			return data, err
		}
		defer response.Body.Close()
		raw, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return data, err
		}
		wrp = gjson.ParseBytes(raw)
	}

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
	var wrp gjson.Result
	if !r.gmSecure {
		resp, err := r.client.R().SetBody(params).Post(r.baseUrl + url)
		if err != nil {
			return data, err
		}
		if !resp.IsSuccess() {
			return data, errors.New(resp.String())
		}
		wrp = gjson.ParseBytes(resp.Body())
	} else {
		body, err := json.Marshal(params)
		if err != nil {
			return data, err
		}
		response, err := r.gmClient.Post(r.baseUrl+url, "application/json", gbytes.NewBuffer(body))
		if err != nil {
			return data, err
		}
		defer response.Body.Close()
		raw, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return data, err
		}
		wrp = gjson.ParseBytes(raw)
	}

	// 数据不存在
	if wrp.Type == gjson.Null {
		return data, errors.New("empty data")
	}
	if !wrp.Get("success").Bool() {
		return data, errors.New(fmt.Sprintf("error code:%d msg:%s", wrp.Get("error.errorCode").Int(), wrp.Get("error.errorMessage").String()))
	}
	return wrp.Get("data"), nil
}

func (r RestyQueryService) GetLedgerHashs() ([]*framework.HashDigest, error) {
	wrp, err := r.query("/ledgers")
	if err != nil {
		return nil, err
	}
	ledgers := wrp.Array()
	hashs := make([]*framework.HashDigest, len(ledgers))
	for i, m := range ledgers {
		hash, err := framework.ParseHashDigest(base58.MustDecode(m.String()))
		if err != nil {
			return nil, err
		}
		hashs[i] = hash
	}
	return hashs, nil
}

func (r RestyQueryService) GetLedger(ledgerHash *framework.HashDigest) (info *ledger_model.LedgerInfo, err error) {
	wrp, err := r.query("/ledgers/" + ledgerHash.ToBase58())
	if err != nil {
		return info, err
	}
	hash, err := framework.ParseHashDigest(base58.MustDecode(wrp.Get("hash").String()))
	if err != nil {
		return nil, err
	}
	info = &ledger_model.LedgerInfo{}
	info.Hash = hash
	digest, err := framework.ParseHashDigest(base58.MustDecode(wrp.Get("latestBlockHash").String()))
	if err != nil {
		return nil, err
	}
	info.LatestBlockHash = digest
	info.LatestBlockHeight = wrp.Get("latestBlockHeight").Int()
	return info, nil
}

func (r RestyQueryService) GetLedgerAdminInfo(ledgerHash *framework.HashDigest) (info *ledger_model.LedgerAdminInfo, err error) {
	wrp, err := r.query("/ledgers/" + ledgerHash.ToBase58() + "/admininfo")
	if err != nil {
		return info, err
	}
	info = &ledger_model.LedgerAdminInfo{}
	info.ParticipantCount = wrp.Get("participantCount").Int()
	metadata := wrp.Get("metadata")
	info.Metadata = ledger_model.LedgerMetadata_V2{
		LedgerMetadata: ledger_model.LedgerMetadata{
			Seed:             []byte(metadata.Get("seed").String()),
			ParticipantsHash: base58.MustDecode(metadata.Get("participantsHash").String()),
			SettingsHash:     base58.MustDecode(metadata.Get("settingsHash").String()),
		},
		RolePrivilegesHash:    base58.MustDecode(metadata.Get("rolePrivilegesHash").String()),
		UserRolesHash:         base58.MustDecode(metadata.Get("userRolesHash").String()),
		GenesisUsers:          parseGenesisUsers(metadata.Get("genesisUsers").Array()),
		ContractRuntimeConfig: parseContractRuntimeConfig(metadata.Get("contractRuntimeConfig")),
	}
	participants := wrp.Get("participants").Array()
	pNodes := make([]ledger_model.ParticipantNode, len(participants))
	for i, node := range participants {
		pNodes[i] = ledger_model.ParticipantNode{
			Id:                   int32(node.Get("id").Int()),
			Name:                 node.Get("name").String(),
			Address:              base58.MustDecode(node.Get("address").String()),
			PubKey:               base58.MustDecode(node.Get("pubKey").String()),
			ParticipantNodeState: ledger_model.READY.GetValueByName(node.Get("participantNodeState").String()).(ledger_model.ParticipantNodeState),
		}
	}
	info.Participants = pNodes
	info.Settings = ledger_model.LedgerSettings{
		ConsensusProvider: wrp.Get("settings.consensusProvider").String(),
		ConsensusSetting:  base58.MustDecode(wrp.Get("settings.consensusSetting").String()),
		CryptoSetting:     parseCryptoSetting(wrp.Get("settings.cryptoSetting")),
	}
	return
}

func (r RestyQueryService) GetLedgerCryptoSetting(ledgerHash *framework.HashDigest) (info ledger_model.CryptoSetting, err error) {
	wrp, err := r.query("/ledgers/" + ledgerHash.ToBase58() + "/settings/crypto")
	if err != nil {
		return info, err
	}
	info = parseCryptoSetting(wrp)
	return
}

func parseCryptoSetting(info gjson.Result) ledger_model.CryptoSetting {
	autoVerifyHash := info.Get("autoVerifyHash").Bool()
	supportedProviders := info.Get("supportedProviders").Array()
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
	return ledger_model.CryptoSetting{
		SupportedProviders: providers,
		HashAlgorithm:      int16(info.Get("hashAlgorithm").Int()),
		AutoVerifyHash:     autoVerifyHash,
	}
}

func (r RestyQueryService) GetConsensusParticipants(ledgerHash *framework.HashDigest) (info []*ledger_model.ParticipantNode, err error) {
	wrp, err := r.query("/ledgers/" + ledgerHash.ToBase58() + "/participants")
	if err != nil {
		return info, err
	}
	info = parseConsensusParticipants(wrp.Array())
	return
}

func parseConsensusParticipants(participants []gjson.Result) []*ledger_model.ParticipantNode {
	info := make([]*ledger_model.ParticipantNode, len(participants))
	for i, node := range participants {
		info[i] = &ledger_model.ParticipantNode{
			Id:                   int32(node.Get("id").Int()),
			Name:                 node.Get("name").String(),
			Address:              base58.MustDecode(node.Get("address").String()),
			PubKey:               base58.MustDecode(node.Get("pubKey").String()),
			ParticipantNodeState: ledger_model.READY.GetValueByName(node.Get("participantNodeState").String()).(ledger_model.ParticipantNodeState),
		}
	}

	return info
}

func (r RestyQueryService) GetLedgerMetadata(ledgerHash *framework.HashDigest) (info *ledger_model.LedgerMetadata, err error) {
	metadata, err := r.query("/ledgers/" + ledgerHash.ToBase58() + "/metadata")
	if err != nil {
		return info, err
	}
	info = &ledger_model.LedgerMetadata{
		Seed:             []byte(metadata.Get("seed").String()),
		ParticipantsHash: base58.MustDecode(metadata.Get("participantsHash").String()),
		SettingsHash:     base58.MustDecode(metadata.Get("settingsHash").String()),
	}

	return
}

func (r RestyQueryService) GetBlockByHeight(ledgerHash *framework.HashDigest, height int64) (info *ledger_model.LedgerBlock, err error) {
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

	info = &ledger_model.LedgerBlock{
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

func (r RestyQueryService) GetBlockByHash(ledgerHash, blockHash *framework.HashDigest) (info *ledger_model.LedgerBlock, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return parseBlock(ledgerHash, wrp)
}

func parseBlock(ledgerHash *framework.HashDigest, block gjson.Result) (info *ledger_model.LedgerBlock, err error) {
	var PreviousHash []byte
	if block.Get("previousHash").Exists() {
		PreviousHash = base58.MustDecode(block.Get("previousHash").String())
	} else {
		PreviousHash = nil
	}

	info = &ledger_model.LedgerBlock{
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

func (r RestyQueryService) GetTransactionCountByHeight(ledgerHash *framework.HashDigest, height int64) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/txs/count", ledgerHash.ToBase58(), height))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetTransactionCountByHash(ledgerHash, blockHash *framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/txs/count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetTransactionTotalCount(ledgerHash *framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/txs/count", ledgerHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetDataAccountCountByHeight(ledgerHash *framework.HashDigest, height int64) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/accounts/count", ledgerHash.ToBase58(), height))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetDataAccountCountByHash(ledgerHash, blockHash *framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/accounts/count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetDataAccountTotalCount(ledgerHash *framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/accounts/count", ledgerHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetUserCountByHeight(ledgerHash *framework.HashDigest, height int64) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/users/count", ledgerHash.ToBase58(), height))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetUserCountByHash(ledgerHash, blockHash *framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/users/count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetUserTotalCount(ledgerHash *framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/users/count", ledgerHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetContractCountByHeight(ledgerHash *framework.HashDigest, height int64) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/contracts/count", ledgerHash.ToBase58(), height))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetContractCountByHash(ledgerHash, blockHash *framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/contracts/count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetContractTotalCount(ledgerHash *framework.HashDigest) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/contracts/count", ledgerHash.ToBase58()))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetUser(ledgerHash *framework.HashDigest, address string) (info *ledger_model.UserInfo, err error) {
	user, err := r.query(fmt.Sprintf("/ledgers/%s/users/address/%s", ledgerHash.ToBase58(), address))
	if err != nil {
		return info, err
	}
	if !user.Exists() {
		return info, errors.New("not exists")
	}
	info = &ledger_model.UserInfo{
		UserAccountHeader: ledger_model.UserAccountHeader{
			BlockchainIdentity: parseBlockchainIdentity(user),
			State:              ledger_model.NORMAL.GetValueByName(user.Get("state").String()).(ledger_model.AccountState),
			Certificate:        user.Get("certificate").String(),
		},
	}
	return
}

func (r RestyQueryService) GetDataAccount(ledgerHash *framework.HashDigest, address string) (info *ledger_model.DataAccountInfo, err error) {
	account, err := r.query(fmt.Sprintf("/ledgers/%s/accounts/address/%s", ledgerHash.ToBase58(), address))
	if err != nil {
		return info, err
	}
	if !account.Exists() {
		return info, errors.New("not exists")
	}
	info = &ledger_model.DataAccountInfo{}
	info.BlockchainIdentity = parseBlockchainIdentity(account.Get("iD"))
	info.DataCount = account.Get("dataset.dataCount").Int()
	info.Permission = parseDataPermission(account.Get("permission"))
	return
}

// KV value 解析
func resolveTypedKVValue(t ledger_model.DataType, v gjson.Result) (interface{}, error) {
	switch t {
	case ledger_model.INT64, ledger_model.TIMESTAMP:
		return v.Int(), nil
	case ledger_model.BYTES, ledger_model.IMG:
		return base64.Decode(v.String())
	default:
		return v.String(), nil
	}
}

func (r RestyQueryService) GetLatestDataEntries(ledgerHash *framework.HashDigest, address string, keys []string) (info []ledger_model.TypedKVEntry, err error) {
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
		t := ledger_model.NIL.GetValueByName(id.Get("type").String()).(ledger_model.DataType)
		v, err := resolveTypedKVValue(t, id.Get("value"))
		if err != nil {
			return nil, err
		}
		info[i] = ledger_model.TypedKVEntry{
			Key:     id.Get("key").String(),
			Value:   v,
			Version: id.Get("version").Int(),
			Type:    t,
		}
	}

	return
}

func (r RestyQueryService) GetDataEntries(ledgerHash *framework.HashDigest, address string, kvInfoVO ledger_model.KVInfoVO) (info []ledger_model.TypedKVEntry, err error) {
	wrp, err := r.queryWithBody(fmt.Sprintf("/ledgers/%s/accounts/%s/entries-version", ledgerHash.ToBase58(), address), kvInfoVO)
	if err != nil {
		return info, err
	}
	kvArray := wrp.Array()
	info = make([]ledger_model.TypedKVEntry, len(kvArray))
	for i, id := range kvArray {
		t := ledger_model.NIL.GetValueByName(id.Get("type").String()).(ledger_model.DataType)
		v, err := resolveTypedKVValue(t, id.Get("value"))
		if err != nil {
			return nil, err
		}
		info[i] = ledger_model.TypedKVEntry{
			Key:     id.Get("key").String(),
			Value:   v,
			Version: id.Get("version").Int(),
			Type:    t,
		}
	}

	return
}

func (r RestyQueryService) GetDataEntriesTotalCount(ledgerHash *framework.HashDigest, address string) (int64, error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/accounts/address/%s/entries/count", ledgerHash.ToBase58(), address))
	if err != nil {
		return 0, err
	}
	return wrp.Int(), nil
}

func (r RestyQueryService) GetLatestDataEntriesByRange(ledgerHash *framework.HashDigest, address string, fromIndex, count int64) (info []ledger_model.TypedKVEntry, err error) {
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
		t := ledger_model.NIL.GetValueByName(id.Get("type").String()).(ledger_model.DataType)
		v, err := resolveTypedKVValue(t, id.Get("value"))
		if err != nil {
			return nil, err
		}
		info[i] = ledger_model.TypedKVEntry{
			Key:     id.Get("key").String(),
			Value:   v,
			Version: id.Get("version").Int(),
			Type:    t,
		}
	}

	return
}

func (r RestyQueryService) GetContract(ledgerHash *framework.HashDigest, address string) (info *ledger_model.ContractInfo, err error) {
	contract, err := r.query(fmt.Sprintf("/ledgers/%s/contracts/address/%s", ledgerHash.ToBase58(), address))
	if err != nil {
		return
	}
	if !contract.Exists() {
		return info, errors.New("not exists")
	}
	info = &ledger_model.ContractInfo{
		BlockchainIdentity: parseBlockchainIdentity(contract),
		MerkleSnapshot: ledger_model.MerkleSnapshot{
			RootHash: base58.MustDecode(contract.Get("rootHash").String()),
		},
		ChainCodeVersion: contract.Get("chainCodeVersion").Int(),
		ChainCode:  bytes.StringToBytes(contract.Get("chainCode").String()),
		State:      ledger_model.NORMAL.GetValueByName(contract.Get("state").String()).(ledger_model.AccountState),
		Permission: parseDataPermission(contract.Get("permission")),
		Lang:       ledger_model.Java.GetValueByName(contract.Get("lang").String()).(ledger_model.ContractLang),
	}

	return
}

func (r RestyQueryService) GetUsers(ledgerHash *framework.HashDigest, fromIndex, count int64) (info []*ledger_model.BlockchainIdentity, err error) {
	params := map[string]string{
		"fromIndex": strconv.FormatInt(fromIndex, 10),
		"count":     strconv.FormatInt(count, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/users", ledgerHash.ToBase58()), params)
	if err != nil {
		return info, err
	}
	idArray := wrp.Array()
	info = make([]*ledger_model.BlockchainIdentity, len(idArray))
	for i, id := range idArray {
		info[i] = parseBlockchainIdentity(id)
	}

	return
}

func (r RestyQueryService) GetDataAccounts(ledgerHash *framework.HashDigest, fromIndex, count int64) (info []*ledger_model.BlockchainIdentity, err error) {
	params := map[string]string{
		"fromIndex": strconv.FormatInt(fromIndex, 10),
		"count":     strconv.FormatInt(count, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/accounts", ledgerHash.ToBase58()), params)
	if err != nil {
		return info, err
	}
	idArray := wrp.Array()
	info = make([]*ledger_model.BlockchainIdentity, len(idArray))
	for i, id := range idArray {
		info[i] = parseBlockchainIdentity(id)
	}

	return
}

func (r RestyQueryService) GetContractAccounts(ledgerHash *framework.HashDigest, fromIndex, count int64) (info []*ledger_model.BlockchainIdentity, err error) {
	params := map[string]string{
		"fromIndex": strconv.FormatInt(fromIndex, 10),
		"count":     strconv.FormatInt(count, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/contracts", ledgerHash.ToBase58()), params)
	if err != nil {
		return info, err
	}
	idArray := wrp.Array()
	info = make([]*ledger_model.BlockchainIdentity, len(idArray))
	for i, id := range idArray {
		info[i] = parseBlockchainIdentity(id)
	}

	return
}

func (r RestyQueryService) GetUserRoles(ledgerHash *framework.HashDigest, userAddress string) (info *ledger_model.RoleSet, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/user-role/%s", ledgerHash.ToBase58(), userAddress))
	if err != nil {
		return info, err
	}

	rolesArray := wrp.Get("roleSet").Array()
	roles := make([]string, len(rolesArray))
	for i, role := range rolesArray {
		roles[i] = role.String()
	}
	info = &ledger_model.RoleSet{
		Policy: ledger_model.UNION.GetValueByName(wrp.Get("policy").String()).(ledger_model.RolesPolicy),
		Roles:  roles,
	}

	return
}

func (r RestyQueryService) GetSystemEvents(ledgerHash *framework.HashDigest, eventName string, fromSequence int64, maxCount int64) (info []*ledger_model.Event, err error) {
	params := map[string]string{
		"fromSequence": strconv.FormatInt(fromSequence, 10),
		"count":        strconv.FormatInt(maxCount, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/events/system/names/%s", ledgerHash.ToBase58(), eventName), params)
	if err != nil {
		return info, err
	}
	idArray := wrp.Array()
	info = make([]*ledger_model.Event, len(idArray))
	for i, item := range idArray {
		info[i] = parseEvent(item)
	}

	return
}

func parseBytesValue(info gjson.Result) ledger_model.BytesValue {
	return ledger_model.BytesValue{
		Type:  ledger_model.NIL.GetValueByName(info.Get("type").String()).(ledger_model.DataType),
		Bytes: base64.MustDecode(info.Get("bytes").String()),
	}
}

func (r RestyQueryService) GetUserEventAccounts(ledgerHash *framework.HashDigest, fromIndex int64, maxCount int64) (info []*ledger_model.BlockchainIdentity, err error) {
	params := map[string]string{
		"fromIndex": strconv.FormatInt(fromIndex, 10),
		"count":     strconv.FormatInt(maxCount, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/events/user/accounts", ledgerHash.ToBase58()), params)
	if err != nil {
		return info, err
	}
	idArray := wrp.Array()
	info = make([]*ledger_model.BlockchainIdentity, len(idArray))
	for i, id := range idArray {
		info[i] = parseBlockchainIdentity(id)
	}

	return
}

func (r RestyQueryService) GetUserEvents(ledgerHash *framework.HashDigest, address string, eventName string, fromSequence int64, maxCount int64) (info []*ledger_model.Event, err error) {
	params := map[string]string{
		"fromSequence": strconv.FormatInt(fromSequence, 10),
		"count":        strconv.FormatInt(maxCount, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/events/user/accounts/%s/names/%s", ledgerHash.ToBase58(), address, eventName), params)
	if err != nil {
		return info, err
	}
	idArray := wrp.Array()
	info = make([]*ledger_model.Event, len(idArray))
	for i, item := range idArray {
		info[i] = parseEvent(item)
	}

	return
}

func parseEvent(event gjson.Result) *ledger_model.Event {
	info := &ledger_model.Event{
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
		info.EventAccount = base58.MustDecode(event.Get("eventAccount").String())
	}
	if !event.Get("content.nil").Bool() {
		info.Content = parseBytesValue(event.Get("content"))
	}

	return info
}

func (r RestyQueryService) GetTransactionsByHeight(ledgerHash *framework.HashDigest, height int64, fromIndex, count int64) (info []*ledger_model.LedgerTransaction, err error) {
	params := map[string]string{
		"fromIndex": strconv.FormatInt(fromIndex, 10),
		"count":     strconv.FormatInt(count, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/blocks/height/%d/txs", ledgerHash.ToBase58(), height), params)
	if err != nil {
		return info, err
	}
	txArray := wrp.Array()
	info = make([]*ledger_model.LedgerTransaction, len(txArray))
	for i, tx := range txArray {
		info[i] = &ledger_model.LedgerTransaction{
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
		BlockHeight:       result.Get("blockHeight").Int(),
		ExecutionState:    ledger_model.SUCCESS.GetValueByName(result.Get("executionState").String()).(ledger_model.TransactionState),
		TransactionHash:   base58.MustDecode(result.Get("transactionHash").String()),
		DataSnapshot:      parseLedgerDataSnapshot(result.Get("dataSnapshot")),
		DerivedOperations: parseOperations(result.Get("derivedOperations").Array()),
	}
}

func (r RestyQueryService) GetTransactionsByHash(ledgerHash, blockHash *framework.HashDigest, fromIndex, count int64) (info []*ledger_model.LedgerTransaction, err error) {
	params := map[string]string{
		"fromIndex": strconv.FormatInt(fromIndex, 10),
		"count":     strconv.FormatInt(count, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/txs", ledgerHash.ToBase58(), blockHash.ToBase58()), params)
	if err != nil {
		return info, err
	}
	txArray := wrp.Array()
	info = make([]*ledger_model.LedgerTransaction, len(txArray))
	for i, tx := range txArray {
		info[i] = &ledger_model.LedgerTransaction{
			Request: parseTxRequest(tx.Get("request")),
			Result:  parseTxResult(tx.Get("result")),
		}
	}

	return
}

func (r RestyQueryService) GetTransactionByContentHash(ledgerHash, contentHash *framework.HashDigest) (info *ledger_model.LedgerTransaction, err error) {
	tx, err := r.query(fmt.Sprintf("/ledgers/%s/txs/hash/%s", ledgerHash.ToBase58(), contentHash.ToBase58()))
	if err != nil {
		return info, err
	}
	return &ledger_model.LedgerTransaction{
		Request: parseTxRequest(tx.Get("request")),
		Result:  parseTxResult(tx.Get("result")),
	}, nil
}

func (r RestyQueryService) GetTransactionStateByContentHash(ledgerHash, contentHash *framework.HashDigest) (info ledger_model.TransactionState, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/txs/state/%s", ledgerHash.ToBase58(), contentHash.ToBase58()))
	if err != nil {
		return info, err
	}
	if !wrp.Exists() {
		return info, errors.New("not exists")
	}
	info, ok := ledger_model.SUCCESS.GetValueByName(wrp.String()).(ledger_model.TransactionState)
	if !ok {
		return ledger_model.SYSTEM_ERROR, errors.New("not TransactionState")
	}
	return info, nil
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

func parseTransactionContent(info gjson.Result) *ledger_model.TransactionContent {
	return &ledger_model.TransactionContent{
		LedgerHash: base58.MustDecode(info.Get("ledgerHash").String()),
		Operations: parseOperations(info.Get("operations").Array()),
		Timestamp:  info.Get("timestamp").Int(),
	}
}

func parseSignatures(info []gjson.Result) []*ledger_model.DigitalSignature {
	signatures := make([]*ledger_model.DigitalSignature, len(info))
	for i, sign := range info {
		signatures[i] = &ledger_model.DigitalSignature{
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
		switch operation.Get("@type").String() {
		case "com.jd.blockchain.ledger.UserRegisterOperation":
			// 注册用户
			dc = parseUserRegisterOperation(operation)
		case "com.jd.blockchain.ledger.DataAccountRegisterOperation":
			// 注册数据账户
			dc = parseDataAccountRegisterOperation(operation.Get("accountID"))
		case "com.jd.blockchain.ledger.DataAccountKVSetOperation":
			// KV写入
			dc = parseDataAccountKVSetOperation(operation)
		case "com.jd.blockchain.ledger.EventAccountRegisterOperation":
			// 事件账户注册
			dc = parseEventAccountRegisterOperation(operation.Get("eventAccountID"))
		case "com.jd.blockchain.ledger.EventPublishOperation":
			// 发布事件
			dc = parseEventPublishOperation(operation)
		case "com.jd.blockchain.ledger.ParticipantRegisterOperation":
			// 注册参与方
			dc = parseParticipantRegisterOperation(operation)
		case "com.jd.blockchain.ledger.ParticipantStateUpdateOperation":
			// 参与方状态变更
			dc = parseParticipantStateUpdateOperation(operation)
		case "com.jd.blockchain.ledger.ContractCodeDeployOperation":
			// 合约部署
			dc = parseContractCodeDeployOperation(operation.Get("chainCode"))
		case "com.jd.blockchain.ledger.ContractEventSendOperation":
			// 合约调用
			dc = parseContractEventSendOperation(operation)
		case "com.jd.blockchain.ledger.RolesConfigureOperation":
			// 角色配置
			dc = parseRolesConfigureOperation(operation)
		case "com.jd.blockchain.ledger.UserAuthorizeOperation":
			// 用户权限配置
			dc = parseUserAuthorizeOperation(operation.Get("userRolesAuthorizations").Array())
		case "com.jd.blockchain.ledger.ConsensusSettingsUpdateOperation":
			// 共识信息变更
			dc = parseConsensusSettingsUpdateOperation(operation)
		case "com.jd.blockchain.ledger.LedgerInitOperation":
			// 账本初始化
			dc = parseLedgerInitOperation(operation.Get("initSetting"))
		case "com.jd.blockchain.ledger.UserCAUpdateOperation":
			// 用户证书更新
			dc = parseUserCAUpdateOperation(operation)
		case "com.jd.blockchain.ledger.UserStateUpdateOperation":
			// 用户状态变更
			dc = parseUserStateUpdateOperation(operation)
		case "com.jd.blockchain.ledger.RootCAUpdateOperation":
			// 根证书更新
			dc = parseRootCAUpdateOperation(operation)
		case "com.jd.blockchain.ledger.AccountPermissionSetOperation":
			// 账户权限配置
			dc = parseAccountPermissionSetOperation(operation)
		case "com.jd.blockchain.ledger.ContractStateUpdateOperation":
			// 合约状态变更
			dc = parseContractStateUpdateOperation(operation)
		case "com.jd.blockchain.ledger.HashAlgorithmUpdateOperation":
			// 哈希算法变更
			dc = parseHashAlgorithmUpdateOperation(operation)
		}

		operations[i] = dc
	}

	return operations
}

func parseUserRegisterOperation(info gjson.Result) binary_proto.DataContract {
	return &ledger_model.UserRegisterOperation{
		UserID:      parseBlockchainIdentity(info.Get("userID")),
		Certificate: info.Get("certificate").String(),
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
		ParticipantName: info.Get("participantName").String(),
		ParticipantID:   parseBlockchainIdentity(info.Get("participantID")),
		Certificate:     info.Get("certificate").String(),
	}
}

func parseParticipantStateUpdateOperation(info gjson.Result) binary_proto.DataContract {
	return &ledger_model.ParticipantStateUpdateOperation{
		State:         ledger_model.READY.GetValueByName(info.Get("state").String()).(ledger_model.ParticipantNodeState),
		ParticipantID: parseBlockchainIdentity(info.Get("participantID")),
	}
}

func parseContractCodeDeployOperation(info gjson.Result) binary_proto.DataContract {
	return &ledger_model.ContractCodeDeployOperation{
		ContractID: parseBlockchainIdentity(info.Get("contractID")),
		ChainCode:  bytes.StringToBytes(info.Get("chainCode").String()),
		Lang:       ledger_model.Java.GetValueByName(info.Get("lang").String()).(ledger_model.ContractLang),
	}
}

func parseContractEventSendOperation(info gjson.Result) binary_proto.DataContract {
	operation := &ledger_model.ContractEventSendOperation{
		ContractAddress: base58.MustDecode(info.Get("contractAddress").String()),
		Event:           info.Get("event").String(),
		Version:         info.Get("version").Int(),
	}
	argsArray := info.Get("args.values").Array()
	args := make([]ledger_model.BytesValue, len(argsArray))
	for i, arg := range argsArray {
		args[i] = ledger_model.BytesValue{
			Bytes: base64.MustDecode(arg.Get("bytes").String()),
			Type:  ledger_model.NIL.GetValueByName(arg.Get("type").String()).(ledger_model.DataType),
		}
	}
	operation.Args = ledger_model.BytesValueList{
		Values: args,
	}

	return operation
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

func parseConsensusSettingsUpdateOperation(info gjson.Result) binary_proto.DataContract {
	properties := ledger_model.Properties{}
	json.Unmarshal([]byte(info.Get("properties").Raw), &properties)
	len := len(properties)
	pss := make([][]byte, len)
	for i := 0; i < len; i++ {
		pss[i] = properties[i].ToBytes()
	}
	return &ledger_model.ConsensusSettingsUpdateOperation{
		Provider:   info.Get("provider").String(),
		Properties: pss,
	}
}

func parseLedgerInitOperation(info gjson.Result) binary_proto.DataContract {
	return &ledger_model.LedgerInitOperation{
		InitSetting: ledger_model.LedgerInitSetting{
			LedgerSeed:             []byte(info.Get("ledgerSeed").String()),
			ConsensusParticipants:  parseConsensusParticipants(info.Get("consensusParticipants").Array()),
			CryptoSetting:          parseCryptoSetting(info.Get("cryptoSetting")),
			ConsensusProvider:      info.Get("consensusProvider").String(),
			ConsensusSettings:      []byte(info.Get("consensusSettings").String()),
			CreatedTime:            info.Get("createdTime").Int(),
			LedgerStructureVersion: info.Get("ledgerStructureVersion").Int(),
			IdentityMode:           ledger_model.KEYPAIR.GetValueByName(info.Get("identityMode").String()).(ledger_model.IdentityMode),
			LedgerCertificates:     parseStringArray(info.Get("ledgerCertificates").Array()),
			GenesisUsers:           parseGenesisUsers(info.Get("genesisUsers").Array()),
			LedgerDataStructure:    ledger_model.MERKLE_TREE.GetValueByName(info.Get("ledgerDataStructure").String()).(ledger_model.LedgerDataStructure),
			ContractRuntimeConfig:  parseContractRuntimeConfig(info.Get("contractRuntimeConfig")),
		},
	}
}

func parseContractRuntimeConfig(info gjson.Result) ledger_model.ContractRuntimeConfig {
	return ledger_model.ContractRuntimeConfig{
		Timeout: info.Get("timeout").Int(),
	}
}

func parseGenesisUsers(info []gjson.Result) []ledger_model.GenesisUser {
	array := make([]ledger_model.GenesisUser, len(info))
	for i, item := range info {
		array[i] = parseGenesisUser(item)
	}

	return array
}

func parseGenesisUser(info gjson.Result) ledger_model.GenesisUser {
	return ledger_model.GenesisUser{
		PubKey:      base58.MustDecode(info.Get("pubKey").String()),
		Certificate: info.Get("certificate").String(),
		Roles:       parseStringArray(info.Get("roles").Array()),
		RolesPolicy: ledger_model.UNION.GetValueByName(info.Get("rolesPolicy").String()).(ledger_model.RolesPolicy),
	}
}

func parseUserCAUpdateOperation(info gjson.Result) binary_proto.DataContract {
	return &ledger_model.UserCAUpdateOperation{
		UserAddress: base58.MustDecode(info.Get("userAddress").String()),
		Certificate: info.Get("certificate").String(),
	}
}

func parseUserStateUpdateOperation(info gjson.Result) binary_proto.DataContract {
	return &ledger_model.UserStateUpdateOperation{
		UserAddress: base58.MustDecode(info.Get("userAddress").String()),
		State:       ledger_model.NORMAL.GetValueByName(info.Get("state").String()).(ledger_model.AccountState),
	}
}

func parseRootCAUpdateOperation(info gjson.Result) binary_proto.DataContract {
	var op ledger_model.RootCAUpdateOperation
	json.Unmarshal([]byte(info.Raw), &op)
	return op
}

func parseAccountPermissionSetOperation(info gjson.Result) binary_proto.DataContract {
	return &ledger_model.AccountPermissionSetOperation{
		Address:     base58.MustDecode(info.Get("address").String()),
		AccountType: ledger_model.DATA.GetValueByName(info.Get("accountType").String()).(ledger_model.AccountType),
		Mode:        int32(info.Get("mode").Int()),
		Role:        info.Get("role").String(),
	}
}

func parseHashAlgorithmUpdateOperation(info gjson.Result) binary_proto.DataContract {
	return &ledger_model.HashAlgorithmUpdateOperation{
		Algorithm: info.Get("algorithm").String(),
	}
}

func parseContractStateUpdateOperation(info gjson.Result) binary_proto.DataContract {
	return &ledger_model.ContractStateUpdateOperation{
		ContractAddress: base58.MustDecode(info.Get("contractAddress").String()),
		State:           ledger_model.NORMAL.GetValueByName(info.Get("state").String()).(ledger_model.AccountState),
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
		AccountAddress: base58.MustDecode(info.Get("accountAddress").String()),
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
		EventAddress: base58.MustDecode(info.Get("eventAddress").String()),
		Events:       writeSet,
	}
}

func parseBlockchainIdentity(id gjson.Result) *ledger_model.BlockchainIdentity {
	return &ledger_model.BlockchainIdentity{
		Address: base58.MustDecode(id.Get("address").String()),
		PubKey:  base58.MustDecode(id.Get("pubKey").String()),
	}
}

func parseDataPermission(permission gjson.Result) ledger_model.DataPermission {
	var p ledger_model.DataPermission
	json.Unmarshal([]byte(permission.Raw), &p)
	return p
}

func (r RestyQueryService) GetSystemEventNameTotalCount(ledgerHash *framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/system/names/count", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetSystemEventNames(ledgerHash *framework.HashDigest, fromIndex, count int64) (info []string, err error) {
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

func (r RestyQueryService) GetSystemEventsTotalCount(ledgerHash *framework.HashDigest, eventName string) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/system/names/%s/count", ledgerHash.ToBase58(), eventName))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetUserEventAccount(ledgerHash *framework.HashDigest, address string) (info *ledger_model.EventAccountInfo, err error) {
	account, err := r.query(fmt.Sprintf("/ledgers/%s/events/user/accounts/%s", ledgerHash.ToBase58(), address))
	if err != nil {
		return info, err
	}
	if !account.Exists() {
		return info, errors.New("not exists")
	}
	info = &ledger_model.EventAccountInfo{}
	info.BlockchainIdentity = parseBlockchainIdentity(account.Get("iD"))
	info.DataCount = account.Get("dataset.dataCount").Int()
	info.Permission = parseDataPermission(account.Get("permission"))
	return info, err
}

func (r RestyQueryService) GetUserEventAccountTotalCount(ledgerHash *framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/user/accounts/count", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetUserEventNames(ledgerHash *framework.HashDigest, address string, fromIndex, count int64) (info []string, err error) {
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

func (r RestyQueryService) GetUserEventNameTotalCount(ledgerHash *framework.HashDigest, address string) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/user/accounts/%s/names/count", ledgerHash.ToBase58(), address))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetUserEventsTotalCount(ledgerHash *framework.HashDigest, address, eventName string) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/user/accounts/%s/names/%s/count", ledgerHash.ToBase58(), address, eventName))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetLatestSystemEvent(ledgerHash *framework.HashDigest, eventName string) (info *ledger_model.Event, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/system/names/%s/latest", ledgerHash.ToBase58(), eventName))
	if err != nil {
		return info, err
	}

	return parseEvent(wrp), nil
}

func (r RestyQueryService) GetLatestUserEvent(ledgerHash *framework.HashDigest, address string, eventName string) (info *ledger_model.Event, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/events/user/accounts/%s/names/%s/latest", ledgerHash.ToBase58(), address, eventName))
	if err != nil {
		return info, err
	}

	return parseEvent(wrp), nil
}

func (r RestyQueryService) GetLatestBlock(ledgerHash *framework.HashDigest) (info *ledger_model.LedgerBlock, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/latest", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return parseBlock(ledgerHash, wrp)
}

func (r RestyQueryService) GetAdditionalTransactionCountByHeight(ledgerHash *framework.HashDigest, blockHeight int64) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/txs/additional-count", ledgerHash.ToBase58(), blockHeight))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalTransactionCountByHash(ledgerHash, blockHash *framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/txs/additional-count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalTransactionCount(ledgerHash *framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/txs/additional-count", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalDataAccountCountByHeight(ledgerHash *framework.HashDigest, blockHeight int64) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/accounts/additional-count", ledgerHash.ToBase58(), blockHeight))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalDataAccountCountByHash(ledgerHash, blockHash *framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/accounts/additional-count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalDataAccountCount(ledgerHash *framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/accounts/additional-count", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalUserCountByHeight(ledgerHash *framework.HashDigest, blockHeight int64) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/users/additional-count", ledgerHash.ToBase58(), blockHeight))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalUserCountByHash(ledgerHash, blockHash *framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/users/additional-count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalUserCount(ledgerHash *framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/users/additional-count", ledgerHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalContractCountByHeight(ledgerHash *framework.HashDigest, blockHeight int64) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/height/%d/contracts/additional-count", ledgerHash.ToBase58(), blockHeight))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalContractCountByHash(ledgerHash, blockHash *framework.HashDigest) (info int64, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/contracts/additional-count", ledgerHash.ToBase58(), blockHash.ToBase58()))
	if err != nil {
		return info, err
	}

	return wrp.Int(), nil
}

func (r RestyQueryService) GetAdditionalContractCount(ledgerHash *framework.HashDigest) (info int64, err error) {
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

func (r RestyQueryService) GetRolePrivileges(ledgerHash *framework.HashDigest, roleName string) (info *ledger_model.RolePrivileges, err error) {
	wrp, err := r.query(fmt.Sprintf("/ledgers/%s/authorization/role/%s", ledgerHash.ToBase58(), roleName))
	if err != nil {
		return info, err
	}

	return parsePrivilegeSet(wrp)
}

func parsePrivilegeSet(rolePrivileges gjson.Result) (info *ledger_model.RolePrivileges, err error) {
	info = &ledger_model.RolePrivileges{
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

func (r RestyQueryService) GetUserPrivileges(ledgerHash *framework.HashDigest, userAddress string) (info *ledger_model.UserRolesPrivileges, err error) {
	userPrivilegeMap, err := r.query(fmt.Sprintf("/ledgers/%s/authorization/user/%s", ledgerHash.ToBase58(), userAddress))
	if err != nil {
		return info, err
	}

	userRoles := parseStringArray(userPrivilegeMap.Get("userRole").Array())
	info = &ledger_model.UserRolesPrivileges{
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

func (r RestyQueryService) GetAdditionalTransactionsByHeight(ledgerHash *framework.HashDigest, height int64, fromIndex, count int64) (info []*ledger_model.LedgerTransaction, err error) {
	params := map[string]string{
		"fromIndex": strconv.FormatInt(fromIndex, 10),
		"count":     strconv.FormatInt(count, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/blocks/height/%d/txs/additional-txs", ledgerHash.ToBase58(), height), params)
	if err != nil {
		return info, err
	}
	txArray := wrp.Array()
	info = make([]*ledger_model.LedgerTransaction, len(txArray))
	for i, tx := range txArray {
		info[i] = &ledger_model.LedgerTransaction{
			Request: parseTxRequest(tx.Get("request")),
			Result:  parseTxResult(tx.Get("result")),
		}
	}

	return
}

func (r RestyQueryService) GetAdditionalTransactionsByHash(ledgerHash, blockHash *framework.HashDigest, fromIndex, count int64) (info []*ledger_model.LedgerTransaction, err error) {
	params := map[string]string{
		"fromIndex": strconv.FormatInt(fromIndex, 10),
		"count":     strconv.FormatInt(count, 10),
	}
	wrp, err := r.queryWithParams(fmt.Sprintf("/ledgers/%s/blocks/hash/%s/txs/additional-txs", ledgerHash.ToBase58(), blockHash.ToBase58()), params)
	if err != nil {
		return info, err
	}
	txArray := wrp.Array()
	info = make([]*ledger_model.LedgerTransaction, len(txArray))
	for i, tx := range txArray {
		info[i] = &ledger_model.LedgerTransaction{
			Request: parseTxRequest(tx.Get("request")),
			Result:  parseTxResult(tx.Get("result")),
		}
	}

	return
}
