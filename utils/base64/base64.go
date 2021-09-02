package base64

import (
	"encoding/base64"
)

func Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func MustDecode(input string) []byte {
	bs, err := Decode(input)
	if err != nil {
		panic(err)
	}

	return bs
}

func Decode(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}
