package ledger_model

import (
	binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
)

/*
 * Author: imuge
 * Date: 2020/5/25 下午7:28
 */

var _ binary_proto.DataContract = (*TransactionContent)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(TransactionContent{})
}

type TransactionContent struct {
	LedgerHash []byte                      `primitiveType:"BYTES"`
	Operations []binary_proto.DataContract `refContract:"768" genericContract:"true" list:"true"`
	Timestamp  int64                       `primitiveType:"INT64"`
}

func NewTransactionContent(ledgerHash framework.HashDigest, operations []binary_proto.DataContract, time int64) TransactionContent {
	return TransactionContent{
		LedgerHash: ledgerHash.ToBytes(),
		Operations: operations,
		Timestamp:  time,
	}
}

func (t TransactionContent) ContractCode() int32 {
	return binary_proto.TX_CONTENT
}

func (t TransactionContent) ContractName() string {
	return "TransactionContent"
}

func (t TransactionContent) Description() string {
	return ""
}
