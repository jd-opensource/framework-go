package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午4:00
 */

var _ binary_proto.DataContract = (*ContractRuntimeConfig)(nil)

func init() {
	binary_proto.RegisterContract(ContractRuntimeConfig{})
}

type ContractRuntimeConfig struct {
	Timeout int64 `primitiveType:"INT64"`
}

func (c ContractRuntimeConfig) ContractCode() int32 {
	return binary_proto.CONTRACT_RUNTIME_CONFIG
}

func (c ContractRuntimeConfig) ContractName() string {
	return "ContractInfo"
}

func (c ContractRuntimeConfig) Description() string {
	return ""
}
