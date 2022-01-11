package biteship

import (
	"fmt"
	"net/http"
)

func (bite *BiteshipImpl) GetCourier() (*ResponseListCourier, *Error) {
	resp := &ResponseListCourier{}
	var url = fmt.Sprintf("%s/v1/couriers", bite.Config.BiteshipUrl)

	errRequest := bite.HttpRequest.Call(http.MethodGet, url, bite.Config.SecretKey, nil, resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}
