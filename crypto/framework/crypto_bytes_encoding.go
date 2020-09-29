package framework

import "github.com/blockchain-jd-com/framework-go/utils/bytes"

/**
 * @Author: imuge
 * @Date: 2020/4/29 3:45 下午
 */

func EncodeBytes(algorithm int16, rawCryptoBytes []byte) []byte {
	return bytes.Concat(bytes.Int16ToBytes(algorithm), rawCryptoBytes)
}

func DecodeAlgorithm(cryptoBytes []byte) int16 {
	return bytes.ToInt16(cryptoBytes[:2])
}

func EncodeKeyBytes(rawKeyBytes []byte, keyType CryptoKeyType) []byte {
	return bytes.Concat([]byte{keyType.Code}, rawKeyBytes)
}

func DecodeKeyType(cryptoBytes bytes.Slice) CryptoKeyType {
	return GetCryptoKeyType(cryptoBytes.GetByte(0))
}
