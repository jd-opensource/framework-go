package sdk

import "github.com/blockchain-jd-com/framework-go/ledger_model"

/*
 * Author: imuge
 * Date: 2020/5/27 下午4:23
 */

type BlockchainService interface {
	ledger_model.BlockchainQueryService
	BlockchainTransactionService
	BlockchainEventService
}
