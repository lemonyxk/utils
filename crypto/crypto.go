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
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

type Bytes []byte

func (c Bytes) Hex() string {
	return hex.EncodeToString(c)
}

func Md5(input []byte) Bytes {
	var byte16 = md5.Sum(input)
	var bytes = make([]byte, 16)
	for i := 0; i < 16; i++ {
		bytes[i] = byte16[i]
	}
	return bytes
}

func Sha1(input []byte) Bytes {
	var sha = sha1.Sum(input)
	var bytes = make([]byte, 20)
	for i := 0; i < 20; i++ {
		bytes[i] = sha[i]
	}
	return bytes
}

func Sha256(input []byte) Bytes {
	var sha = sha256.Sum256(input)
	var bytes = make([]byte, 32)
	for i := 0; i < 32; i++ {
		bytes[i] = sha[i]
	}
	return bytes
}

func Sha512(input []byte) Bytes {
	var sha = sha512.Sum512(input)
	var bytes = make([]byte, 64)
	for i := 0; i < 64; i++ {
		bytes[i] = sha[i]
	}
	return bytes
}

func HMacSha1(fn func() hash.Hash, key, input []byte) Bytes {
	var h = hmac.New(fn, key)
	h.Write(input)
	return h.Sum(nil)
}
