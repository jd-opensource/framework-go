package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午2:43
 */

var _ binary_proto.DataContract = (*EndpointRequest)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(EndpointRequest{})
}

type EndpointRequest struct {
	Hash []byte `primitiveType:"BYTES"`
	// 交易内容
	TransactionContent TransactionContent `refContract:"528"`
	// 终端用户的签名列表
	EndpointSignatures []DigitalSignature `refContract:"2864" list:"true"`
}

func (e EndpointRequest) ContractCode() int32 {
	return binary_proto.REQUEST_ENDPOINT
}

func (e EndpointRequest) ContractName() string {
	return "EndpointRequest"
}

func (e EndpointRequest) Description() string {
	return ""
}
