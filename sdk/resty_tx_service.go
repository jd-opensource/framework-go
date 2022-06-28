package sdk

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"

	gbytes "bytes"

	binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"
	"github.com/blockchain-jd-com/framework-go/gmsm/gmtls"
	gmx509 "github.com/blockchain-jd-com/framework-go/gmsm/x509"
	"github.com/blockchain-jd-com/framework-go/ledger_model"
	"github.com/go-resty/resty/v2"
)

/*
 * Author: imuge
 * Date: 2020/5/29 下午1:55
 */

var _ ledger_model.TransactionService = (*RestyTxService)(nil)

type RestyTxService struct {
	host       string
	port       int
	secure     bool
	gmSecure   bool
	url        string
	security   *SSLSecurity
	gmSecurity *GMSSLSecurity
}

func NewRestyTxService(host string, port int) *RestyTxService {
	url := fmt.Sprintf("http://%s:%d/rpc/tx", host, port)
	return &RestyTxService{
		host:   host,
		port:   port,
		secure: false,
		url:    url,
	}
}

func NewSecureRestyTxService(host string, port int, security *SSLSecurity) *RestyTxService {
	url := fmt.Sprintf("https://%s:%d/rpc/tx", host, port)
	return &RestyTxService{
		host:     host,
		port:     port,
		secure:   true,
		url:      url,
		security: security,
	}
}

func NewGMSecureRestyTxService(host string, port int, security *GMSSLSecurity) *RestyTxService {
	url := fmt.Sprintf("https://%s:%d/rpc/tx", host, port)
	return &RestyTxService{
		host:       host,
		port:       port,
		secure:     true,
		gmSecure:   true,
		url:        url,
		gmSecurity: security,
	}
}

func (r *RestyTxService) Process(txRequest *ledger_model.TransactionRequest) (response *ledger_model.TransactionResponse, err error) {
	msg, err := binary_proto.NewCodec().Encode(txRequest)
	if err != nil {
		return nil, err
	}
	if !r.gmSecure {
		return r.process(msg)
	} else {
		return r.gmProcess(msg)
	}
}

func (r *RestyTxService) process(msg []byte) (response *ledger_model.TransactionResponse, err error) {
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
	resp, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/bin-obj").
		SetBody(msg).
		Post(r.url)
	if !resp.IsSuccess() {
		err = errors.New(resp.String())
		return
	}
	if tresp, err := binary_proto.NewCodec().Decode(resp.Body()); err != nil {
		return nil, err
	} else {
		resp, ok := tresp.(ledger_model.TransactionResponse)
		if ok {
			return &resp, nil
		} else {
			return nil, errors.New("请求失败")
		}
	}
}

func (r *RestyTxService) gmProcess(msg []byte) (response *ledger_model.TransactionResponse, err error) {
	var certPool *gmx509.CertPool
	if r.gmSecurity != nil {
		certPool = r.gmSecurity.RootCerts
	}
	var certificates []gmtls.Certificate
	if nil != r.gmSecurity && nil != r.gmSecurity.EncCert && nil != r.gmSecurity.SigCert {
		certificates = []gmtls.Certificate{*r.gmSecurity.SigCert, *r.gmSecurity.EncCert}
	} else {
		certificates = []gmtls.Certificate{}
	}
	config := &gmtls.Config{
		GMSupport:          &gmtls.GMSupport{},
		RootCAs:            certPool,
		ClientAuth:         gmtls.NoClientCert,
		Certificates:       certificates,
		InsecureSkipVerify: r.security == nil || r.security.RootCerts == nil,
	}

	gmClient := gmtls.NewCustomHTTPSClient(config)
	result, err := gmClient.Post(r.url, "application/bin-obj", gbytes.NewBuffer(msg))
	defer result.Body.Close()
	if err != nil {
		return nil, err
	}
	raw, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	if tresp, err := binary_proto.NewCodec().Decode(raw); err != nil {
		return nil, err
	} else {
		resp, ok := tresp.(ledger_model.TransactionResponse)
		if ok {
			return &resp, nil
		} else {
			return nil, errors.New("请求失败")
		}
	}
}
