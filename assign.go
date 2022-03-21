/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2020-04-27 03:33
**/

package utils

import (
	"errors"
	"reflect"
	"strconv"
)

type ass int

const Assign ass = iota

type destination struct {
	dst       any
	src       any
	tag       string
	allowZero bool
	allowTag  bool
	allowWeak bool
}

type field struct {
	data *destination
	name string
}

type source struct {
	data *destination
}

func (a ass) Dest(dst any) *destination {
	return &destination{
		dst:       dst,
		src:       nil,
		tag:       "json",
		allowZero: false,
		allowTag:  false,
		allowWeak: false,
	}
}

func (d *destination) Src(src any) *source {
	d.src = src
	return &source{data: d}
}

func (d *destination) Field(name string) *field {
	return &field{data: d, name: name}
}

func (s *source) Do() error {
	var d = s.data
	return doAssign(d.dst, d.src, d.tag, d.allowZero, d.allowTag, d.allowWeak)
}

func (s *source) AllowZero() *source {
	s.data.allowZero = true
	return s
}

func (s *source) AllowTag() *source {
	s.data.allowTag = true
	return s
}

func (s *source) AllowWeak() *source {
	s.data.allowWeak = true
	return s
}

func (s *source) SetTag(tag string) *source {
	s.data.tag = tag
	return s
}

func (f *field) Set(v any) error {
	var d = f.data
	d.src = map[string]any{f.name: v}
	d.allowZero = true
	return doAssign(d.dst, d.src, d.tag, d.allowZero, d.allowTag, d.allowWeak)
}

func (f *field) AllowTag() *field {
	f.data.allowTag = true
	return f
}

func (f *field) SetTag(tag string) *field {
	f.data.tag = tag
	return f
}

func doAssign(dst, src any, tag string, allowZero, allowTag bool, allowWeak bool) error {

	var dstValue = reflect.ValueOf(dst)
	var srcValue = reflect.ValueOf(src)
	var dstType = reflect.TypeOf(dst)
	var srcType = reflect.TypeOf(src)

	if dstValue.IsNil() {
		return errors.New("dst is nil")
	}

	var dstValueElem = dstValue.Elem()
	var srcValueElem = srcValue
	var dstTypeElem = dstType.Elem()
	var srcTypeElem = srcType

	if dstType.Kind() != reflect.Ptr {
		return errors.New("kind of dst is not ptr")
	}

	if dstTypeElem.Kind() != reflect.Struct {
		return errors.New("kind of dst is not struct")
	}

	if srcType.Kind() == reflect.Ptr {
		// nothing to copy
		if srcValue.IsNil() {
			return nil
		}
		srcValueElem = srcValue.Elem()
		srcTypeElem = srcType.Elem()
	}

	// if !allowDiffType {
	// 	if dstValueElem.Type() != srcValueElem.Type() {
	// 		return errors.New("dst and src has different kind")
	// 	}
	// }

	switch srcTypeElem.Kind() {
	case reflect.Struct:
		for i := 0; i < srcTypeElem.NumField(); i++ {

			var t = srcValueElem.Field(i)

			if !allowZero {
				if t.IsZero() {
					continue
				}
			}

			var name = srcTypeElem.Field(i).Name

			var s, ok = dstTypeElem.FieldByName(name)
			if !ok {
				continue
			}

			if s.Type.Kind() != t.Kind() {
				if !allowWeak {
					continue
				}
			}

			v := dstValueElem.FieldByIndex(s.Index)

			weakFunc(v, t)

			// dstValueElem.FieldByIndex(s.Index).Set(t)
		}
	case reflect.Map:
		var keys = srcValueElem.MapKeys()
		for i := 0; i < len(keys); i++ {

			if keys[i].Kind() != reflect.String {
				continue
			}

			var name = keys[i]
			var t = srcValueElem.MapIndex(name)

			if !t.IsValid() {
				continue
			}

			if !allowZero {
				if t.IsZero() {
					continue
				}
			}

			srcKey := name.String()

			var s, ok = dstTypeElem.FieldByName(srcKey)
			if !ok {
				if !allowTag {
					continue
				}
				k, ok := hasTag(dstTypeElem, tag, srcKey)
				if !ok {
					continue
				}

				srcKey = k.Name
				s = k

			}

			var it = reflect.TypeOf(t.Interface())

			if s.Type.Kind() != it.Kind() {
				if !allowWeak {
					continue
				}
			}

			v := dstValueElem.FieldByName(srcKey)

			weakFunc(v, t)

		}
	default:
		return errors.New("kind of src is not struct or map")
	}

	return nil
}

func weakFunc(v, t reflect.Value) {

	switch v.Interface().(type) {
	case int:
		switch t.Interface().(type) {
		case int:
			v.SetInt(int64(t.Interface().(int)))
		case uint64:
			v.SetInt(int64(t.Interface().(uint64)))
		case float64:
			v.SetInt(int64(t.Interface().(float64)))
		case bool:
			if t.Interface().(bool) {
				v.SetInt(1)
			} else {
				v.SetInt(0)
			}
		case []byte:
			v.SetInt(0)
		case string:
			var s = t.Interface().(string)
			r, _ := strconv.Atoi(s)
			v.SetInt(int64(r))
		}
	case uint64:
		switch t.Interface().(type) {
		case int:
			v.SetUint(uint64(t.Interface().(int)))
		case uint64:
			v.SetUint(uint64(t.Interface().(uint64)))
		case float64:
			v.SetUint(uint64(t.Interface().(float64)))
		case bool:
			if t.Interface().(bool) {
				v.SetUint(1)
			} else {
				v.SetUint(0)
			}
		case []byte:
			v.SetUint(0)
		case string:
			var s = t.Interface().(string)
			r, _ := strconv.Atoi(s)
			v.SetUint(uint64(r))
		}

	case float64:
		switch t.Interface().(type) {
		case int:
			v.SetFloat(float64(t.Interface().(int)))
		case uint64:
			v.SetFloat(float64(t.Interface().(uint64)))
		case float64:
			v.SetFloat(t.Interface().(float64))
		case bool:
			if t.Interface().(bool) {
				v.SetFloat(1)
			} else {
				v.SetFloat(0)
			}
		case []byte:
			v.SetFloat(0)
		case string:
			var s = t.Interface().(string)
			r, _ := strconv.ParseFloat(s, 64)
			v.SetFloat(r)
		}

	case bool:
		switch t.Interface().(type) {
		case int:
			v.SetBool(t.Interface().(int) > 0)
		case uint64:
			v.SetBool(t.Interface().(uint64) > 0)
		case float64:
			v.SetBool(t.Interface().(float64) > 0)
		case bool:
			if t.Interface().(bool) {
				v.SetBool(true)
			} else {
				v.SetBool(false)
			}
		case []byte:
			v.SetBool(false)
		case string:
			v.SetBool(t.Interface().(string) != "")
		}

	case []byte:
		switch t.Interface().(type) {
		case int:
			v.SetBytes([]byte(nil))
		case uint64:
			v.SetBytes([]byte(nil))
		case float64:
			v.SetBytes([]byte(nil))
		case bool:
			v.SetBytes([]byte(nil))
		case []byte:
			v.SetBytes(t.Interface().([]byte))
		case string:
			v.SetBytes([]byte(t.Interface().(string)))
		}

	case string:
		switch t.Interface().(type) {
		case int:
			var s = strconv.Itoa(t.Interface().(int))
			v.SetString(s)
		case uint64:
			var s = strconv.FormatUint(t.Interface().(uint64), 10)
			v.SetString(s)
		case float64:
			var s = strconv.FormatFloat(t.Interface().(float64), 'f', -1, 64)
			v.SetString(s)
		case bool:
			if t.Interface().(bool) {
				v.SetString("TRUE")
			} else {
				v.SetString("FALSE")
			}
		case []byte:
			v.SetString(string(t.Interface().([]byte)))
		case string:
			v.SetString(t.Interface().(string))
		}
	}

}

func hasTag(s reflect.Type, t, k string) (reflect.StructField, bool) {
	var n = s.NumField()
	for i := 0; i < n; i++ {
		if s.Field(i).Tag.Get(t) == k {
			return s.Field(i), true
		}
	}
	return reflect.StructField{}, false
}
