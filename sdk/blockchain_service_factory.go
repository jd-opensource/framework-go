package sdk

/*
 * Author: imuge
 * Date: 2020/5/27 下午4:33
 */

type BlockchainServiceFactory interface {
	GetBlockchainService() BlockchainService
}
