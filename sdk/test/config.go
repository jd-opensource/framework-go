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
	GATEWAY_HOST = "127.0.0.1"
	// 网关服务端口
	GATEWAY_PORT = 8080
	// 是否建立安全连接
	SECURE = true
	// 是否是国密
	GM_SECURE = true
	// SSL配置
	SSL_ROOT_CERT        = "/home/imuge/jd/nodes/peer0/config/certs/root.crt"
	SSL_CLIENT_CERT      = "/home/imuge/jd/nodes/peer0/config/certs/tls/gw0.crt"
	SSL_CLIENT_KEY       = "/home/imuge/jd/nodes/peer0/config/keys/gw0.key"
	SSL_CLIENT_ENC_KEY   = "/home/imuge/jd/nodes/peer0/config/keys/gw0.enc.key"
	SSL_CLIENT_ENC_CERT  = "/home/imuge/jd/nodes/peer0/config/certs/tls/gw0.enc.crt"
	SSL_CLIENT_SIGN_KEY  = "/home/imuge/jd/nodes/peer0/config/keys/gw0.sign.key"
	SSL_CLIENT_SIGN_CERT = "/home/imuge/jd/nodes/peer0/config/certs/tls/gw0.sign.crt"

	// KEYPAIR身份模式，区块链上已存在的有操作权限的用户公私钥信息
	NODE_PRIVITE_KEY = crypto.MustDecodePrivKey(
		"177gjsSFEGxuqp8M8yzFy3QjGarKeJ9hyb2vWkUYBXN5yAUwcSamDZpXS4NpynXKeQJd9Pc",
		base58.MustDecode("AXhhKihAa2LaRwY5mftnngSPKDF4N9JignnQ4skynY8y"))
	NODE_PUBLIC_KEY = crypto.MustDecodePubKey("7VeRJbVfA9wNM2RnQwYuzqzvwMqCmHPqHqANxWLhFQZrvCjB")

	// CA身份模式，区块链上已存在的有操作权限的用户证书和私钥信息
	//cert, _             = ca.RetrieveCertificateFile("/home/imuge/jd/nodes/peer0/config/certs/sign/gw0.crt")
	//NODE_PUBLIC_KEY, _  = ca.RetrievePubKey(cert)
	//NODE_PRIVITE_KEY, _ = ca.RetrievePrivKeyFile(cert.Algorithm, "/home/imuge/jd/nodes/peer0/config/keys/gw0.key")

	NODE_KEY, _ = ledger_model.NewBlockchainKeypair(NODE_PUBLIC_KEY, NODE_PRIVITE_KEY)
)
