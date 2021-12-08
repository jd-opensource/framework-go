package sdk

import (
	"github.com/blockchain-jd-com/framework-go/ledger_model"
)

/*
 * Author: imuge
 * Date: 2020/5/27 下午4:35
 */

var _ BlockchainServiceFactory = (*GatewayServiceFactory)(nil)

type GatewayServiceFactory struct {
	blockchainService BlockchainService
}

func NewGatewayServiceFactory(blockchainService BlockchainService) *GatewayServiceFactory {
	return &GatewayServiceFactory{
		blockchainService: blockchainService,
	}
}

func MustConnect(gatewayHost string, gatewayPort int, secure bool, userKey ledger_model.BlockchainKeypair) *GatewayServiceFactory {
	queryService := NewRestyQueryService(gatewayHost, gatewayPort, secure)
	txService := NewEndpointAutoSigner(userKey, NewRestyTxService(gatewayHost, gatewayPort, secure))
	ledgerHashs, err := queryService.GetLedgerHashs()
	if err != nil {
		panic(err)
	}
	cryptoSettings := make([]ledger_model.CryptoSetting, len(ledgerHashs))
	for i, ledger := range ledgerHashs {
		ledgerAdminInfo, err := queryService.GetLedgerAdminInfo(ledger)
		if err != nil {
			panic(err)
		}
		cryptoSettings[i] = ledgerAdminInfo.Settings.CryptoSetting
	}
	service := NewGatewayBlockchainService(ledgerHashs, cryptoSettings, txService, queryService)
	return NewGatewayServiceFactory(service)
}

func Connect(gatewayHost string, gatewayPort int, secure bool, userKey ledger_model.BlockchainKeypair) (*GatewayServiceFactory, error) {
	queryService := NewRestyQueryService(gatewayHost, gatewayPort, secure)
	txService := NewEndpointAutoSigner(userKey, NewRestyTxService(gatewayHost, gatewayPort, secure))
	ledgerHashs, err := queryService.GetLedgerHashs()
	if err != nil {
		return nil, err
	}
	cryptoSettings := make([]ledger_model.CryptoSetting, len(ledgerHashs))
	for i, ledger := range ledgerHashs {
		ledgerAdminInfo, err := queryService.GetLedgerAdminInfo(ledger)
		if err != nil {
			return nil, err
		}
		cryptoSettings[i] = ledgerAdminInfo.Settings.CryptoSetting
	}
	service := NewGatewayBlockchainService(ledgerHashs, cryptoSettings, txService, queryService)
	return NewGatewayServiceFactory(service), nil
}

func MustConnectWithoutUserKey(gatewayHost string, gatewayPort int, secure bool) *GatewayServiceFactory {
	queryService := NewRestyQueryService(gatewayHost, gatewayPort, secure)
	txService := NewRestyTxService(gatewayHost, gatewayPort, secure)
	ledgerHashs, err := queryService.GetLedgerHashs()
	if err != nil {
		panic(err)
	}
	cryptoSettings := make([]ledger_model.CryptoSetting, len(ledgerHashs))
	for i, ledger := range ledgerHashs {
		ledgerAdminInfo, err := queryService.GetLedgerAdminInfo(ledger)
		if err != nil {
			panic(err)
		}
		cryptoSettings[i] = ledgerAdminInfo.Settings.CryptoSetting
	}
	service := NewGatewayBlockchainService(ledgerHashs, cryptoSettings, txService, queryService)
	return NewGatewayServiceFactory(service)
}

func ConnectWithoutUserKey(gatewayHost string, gatewayPort int, secure bool) (*GatewayServiceFactory, error) {
	queryService := NewRestyQueryService(gatewayHost, gatewayPort, secure)
	txService := NewRestyTxService(gatewayHost, gatewayPort, secure)
	ledgerHashs, err := queryService.GetLedgerHashs()
	if err != nil {
		return nil, err
	}
	cryptoSettings := make([]ledger_model.CryptoSetting, len(ledgerHashs))
	for i, ledger := range ledgerHashs {
		ledgerAdminInfo, err := queryService.GetLedgerAdminInfo(ledger)
		if err != nil {
			return nil, err
		}
		cryptoSettings[i] = ledgerAdminInfo.Settings.CryptoSetting
	}
	service := NewGatewayBlockchainService(ledgerHashs, cryptoSettings, txService, queryService)
	return NewGatewayServiceFactory(service), nil
}

func (g GatewayServiceFactory) GetBlockchainService() BlockchainService {
	return g.blockchainService
}
