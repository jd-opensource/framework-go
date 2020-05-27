package sdk

import (
	"errors"
	"fmt"
	binary_proto "framework-go/binary-proto"
	"framework-go/ledger_model"
	"github.com/go-resty/resty/v2"
)

/*
 * Author: imuge
 * Date: 2020/5/29 下午1:55
 */

var _ ledger_model.TransactionService = (*RestyTxService)(nil)

type RestyTxService struct {
	host   string
	port   int
	secure bool
	client *resty.Client
	url    string
}

func NewRestyTxService(host string, port int, secure bool) *RestyTxService {
	var url string
	if secure {
		url = fmt.Sprintf("https://%s:%d/rpc/tx", host, port)
	} else {
		url = fmt.Sprintf("http://%s:%d/rpc/tx", host, port)
	}
	return &RestyTxService{
		host:   host,
		port:   port,
		secure: secure,
		client: resty.New(),
		url:    url,
	}
}

func (r *RestyTxService) Process(txRequest ledger_model.TransactionRequest) (response ledger_model.TransactionResponse, err error) {
	msg, _ := binary_proto.Cdc.Encode(txRequest)

	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/bin-obj").
		SetBody(msg).
		Post(r.url)
	fmt.Println(fmt.Sprintf("%s \n %v \n", r.url, resp))
	if !resp.IsSuccess() {
		err = errors.New(resp.String())
		return
	}
	tresp, err := binary_proto.Cdc.Decode(resp.Body())
	if err != nil {
		return
	}

	return tresp.(ledger_model.TransactionResponse), nil
}
