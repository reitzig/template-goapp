package internal

import "fmt"

type ConfigValue[V any] struct {
	Value  V
	Source string
}

type PrintableConfigValue struct {
	name   string
	value  string
	source string
}

type Config interface {
	values() []PrintableConfigValue
}

type RootConfig struct {
	Debug ConfigValue[bool]

	Hello HelloConfig
}

func (rootConfig *RootConfig) values() []PrintableConfigValue {
	return append([]PrintableConfigValue{
		PrintableConfigValue{
			name:   "debug",
			value:  fmt.Sprintf("%v", rootConfig.Debug.Value),
			source: rootConfig.Debug.Source,
		},
	}, rootConfig.Hello.values()...)
}

type HelloConfig struct {
	Dear ConfigValue[bool]
}

func (helloConfig *HelloConfig) values() []PrintableConfigValue {
	return []PrintableConfigValue{
		PrintableConfigValue{
			name:   "dear",
			value:  fmt.Sprintf("%v", helloConfig.Dear.Value),
			source: helloConfig.Dear.Source,
		},
	}
}
