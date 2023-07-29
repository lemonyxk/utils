/**
* @program: lemon
*
* @description:
*
* @author: lemon
*
* @create: 2019-11-05 11:42
**/

package structure

import (
	"github.com/mitchellh/mapstructure"
)

type Config struct {
	TagName string
}

var defaultConfig = Config{
	TagName: "json",
}

func SetConfig(config Config) {
	defaultConfig = config
}

func WeakDecode(input any, output any) error {
	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           output,
		TagName:          defaultConfig.TagName,
		WeaklyTypedInput: true,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func Decode(input any, output any) error {
	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   output,
		TagName:  defaultConfig.TagName,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}
