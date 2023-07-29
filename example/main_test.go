/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2023-07-29 19:16
**/

package main

import (
	"github.com/lemonyxk/utils/crypto"
	"testing"
)

var secretKey = crypto.GenerateKey(1024)

var publicKey = crypto.GeneratePublicKey(secretKey)

func BenchmarkName(b *testing.B) {
	var bts, _ = crypto.RsaEncrypt([]byte(publicKey), []byte("hello world"))

	for i := 0; i < b.N; i++ {

		_, _ = crypto.RsaDecrypt([]byte(secretKey), bts)

	}
}
