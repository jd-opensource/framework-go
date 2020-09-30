package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/29 下午2:54
 */

type ActiveParticipantService interface {
	/**
	ledgerHash	账本HASH
	ip 待激活节点IP
	port 待激活节点端口
	remoteManageHost	其他任一非拜占庭节点IP
	remoteManagePort	其他任一非拜占庭节点管理端口
	 */
	ActivateParticipant(ledgerHash, ip string, port int, remoteManageHost string, remoteManagePort int) (bool, error)

	/**
	ledgerHash	账本HASH
	participantAddress 待移除节点地址
	remoteManageHost	其他任一非拜占庭节点IP
	remoteManagePort	其他任一非拜占庭节点管理端口
	*/
	InactivateParticipant(ledgerHash, participantAddress, remoteManageHost string, remoteManagePort int) (bool, error)
}
