package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/25 下午6:50
 */

var _ binary_proto.DataContract = (*Operation)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(Operation{})
}

type Operation struct {
}

func (o Operation) ContractCode() int32 {
	return binary_proto.TX_OP
}

func (o Operation) ContractName() string {
	return "Operation"
}

func (o Operation) Description() string {
	return ""
}
