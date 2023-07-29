/**
* @program: lemon
*
* @description:
*
* @author: lemon
*
* @create: 2019-12-16 21:24
**/

package conv

import (
	"strconv"
	"unsafe"
)

func Itoa(i int) string {
	return strconv.Itoa(i)
}

func Atoi(i string) int {
	var n, _ = strconv.Atoi(i)
	return n
}

func StringToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Float64ToString(i float64) string {
	return strconv.FormatFloat(i, 'f', -1, 64)
}

func Float32ToString(i float64) string {
	return strconv.FormatFloat(i, 'f', -1, 32)
}

func StringToFloat64(i string) float64 {
	var n, _ = strconv.ParseFloat(i, 64)
	return n
}

func StringToFloat32(i string) float64 {
	var n, _ = strconv.ParseFloat(i, 32)
	return n
}
