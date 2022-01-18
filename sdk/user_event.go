package sdk

import (
	"fmt"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/ledger_model"
	"time"
)

/*
 * Author: imuge
 * Date: 2020/7/3 下午3:29
 */

type UserEventPoint struct {
	EventAccount string // 事件账户地址
	EventName    string // 事件名
	Sequence     int64  // 序列
}

func NewUserEventPoint(eventAccount, eventName string, sequence int64) UserEventPoint {
	return UserEventPoint{
		EventAccount: eventAccount,
		EventName:    eventName,
		Sequence:     sequence,
	}
}

// 事件监听器
type UserEventListener interface {
	OnEvent(event *ledger_model.Event, context *UserEventContext)
}

type UserEventContext struct {
	LedgerHash   *framework.HashDigest
	EventHandler *UserEventListenerHandle
}

func NewUserEventContext(ledgerHash *framework.HashDigest, eventHandler *UserEventListenerHandle) *UserEventContext {
	return &UserEventContext{
		LedgerHash:   ledgerHash,
		EventHandler: eventHandler,
	}
}

// 事件监听处理器
type UserEventListenerHandle struct {
	queryService ledger_model.BlockchainQueryService // 区块链查询器
	ledgerHash   *framework.HashDigest               // 账本hash

	eventPoints []UserEventPoint  // 监听事件列表
	listener    UserEventListener // 事件监听器

	eventSequences map[string]int64 // 事件当前监听起始序号

	stop bool
}

func NewUserEventListenerHandle(queryService ledger_model.BlockchainQueryService, ledgerHash *framework.HashDigest, eventPoints []UserEventPoint, listener UserEventListener) UserEventListenerHandle {
	// init event sequences
	eventSequences := make(map[string]int64, len(eventPoints))
	for _, point := range eventPoints {
		eventSequences[point.EventAccount+point.EventName] = point.Sequence
	}

	handler := UserEventListenerHandle{
		queryService: queryService,
		ledgerHash:   ledgerHash,

		eventPoints: eventPoints,
		listener:    listener,

		eventSequences: eventSequences,
	}

	// start events request
	go handler.start()

	return handler
}

func (h *UserEventListenerHandle) EventPoints() []UserEventPoint {
	return h.eventPoints
}

func (h *UserEventListenerHandle) Cancel() {
	h.stop = true
}

func (h *UserEventListenerHandle) start() {
	for !h.stop {
		// 每隔一秒监听一次
		time.Sleep(time.Second)
		h.loadEvents()
	}
}

func (h *UserEventListenerHandle) loadEvents() {
	for _, point := range h.eventPoints {

		startSequence := h.eventSequences[point.EventAccount+point.EventName]
		if startSequence < 0 {
			startSequence = 0
		}
		events, err := h.queryService.GetUserEvents(h.ledgerHash, point.EventAccount, point.EventName, startSequence, 10)
		if err != nil {
			fmt.Println(err)
			break
		}

		maxSequence := startSequence
		for _, event := range events {
			if event.Sequence > maxSequence {
				maxSequence = event.Sequence
			}
			h.listener.OnEvent(event, NewUserEventContext(h.ledgerHash, h))
		}

		if len(events) > 0 {
			h.eventSequences[point.EventAccount+point.EventName] = maxSequence + 1
		}
	}
}
