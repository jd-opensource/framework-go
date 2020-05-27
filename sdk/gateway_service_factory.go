package sdk

import (
	"framework-go/ledger_model"
)

/*
 * Author: imuge
 * Date: 2020/5/27 下午4:35
 */

var _ BlockchainServiceFactory = (*GatewayServiceFactory)(nil)

type GatewayServiceFactory struct {
	userKey ledger_model.BlockchainKeypair

	blockchainService BlockchainService
}

func NewGatewayServiceFactory(userKey ledger_model.BlockchainKeypair, blockchainService BlockchainService) GatewayServiceFactory {
	return GatewayServiceFactory{
		userKey:           userKey,
		blockchainService: blockchainService,
	}
}

func Connect(gatewayHost string, gatewayPort int32, secure bool, userKey ledger_model.BlockchainKeypair) GatewayServiceFactory {
	service := NewGatewayBlockchainService(NewRestyBlockchainService(gatewayHost, gatewayPort, secure))
	return NewGatewayServiceFactory(userKey, service)
}

func (g GatewayServiceFactory) GetBlockchainService() BlockchainService {
	return g.blockchainService
}
