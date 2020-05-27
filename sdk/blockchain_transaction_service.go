package sdk

import (
	"framework-go/crypto/framework"
	"framework-go/ledger_model"
)

/*
 * Author: imuge
 * Date: 2020/5/28 下午3:58
 */

type BlockchainTransactionService interface {

	// 发起新交易
	NewTransaction(ledgerHash framework.HashDigest) ledger_model.TransactionTemplate

	// 根据交易内容准备交易实例
	PrepareTransaction(content ledger_model.TransactionContent) ledger_model.PreparedTransaction
}
