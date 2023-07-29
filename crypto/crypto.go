/**
* @program: lemon
*
* @description:
*
* @author: lemon
*
* @create: 2019-11-05 20:45
**/

package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"hash"
)

type Result struct {
	bts []byte
}

func (c *Result) Bytes() []byte {
	return c.bts
}

func (c *Result) Hex() string {
	return hex.EncodeToString(c.bts)
}

func Md5(input []byte) *Result {
	var byte16 = md5.Sum(input)
	var bytes = make([]byte, 16)
	for i := 0; i < 16; i++ {
		bytes[i] = byte16[i]
	}
	return &Result{bts: bytes}
}

func GenerateKey(bit int) []byte {
	privateKey, err := rsa.GenerateKey(rand.Reader, bit)
	if err != nil {
		panic(err)
	}

	var bytes = x509.MarshalPKCS1PrivateKey(privateKey)
	var pemBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
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
		Type:  "RSA PUBLIC KEY",
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

func Sha1(input []byte) *Result {
	var sha = sha1.Sum(input)
	var bytes = make([]byte, 20)
	for i := 0; i < 20; i++ {
		bytes[i] = sha[i]
	}
	return &Result{bts: bytes}
}

func Sha256(input []byte) *Result {
	var sha = sha256.Sum256(input)
	var bytes = make([]byte, 32)
	for i := 0; i < 32; i++ {
		bytes[i] = sha[i]
	}
	return &Result{bts: bytes}
}

func Sha512(input []byte) *Result {
	var sha = sha512.Sum512(input)
	var bytes = make([]byte, 64)
	for i := 0; i < 64; i++ {
		bytes[i] = sha[i]
	}
	return &Result{bts: bytes}
}

func HMacSha1(fn func() hash.Hash, key, input []byte) *Result {
	var h = hmac.New(fn, key)
	h.Write(input)
	return &Result{bts: h.Sum(nil)}
}

func Base64Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func Base64Decode(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}
