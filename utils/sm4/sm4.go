package sm4

/**
 * 基于github.com/ZZMarquis/gm/sm4扩展实现CBC/PKCS7Padding
 */

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"github.com/ZZMarquis/gm/sm4"
)

var (
	KEY_SIZE = 16
	IV_SIZE  = 16
)

func GenerateSymmetricKey() []byte {
	bytes := make([]byte, KEY_SIZE)
	rand.Read(bytes)
	return bytes
}

func Encrypt(key, data []byte) []byte {
	iv := make([]byte, IV_SIZE)
	rand.Reader.Read(iv[:])
	encrypt, err := encrypt(key, iv, data, true)
	if err != nil {
		panic(err)
	}

	return append(iv, encrypt...)
}

func Decrypt(key, data []byte) []byte {
	iv := data[:IV_SIZE]
	decrypt, err := decrypt(key, iv, data[IV_SIZE:], true)
	if err != nil {
		panic(err)
	}

	return decrypt
}

func encrypt(key, iv, plantText []byte, paddingStatus bool) ([]byte, error) {
	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if paddingStatus {
		plantText = pKCS7Padding(plantText, block.BlockSize())
	}

	blockModel := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plantText))
	blockModel.CryptBlocks(ciphertext, plantText)
	return ciphertext, nil
}

func decrypt(key, iv, ciphertext []byte, paddingStatus bool) ([]byte, error) {
	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < block.BlockSize() {
		return nil, errors.New("crypto/cipher: ciphertext too short")
	}

	if len(ciphertext)%block.BlockSize() != 0 {
		return nil, errors.New("crypto/cipher: ciphertext is not a multiple of the block size")
	}

	blockModel := cipher.NewCBCDecrypter(block, iv)
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	if paddingStatus {
		plantText, err = pKCS7UnPadding(plantText)
	}
	if err != nil {
		return nil, err
	}
	return plantText, nil
}

func pKCS7UnPadding(plantText []byte) ([]byte, error) {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	if length-unpadding < 0 || length-unpadding > length {
		return nil, errors.New("un padding error")
	}
	return plantText[:(length - unpadding)], nil
}

func pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
