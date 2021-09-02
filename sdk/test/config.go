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
	GATEWAY_PORT = 8080
	SECURE       = false

	// 区块链上已存在的有操作权限的用户公私钥信息
	NODE_PRIVITE_KEY = crypto.DecodePrivKey(
		"177gk2TroY9kHetQu9NGxsLhNFdq65pdAUDrVxxEx6S15LhbYMW6uD5VD4Xj9HifE8cEtqS",
		base58.MustDecode("8EjkXVSTxMFjCvNNsTo8RBMDEVQmk7gYkW4SCDuvdsBG"))
	NODE_PUBLIC_KEY = crypto.DecodePubKey("7VeREdd9ah9pdJH4vnVR7pwjyn4RttCbanhULtPnLyRayrMz")
	NODE_KEY        = ledger_model.NewBlockchainKeypair(NODE_PUBLIC_KEY, NODE_PRIVITE_KEY)
)
