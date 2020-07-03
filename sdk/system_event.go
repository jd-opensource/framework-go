package sdk

import (
	"fmt"
	"framework-go/crypto/framework"
	"framework-go/ledger_model"
	"framework-go/utils/bytes"
	"time"
)

/*
 * Author: imuge
 * Date: 2020/7/3 下午3:29
 */

type SystemEventPoint struct {
	EventName    string // 事件名
	Sequence     int64  // 序列
	MaxBatchSize int64  // 一次监听处理接受最大事件数量
}

func NewSystemEventPoint(eventName string, sequence int64) SystemEventPoint {
	return SystemEventPoint{
		EventName:    eventName,
		Sequence:     sequence,
		MaxBatchSize: 10,
	}
}

func NewSystemEventPointWithCustomBatchSize(eventName string, sequence int64, batchSize int64) SystemEventPoint {
	return SystemEventPoint{
		EventName:    eventName,
		Sequence:     sequence,
		MaxBatchSize: batchSize,
	}
}

// 事件监听器
type SystemEventListener interface {
	OnEvents(events []ledger_model.Event, context SystemEventContext)
}

type SystemEventContext struct {
	LedgerHash   framework.HashDigest
	EventHandler *SystemEventListenerHandle
}

func NewSystemEventContext(ledgerHash framework.HashDigest, eventHandler *SystemEventListenerHandle) SystemEventContext {
	return SystemEventContext{
		LedgerHash:   ledgerHash,
		EventHandler: eventHandler,
	}
}

// 事件监听处理器
type SystemEventListenerHandle struct {
	queryService ledger_model.BlockchainQueryService // 区块链查询器
	ledgerHash   framework.HashDigest                // 账本hash

	eventPoint SystemEventPoint    // 监听事件
	listener   SystemEventListener // 事件监听器

	eventSequence int64 // 事件当前监听起始序号

	stop bool
}

func NewSystemEventListenerHandle(queryService ledger_model.BlockchainQueryService, ledgerHash framework.HashDigest, eventPoint SystemEventPoint, listener SystemEventListener) SystemEventListenerHandle {
	handler := SystemEventListenerHandle{
		queryService: queryService,
		ledgerHash:   ledgerHash,

		eventPoint: eventPoint,
		listener:   listener,

		eventSequence: eventPoint.Sequence,
	}

	// start events request
	go handler.start()

	return handler
}

func (h *SystemEventListenerHandle) EventPoint() SystemEventPoint {
	return h.eventPoint
}

func (h *SystemEventListenerHandle) Cancel() {
	h.stop = true
}

func (h *SystemEventListenerHandle) start() {
	for !h.stop {
		// 每隔一秒监听一次
		time.Sleep(time.Second)
		h.loadEvents()
	}
}

func (h *SystemEventListenerHandle) loadEvents() {
	startSequence := h.eventSequence

	if h.eventPoint.EventName == "new_block" {
		ledgerInfo, err := h.queryService.GetLedger(h.ledgerHash)
		if err != nil {
			fmt.Println(err)
			return
		}
		events := make([]ledger_model.Event, 0)
		for i := startSequence; i < startSequence+h.eventPoint.MaxBatchSize && i <= ledgerInfo.LatestBlockHeight; i++ {
			info := ledger_model.Event{
				Name:     h.eventPoint.EventName,
				Sequence: i,
				Content: ledger_model.BytesValue{
					Type:  ledger_model.INT64,
					Bytes: bytes.Int64ToBytes(i),
				},
				BlockHeight: i,
			}
			events = append(events, info)
		}

		maxSequence := startSequence
		for _, event := range events {
			if event.Sequence > maxSequence {
				maxSequence = event.Sequence
			}
		}
		if len(events) > 0 {
			h.listener.OnEvents(events, NewSystemEventContext(h.ledgerHash, h))
		}

		if maxSequence > startSequence {
			h.eventSequence = maxSequence + 1
		}
	}
}
