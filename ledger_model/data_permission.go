package ledger_model

type DataPermission struct {
	// 所属角色
	Role string `json:"role"`
	// 所属用户列表
	Owners []string `json:"owners"`
	// 权限值
	Mode string `json:"modeBits"`
}
