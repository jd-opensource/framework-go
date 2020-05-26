package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/25 下午7:28
 */

var _ binary_proto.DataContract = (*TransactionContentBody)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(TransactionContentBody{})
}

type TransactionContentBody struct {
	LedgerHash []byte                      `primitiveType:"BYTES"`
	Operations []binary_proto.DataContract `refContract:"768" genericContract:"true" repeatable:"true"`
	Timestamp  int64                       `primitiveType:"INT64"`
}

func (t TransactionContentBody) Code() int32 {
	return binary_proto.TX_CONTENT_BODY
}

func (t TransactionContentBody) Name() string {
	return "TransactionContent"
}

func (t TransactionContentBody) Description() string {
	return ""
}
