package crypto

import (
	"framework-go/crypto/framework"
	"framework-go/utils/aes"
	"framework-go/utils/base58"
	"framework-go/utils/sha"
)

var (
	PubKeyFileMagicNum  = []byte{255, 112, 117, 98}
	PrivKeyFileMagicNum = []byte{0, 112, 114, 118}
)

func EncodePubKey(pubKey framework.PubKey) string {
	return base58.Encode(append(PubKeyFileMagicNum, pubKey.ToBytes()...))
}

func DecodePubKey(base58PubKey string) framework.PubKey {
	key, err := base58.Decode(base58PubKey)
	if err != nil {
		panic(err)
	}

	return framework.ParsePubKey(key[len(PubKeyFileMagicNum):])
}

func EncodePrivKey(privKey framework.PrivKey, pwdBytes []byte) string {
	return base58.Encode(append(PrivKeyFileMagicNum, EncryptPrivKey(privKey, pwdBytes)...))
}

func EncodePrivKeyWithRawPwd(privKey framework.PrivKey, pwd string) string {
	return EncodePrivKey(privKey, sha.Sha256([]byte(pwd)))
}

func DecodePrivKey(base58PrivKey string, pwdBytes []byte) framework.PrivKey {
	key, err := base58.Decode(base58PrivKey)
	if err != nil {
		panic(err)
	}
	return framework.ParsePrivKey(DecryptPrivKey(key[len(PrivKeyFileMagicNum):], pwdBytes))
}

func DecodePrivKeyWithRawPwd(base58PrivKey string, pwd string) framework.PrivKey {
	return DecodePrivKey(base58PrivKey, sha.Sha256([]byte(pwd)))
}

func EncryptPrivKey(privKey framework.PrivKey, pwdBytes []byte) []byte {
	cipherText, err := aes.Encrypt(sha.Sha128(pwdBytes), privKey.ToBytes())
	if err != nil {
		panic(err)
	}
	return cipherText
}

func DecryptPrivKey(encPrivKey []byte, pwdBytes []byte) []byte {
	originText, err := aes.Decrypt(sha.Sha128(pwdBytes), encPrivKey)
	if err != nil {
		panic(err)
	}
	return originText
}
