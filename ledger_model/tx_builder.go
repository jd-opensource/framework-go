package ledger_model

import (
	binary_proto "framework-go/binary-proto"
	"framework-go/crypto"
	"framework-go/crypto/framework"
	"time"
)

/*
 * Author: imuge
 * Date: 2020/5/29 上午9:17
 */

var _ TransactionBuilder = (*TxBuilder)(nil)

type TxBuilder struct {
	ledgerHash framework.HashDigest
	opFactory  BlockchainOperationFactory
}

func NewTxBuilder(ledgerHash framework.HashDigest) *TxBuilder {
	return &TxBuilder{
		ledgerHash: ledgerHash,
		opFactory:  BlockchainOperationFactory{},
	}
}

func (t *TxBuilder) Security() *SecurityOperationBuilder {
	return t.opFactory.Security()
}

func (t *TxBuilder) GetLedgerHash() framework.HashDigest {
	return t.ledgerHash
}

func (t *TxBuilder) Users() *UserRegisterOperationBuilder {
	return t.opFactory.Users()
}

func (t *TxBuilder) DataAccounts() *DataAccountRegisterOperationBuilder {
	return t.opFactory.DataAccounts()
}

func (t *TxBuilder) DataAccount(accountAddress []byte) *DataAccountKVSetOperationBuilder {
	return t.opFactory.DataAccount(accountAddress)
}

func (t *TxBuilder) PrepareRequestNow() TransactionRequestBuilder {
	return NewTxRequestBuilder(t.PrepareContentNow())
}

func (t *TxBuilder) PrepareRequest(time int64) TransactionRequestBuilder {
	return NewTxRequestBuilder(t.PrepareContent(time))
}

func (t *TxBuilder) PrepareContentNow() TransactionContent {
	return t.PrepareContent(time.Now().Unix())
}

func (t *TxBuilder) PrepareContent(time int64) TransactionContent {
	txContent := NewTransactionContent(t.ledgerHash, t.opFactory.operationList, time)
	contentHash := computeTxContentHash(txContent.TransactionContentBody)
	txContent.Hash = contentHash.ToBytes()

	return txContent
}

func computeTxContentHash(txContent TransactionContentBody) framework.HashDigest {
	contentBodyBytes, err := binary_proto.Cdc.Encode(txContent)
	if err != nil {
		panic(err)
	}
	contentHash := crypto.GetHashFunctionByName(DEFAULT_HASH_ALGORITHM).Hash(contentBodyBytes)
	return contentHash
}

func verifyTxContentHash(txContent TransactionContentBody, verifiedHash framework.HashDigest) bool {
	hash := computeTxContentHash(txContent)
	return hash.Equals(verifiedHash)
}
