/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2020-08-21 16:27
**/

package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

type ex int

const Extract ex = iota

type meta struct {
	src []interface{}
}

type extract struct {
	src   []interface{}
	field string
}

func (e ex) Src(src ...interface{}) *meta {
	return &meta{src: src}
}

func (m *meta) Field(name string) *extract {
	return &extract{src: m.src, field: name}
}

func (e *extract) Int() []int {
	return doInt(e.src, e.field)
}

func (e *extract) Int32() []int32 {
	return doInt32(e.src, e.field)
}

func (e *extract) Int64() []int64 {
	return doInt64(e.src, e.field)
}

func (e *extract) Float32() []float32 {
	return doFloat32(e.src, e.field)
}

func (e *extract) Float64() []float64 {
	return doFloat64(e.src, e.field)
}

func (e *extract) String() []string {
	return doString(e.src, e.field)
}

func doInt(src []interface{}, field string) []int {
	var data = doExtract(src, field)
	var res []int
	for i := 0; i < len(data); i++ {
		switch data[i].(type) {
		case int:
			res = append(res, data[i].(int))
		case int32:
			res = append(res, int(data[i].(int32)))
		case int64:
			res = append(res, int(data[i].(int64)))
		case float32:
			res = append(res, int(data[i].(float32)))
		case float64:
			res = append(res, int(data[i].(float64)))
		case string:
			var r, _ = strconv.Atoi(data[i].(string))
			res = append(res, r)
		default:
			var r, _ = strconv.Atoi(fmt.Sprintf("%v", data[i]))
			res = append(res, r)
		}
	}
	return res
}

func doInt32(src []interface{}, field string) []int32 {
	var data = doExtract(src, field)
	var res []int32
	for i := 0; i < len(data); i++ {
		switch data[i].(type) {
		case int:
			res = append(res, int32(data[i].(int)))
		case int32:
			res = append(res, data[i].(int32))
		case int64:
			res = append(res, int32(data[i].(int64)))
		case float32:
			res = append(res, int32(data[i].(float32)))
		case float64:
			res = append(res, int32(data[i].(float64)))
		case string:
			var r, _ = strconv.Atoi(data[i].(string))
			res = append(res, int32(r))
		default:
			var r, _ = strconv.Atoi(fmt.Sprintf("%v", data[i]))
			res = append(res, int32(r))
		}
	}
	return res
}

func doInt64(src []interface{}, field string) []int64 {
	var data = doExtract(src, field)
	var res []int64
	for i := 0; i < len(data); i++ {
		switch data[i].(type) {
		case int:
			res = append(res, int64(data[i].(int)))
		case int32:
			res = append(res, int64(data[i].(int32)))
		case int64:
			res = append(res, data[i].(int64))
		case float32:
			res = append(res, int64(data[i].(float32)))
		case float64:
			res = append(res, int64(data[i].(float64)))
		case string:
			var r, _ = strconv.Atoi(data[i].(string))
			res = append(res, int64(r))
		default:
			var r, _ = strconv.Atoi(fmt.Sprintf("%v", data[i]))
			res = append(res, int64(r))
		}
	}
	return res
}

func doFloat32(src []interface{}, field string) []float32 {
	var data = doExtract(src, field)
	var res []float32
	for i := 0; i < len(data); i++ {
		switch data[i].(type) {
		case int:
			res = append(res, float32(data[i].(int)))
		case int32:
			res = append(res, float32(data[i].(int32)))
		case int64:
			res = append(res, float32(data[i].(int64)))
		case float32:
			res = append(res, data[i].(float32))
		case float64:
			res = append(res, float32(data[i].(float64)))
		case string:
			var r, _ = strconv.ParseFloat(data[i].(string), 32)
			res = append(res, float32(r))
		default:
			var r, _ = strconv.ParseFloat(fmt.Sprintf("%v", data[i]), 32)
			res = append(res, float32(r))
		}
	}
	return res
}

func doFloat64(src []interface{}, field string) []float64 {
	var data = doExtract(src, field)
	var res []float64
	for i := 0; i < len(data); i++ {
		switch data[i].(type) {
		case int:
			res = append(res, float64(data[i].(int)))
		case int32:
			res = append(res, float64(data[i].(int32)))
		case int64:
			res = append(res, float64(data[i].(int64)))
		case float32:
			res = append(res, float64(data[i].(float32)))
		case float64:
			res = append(res, data[i].(float64))
		case string:
			var r, _ = strconv.ParseFloat(data[i].(string), 64)
			res = append(res, r)
		default:
			var r, _ = strconv.ParseFloat(fmt.Sprintf("%v", data[i]), 64)
			res = append(res, r)
		}
	}
	return res
}

func doString(src []interface{}, field string) []string {
	var data = doExtract(src, field)
	var res []string
	for i := 0; i < len(data); i++ {
		switch data[i].(type) {
		case int:
			var r = strconv.Itoa(data[i].(int))
			res = append(res, r)
		case int32:
			var r = strconv.Itoa(int(data[i].(int32)))
			res = append(res, r)
		case int64:
			var r = strconv.Itoa(int(data[i].(int64)))
			res = append(res, r)
		case float32:
			var r = strconv.FormatFloat(float64(data[i].(float32)), 'E', -1, 32)
			res = append(res, r)
		case float64:
			var r = strconv.FormatFloat(float64(data[i].(float32)), 'E', -1, 64)
			res = append(res, r)
		case string:
			res = append(res, data[i].(string))
		default:
			res = append(res, fmt.Sprintf("%v", data[i]))
		}
	}
	return res
}

func doExtract(src []interface{}, field string) []interface{} {

	if len(src) == 0 {
		return src
	}

	var res []interface{}

	var fn func(source interface{})
	fn = func(source interface{}) {
		var srcValue = reflect.ValueOf(source)
		var srcType = reflect.TypeOf(source)

		var srcValueElem = srcValue
		var srcTypeElem = srcType

		if !srcValueElem.IsValid() {
			return
		}

		if srcType.Kind() == reflect.Ptr {
			if srcValue.IsNil() {
				return
			}
			srcValueElem = srcValue.Elem()
			srcTypeElem = srcType.Elem()
		}

		switch srcTypeElem.Kind() {
		case reflect.Struct:
			var s, ok = srcTypeElem.FieldByName(field)
			if !ok {
				return
			}
			var v = srcValueElem.FieldByIndex(s.Index)
			if !v.IsValid() {
				res = append(res, nil)
				return
			}
			res = append(res, v.Interface())
		case reflect.Map:
			var keys = srcValueElem.MapKeys()
			for j := 0; j < len(keys); j++ {
				if keys[j].String() != field {
					continue
				}
				var v = srcValueElem.MapIndex(keys[j])
				if !v.IsValid() {
					res = append(res, nil)
					return
				}
				res = append(res, v.Interface())
				return
			}
		case reflect.Slice:
			for j := 0; j < srcValueElem.Len(); j++ {
				var v = srcValueElem.Index(j)
				if !v.IsValid() {
					fn(nil)
					return
				}
				fn(v.Interface())
			}
		default:
			panic("kind of src is not struct or map")
		}
	}

	for i := 0; i < len(src); i++ {
		fn(src[i])
	}

	return res
}
