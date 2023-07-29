/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2023-07-30 02:05
**/

package crypto

import "encoding/base64"

func Base64Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func Base64Decode(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}
