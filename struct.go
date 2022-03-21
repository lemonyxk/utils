/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2019-11-05 11:42
**/

package utils

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type structure int

const Structure structure = iota

func (d structure) StructToMap(input any) map[string]any {
	var output = make(map[string]any)

	if input == nil {
		return output
	}

	var kf = reflect.TypeOf(input)

	var vf = reflect.ValueOf(input)

	if kf.Kind() == reflect.Ptr {
		kf = kf.Elem()
		vf = vf.Elem()
	}

	if kf.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < kf.NumField(); i++ {
		var v = vf.Field(i)
		if !v.IsValid() {
			output[kf.Field(i).Tag.Get("json")] = nil
			continue
		}
		output[kf.Field(i).Tag.Get("json")] = v.Interface()
	}

	return output
}

func (d structure) MapToStruct(input any, output any) error {
	return mapstructure.WeakDecode(input, output)
}

func (d structure) GetTags(src any) []string {
	if src == nil {
		return []string{}
	}

	var kf = reflect.TypeOf(src)
	if kf.Kind() == reflect.Ptr {
		kf = kf.Elem()
	}

	if kf.Kind() != reflect.Struct {
		return []string{}
	}

	var res []string

	for i := 0; i < kf.NumField(); i++ {
		var tag = kf.Field(i).Tag.Get("json")
		if tag == "" {
			tag = kf.Field(i).Name
		}
		res = append(res, tag)
	}

	return res
}
