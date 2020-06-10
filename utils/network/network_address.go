package network

import "framework-go/utils/bytes"

/*
 * Author: imuge
 * Date: 2020/6/9 下午5:27
 */

type Address struct {
	Host   string
	Port   int32
	Secure bool
}

func NewAddress(host string, port int32, secure bool) *Address {
	return &Address{
		Host:   host,
		Port:   port,
		Secure: secure,
	}
}

func (a *Address) ToBytes() []byte {
	bf := make([]byte, 0)
	if a.Secure {
		bf = append(bf, 1)
	} else {
		bf = append(bf, 0)
	}
	bf = append(bf, bytes.Int32ToBytes(a.Port)...)
	bf = append(bf, bytes.StringToBytes(a.Host)...)

	return bf
}

func FromBytes(bs []byte) *Address {
	address := &Address{
		Secure: bytes.ToBoolean(bs[0]),
		Port: bytes.ToInt32(bs[1:5]),
		Host: bytes.ToString(bs[5:]),
	}

	return address
}
