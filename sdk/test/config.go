package test

import (
	"github.com/blockchain-jd-com/framework-go/crypto"
	"github.com/blockchain-jd-com/framework-go/ledger_model"
	"github.com/blockchain-jd-com/framework-go/utils/base58"
)

/*
 * Author: imuge
 * Date: 2020/5/29 下午3:21
 */

var (
	// 网关服务IP
	GATEWAY_HOST = "localhost"
	// 网关服务端口
	GATEWAY_PORT = 11000
	SECURE       = false

	// 区块链上已存在的有操作权限的用户公私钥信息
	NODE_PRIVITE_KEY = crypto.DecodePrivKey(
		"177gjzHTznYdPgWqZrH43W3yp37onm74wYXT4v9FukpCHBrhRysBBZh7Pzdo5AMRyQGJD7x",
		base58.MustDecode("DYu3G8aGTMBW1WrTw76zxQJQU4DHLw9MLyy7peG4LKkY"))
	NODE_PUBLIC_KEY = crypto.DecodePubKey("3snPdw7i7PjVKiTH2VnXZu5H8QmNaSXpnk4ei533jFpuifyjS5zzH9")
	NODE_KEY        = ledger_model.NewBlockchainKeypair(NODE_PUBLIC_KEY, NODE_PRIVITE_KEY)
)
