package sdk

import (
	"errors"
	"fmt"
	"framework-go/ledger_model"
	"framework-go/utils/base58"
	"github.com/go-resty/resty/v2"
	"strconv"
)

/*
 * Author: imuge
 * Date: 2020/6/29 下午2:56
 */

var _ ledger_model.ActiveParticipantService = (*RestyConsensusService)(nil)

type RestyConsensusService struct {
	host    string
	port    int
	secure  bool
	client  *resty.Client
	baseUrl string
}

func NewRestyConsensusService(host string, port int, secure bool) *RestyConsensusService {
	var baseUrl string
	if secure {
		baseUrl = fmt.Sprintf("https://%s:%d", host, port)
	} else {
		baseUrl = fmt.Sprintf("http://%s:%d", host, port)
	}
	return &RestyConsensusService{
		host:    host,
		port:    port,
		secure:  secure,
		client:  resty.New(),
		baseUrl: baseUrl,
	}
}

func (r RestyConsensusService) ActivateParticipant(ledgerHash, host string, port int) (info ledger_model.TransactionResponse, err error) {
	url := "/management/delegate/activeparticipant"
	params := map[string]string{
		"ledgerHash": ledgerHash,
		"consensusHost": host,
		"consensusPort": strconv.Itoa(port),
	}
	resp, err := r.client.R().SetFormData(params).SetResult(ActivateParticipantResponse{}).Post(r.baseUrl + url)
	if err != nil {
		return
	}
	fmt.Println(fmt.Sprintf("%s \n %v \n", url, resp))
	if !resp.IsSuccess() {
		return info, errors.New(resp.String())
	}
	wrp, ok := resp.Result().(*ActivateParticipantResponse)
	if !ok {
		return info, errors.New("unparseable response")
	}

	data := wrp.Data.(map[string]interface{})
	success := wrp.Success && data["success"].(bool) && data["executionState"].(string) == "SUCCESS"
	if !success {
		return info, errors.New(data["executionState"].(string))
	}

	blockHash := data["blockHash"].(map[string]interface{})["value"].(string)
	blockHeight := int64(data["blockHeight"].(float64))
	contentHash := data["contentHash"].(map[string]interface{})["value"].(string)

	return ledger_model.TransactionResponse{
		Success:        true,
		BlockHeight:    blockHeight,
		ExecutionState: ledger_model.SUCCESS,
		BlockHash:      base58.MustDecode(blockHash),
		ContentHash:    base58.MustDecode(contentHash),
	}, nil
}

type ActivateParticipantResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
