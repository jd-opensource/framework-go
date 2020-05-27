package test

import (
	"framework-go/crypto"
	"framework-go/ledger_model"
	"framework-go/utils/base58"
)

/*
 * Author: imuge
 * Date: 2020/5/29 下午3:21
 */

var (
	GATEWAY_HOST = "localhost"
	GATEWAY_PORT = 8081
	SECURE       = false

	NODE_PRIVITE_KEY = crypto.DecodePrivKey("177gjtFXNyXXGZFfHii2kJn4HEfk6KKtGHr9iEAYZUZpCdKrNBZjsjzGbzXVoB2Sohd4x7A", base58.MustDecode("8EjkXVSTxMFjCvNNsTo8RBMDEVQmk7gYkW4SCDuvdsBG"))
	NODE_PUBLIC_KEY  = crypto.DecodePubKey("3snPdw7i7Pm3vaGUsyBe5Vkce8jMFjcu4dJFSQY3A81P6VJh7iC1d8")
	NODE_KEY         = ledger_model.NewBlockchainKeypair(NODE_PUBLIC_KEY, NODE_PRIVITE_KEY)
)
