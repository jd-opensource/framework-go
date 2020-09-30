package sdk

import (
	"errors"
	"fmt"
	"github.com/blockchain-jd-com/framework-go/ledger_model"
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

func (r RestyConsensusService) ActivateParticipant(ledgerHash, host string, port int, remoteManageHost string, remoteManagePort int) (info bool, err error) {
	url := "/management/delegate/activeparticipant"
	params := map[string]string{
		"ledgerHash":       ledgerHash,
		"consensusHost":    host,
		"consensusPort":    strconv.Itoa(port),
		"remoteManageHost": remoteManageHost,
		"remoteManagePort": strconv.Itoa(remoteManagePort),
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
	} else {
		return wrp.Success, nil
	}
}

func (r RestyConsensusService) InactivateParticipant(ledgerHash, participantAddress, remoteManageHost string, remoteManagePort int) (info bool, err error) {
	url := "/management/delegate/deactiveparticipant"
	params := map[string]string{
		"ledgerHash":         ledgerHash,
		"participantAddress": participantAddress,
		"remoteManageHost":   remoteManageHost,
		"remoteManagePort":   strconv.Itoa(remoteManagePort),
	}
	resp, err := r.client.R().SetFormData(params).SetResult(InactivateParticipantResponse{}).Post(r.baseUrl + url)
	if err != nil {
		return
	}
	fmt.Println(fmt.Sprintf("%s \n %v \n", url, resp))
	if !resp.IsSuccess() {
		return info, errors.New(resp.String())
	}
	wrp, ok := resp.Result().(*InactivateParticipantResponse)
	if !ok {
		return info, errors.New("unparseable response")
	} else {
		return wrp.Success, nil
	}
}

type ActivateParticipantResponse struct {
	Success bool `json:"success"`
}

type InactivateParticipantResponse struct {
	Success bool `json:"success"`
}
