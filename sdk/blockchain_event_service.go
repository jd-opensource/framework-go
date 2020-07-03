package sdk

import "framework-go/crypto/framework"

/*
 * Author: imuge
 * Date: 2020/7/3 下午3:06
 */

type BlockchainEventService interface {

	/**
	 * 监听系统事件；
	 *
	 * @param ledgerHash
	 * @param eventPoint
	 * @param listener
	 * @return
	 */
	MonitorSystemEvent(ledgerHash framework.HashDigest, eventPoint SystemEventPoint, listener SystemEventListener) SystemEventListenerHandle

	/**
	 * 监听用户事件；
	 *
	 * @param ledgerHash
	 * @param eventAccount  事件账户地址；
	 * @param eventName
	 * @param startSequence
	 * @param listener
	 * @return
	 */
	MonitorUserEvent(ledgerHash framework.HashDigest, eventAccount, eventName string, startSequence int64, listener UserEventListener) UserEventListenerHandle

	/**
	* 监听用户事件列表
	*
	* @param ledgerHash
	* @param startingEventPoints
	* @param listener
	* @return
	 */
	MonitorUserEvents(ledgerHash framework.HashDigest, startingEventPoints []UserEventPoint, listener UserEventListener) UserEventListenerHandle
}
