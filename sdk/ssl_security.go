package sdk

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

type SSLSecurity struct {
	RootCerts   *x509.CertPool
	ClientCerts []tls.Certificate
}

func NewSSLSecurity(rootCertFile string, clientCertFile, clientKeyFile string) (*SSLSecurity, error) {
	rootCerts := x509.NewCertPool()
	rootPemData, err := ioutil.ReadFile(rootCertFile)
	if err != nil {
		return nil, err
	}
	rootCerts.AppendCertsFromPEM(rootPemData)
	clientCert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		return nil, err
	}
	return &SSLSecurity{
		RootCerts:   rootCerts,
		ClientCerts: []tls.Certificate{clientCert},
	}, nil
}
