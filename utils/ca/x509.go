package ca

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/blockchain-jd-com/framework-go/crypto"
	"github.com/blockchain-jd-com/framework-go/crypto/ca"
	"github.com/blockchain-jd-com/framework-go/crypto/classic"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/crypto/sm"
	x5092 "github.com/blockchain-jd-com/framework-go/gmsm/x509"
	bytes2 "github.com/blockchain-jd-com/framework-go/utils/bytes"
	"io/ioutil"
	"strings"
)

var caSignatureAlgorithms = map[string]framework.CryptoAlgorithm{
	"SM2-SM3":      sm.SM2_ALGORITHM,
	"ED25519":      classic.ED25519_ALGORITHM,
	"SHA256-RSA":   classic.RSA_ALGORITHM,
	"ECDSA-SHA256": classic.ECDSA_ALGORITHM,
}

// 解析X509证书，支持RSA 2048,ECDSA P-256,ED25519,SM2 SM3WithSM2
func RetrieveCertificate(cert string) (*ca.Certificate, error) {
	var algorithm string
	var smCert *x5092.Certificate
	classicCert, err := resolveClassicCertificate(cert)
	if err != nil {
		smCert, err = resolveSMCertificate(cert)
		if err != nil {
			return nil, errors.New("can not parse certificate")
		}
		algorithm = smCert.SignatureAlgorithm.String()
	} else {
		algorithm = classicCert.SignatureAlgorithm.String()
	}
	cryptoAlgorithm, ok := caSignatureAlgorithms[strings.ToUpper(algorithm)]
	if !ok {
		return nil, errors.New("can not parse certificate with algorithm: " + algorithm)
	}

	return &ca.Certificate{
		Algorithm:   cryptoAlgorithm.Name,
		ClassicCert: classicCert,
		SMCert:      smCert,
	}, nil
}

func RetrieveCertificateFile(certFile string) (*ca.Certificate, error) {
	bytes, err := ioutil.ReadFile(certFile)
	if err != nil {
		return nil, err
	}
	cert := string(bytes)
	var algorithm string
	var smCert *x5092.Certificate
	classicCert, err := resolveClassicCertificate(cert)
	if err != nil {
		smCert, err = resolveSMCertificate(cert)
		if err != nil {
			return nil, errors.New("can not parse certificate")
		}
		algorithm = smCert.SignatureAlgorithm.String()
	} else {
		algorithm = classicCert.SignatureAlgorithm.String()
	}
	cryptoAlgorithm, ok := caSignatureAlgorithms[strings.ToUpper(algorithm)]
	if !ok {
		return nil, errors.New("can not parse certificate with algorithm: " + algorithm)
	}

	return &ca.Certificate{
		Algorithm:   cryptoAlgorithm.Name,
		ClassicCert: classicCert,
		SMCert:      smCert,
	}, nil
}

func resolveClassicCertificate(cert string) (*x509.Certificate, error) {
	pemCert := []byte(cert)
	certPool := x509.NewCertPool()
	ok := certPool.AppendCertsFromPEM(pemCert)
	if !ok {
		return nil, errors.New("can not parse certificate")
	}
	block, pemCert := pem.Decode(pemCert)
	if block == nil {
		return nil, errors.New("can not parse certificate")
	}
	if block.Type != "CERTIFICATE" || len(block.Headers) != 0 {
		return nil, errors.New("can not parse certificate")
	}
	return x509.ParseCertificate(block.Bytes)
}

func resolveSMCertificate(cert string) (*x5092.Certificate, error) {
	return x5092.ReadCertificateFromPem(bytes2.StringToBytes(cert))
}

// 解析证书公钥信息
func RetrievePubKey(certificate *ca.Certificate) (*framework.PubKey, error) {
	return crypto.GetSignatureFunctionByName(certificate.Algorithm).RetrievePubKeyFromCA(certificate)
}

// 解析X509私钥文件
func RetrievePrivKey(algorithm string, privkey string) (*framework.PrivKey, error) {
	return crypto.GetSignatureFunctionByName(algorithm).RetrievePrivKey(privkey)
}

func RetrievePrivKeyFile(algorithm string, privkeyFile string) (*framework.PrivKey, error) {
	bytes, err := ioutil.ReadFile(privkeyFile)
	if err != nil {
		return nil, err
	}
	return crypto.GetSignatureFunctionByName(algorithm).RetrievePrivKey(string(bytes))
}

// 解析X509加密私钥文件
func RetrieveEncrypedPrivKey(algorithm string, privkey string, pwd []byte) (*framework.PrivKey, error) {
	return crypto.GetSignatureFunctionByName(algorithm).RetrieveEncrypedPrivKey(privkey, pwd)
}
