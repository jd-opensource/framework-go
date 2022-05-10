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

func MustConnect(gatewayHost string, gatewayPort int, userKey *ledger_model.BlockchainKeypair) *GatewayServiceFactory {
	queryService := NewRestyQueryService(gatewayHost, gatewayPort)
	txService := NewEndpointAutoSigner(userKey, NewRestyTxService(gatewayHost, gatewayPort))
	ledgerHashs, err := queryService.GetLedgerHashs()
	if err != nil {
		panic(err)
	}
	cryptoSettings := make([]ledger_model.CryptoSetting, len(ledgerHashs))
	for i, ledger := range ledgerHashs {
		cryptoSetting, err := queryService.GetLedgerCryptoSetting(ledger)
		if err != nil {
			panic(err)
		}
		cryptoSettings[i] = cryptoSetting
	}
	service := NewGatewayBlockchainService(ledgerHashs, cryptoSettings, txService, queryService)
	return NewGatewayServiceFactory(service)
}

func MustSecureConnect(gatewayHost string, gatewayPort int, userKey *ledger_model.BlockchainKeypair, security *SSLSecurity) *GatewayServiceFactory {
	queryService := NewSecureRestyQueryService(gatewayHost, gatewayPort, security)
	txService := NewEndpointAutoSigner(userKey, NewSecureRestyTxService(gatewayHost, gatewayPort, security))
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

func MustGMSecureConnect(gatewayHost string, gatewayPort int, userKey *ledger_model.BlockchainKeypair, security *GMSSLSecurity) *GatewayServiceFactory {
	queryService := NewGMSecureRestyQueryService(gatewayHost, gatewayPort, security)
	txService := NewEndpointAutoSigner(userKey, NewGMSecureRestyTxService(gatewayHost, gatewayPort, security))
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

func Connect(gatewayHost string, gatewayPort int, userKey *ledger_model.BlockchainKeypair) (*GatewayServiceFactory, error) {
	queryService := NewRestyQueryService(gatewayHost, gatewayPort)
	txService := NewEndpointAutoSigner(userKey, NewRestyTxService(gatewayHost, gatewayPort))
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

func SecureConnect(gatewayHost string, gatewayPort int, userKey *ledger_model.BlockchainKeypair, security *SSLSecurity) (*GatewayServiceFactory, error) {
	queryService := NewRestyQueryService(gatewayHost, gatewayPort)
	txService := NewEndpointAutoSigner(userKey, NewSecureRestyTxService(gatewayHost, gatewayPort, security))
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

func MustConnectWithoutUserKey(gatewayHost string, gatewayPort int) *GatewayServiceFactory {
	queryService := NewRestyQueryService(gatewayHost, gatewayPort)
	txService := NewRestyTxService(gatewayHost, gatewayPort)
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

func MustSecureConnectWithoutUserKey(gatewayHost string, gatewayPort int, security *SSLSecurity) *GatewayServiceFactory {
	queryService := NewRestyQueryService(gatewayHost, gatewayPort)
	txService := NewSecureRestyTxService(gatewayHost, gatewayPort, security)
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

func MustGMSecureConnectWithoutUserKey(gatewayHost string, gatewayPort int, security *GMSSLSecurity) *GatewayServiceFactory {
	queryService := NewGMSecureRestyQueryService(gatewayHost, gatewayPort, security)
	txService := NewGMSecureRestyTxService(gatewayHost, gatewayPort, security)
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

func ConnectWithoutUserKey(gatewayHost string, gatewayPort int, security *SSLSecurity) (*GatewayServiceFactory, error) {
	queryService := NewSecureRestyQueryService(gatewayHost, gatewayPort, security)
	txService := NewSecureRestyTxService(gatewayHost, gatewayPort, security)
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

func SecureConnectWithoutUserKey(gatewayHost string, gatewayPort int, security *SSLSecurity) (*GatewayServiceFactory, error) {
	queryService := NewSecureRestyQueryService(gatewayHost, gatewayPort, security)
	txService := NewSecureRestyTxService(gatewayHost, gatewayPort, security)
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
