package test

import "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/22 下午3:55
 */

var _ binary_proto.DataContract = (*RefContract)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(RefContract{})
}

type RefContract struct {
	RefContractInner
}

func NewRefContract() RefContract {
	return RefContract{NewRefContractInner()}
}

func (J RefContract) ContractCode() int32 {
	return 0x03
}

func (J RefContract) ContractName() string {
	return ""
}

func (J RefContract) Description() string {
	return ""
}

func (J RefContract) Equals(contract RefContract) bool {
	return J.I8 == contract.I8
}

type RefContractInner struct {
	I8 int8 `primitiveType:"INT8"`
}

func (r RefContractInner) ContractCode() int32 {
	return 0x05
}

func (r RefContractInner) ContractName() string {
	return ""
}

func (r RefContractInner) Description() string {
	return ""
}

var _ binary_proto.DataContract = (*RefContractInner)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(RefContractInner{})
}

func NewRefContractInner() RefContractInner {
	return RefContractInner{1}
}
