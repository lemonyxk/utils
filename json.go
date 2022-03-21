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

func (j json) Encode(v any) []byte {
	res, err := jsoniter.Marshal(v)
	if err != nil {
		return nil
	}
	return res
}

func (j json) Decode(data []byte, output any) error {
	return jsoniter.Unmarshal(data, output)
}

type anyRes struct {
	any jsoniter.Any
}

type res struct {
	any jsoniter.Any
}

func (j json) New(v any) anyRes {
	return anyRes{any: jsoniter.Get(j.Encode(v))}
}

func (j json) Bytes(v []byte) anyRes {
	return anyRes{any: jsoniter.Get(v)}
}

func (j json) String(v string) anyRes {
	return anyRes{any: jsoniter.Get([]byte(v))}
}

func (r anyRes) Any() jsoniter.Any {
	return r.any
}

func (r anyRes) Get(path ...any) res {
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

func (r res) Interface() any {
	return r.any.GetInterface()
}

func (r res) Array() arrRes {
	var result []jsoniter.Any
	var val = r.any
	for i := 0; i < val.Size(); i++ {
		result = append(result, val.Get(i))
	}
	return result
}

type arrRes []jsoniter.Any

func (a arrRes) String() []string {
	var result []string
	for i := 0; i < len(a); i++ {
		result = append(result, a[i].ToString())
	}
	return result
}

func (a arrRes) Int() []int {
	var result []int
	for i := 0; i < len(a); i++ {
		result = append(result, a[i].ToInt())
	}
	return result
}

func (a arrRes) Float64() []float64 {
	var result []float64
	for i := 0; i < len(a); i++ {
		result = append(result, a[i].ToFloat64())
	}
	return result
}

func (a arrRes) Get(path ...any) arrRes {
	var result []jsoniter.Any
	for i := 0; i < len(a); i++ {
		result = append(result, a[i].Get(path...))
	}
	return result
}
