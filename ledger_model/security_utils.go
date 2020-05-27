package ledger_model

import "strings"

/*
 * Author: imuge
 * Date: 2020/5/29 下午5:47
 */

const MAX_ROLE_NAMES = 20

// 校验角色名称的有效性，并格式化角色名称：去掉两端空白字符，统一为大写字符
func FormatRoleName(roleName string) string {
	if len(roleName) == 0 {
		panic("Role name is empty!")
	}
	roleName = strings.TrimSpace(roleName)
	if len(roleName) > MAX_ROLE_NAMES {
		panic("Role name exceeds max length!")
	}
	if len(roleName) == 0 {
		panic("Role name is empty!")
	}

	return strings.ToUpper(roleName)
}
