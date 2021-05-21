/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2019-11-04 18:53
**/

package utils

import (
	"github.com/json-iterator/go"
)

type json int

const Json json = iota

func (j json) Encode(v interface{}) []byte {
	res, err := jsoniter.Marshal(v)
	if err != nil {
		return nil
	}
	return res
}

func (j json) Decode(data []byte, output interface{}) error {
	return jsoniter.Unmarshal(data, output)
}

type any struct {
	any jsoniter.Any
}

type res struct {
	any jsoniter.Any
}

func (j json) New(v interface{}) any {
	return any{any: jsoniter.Get(j.Encode(v))}
}

func (j json) Bytes(v []byte) any {
	return any{any: jsoniter.Get(v)}
}

func (j json) String(v string) any {
	return any{any: jsoniter.Get([]byte(v))}
}

func (r any) Any() jsoniter.Any {
	return r.any
}

func (r any) Get(path ...interface{}) res {
	return res{any: r.any.Get(path...)}
}

func (r res) Exists() bool {
	return r.any.LastError() == nil
}

func (r res) String() string {
	return r.any.ToString()
}

func (r res) Bytes() []byte {
	return []byte(r.any.ToString())
}

func (r res) Size() int {
	return r.any.Size()
}

func (r res) Int() int {
	return r.any.ToInt()
}

func (r res) Float64() float64 {
	return r.any.ToFloat64()
}

func (r res) Bool() bool {
	return r.any.ToBool()
}

func (r res) Interface() interface{} {
	return r.any.GetInterface()
}

func (r res) Array() resArr {
	var result []jsoniter.Any
	var val = r.any
	for i := 0; i < val.Size(); i++ {
		result = append(result, val.Get(i))
	}
	return result
}

type resArr []jsoniter.Any

func (a resArr) String() []string {
	var result []string
	for i := 0; i < len(a); i++ {
		result = append(result, a[i].ToString())
	}
	return result
}

func (a resArr) Int() []int {
	var result []int
	for i := 0; i < len(a); i++ {
		result = append(result, a[i].ToInt())
	}
	return result
}

func (a resArr) Float64() []float64 {
	var result []float64
	for i := 0; i < len(a); i++ {
		result = append(result, a[i].ToFloat64())
	}
	return result
}

func (a resArr) Get(path ...interface{}) resArr {
	var result []jsoniter.Any
	for i := 0; i < len(a); i++ {
		result = append(result, a[i].Get(path...))
	}
	return result
}
