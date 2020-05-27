package ledger_model

import "fmt"

/*
 * Author: imuge
 * Date: 2020/5/28 下午6:39
 */

type State byte

const (
	// 可操作
	OPERABLE State = iota
	// 就绪
	PREPARED
	// 已提交
	COMMITTED
	// 已关闭
	CLOSED
)

type TxStateManager struct {
	state State
}

func NewTxStateManager() *TxStateManager {
	return &TxStateManager{
		state: OPERABLE,
	}
}

func (t *TxStateManager) operate() {
	if t.state != OPERABLE {
		panic(fmt.Sprintf("Cannot define operations in %v state!", t.state))
	}
}

func (t *TxStateManager) prepare() {
	if t.state != OPERABLE {
		panic(fmt.Sprintf("Cannot switch to %v state in %v state!", PREPARED, t.state))
	}
	t.state = PREPARED
}

func (t *TxStateManager) commit() {
	if t.state != PREPARED {
		panic(fmt.Sprintf("Cannot switch to %v state in %v state!", COMMITTED, t.state))
	}
	t.state = COMMITTED
}

func (t *TxStateManager) complete() {
	if t.state != COMMITTED {
		panic(fmt.Sprintf("Cannot complete normally in %v state!", t.state))
	}
	t.state = CLOSED
}

// 关闭交易
// 返回 此次操作前是否已经处于关闭状态
func (t *TxStateManager) close() bool {
	if t.state == CLOSED {
		return true
	}
	t.state = CLOSED
	return false
}
