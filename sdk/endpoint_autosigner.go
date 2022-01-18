package sdk

import "github.com/blockchain-jd-com/framework-go/ledger_model"

/*
 * Author: imuge
 * Date: 2020/5/29 下午1:18
 */

var _ ledger_model.TransactionService = (*EndpointAutoSigner)(nil)

type EndpointAutoSigner struct {
	innerService ledger_model.TransactionService
	userKey      *ledger_model.BlockchainKeypair
}

func NewEndpointAutoSigner(userKey *ledger_model.BlockchainKeypair, service ledger_model.TransactionService) *EndpointAutoSigner {
	return &EndpointAutoSigner{
		innerService: service,
		userKey:      userKey,
	}
}

func (e *EndpointAutoSigner) Process(txRequest *ledger_model.TransactionRequest) (*ledger_model.TransactionResponse, error) {
	// TODO: 未实现按不同的账本的密码参数配置，采用不同的哈希算法和签名算法；
	if !txRequest.ContainsEndpointSignature(e.userKey.GetIdentity().PubKey) {
		// TODO: 优化上下文对此 TransactionContent 的多次序列化带来的额外性能开销；
		signature, err := ledger_model.SignBytes(txRequest.TransactionHash, e.userKey.AsymmetricKeypair)
		if err != nil {
			return nil, err
		}
		txRequest.AddEndpointSignatures(signature)
	}

	return e.innerService.Process(txRequest)
}
