package ledger_model

import "sync"

/*
 * Author: imuge
 * Date: 2020/5/29 下午5:38
 */

var _ RolesConfigurer = (*RolesConfigureOpTemplate)(nil)

type RolesConfigureOpTemplate struct {
	*RolesConfigureOperation

	rolesMap map[string]RolePrivilegeConfigurer
	mutex    sync.Mutex
}

func NewRolesConfigureOpTemplate() *RolesConfigureOpTemplate {
	return &RolesConfigureOpTemplate{
		RolesConfigureOperation: &RolesConfigureOperation{},
		rolesMap:                make(map[string]RolePrivilegeConfigurer),
	}
}

func (r *RolesConfigureOpTemplate) Configure(roleName string) RolePrivilegeConfigurer {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	roleName = FormatRoleName(roleName)
	roleConfig, ok := r.rolesMap[roleName]
	if !ok {
		roleConfig = NewRolePrivilegeConfig(r, roleName)
		r.rolesMap[roleName] = roleConfig
	}
	return roleConfig
}

func (r *RolesConfigureOpTemplate) getOperation() *RolesConfigureOperation {
	return r.RolesConfigureOperation
}
