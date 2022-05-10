package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/29 下午2:54
 */

type ParticipantService interface {
	/**
	激活参与方
	*/
	ActivateParticipant(params ActivateParticipantParams) (bool, error)

	/**
	移除参与方
	*/
	InactivateParticipant(params InactivateParticipantParams) (bool, error)
}

type ActivateParticipantParams struct {
	LedgerHash         string
	ConsensusHost      string // 待激活节点共识地址
	ConsensusPort      int    // 待激活节点共识端口
	ConsensusSecure    bool   // 待激活节点共识服务是否启动安全连接
	RemoteManageHost   string // 数据同步节点地址
	RemoteManagePort   int    // 数据同步节点端口
	RemoteManageSecure bool   // 数据同步节点服务是否启动安全连接
	Shutdown           bool   // 是否停止旧的节点服务
}

type InactivateParticipantParams struct {
	LedgerHash         string
	ParticipantAddress string
}
