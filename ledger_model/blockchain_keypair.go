package ledger_model

import (
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
)

/*
 * Author: imuge
 * Date: 2020/5/27 下午4:39
 */

type BlockchainKeypair struct {
	framework.AsymmetricKeypair
	id BlockchainIdentity
}

func NewBlockchainKeypair(pubKey framework.PubKey, privKey framework.PrivKey) BlockchainKeypair {
	if pubKey.GetAlgorithm() != privKey.GetAlgorithm() {
		panic("The PublicKey's algorithm is different from the PrivateKey's!")
	}
	return BlockchainKeypair{
		framework.NewAsymmetricKeypair(pubKey, privKey),
		NewBlockchainIdentity(pubKey),
	}
}

func (pair BlockchainKeypair) GetAddress() []byte {
	return pair.id.Address
}

func (pair BlockchainKeypair) GetIdentity() BlockchainIdentity {
	return pair.id
}
