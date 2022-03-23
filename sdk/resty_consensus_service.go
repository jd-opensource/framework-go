package sdk

import (
	"crypto/tls"
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

var _ ledger_model.ParticipantService = (*RestyConsensusService)(nil)

type RestyConsensusService struct {
	host     string
	port     int
	secure   bool
	baseUrl  string
	security *SSLSecurity
}

func NewRestyConsensusService(host string, port int) *RestyConsensusService {
	baseUrl := fmt.Sprintf("http://%s:%d", host, port)
	return &RestyConsensusService{
		host:    host,
		port:    port,
		secure:  false,
		baseUrl: baseUrl,
	}
}

func NewSecureRestyConsensusService(host string, port int, security *SSLSecurity) *RestyConsensusService {
	baseUrl := fmt.Sprintf("https://%s:%d", host, port)
	return &RestyConsensusService{
		host:     host,
		port:     port,
		secure:   true,
		baseUrl:  baseUrl,
		security: security,
	}
}

func (r RestyConsensusService) ActivateParticipant(params ledger_model.ActivateParticipantParams) (info bool, err error) {
	url := "/management/delegate/activeparticipant"
	args := map[string]string{
		"ledgerHash":         params.LedgerHash,
		"consensusHost":      params.ConsensusHost,
		"consensusPort":      strconv.Itoa(params.ConsensusPort),
		"consensusStorage":   params.ConsensusStorage,
		"consensusSecure":    strconv.FormatBool(params.ConsensusSecure),
		"remoteManageHost":   params.RemoteManageHost,
		"remoteManagePort":   strconv.Itoa(params.RemoteManagePort),
		"remoteManageSecure": strconv.FormatBool(params.RemoteManageSecure),
		"shutdown":           strconv.FormatBool(params.Shutdown),
	}
	client := resty.New()
	if r.secure {
		if r.security != nil {
			client.SetTLSClientConfig(&tls.Config{
				RootCAs:      r.security.RootCerts,
				Certificates: r.security.ClientCerts,
			})
		} else {
			client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		}
	}
	resp, err := client.R().SetFormData(args).SetResult(ActivateParticipantResponse{}).Post(r.baseUrl + url)
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

func (r RestyConsensusService) InactivateParticipant(params ledger_model.InactivateParticipantParams) (info bool, err error) {
	url := "/management/delegate/deactiveparticipant"
	args := map[string]string{
		"ledgerHash":         params.LedgerHash,
		"participantAddress": params.ParticipantAddress,
	}
	client := resty.New()
	if r.secure {
		if r.security != nil {
			client.SetTLSClientConfig(&tls.Config{
				RootCAs:      r.security.RootCerts,
				Certificates: r.security.ClientCerts,
			})
		} else {
			client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		}
	}
	resp, err := client.R().SetFormData(args).SetResult(InactivateParticipantResponse{}).Post(r.baseUrl + url)
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
