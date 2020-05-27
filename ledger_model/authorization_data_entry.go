package ledger_model

import (
	"sync"
	"unsafe"
)

/*
 * Author: imuge
 * Date: 2020/6/1 下午4:28
 */

var _ UserRolesAuthorizer = (*AuthorizationDataEntry)(nil)

type AuthorizationDataEntry struct {
	UserRolesEntry

	opTemplate  *UserAuthorizeOpTemplate
	authRoles   map[string]bool
	unauthRoles map[string]bool

	mutex sync.Mutex
}

func NewAuthorizationDataEntry(opTemplate *UserAuthorizeOpTemplate, users [][]byte) *AuthorizationDataEntry {
	return &AuthorizationDataEntry{
		opTemplate: opTemplate,
		UserRolesEntry: UserRolesEntry{
			Addresses: users,
		},
		authRoles:   make(map[string]bool),
		unauthRoles: make(map[string]bool),
	}
}

func (a *AuthorizationDataEntry) ForUser(users [][]byte) UserRolesAuthorizer {
	return a.opTemplate.ForUser(users)
}

func (a *AuthorizationDataEntry) Authorize(role string) UserRolesAuthorizer {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	roleName := FormatRoleName(role)
	delete(a.unauthRoles, roleName)
	a.authRoles[roleName] = true
	a.updateAuthorization()

	return a
}

func (a *AuthorizationDataEntry) Unauthorize(role string) UserRolesAuthorizer {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	roleName := FormatRoleName(role)
	delete(a.authRoles, roleName)
	a.unauthRoles[roleName] = true
	a.updateAuthorization()

	return a
}

func (a *AuthorizationDataEntry) SetPolicy(rolePolicy RolesPolicy) UserRolesAuthorizer {
	a.Policy = rolePolicy
	a.updateAuthorization()

	return a
}

func (a *AuthorizationDataEntry) updateAuthorization() {
	var ars []string
	var uars []string
	for p, ok := range a.authRoles {
		if ok {
			ars = append(ars, p)
		}
	}
	for p, ok := range a.unauthRoles {
		if ok {
			uars = append(uars, p)
		}
	}
	a.AuthorizedRoles = ars
	a.UnauthorizedRoles = uars

	added := false
	index := 0
	var auth UserRolesEntry
	for index, auth = range a.opTemplate.UserRolesAuthorizations {
		a1 := *(*string)(unsafe.Pointer(&auth.Addresses))
		a2 := *(*string)(unsafe.Pointer(&a.Addresses))
		if a1 == a2 {
			added = true
			break
		}
	}
	if added {
		a.opTemplate.UserRolesAuthorizations[index] = a.UserRolesEntry
	} else {
		a.opTemplate.UserRolesAuthorizations = append(a.opTemplate.UserRolesAuthorizations, a.UserRolesEntry)
	}
}
