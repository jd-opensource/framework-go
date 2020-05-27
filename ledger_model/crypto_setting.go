package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午1:49
 */

var _ binary_proto.DataContract = (*CryptoSetting)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(CryptoSetting{})
}

type CryptoSetting struct {
	// 系统支持的密码服务提供者
	SupportedProviders []CryptoProvider `refContract:"1603" list:"true"`
	// 系统中使用的 Hash 算法
	HashAlgorithm int16 `primitiveType:"INT16"`
	// 当有加载附带哈希摘要的数据时，是否重新计算哈希摘要进行完整性校验
	AutoVerifyHash bool `primitiveType:"BOOLEAN"`
}

func (c CryptoSetting) ContractCode() int32 {
	return binary_proto.METADATA_CRYPTO_SETTING
}

func (c CryptoSetting) ContractName() string {
	return "CryptoSetting"
}

func (c CryptoSetting) Description() string {
	return ""
}
