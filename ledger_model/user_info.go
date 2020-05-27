package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午3:35
 */

var _ binary_proto.DataContract = (*UserInfo)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(UserInfo{})
}

type UserInfo struct {
	UserAccountHeader
}

func (b UserInfo) ContractCode() int32 {
	return binary_proto.USER_INFO
}

func (b UserInfo) ContractName() string {
	return "UserInfo"
}

func (b UserInfo) Description() string {
	return ""
}
