package ca

import (
	"fmt"
	"github.com/blockchain-jd-com/framework-go/crypto"
	"github.com/blockchain-jd-com/framework-go/crypto/classic"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/crypto/sm"
	"github.com/blockchain-jd-com/framework-go/utils/base58"
	"github.com/stretchr/testify/require"
	"testing"
)

// cert privatekey pwd privkey pwdOfprivkey pubkey
var certs = [][7]string{
	{classic.ED25519_ALGORITHM.Name,
		"-----BEGIN CERTIFICATE-----\nMIIBoDCCAVKgAwIBAgIEMthl+TAFBgMrZXAwcDEMMAoGA1UECgwDSkRUMQ0wCwYDVQQLDARST09U\nMQswCQYDVQQGEwJDTjELMAkGA1UECAwCQkoxCzAJBgNVBAcMAkJKMQ0wCwYDVQQDDARyb290MRsw\nGQYJKoZIhvcNAQkBFgxpbXVnZUBqZC5jb20wHhcNMjEwOTE0MDM0OTUxWhcNMzEwOTEyMDM0OTUx\nWjBxMQwwCgYDVQQKDANKRFQxDTALBgNVBAsMBFVTRVIxCzAJBgNVBAYTAkNOMQswCQYDVQQIDAJC\nSjELMAkGA1UEBwwCQkoxDjAMBgNVBAMMBXVzZXIxMRswGQYJKoZIhvcNAQkBFgxpbXVnZUBqZC5j\nb20wKjAFBgMrZXADIQCy397VCIes8I2VT7JEOePnHO8+txh+R804J4o9wKQw1qMNMAswCQYDVR0T\nBAIwADAFBgMrZXADQQDqw3Y1T7BtKnD87Mblu6hjqCvK2Wj2k/RPHg1hZCJxR7H1J26a5gFaM07z\n6Geq0wBE3nYjjDqO42tsbPsBSCkO\n-----END CERTIFICATE-----",
		"-----BEGIN PRIVATE KEY-----\nMC4CAQAwBQYDK2VwBCIEIHAWoSerB+AQ0nBqMyBMyq33W4k2/NBJdzzWWOyuthER\n-----END PRIVATE KEY-----",
		"",
		"177gjvd2GpfwkXRCUs6RdSs3fe3aST7QccvEFXAKncVFP3NWhJuAvo6RfY7UXgE4kt3LTsZ",
		"8EjkXVSTxMFjCvNNsTo8RBMDEVQmk7gYkW4SCDuvdsBG",
		"7VeRK5XdXATZKZDN5XsHWvJaLcWFNLmxnT3L41H8mYHJFGWh"},
	{sm.SM2_ALGORITHM.Name,
		"-----BEGIN CERTIFICATE-----\nMIIB3zCCAYagAwIBAgIELysKCTAKBggqgRzPVQGDdTBwMQwwCgYDVQQKDANKRFQxDTALBgNVBAsM\nBFJPT1QxCzAJBgNVBAYTAkNOMQswCQYDVQQIDAJCSjELMAkGA1UEBwwCQkoxDTALBgNVBAMMBHJv\nb3QxGzAZBgkqhkiG9w0BCQEWDGltdWdlQGpkLmNvbTAeFw0yMTA5MTQwMzQ1MjhaFw0zMTA5MTIw\nMzQ1MjhaMHExDDAKBgNVBAoMA0pEVDENMAsGA1UECwwEVVNFUjELMAkGA1UEBhMCQ04xCzAJBgNV\nBAgMAkJKMQswCQYDVQQHDAJCSjEOMAwGA1UEAwwFdXNlcjExGzAZBgkqhkiG9w0BCQEWDGltdWdl\nQGpkLmNvbTBZMBMGByqGSM49AgEGCCqBHM9VAYItA0IABJkY/Y3Ug6u4HTD/pvERvERaWfxtJe7+\nJIffDOVlIhxJYDXiKqE9WQR1Q5LbnIcw1Ou1/gz7TcCPVu3vEgPlJUWjDTALMAkGA1UdEwQCMAAw\nCgYIKoEcz1UBg3UDRwAwRAIgIhawFl2aKJnHUnBiDtR4o13iY1gEtLOM2jTZqJ93aiUCIGYI6EgZ\nPNGkV6VjrOpAyt5DpSTZ/ILCKGJkNNdcwJab\n-----END CERTIFICATE-----",
		"-----BEGIN EC PARAMETERS-----\nBggqgRzPVQGCLQ==\n-----END EC PARAMETERS-----\n-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIFCnh1/MuJaXSLmRsH52xzw4KdCWaIKhYb9uzsbkkmWuoAoGCCqBHM9V\nAYItoUQDQgAEmRj9jdSDq7gdMP+m8RG8RFpZ/G0l7v4kh98M5WUiHElgNeIqoT1Z\nBHVDktuchzDU67X+DPtNwI9W7e8SA+UlRQ==\n-----END EC PRIVATE KEY-----",
		"",
		"177gjtiZQXwKMZiQipeKiVhuuCiyq4dzdrTiDh9LA1JGtGozFQwToUBRu6PMwvp9enNojDR",
		"8EjkXVSTxMFjCvNNsTo8RBMDEVQmk7gYkW4SCDuvdsBG",
		"SFZ6LkFNrRcB24VBCcrDerTeKEC66wXEUwGKT3ooWtvhZTsh3y4wMFDRsLkEeN6sfVMqHcS8S6DGiumpxrjFx6GATeSip"},
	{classic.RSA_ALGORITHM.Name,
		"-----BEGIN CERTIFICATE-----\nMIIDbDCCAlSgAwIBAgIEQVFAnzANBgkqhkiG9w0BAQsFADBwMQwwCgYDVQQKDANKRFQxDTALBgNV\nBAsMBFJPT1QxCzAJBgNVBAYTAkNOMQswCQYDVQQIDAJCSjELMAkGA1UEBwwCQkoxDTALBgNVBAMM\nBHJvb3QxGzAZBgkqhkiG9w0BCQEWDGltdWdlQGpkLmNvbTAeFw0yMTA5MTQwMzUxMDNaFw0zMTA5\nMTIwMzUxMDNaMHExDDAKBgNVBAoMA0pEVDENMAsGA1UECwwEVVNFUjELMAkGA1UEBhMCQ04xCzAJ\nBgNVBAgMAkJKMQswCQYDVQQHDAJCSjEOMAwGA1UEAwwFdXNlcjExGzAZBgkqhkiG9w0BCQEWDGlt\ndWdlQGpkLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAM77c2p0YuC2lV7ZTTsu\nwm7IEdoJHIg7YcoszVcgtbwdOMWXJi9ysiRSHCKHc/vCU24PVFCtJiHp+VG+95S0TwTB4DG8mGGV\n7i8NjSxZrcQHB5j3cj8csPjn10U9hQvHDHHuNQ8gIb0gpVP4tRJ3LeXjXMXs/Z7jzjxKPtEPqhho\n+AUTE9OON9fpPr1df8EcF+H1TT8yQHf6xv71eLJJFdsYsXw7MQ6TvufSEVpJxxTD0RBJwrakfmkI\n8gyydvO0U8BClVEmb+7LCVYYdNxPMj2BgnZpU+ORh+d9Y9H0BFYC6Rvd8PzNS9otPmv7uD/7X9zK\nQA7+3/yR3lAvpYwHxKkCAwEAAaMNMAswCQYDVR0TBAIwADANBgkqhkiG9w0BAQsFAAOCAQEAN+f0\nW8AS8kvK0dDH7Z/XM/+lhcI+BtIG3jTT/5iBvLZRP7fh0U53fdKA5pX7jVvPV+ckOVjWYkljXVAp\niqhVQtJVSsW4bCc0qxnMehQf30bPspY9Hn2gu5aniTCheU72t0i7rW80Ph9X0Tfy5Mow+3uTJ/Pg\nIVuD6cVOR0NiRIe9TrlV1h5O/vXsKsCozneoXM8p6Vvzwv/TE0B7mbvgiTrUz3kEUkeIIJByd6NA\nr70oYLrdqmy0EklKX8G3kFzGTHwSDNo3l1iwqRhd7Npv+hlWASR6tutWPbep4QdMwZ9Vgv3NMr7o\nXiKJR6tdxj250KZMtSixXae0yFI73lyJeA==\n-----END CERTIFICATE-----",
		"-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAzvtzanRi4LaVXtlNOy7CbsgR2gkciDthyizNVyC1vB04xZcm\nL3KyJFIcIodz+8JTbg9UUK0mIen5Ub73lLRPBMHgMbyYYZXuLw2NLFmtxAcHmPdy\nPxyw+OfXRT2FC8cMce41DyAhvSClU/i1Enct5eNcxez9nuPOPEo+0Q+qGGj4BRMT\n04431+k+vV1/wRwX4fVNPzJAd/rG/vV4skkV2xixfDsxDpO+59IRWknHFMPREEnC\ntqR+aQjyDLJ287RTwEKVUSZv7ssJVhh03E8yPYGCdmlT45GH531j0fQEVgLpG93w\n/M1L2i0+a/u4P/tf3MpADv7f/JHeUC+ljAfEqQIDAQABAoIBABOgVgtwMzE4O7kT\ny6Lcf+hHeIkSfk4DeQ+ijusyxh8q+cdwQFku1YOwweU0r3bd/TfW8UZGUAxwMxY6\nw7j26sIHOW4atXivYeq8eYUmkPdt0tBNsMIZdnh4apOwtGBvaxeUhrVDKGyjRboA\nJiwbK+yvOFlP0j64pPiKvu+A8XHBIQTXtHRmIIEZBHVlzCHmu8lxuhWImWZdMuH3\nDTGHfC7gqy2dq0+PUeEKIv0Abhorg+Jzsw8p4od1WA6t3oxmyGOOvgx6AZDrYCsj\n25fdJ/zIl2jw3Plb68SWv4V4taKD09UHK0iMBaOWjMDpnG+Wyoyn2xmmcTDqpTqZ\n2o+kgF8CgYEA8Wnlm/yqnavPbDEwGr/yX3ifUbAqY/batwMzhgPG+/ynmdhkIXQu\ntzR7PUJ2EpLNjxvsX2jeQPi9oa2Xe2jA13/CNNNMXI1fKRL6ostJm1zPSFpYBYDh\n4eqAMEyRJDPar6/l0kveO+F771Deknjwd/rDTtFI5CF0r4TP9L0APRcCgYEA23z6\n//cMAJlpZrbP6bUAowLsfALtGGqO+peaJ68xZASUCCrtAB7m93rArb4m+zPzHcYp\nKgXz4W1ZjwMbYH46jUF20s7MYn6+E+AITphK9tpYQbeSoDov7fGU2EoCmhFDtTcq\nPvE6JUG0Ga91qXRFmWqRmW3L4ayDKf04iC6apD8CgYEAtqjsfTkRExmzaOZSwnqn\ndcs7qMBFYrudw0mdy3HCNll1qrcLFDDnQ+FmufQ2iFkhRX3YPFyJhdlvCgzhiBO7\njZJyLCwQJBsnfFmK4HA2MmJnyBPrc8aPorMe6OyWCTFe7v2FQ7f53479ihbDQUpW\nkEFhU5qQr2QM+NzhyAjVTGMCgYEAplF9bYrJqIaXjQLIZ+MFeYDUrGAXQ6IzeAZ7\nBMlHlu+1ML8+WhIQmMWGzeFCbqX9+rjXJoXeORsAe1MyYpskSTerD7EuxRAffrYL\n9WqHm6j2qc2uKQYOnbKrRH5InHCqqt4DgDCRC/xOugvwEBkQSGGttOKzVO3Bcob0\nWJVgD0cCgYAsL4V1ETQ2YAt/87lfJK/70d1LbaZU1ALGN7/JMJ7tEB4k6wGwN5At\n9dMxW6P4XPqWRSELR9wBLLbBFRMeZHrI8mkXfVVKUS1t7lQXdcfC32XH/oCabas7\nGBF9xRdBXiNoJCIOO31q7ycvnYn3HEjxrZLw1RczvFHlb2ViudhWOg==\n-----END RSA PRIVATE KEY-----",
		"",
		"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"},
	{classic.ECDSA_ALGORITHM.Name,
		"-----BEGIN CERTIFICATE-----\nMIIB4TCCAYagAwIBAgIEdWn0mzAKBggqhkjOPQQDAjBwMQwwCgYDVQQKDANKRFQxDTALBgNVBAsM\nBFJPT1QxCzAJBgNVBAYTAkNOMQswCQYDVQQIDAJCSjELMAkGA1UEBwwCQkoxDTALBgNVBAMMBHJv\nb3QxGzAZBgkqhkiG9w0BCQEWDGltdWdlQGpkLmNvbTAeFw0yMTA5MTQwNTMwMjVaFw0zMTA5MTIw\nNTMwMjVaMHExDDAKBgNVBAoMA0pEVDENMAsGA1UECwwEVVNFUjELMAkGA1UEBhMCQ04xCzAJBgNV\nBAgMAkJKMQswCQYDVQQHDAJCSjEOMAwGA1UEAwwFdXNlcjExGzAZBgkqhkiG9w0BCQEWDGltdWdl\nQGpkLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABOtRinwy/KLIZ/o9abKK+mrPm2Uh272d\nO1XwVduP9Re+nOAshSYvtvkcsLo30Xd9a9FgJKrtGEtNmcUr287pBzOjDTALMAkGA1UdEwQCMAAw\nCgYIKoZIzj0EAwIDSQAwRgIhAJwY6pc4eCWYuEmaTz/ogLk4EUkn/D5zmmVuXuyCWxzyAiEA2/yr\nlJcBqIcazi7vES0mvTej/o7Dox2ys+/indZxG7Q=\n-----END CERTIFICATE-----",
		"-----BEGIN EC PARAMETERS-----\nBggqhkjOPQMBBw==\n-----END EC PARAMETERS-----\n-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIOIyqdM4PTaRPluesMVyHNyYVoBzQGvyekdqSQYY5jProAoGCCqGSM49\nAwEHoUQDQgAE61GKfDL8oshn+j1psor6as+bZSHbvZ07VfBV24/1F76c4CyFJi+2\n+RywujfRd31r0WAkqu0YS02ZxSvbzukHMw==\n-----END EC PRIVATE KEY-----",
		"",
		"177gjwaK7t3mJa5KrMH7xDsgWYj7qkBxuoDzLRc4JByDot52UMFkXpCcBBjMBF6aJKB5oZq",
		"8EjkXVSTxMFjCvNNsTo8RBMDEVQmk7gYkW4SCDuvdsBG",
		"9WsnNBZmtLA6Cr3jpNMsRC4NLuFQhuEvwgiG74uA3i8GWBxVC1rPh2hCNRqeMdPURWhFf2MjAgs9JV2Zijfpakr68yEZG"},
}

func TestResolveCertificate(t *testing.T) {
	for _, cert := range certs {
		certificate, err := RetrieveCertificate(cert[1])
		require.Nil(t, err)
		require.Equal(t, cert[0], certificate.Algorithm)
	}
}

func TestResolvePubKey(t *testing.T) {
	for _, cert := range certs {
		certificate, err := RetrieveCertificate(cert[1])
		require.Nil(t, err)
		require.Equal(t, cert[0], certificate.Algorithm)
		key := RetrievePubKey(certificate)
		require.Equal(t, crypto.EncodePubKey(key), cert[6])
	}
}

func TestResolvePrivKey(t *testing.T) {
	for _, cert := range certs {
		certificate, err := RetrieveCertificate(cert[1])
		require.Nil(t, err)
		var key framework.PrivKey
		if len(cert[3]) == 0 {
			key, err = RetrievePrivKey(certificate.Algorithm, cert[2])
			require.Nil(t, err)
		} else {
			key, err = RetrieveEncrypedPrivKey(certificate.Algorithm, cert[2], []byte(cert[3]))
			require.Nil(t, err)
		}
		require.Equal(t, cert[4], crypto.EncodePrivKey(key, base58.MustDecode(cert[5])))

		fmt.Println(certificate.Algorithm)
		function := crypto.GetSignatureFunctionByName(certificate.Algorithm)
		sign := function.Sign(key, []byte("imuge"))
		pubKey := RetrievePubKey(certificate)
		require.True(t, function.Verify(pubKey, []byte("imuge"), sign))
	}
}