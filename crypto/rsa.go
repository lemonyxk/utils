/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2023-07-30 02:04
**/

package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func GenerateRSAKey(bit int) []byte {
	privateKey, err := rsa.GenerateKey(rand.Reader, bit)
	if err != nil {
		panic(err)
	}

	var bytes = x509.MarshalPKCS1PrivateKey(privateKey)
	var pemBlock = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: bytes,
	}

	return pem.EncodeToMemory(pemBlock)
}

func GeneratePublicKey(privateKey []byte) []byte {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		panic(errors.New("private key error"))
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	bytes := x509.MarshalPKCS1PublicKey(&key.PublicKey)
	var pemBlock = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: bytes,
	}

	return pem.EncodeToMemory(pemBlock)
}

func RsaEncrypt(publicKey, input []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.EncryptPKCS1v15(rand.Reader, pub, input)
}

func RsaDecrypt(privateKey, input []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, key, input)
}
