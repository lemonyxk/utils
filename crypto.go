/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2019-11-05 20:45
**/

package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"hash"
)

type crypto int

const Crypto crypto = iota

type cs struct {
	bts []byte
}

func (c *cs) Bytes() []byte {
	return c.bts
}

func (c *cs) Hex() string {
	return hex.EncodeToString(c.bts)
}

func (c crypto) Md5(input []byte) *cs {
	var byte16 = md5.Sum(input)
	var bytes = make([]byte, 16)
	for i := 0; i < 16; i++ {
		bytes[i] = byte16[i]
	}
	return &cs{bts: bytes}
}

func (c crypto) Sha1(input []byte) *cs {
	var sha = sha1.Sum(input)
	var bytes = make([]byte, 20)
	for i := 0; i < 20; i++ {
		bytes[i] = sha[i]
	}
	return &cs{bts: bytes}
}

func (c crypto) Sha256(input []byte) *cs {
	var sha = sha256.Sum256(input)
	var bytes = make([]byte, 32)
	for i := 0; i < 32; i++ {
		bytes[i] = sha[i]
	}
	return &cs{bts: bytes}
}

func (c crypto) Sha512(input []byte) *cs {
	var sha = sha512.Sum512(input)
	var bytes = make([]byte, 64)
	for i := 0; i < 64; i++ {
		bytes[i] = sha[i]
	}
	return &cs{bts: bytes}
}

func (c crypto) HmacSha1(fn func() hash.Hash, key, input []byte) *cs {
	var h = hmac.New(fn, key)
	h.Write(input)
	return &cs{bts: h.Sum(nil)}
}

func (c crypto) Base64Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func (c crypto) Base64Decode(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}
