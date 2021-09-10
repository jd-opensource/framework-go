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
		"177gk27mzG6iW3CUizaGr8TBQrU1uzhy2GUb6ha2gzKVgKMmgKzcDtcJRAs6BnQQemdT8nu",
		base58.MustDecode("8EjkXVSTxMFjCvNNsTo8RBMDEVQmk7gYkW4SCDuvdsBG"))
	NODE_PUBLIC_KEY = crypto.DecodePubKey("7VeRLU7adGErErcMx78qwqQWcwXk11DSsxgBtdtbuWcNBsY8")

	// 区块链上已存在的有操作权限的用户证书和私钥信息
	// SM2
	//cert, _             = ca.RetrieveCertificate("-----BEGIN CERTIFICATE-----\nMIIC5DCCAomgAwIBAgIUJitW9Yyrd6Ogi5KcautPbJtMfmowCgYIKoEcz1UBg3Uw\nXzELMAkGA1UEBhMCVVMxFzAVBgNVBAgTDk5vcnRoIENhcm9saW5hMRQwEgYDVQQK\nEwtIeXBlcmxlZGdlcjEPMA0GA1UECxMGRmFicmljMRAwDgYDVQQDEwdjYS1vcmcx\nMB4XDTIxMDkwMTA2NDEwMFoXDTMxMDgzMDA2NDYwMFowcjEVMBMGA1UEChMMamRj\naGFpbmRlbW8yMTAwDQYDVQQLEwZjbGllbnQwCwYDVQQLEwRvcmcxMBIGA1UECxML\nZGVwYXJ0bWVudDExJzAlBgNVBAMMHkFkbWluQHRlc3RyYWZ0MDkwMWpkY2hhaW5k\nZW1vMjBZMBMGByqGSM49AgEGCCqBHM9VAYItA0IABID4foj32EAKKZqOWkxVEaiQ\nsoqZOawZJ42rPqJkJhPojjRUFMNePpFjixMUiJfjjurVo+qnL3ybGAA1vI6jw/Sj\nggEOMIIBCjAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQU\ngkai9Lsnuyt5Vm7FEthIriUDOI8wHwYDVR0jBBgwFoAUv1GYoY0PhbTyZqNu3yYh\njbKDYx0wJgYDVR0RBB8wHYIbZmFicmljdG9vbC02ZjU3ZGM5YjY4LWNtY2s3MIGB\nBggqAwQFBgcIAQR1eyJhdHRycyI6eyJoZi5BZmZpbGlhdGlvbiI6Im9yZzEuZGVw\nYXJ0bWVudDEiLCJoZi5FbnJvbGxtZW50SUQiOiJBZG1pbkB0ZXN0cmFmdDA5MDFq\nZGNoYWluZGVtbzIiLCJoZi5UeXBlIjoiY2xpZW50In19MAoGCCqBHM9VAYN1A0kA\nMEYCIQCtBche+o3gcjHp6ci2HiH7zY0TMpf9EgaH15OqkFJX7AIhAOlGGCKJ51Iw\nmEOgaftJbgN1pohVxO6K2DybmKc8pu0R\n-----END CERTIFICATE-----")
	//NODE_PUBLIC_KEY     = ca.RetrievePubKey(cert)
	//NODE_PRIVITE_KEY, _ = ca.RetrievePrivKey(cert.Algorithm, "-----BEGIN PRIVATE KEY-----\nMIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgfN3XcTsTbwl9Q6/Y\nsRX6kUxJZDXTtArudne/FdjJuzmhRANCAASA+H6I99hACimajlpMVRGokLKKmTms\nGSeNqz6iZCYT6I40VBTDXj6RY4sTFIiX447q1aPqpy98mxgANbyOo8P0\n-----END PRIVATE KEY-----")

	NODE_KEY = ledger_model.NewBlockchainKeypair(NODE_PUBLIC_KEY, NODE_PRIVITE_KEY)
)
