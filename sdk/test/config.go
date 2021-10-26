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
		"177gjvG9ZKkGwdzKfrK2YguhS2Wthu6EdbVSF9WqCxfmqdJuVz82BfFwt53XaGYEbp8RqRW",
		base58.MustDecode("8EjkXVSTxMFjCvNNsTo8RBMDEVQmk7gYkW4SCDuvdsBG"))
	NODE_PUBLIC_KEY = crypto.DecodePubKey("3snPdw7i7PhgdrXp9UxgTMr5PAYFxrEWdRdAdn9hsBA4pvp1iVYXM6")

	// 区块链上已存在的有操作权限的用户证书和私钥信息
	//cert, _             = ca.RetrieveCertificate("-----BEGIN CERTIFICATE-----\nMIIB4DCCAYagAwIBAgIEbHiebzAKBggqhkjOPQQDAjBwMQwwCgYDVQQKDANKRFQxDTALBgNVBAsM\nBFJPT1QxCzAJBgNVBAYTAkNOMQswCQYDVQQIDAJCSjELMAkGA1UEBwwCQkoxDTALBgNVBAMMBHJv\nb3QxGzAZBgkqhkiG9w0BCQEWDGltdWdlQGpkLmNvbTAeFw0yMTEwMTUxMTQxMjVaFw0zMTEwMTMx\nMTQxMjVaMHExDDAKBgNVBAoMA0pEVDENMAsGA1UECwwEUEVFUjELMAkGA1UEBhMCQ04xCzAJBgNV\nBAgMAkJKMQswCQYDVQQHDAJCSjEOMAwGA1UEAwwFcGVlcjAxGzAZBgkqhkiG9w0BCQEWDGltdWdl\nQGpkLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABMUES60oHIZUf1RsO0DyGV/Or3dYw+qU\nHszonKIfJxRuLJtfMxE80r2114gjWCcVu8tRCQN/3Gnz9GHIUCWWNGKjDTALMAkGA1UdEwQCMAAw\nCgYIKoZIzj0EAwIDSAAwRQIhAJAKY6xAUwbH0aYmHd/o640n9Rw6O4Mfg55jsU+SzX8GAiBn7Cb8\n3bo4qLjI6/LWRTj2zMMqRJDo3Pakf+WyyoR4Yg==\n-----END CERTIFICATE-----")
	//NODE_PUBLIC_KEY     = ca.RetrievePubKey(cert)
	//NODE_PRIVITE_KEY, _ = ca.RetrievePrivKey(cert.Algorithm, "-----BEGIN EC PARAMETERS-----\nBggqhkjOPQMBBw==\n-----END EC PARAMETERS-----\n-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIMT+m6VpaAmPjIzTWKnL1zjm9J3pPOdWm6B3uoKosJuJoAoGCCqGSM49\nAwEHoUQDQgAExQRLrSgchlR/VGw7QPIZX86vd1jD6pQezOicoh8nFG4sm18zETzS\nvbXXiCNYJxW7y1EJA3/cafP0YchQJZY0Yg==\n-----END EC PRIVATE KEY-----")

	NODE_KEY = ledger_model.NewBlockchainKeypair(NODE_PUBLIC_KEY, NODE_PRIVITE_KEY)
)
