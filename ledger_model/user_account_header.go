package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午3:36
 */

var _ binary_proto.DataContract = (*UserAccountHeader)(nil)

func init() {
	binary_proto.RegisterContract(UserAccountHeader{})
}

type UserAccountHeader struct {
	*BlockchainIdentity
	State       AccountState `refEnum:"788"`
	Certificate string       `primitiveType:"TEXT"`
}

func (b UserAccountHeader) ContractCode() int32 {
	return binary_proto.USER_ACCOUNT_HEADER
}

func (b UserAccountHeader) ContractName() string {
	return "UserAccountHeader"
}

func (b UserAccountHeader) Description() string {
	return ""
}
