package sm4

/**
 * 基于github.com/ZZMarquis/gm/sm4扩展实现CBC/PKCS7Padding
 */

import (
	"bytes"
	"crypto/cipher"
	"errors"
	"github.com/ZZMarquis/gm/sm4"
)

func Sm4Enc(key, iv, plantText []byte, paddingStatus bool) ([]byte, error) {
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

func Sm4Dec(key, iv, ciphertext []byte, paddingStatus bool) ([]byte, error) {
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
