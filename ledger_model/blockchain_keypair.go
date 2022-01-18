package ledger_model

import (
	"errors"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
)

/*
 * Author: imuge
 * Date: 2020/5/27 下午4:39
 */

type BlockchainKeypair struct {
	*framework.AsymmetricKeypair
	id *BlockchainIdentity
}

func NewBlockchainKeypair(pubKey *framework.PubKey, privKey *framework.PrivKey) (*BlockchainKeypair, error) {
	if pubKey.GetAlgorithm() != privKey.GetAlgorithm() {
		return nil, errors.New("The PublicKey's algorithm is different from the PrivateKey's!")
	}
	keypair, err := framework.NewAsymmetricKeypair(pubKey, privKey)
	if err != nil {
		return nil, err
	}
	return &BlockchainKeypair{
		keypair,
		NewBlockchainIdentity(pubKey),
	}, nil
}

func (pair BlockchainKeypair) GetAddress() []byte {
	return pair.id.Address
}

func (pair BlockchainKeypair) GetIdentity() *BlockchainIdentity {
	return pair.id
}
