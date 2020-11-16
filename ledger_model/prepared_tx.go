package ledger_model

import (
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
)

/*
 * Author: imuge
 * Date: 2020/5/29 上午10:25
 */

var _ PreparedTransaction = (*PreparedTx)(nil)

type PreparedTx struct {
	txReqBuilder TransactionRequestBuilder
	txService    TransactionService
}

func NewPreparedTx(txReqBuilder TransactionRequestBuilder, txService TransactionService) *PreparedTx {
	return &PreparedTx{
		txReqBuilder: txReqBuilder,
		txService:    txService,
	}
}

func (p *PreparedTx) GetHash() framework.HashDigest {
	return p.txReqBuilder.GetTransactionHash()
}

func (p *PreparedTx) GetTransactionContent() TransactionContent {
	return p.txReqBuilder.GetTransactionContent()
}

func (p *PreparedTx) Sign(keyPair framework.AsymmetricKeypair) DigitalSignature {
	signature := Sign(p.txReqBuilder.GetTransactionHash(), keyPair)
	p.AddSignature(signature)
	return signature
}

func (p *PreparedTx) AddSignature(signature DigitalSignature) {
	p.txReqBuilder.AddEndpointSignature(signature)
}

func (p *PreparedTx) Commit() (TransactionResponse, error) {
	// 生成请求；
	txReq := p.txReqBuilder.BuildRequest()
	// 发起交易请求；
	return p.txService.Process(txReq)
}
