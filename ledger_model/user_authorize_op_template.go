package ledger_model

import (
	"sync"
	"unsafe"
)

/*
 * Author: imuge
 * Date: 2020/6/1 下午1:40
 */

var _ UserAuthorizer = (*UserAuthorizeOpTemplate)(nil)

type UserAuthorizeOpTemplate struct {
	*UserAuthorizeOperation

	userAuthMap map[string]UserRolesAuthorizer
	mutex       sync.Mutex
}

func NewUserAuthorizeOpTemplate() *UserAuthorizeOpTemplate {
	return &UserAuthorizeOpTemplate{
		UserAuthorizeOperation: &UserAuthorizeOperation{},
		userAuthMap:            make(map[string]UserRolesAuthorizer),
	}
}

func (u *UserAuthorizeOpTemplate) ForUser(users [][]byte) UserRolesAuthorizer {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	address := *(*string)(unsafe.Pointer(&users))
	userAuth, ok := u.userAuthMap[address]
	if !ok {
		userAuth = NewAuthorizationDataEntry(u, users)
		u.userAuthMap[address] = userAuth
	}
	return userAuth
}

func (u *UserAuthorizeOpTemplate) getOperation() *UserAuthorizeOperation {
	return u.UserAuthorizeOperation
}
