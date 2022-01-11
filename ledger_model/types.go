package ledger_model

const (
	// 操作类型名称 TODO 不全
	OperationTypeUserRegister                     = "UserRegisterOperation"            // 用户注册
	OperationTypeDataAccountRegisterOperation     = "DataAccountRegisterOperation"     // 数据账户注册
	OperationTypeDataAccountKVSetOperation        = "DataAccountKVSetOperation"        // KV写入
	OperationTypeEventAccountRegisterOperation    = "EventAccountRegisterOperation"    // 事件账户注册
	OperationTypeEventPublishOperation            = "EventPublishOperation"            // 事件发布
	OperationTypeParticipantRegisterOperation     = "ParticipantRegisterOperation"     // 参与方注册
	OperationTypeParticipantStateUpdateOperation  = "ParticipantStateUpdateOperation"  // 参与方变更
	OperationTypeContractCodeDeployOperation      = "ContractCodeDeployOperation"      // 合约部署
	OperationTypeContractEventSendOperation       = "ContractEventSendOperation"       // 合约调用
	OperationTypeRolesConfigureOperation          = "RolesConfigureOperation"          // 角色配置
	OperationTypeUserAuthorizeOperation           = "UserAuthorizeOperation"           // 用户授权
	OperationTypeUserStateUpdate                  = "UserStateUpdateOperation"         // 用户状态变更
	OperationTypeContractStateUpdate              = "ContractStateUpdateOperation"     // 合约状态变更
	OperationTypeAccountPermissionSetOperation    = "AccountPermissionSetOperation"    // 账户权限变更
	OperationTypeUserCAUpdateOperation            = "UserCAUpdateOperation"            // 用户证书更新
	OperationTypeRootCAUpdateOperation            = "RootCAUpdateOperation"            // 根证书更新
	OperationTypeLedgerInitOperation              = "LedgerInitOperation"              // 账本初始化
	OperationTypeConsensusSettingsUpdateOperation = "ConsensusSettingsUpdateOperation" // 共识信息变更
	OperationTypeConsensusTypeUpdateOperation     = "ConsensusTypeUpdateOperation"     // 共识变更
	OperationTypeCryptoHashAlgoUpdateOperation    = "CryptoHashAlgoUpdateOperation"    // Hash算法变更
)
