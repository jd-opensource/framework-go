package ledger_model

type DataAccountInfo struct {
	*BlockchainIdentity
	// 数据权限
	Permission DataPermission `json:"permission"`
	// KV总数
	DataCount int64
}
