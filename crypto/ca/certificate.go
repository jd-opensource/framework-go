package ca

import (
	"crypto/x509"
	x5092 "github.com/tjfoc/gmsm/x509"
)

type Certificate struct {
	Algorithm   string
	ClassicCert *x509.Certificate
	SMCert      *x5092.Certificate
}
