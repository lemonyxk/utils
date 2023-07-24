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
	"github.com/mitchellh/mapstructure"
)

type structure struct {
	TagName string
}

var Struct = structure{
	TagName: "json",
}

func (d structure) WeakDecode(input any, output any) error {
	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           output,
		TagName:          d.TagName,
		WeaklyTypedInput: true,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func (d structure) Decode(input any, output any) error {
	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   output,
		TagName:  d.TagName,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}
