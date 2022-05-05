package sdk

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/blockchain-jd-com/framework-go/gmsm/gmtls"
	gmx509 "github.com/blockchain-jd-com/framework-go/gmsm/x509"
	"github.com/blockchain-jd-com/framework-go/ledger_model"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/url"
	"strconv"
)

/*
 * Author: imuge
 * Date: 2020/6/29 下午2:56
 */

var _ ledger_model.ParticipantService = (*RestyConsensusService)(nil)

type RestyConsensusService struct {
	host       string
	port       int
	secure     bool
	baseUrl    string
	security   *SSLSecurity
	gmSecure   bool
	gmSecurity *GMSSLSecurity
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

func NewGMSecureConsensusService(host string, port int, security *GMSSLSecurity) *RestyConsensusService {
	baseUrl := fmt.Sprintf("https://%s:%d", host, port)
	return &RestyConsensusService{
		host:       host,
		port:       port,
		secure:     true,
		gmSecure:   true,
		baseUrl:    baseUrl,
		gmSecurity: security,
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

	if !r.gmSecure {
		return r.active(url, args)
	} else {
		return r.gmActive(url, args)
	}
}

func (r RestyConsensusService) active(url string, args map[string]string) (info bool, err error) {
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

func (r RestyConsensusService) gmActive(activeUrl string, args map[string]string) (info bool, err error) {
	var certPool *gmx509.CertPool
	if r.gmSecurity != nil {
		certPool = r.gmSecurity.RootCerts
	}
	config := &gmtls.Config{
		GMSupport:          &gmtls.GMSupport{},
		RootCAs:            certPool,
		ClientAuth:         gmtls.NoClientCert,
		InsecureSkipVerify: r.security == nil || r.security.RootCerts == nil,
	}

	gmClient := gmtls.NewCustomHTTPSClient(config)
	queryForm := url.Values{}
	for k, v := range args {
		queryForm.Add(k, v)
	}
	result, err := gmClient.PostForm(r.baseUrl+activeUrl, queryForm)
	defer result.Body.Close()
	if err != nil {
		return
	}
	raw, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return
	}

	tresp := gjson.ParseBytes(raw)
	if !tresp.Get("success").Bool() {
		return info, errors.New("active error")
	} else {
		return true, nil
	}
}

func (r RestyConsensusService) InactivateParticipant(params ledger_model.InactivateParticipantParams) (info bool, err error) {
	url := "/management/delegate/deactiveparticipant"
	args := map[string]string{
		"ledgerHash":         params.LedgerHash,
		"participantAddress": params.ParticipantAddress,
	}
	if !r.gmSecure {
		return r.inactive(url, args)
	} else {
		return r.gmInactive(url, args)
	}
}

func (r RestyConsensusService) inactive(url string, args map[string]string) (info bool, err error) {
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

func (r RestyConsensusService) gmInactive(inactiveUrl string, args map[string]string) (info bool, err error) {
	var certPool *gmx509.CertPool
	if r.gmSecurity != nil {
		certPool = r.gmSecurity.RootCerts
	}
	config := &gmtls.Config{
		GMSupport:          &gmtls.GMSupport{},
		RootCAs:            certPool,
		ClientAuth:         gmtls.NoClientCert,
		InsecureSkipVerify: r.security == nil || r.security.RootCerts == nil,
	}

	gmClient := gmtls.NewCustomHTTPSClient(config)
	queryForm := url.Values{}
	for k, v := range args {
		queryForm.Add(k, v)
	}
	result, err := gmClient.PostForm(r.baseUrl+inactiveUrl, queryForm)
	defer result.Body.Close()
	if err != nil {
		return
	}
	raw, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return
	}

	tresp := gjson.ParseBytes(raw)
	if !tresp.Get("success").Bool() {
		return info, errors.New("inactive error")
	} else {
		return true, nil
	}
}

type ActivateParticipantResponse struct {
	Success bool `json:"success"`
}

type InactivateParticipantResponse struct {
	Success bool `json:"success"`
}
