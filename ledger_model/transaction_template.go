package ledger_model

import "github.com/blockchain-jd-com/framework-go/crypto/framework"

/*
 * Author: imuge
 * Date: 2020/5/28 下午6:19
 */

type TransactionTemplate interface {
	ClientOperator
	GetLedgerHash() framework.HashDigest
	Prepare() PreparedTransaction
}
