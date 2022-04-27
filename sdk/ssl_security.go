package sdk

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/blockchain-jd-com/framework-go/gmsm/gmtls"
	gmx509 "github.com/blockchain-jd-com/framework-go/gmsm/x509"
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

// 国密TLS 暂时仅支持忽略和单向认证
type GMSSLSecurity struct {
	RootCerts *gmx509.CertPool
	SigCert   *gmtls.Certificate
	EncCert   *gmtls.Certificate
}

func NewGMSSLSecurity(rootCertFile string) (*GMSSLSecurity, error) {
	rootPemData, err := ioutil.ReadFile(rootCertFile)
	if err != nil {
		return nil, err
	}
	rootCerts := gmx509.NewCertPool()
	rootCerts.AppendCertsFromPEM(rootPemData)

	return &GMSSLSecurity{
		RootCerts: rootCerts,
	}, nil
}

func NewTwoWayGMSSLSecurity(rootCertFile, signCertPath, signKeyFile, encCertFile, encKeyFile string) (*GMSSLSecurity, error) {
	rootPemData, err := ioutil.ReadFile(rootCertFile)
	if err != nil {
		return nil, err
	}
	rootCerts := gmx509.NewCertPool()
	rootCerts.AppendCertsFromPEM(rootPemData)

	sigCert, err := gmtls.LoadX509KeyPair(signCertPath, signKeyFile)
	if err != nil {
		return nil, err
	}
	encCert, err := gmtls.LoadX509KeyPair(encCertFile, encKeyFile)
	if err != nil {
		return nil, err

	}

	return &GMSSLSecurity{
		RootCerts: rootCerts,
		SigCert:   &sigCert,
		EncCert:   &encCert,
	}, nil
}
