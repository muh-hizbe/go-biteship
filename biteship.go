package biteship

type Biteship interface {
	CreateOrder(request *CreateOrderRequestParam) (*ResponseCreateOrder, *Error)
	GetCourier() (*ResponseListCourier, *Error)
	GetRatesCouriers(request *RequestCourierRates) (*ResponseListRatesCouriers, *Error)
}

type BiteshipImpl struct {
	Config      *ConfigOption
	HttpRequest *HttpRequestImpl
}

func New(key string, config ...ConfigOption) Biteship {
	defaultConfig := DefaultConfig(key)

	if len(config) > 0 {
		defaultConfig = &config[0]
		defaultConfig.SecretKey = key
	}

	//fmt.Println(defaultConfig)

	return &BiteshipImpl{
		Config: defaultConfig,
	}
}
