/**
* @program: lemon
*
* @description:
*
* @author: lemon
*
* @create: 2019-11-04 18:53
**/

package json

import (
	json "github.com/json-iterator/go"
)

func Encode(v any) []byte {
	res, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	return res
}

func Decode(data []byte, output any) error {
	return json.Unmarshal(data, output)
}

type Result struct {
	any json.Any
}

func New(v any) Result {
	return Result{any: json.Get(Encode(v))}
}

func Bytes(v []byte) Result {
	return Result{any: json.Get(v)}
}

func String(v string) Result {
	return Result{any: json.Get([]byte(v))}
}

func (r Result) Any() json.Any {
	return r.any
}

func (r Result) Get(path ...any) Result {
	return Result{any: r.any.Get(path...)}
}

func (r Result) Exists() bool {
	return r.any.LastError() == nil
}

func (r Result) String() string {
	return r.any.ToString()
}

func (r Result) Bytes() []byte {
	return []byte(r.any.ToString())
}

func (r Result) Size() int {
	return r.any.Size()
}

func (r Result) Int() int {
	return r.any.ToInt()
}

func (r Result) Float64() float64 {
	return r.any.ToFloat64()
}

func (r Result) Bool() bool {
	return r.any.ToBool()
}

func (r Result) Interface() any {
	return r.any.GetInterface()
}

func (r Result) Array() Results {
	var results []json.Any
	var val = r.any
	for i := 0; i < val.Size(); i++ {
		results = append(results, val.Get(i))
	}
	return results
}

type Results []json.Any

func (a Results) String() []string {
	var result []string
	for i := 0; i < len(a); i++ {
		result = append(result, a[i].ToString())
	}
	return result
}

func (a Results) Int() []int {
	var result []int
	for i := 0; i < len(a); i++ {
		result = append(result, a[i].ToInt())
	}
	return result
}

func (a Results) Float64() []float64 {
	var result []float64
	for i := 0; i < len(a); i++ {
		result = append(result, a[i].ToFloat64())
	}
	return result
}

func (a Results) Get(path ...any) Results {
	var result []json.Any
	for i := 0; i < len(a); i++ {
		result = append(result, a[i].Get(path...))
	}
	return result
}
