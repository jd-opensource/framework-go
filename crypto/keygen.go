package crypto

import (
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/utils/aes"
	"github.com/blockchain-jd-com/framework-go/utils/base58"
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
	"github.com/blockchain-jd-com/framework-go/utils/sha"
)

var (
	PubKeyFileMagicNum  = []byte{255, 112, 117, 98}
	PrivKeyFileMagicNum = []byte{0, 112, 114, 118}
)

func EncodePubKey(pubKey *framework.PubKey) string {
	return base58.Encode(pubKey.ToBytes())
}

func DecodePubKey(base58PubKey string) (*framework.PubKey, error) {
	key, err := base58.Decode(base58PubKey)
	if err != nil {
		return nil, err
	}

	// 兼容1.4.0之前版本公钥输出格式
	if bytes.StartsWith(key, PubKeyFileMagicNum) {
		return framework.ParsePubKey(key[len(PubKeyFileMagicNum):])
	} else {
		return framework.ParsePubKey(key)
	}
}

func MustDecodePubKey(base58PubKey string) *framework.PubKey {
	key, err := DecodePubKey(base58PubKey)
	if err != nil {
		panic(err)
	}

	return key
}

func EncodePrivKey(privKey *framework.PrivKey, pwdBytes []byte) (string, error) {
	encryptPrivKey, err := EncryptPrivKey(privKey, pwdBytes)
	if err != nil {
		return "", err
	}
	return base58.Encode(append(PrivKeyFileMagicNum, encryptPrivKey...)), nil
}

func EncodePrivKeyWithRawPwd(privKey *framework.PrivKey, pwd string) (string, error) {
	return EncodePrivKey(privKey, sha.Sha256([]byte(pwd)))
}

func DecodePrivKey(base58PrivKey string, pwdBytes []byte) (*framework.PrivKey, error) {
	key, err := base58.Decode(base58PrivKey)
	if err != nil {
		return nil, err
	}
	privKey, err := DecryptPrivKey(key[len(PrivKeyFileMagicNum):], pwdBytes)
	return framework.ParsePrivKey(privKey)
}

func MustDecodePrivKey(base58PrivKey string, pwdBytes []byte) *framework.PrivKey {
	key, err := DecodePrivKey(base58PrivKey, pwdBytes)
	if err != nil {
		panic(err)
	}

	return key
}

func DecodePrivKeyWithRawPwd(base58PrivKey string, pwd string) (*framework.PrivKey, error) {
	return DecodePrivKey(base58PrivKey, sha.Sha256([]byte(pwd)))
}

func EncryptPrivKey(privKey *framework.PrivKey, pwdBytes []byte) ([]byte, error) {
	cipherText, err := aes.Encrypt(sha.Sha128(pwdBytes), privKey.ToBytes())
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}

func DecryptPrivKey(encPrivKey []byte, pwdBytes []byte) ([]byte, error) {
	originText, err := aes.Decrypt(sha.Sha128(pwdBytes), encPrivKey)
	if err != nil {
		return nil, err
	}
	return originText, nil
}
