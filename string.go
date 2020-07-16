/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2019-12-10 21:23
**/

package utils

import (
	"bytes"
	"fmt"
	"reflect"
	"unsafe"
)

type str int

const String str = iota

func (s str) JoinInterface(v []interface{}, sep string) string {
	var buf bytes.Buffer
	for i := 0; i < len(v); i++ {
		switch v[i].(type) {
		case string:
			buf.WriteString(v[i].(string))
		default:
			buf.WriteString(fmt.Sprintf("%v", v[i]))
		}
		if i != len(v)-1 {
			buf.WriteString(sep)
		}
	}
	return buf.String()
}

func (s str) Join(v []string, sep string) string {
	var buf bytes.Buffer
	for i := 0; i < len(v); i++ {
		buf.WriteString(v[i])
		if i != len(v)-1 {
			buf.WriteString(sep)
		}
	}
	return buf.String()
}

func (s str) ToIntInterface(v []interface{}) []int {
	var res []int
	for i := 0; i < len(v); i++ {
		switch v[i].(type) {
		case string:
			res = append(res, Conv.Atoi(v[i].(string)))
		case int:
			res = append(res, v[i].(int))
		case float64:
			res = append(res, int(v[i].(float64)))
		default:
			res = append(res, Conv.Atoi(fmt.Sprintf("%v", v[i])))
		}
	}
	return res
}

func (s str) ToInt(v []string) []int {
	var res []int
	for i := 0; i < len(v); i++ {
		res = append(res, Conv.Atoi(v[i]))
	}
	return res
}

func (s str) ToFloat64(v []string) []float64 {
	var res []float64
	for i := 0; i < len(v); i++ {
		res = append(res, Conv.StringToFloat64(v[i]))
	}
	return res
}

func (s str) BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func (s str) StringToBytes(str string) (bs []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return
}
