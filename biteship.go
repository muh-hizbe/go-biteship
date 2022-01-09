package biteship

import "fmt"

type Biteship interface {
	CreateOrder(request *RequestParam) (*ResponseCreateOrder, error)
}

type BiteshipImpl struct {
	Config *ConfigOption
}

func New(key string, config ...ConfigOption) Biteship {
	defaultConfig := DefaultConfig(key)

	if len(config) > 0 {
		defaultConfig = &config[0]
		defaultConfig.SecretKey = key
	}

	fmt.Println(defaultConfig)

	return &BiteshipImpl{
		Config: defaultConfig,
	}
}
