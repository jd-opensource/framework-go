package ledger_model

import "sync"

/*
 * Author: imuge
 * Date: 2020/5/29 下午5:52
 */

var _ RolePrivilegeConfigurer = (*RolePrivilegeConfig)(nil)

type RolePrivilegeConfig struct {
	RolePrivilegeEntry
	opTemplate               *RolesConfigureOpTemplate
	enableLedgerPermissions  map[LedgerPermission]bool
	disableLedgerPermissions map[LedgerPermission]bool
	enableTxPermissions      map[TransactionPermission]bool
	disableTxPermissions     map[TransactionPermission]bool
	mutex                    sync.Mutex
}

func NewRolePrivilegeConfig(opTemplate *RolesConfigureOpTemplate, name string) *RolePrivilegeConfig {
	return &RolePrivilegeConfig{
		opTemplate: opTemplate,
		RolePrivilegeEntry: RolePrivilegeEntry{
			RoleName: name,
		},
		enableLedgerPermissions:  make(map[LedgerPermission]bool),
		disableLedgerPermissions: make(map[LedgerPermission]bool),
		enableTxPermissions:      make(map[TransactionPermission]bool),
		disableTxPermissions:     make(map[TransactionPermission]bool),
	}
}

func (r *RolePrivilegeConfig) Configure(roleName string) RolePrivilegeConfigurer {
	return r.opTemplate.Configure(roleName)
}

func (r *RolePrivilegeConfig) RoleName() string {
	return r.RolePrivilegeEntry.RoleName
}

func (r *RolePrivilegeConfig) DisableTransactionPermission(permission TransactionPermission) RolePrivilegeConfigurer {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.enableTxPermissions[permission] = false
	r.disableTxPermissions[permission] = true
	r.updatePermissions()

	return r
}

func (r *RolePrivilegeConfig) EnableTransactionPermission(permission TransactionPermission) RolePrivilegeConfigurer {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.enableTxPermissions[permission] = true
	r.disableTxPermissions[permission] = false
	r.updatePermissions()

	return r
}

func (r *RolePrivilegeConfig) DisableLedgerPermission(permission LedgerPermission) RolePrivilegeConfigurer {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.enableLedgerPermissions[permission] = false
	r.disableLedgerPermissions[permission] = true
	r.updatePermissions()

	return r
}

func (r *RolePrivilegeConfig) EnableLedgerPermission(permission LedgerPermission) RolePrivilegeConfigurer {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.enableLedgerPermissions[permission] = true
	r.disableLedgerPermissions[permission] = false
	r.updatePermissions()

	return r
}

func (r *RolePrivilegeConfig) updatePermissions() {
	var elps []LedgerPermission
	var dlps []LedgerPermission
	var etps []TransactionPermission
	var dtps []TransactionPermission
	for p, ok := range r.enableLedgerPermissions {
		if ok {
			elps = append(elps, p)
		}
	}
	for p, ok := range r.disableLedgerPermissions {
		if ok {
			dlps = append(dlps, p)
		}
	}
	for p, ok := range r.enableTxPermissions {
		if ok {
			etps = append(etps, p)
		}
	}
	for p, ok := range r.disableTxPermissions {
		if ok {
			dtps = append(dtps, p)
		}
	}
	r.EnableLedgerPermissions = elps
	r.DisableLedgerPermissions = dlps
	r.EnableTransactionPermissions = etps
	r.DisableTransactionPermissions = dtps

	added := false
	index := 0
	var role RolePrivilegeEntry
	for index, role = range r.opTemplate.Roles {
		if role.RoleName == r.RoleName() {
			added = true
			break
		}
	}
	if added {
		r.opTemplate.Roles[index] = r.RolePrivilegeEntry
	} else {
		r.opTemplate.Roles = append(r.opTemplate.Roles, r.RolePrivilegeEntry)
	}

}
