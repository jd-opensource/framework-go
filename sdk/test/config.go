package test

import (
	"github.com/blockchain-jd-com/framework-go/ledger_model"
	"github.com/blockchain-jd-com/framework-go/utils/ca"
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
	SSL_CLIENT_CERT      = ""
	SSL_CLIENT_KEY       = ""
	SSL_CLIENT_ENC_KEY   = "/home/imuge/jd/nodes/peer0/config/keys/gw0.enc.key"
	SSL_CLIENT_ENC_CERT  = "/home/imuge/jd/nodes/peer0/config/certs/tls/gw0.enc.crt"
	SSL_CLIENT_SIGN_KEY  = "/home/imuge/jd/nodes/peer0/config/keys/gw0.sign.key"
	SSL_CLIENT_SIGN_CERT = "/home/imuge/jd/nodes/peer0/config/certs/tls/gw0.sign.crt"

	// KEYPAIR身份模式，区块链上已存在的有操作权限的用户公私钥信息
	//NODE_PRIVITE_KEY = crypto.MustDecodePrivKey(
	//	"177gjyiJjdZNfEu4kgR97BftoUtJ54ixiQZS9uktxqtDue6bGwwBLz4hEXw7Gu5fgGpGceG",
	//	base58.MustDecode("AXhhKihAa2LaRwY5mftnngSPKDF4N9JignnQ4skynY8y"))
	//NODE_PUBLIC_KEY = crypto.MustDecodePubKey("7VeRKo8hU8mYfvGUGbNuCrfDqSY6PqgBt1dWWJwP6ofJnzkR")

	// CA身份模式，区块链上已存在的有操作权限的用户证书和私钥信息
	cert, _             = ca.RetrieveCertificate("-----BEGIN CERTIFICATE-----\nMIICFTCCAbugAwIBAgIEB73U1TAKBggqgRzPVQGDdTByMQwwCgYDVQQKDANKRFQx\nDTALBgNVBAsMBFJPT1QxCzAJBgNVBAYTAkNOMQswCQYDVQQIDAJCSjELMAkGA1UE\nBwwCQkoxDTALBgNVBAMMBHJvb3QxHTAbBgkqhkiG9w0BCQEWDmpkY2hhaW5AamQu\nY29tMB4XDTIyMDQyNjEyNTk0MloXDTMyMDQyMzEyNTk0MlowbzEMMAoGA1UECgwD\nSkRUMQswCQYDVQQLDAJHVzELMAkGA1UEBhMCQ04xCzAJBgNVBAgMAkJKMQswCQYD\nVQQHDAJCSjEMMAoGA1UEAwwDZ3cwMR0wGwYJKoZIhvcNAQkBFg5qZGNoYWluQGpk\nLmNvbTBZMBMGByqGSM49AgEGCCqBHM9VAYItA0IABFmmcV/tfDNgtVxhtUImWL+S\nnw7kE1ACG2MCAJvncBOF091l9ZAVEBGoMCuNqg3ALWb0jZmFHOvO/KnB9QYRcDqj\nQjBAMA4GA1UdDwEB/wQEAwIHgDAgBgNVHQ4BAf8EFgQUQo+ujig/quzYPTb/effw\n9yCJ590wDAYDVR0TAQH/BAIwADAKBggqgRzPVQGDdQNIADBFAiAZSeOrrZANqT10\ne3uT8SI83SD2+EnefUZQmE7EVvb/xQIhAIA79SsRcxqT7j2ta9hpuDbq/wI8nAL9\nCtwMSGp4ysiG\n-----END CERTIFICATE-----")
	NODE_PUBLIC_KEY, _     = ca.RetrievePubKey(cert)
	NODE_PRIVITE_KEY, _ = ca.RetrievePrivKey(cert.Algorithm, "-----BEGIN EC PARAMETERS-----\nBggqgRzPVQGCLQ==\n-----END EC PARAMETERS-----\n-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEICZfSOXLlosNMOP9l1yeNO1Xi9MFvPCxRBC0EHZh35IDoAoGCCqBHM9V\nAYItoUQDQgAEWaZxX+18M2C1XGG1QiZYv5KfDuQTUAIbYwIAm+dwE4XT3WX1kBUQ\nEagwK42qDcAtZvSNmYUc6878qcH1BhFwOg==\n-----END EC PRIVATE KEY-----")

	NODE_KEY, _ = ledger_model.NewBlockchainKeypair(NODE_PUBLIC_KEY, NODE_PRIVITE_KEY)
)
