package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

var _ binary_proto.DataContract = (*LedgerInitOperation)(nil)

func init() {
	binary_proto.RegisterContract(LedgerInitOperation{})
}

type LedgerInitOperation struct {
	InitSetting LedgerInitSetting `refContract:"1552"`
}

func (lso LedgerInitOperation) ContractCode() int32 {
	return binary_proto.TX_OP_LEDGER_INIT
}

func (lso LedgerInitOperation) ContractName() string {
	return OperationTypeLedgerInitOperation
}

func (lso LedgerInitOperation) Description() string {
	return ""
}
