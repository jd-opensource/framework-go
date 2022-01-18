package ledger_model

type EventAccountInfo struct {
	*BlockchainIdentity
	// 数据权限
	Permission DataPermission `json:"permission"`
	// KV总数
	DataCount int64
}
