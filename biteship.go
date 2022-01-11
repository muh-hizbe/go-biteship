package biteship

type Biteship interface {
	GetCourier() (*ResponseListCourier, *Error)
	GetRatesCouriers(request *RequestCourierRates) (*ResponseListRatesCouriers, *Error)

	CreateOrder(request *CreateOrderRequestParam) (*ResponseCreateOrder, *Error)
	RetrieveOrder(orderId string) (*ResponseRetrieveOrder, *Error)
	UpdateOrder(orderId string, request interface{}) (*ResponseCreateOrder, *Error)
	ConfirmOrder(orderId string) (*ResponseCreateOrder, *Error)
	CancelOrder(orderId string, reason string) (*ResponseCancelOrder, *Error)

	TrackingOrder(orderId string) (*ResponseTrackingOrder, *Error)
	TrackingOrderByWaybill(waybillId string, courierCode string) (*ResponseTrackingOrder, *Error)
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
