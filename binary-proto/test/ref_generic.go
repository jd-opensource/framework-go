package test

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/25 上午11:58
 */

var _ binary_proto.DataContract = (*RefGeneric)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(RefGeneric{})
}

type RefGeneric struct {
}

func (r RefGeneric) ContractCode() int32 {
	return 0x04
}

func (r RefGeneric) ContractName() string {
	return ""
}

func (r RefGeneric) Description() string {
	return ""
}
