package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午3:49
 */

var _ binary_proto.DataContract = (*RoleSet)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(RoleSet{})
}

type RoleSet struct {
	Policy RolesPolicy `refEnum:"1027"`
	Roles  []string    `primitiveType:"TEXT" list:"true"`
}

func (r RoleSet) ContractCode() int32 {
	return binary_proto.ROLE_SET
}

func (r RoleSet) ContractName() string {
	return "RoleSet"
}

func (r RoleSet) Description() string {
	return ""
}
