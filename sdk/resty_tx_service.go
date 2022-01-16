package sdk

import (
	"crypto/tls"
	"errors"
	"fmt"
	binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"
	"github.com/blockchain-jd-com/framework-go/ledger_model"
	"github.com/go-resty/resty/v2"
)

/*
 * Author: imuge
 * Date: 2020/5/29 下午1:55
 */

var _ ledger_model.TransactionService = (*RestyTxService)(nil)

type RestyTxService struct {
	host     string
	port     int
	secure   bool
	url      string
	security *SSLSecurity
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

func (r *RestyTxService) Process(txRequest ledger_model.TransactionRequest) (response ledger_model.TransactionResponse, err error) {
	msg, _ := binary_proto.NewCodec().Encode(txRequest)

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
		return ledger_model.TransactionResponse{}, err
	} else {
		return tresp.(ledger_model.TransactionResponse), nil
	}
}
