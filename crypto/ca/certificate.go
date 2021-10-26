package ca

import (
	"crypto/x509"
	"encoding/pem"
	x5092 "github.com/tjfoc/gmsm/x509"
)

type Certificate struct {
	Algorithm   string
	ClassicCert *x509.Certificate
	SMCert      *x5092.Certificate
}

func (cert *Certificate) ToPEMString() string {
	if cert.ClassicCert != nil {
		return string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert.ClassicCert.Raw}))
	} else {
		return string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert.SMCert.Raw}))
	}
}
