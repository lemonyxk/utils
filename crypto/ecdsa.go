/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2023-07-30 02:08
**/

package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"time"
)

func GenerateECDSAKey(bit int) *ecdsa.PrivateKey {
	var curve elliptic.Curve
	switch bit {
	case 224:
		curve = elliptic.P224()
	case 256:
		curve = elliptic.P256()
	case 384:
		curve = elliptic.P384()
	case 521:
		curve = elliptic.P521()
	default:
		panic("invalid bit")
	}
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	return privateKey
}

func GenerateCertificate(organization []string, dnsNames []string, duration time.Duration) ([]byte, []byte) {

	var privateKey = GenerateECDSAKey(256)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		panic(err)
	}

	template := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: organization,
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(duration),
		DNSNames:  dnsNames,

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
	if err != nil {
		panic(err)
	}

	pemCert := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if pemCert == nil {
		panic("pemCert is nil")
	}

	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		panic(err)
	}

	pemKey := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privateKeyBytes})
	if pemKey == nil {
		panic("pemKey is nil")
	}

	return pemCert, pemKey
}
