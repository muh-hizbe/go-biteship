package biteship

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

func (bite *BiteshipImpl) CreateOrder(request *CreateOrderRequestParam) (*ResponseCreateOrder, *Error) {

	resp := &ResponseCreateOrder{}
	var url = fmt.Sprintf("%s/v1/orders", bite.Config.BiteshipUrl)
	var errMarshal error
	jsonRequest := []byte("")

	isParamsNil := reflect.ValueOf(request).Kind() == reflect.Ptr && reflect.ValueOf(request).IsNil()

	if !isParamsNil {
		jsonRequest, errMarshal = json.Marshal(request)
		if errMarshal != nil {
			log.Println(errMarshal)
			return resp, ErrorGo(errMarshal)
		}
	}

	errRequest := bite.HttpRequest.Call(http.MethodPost, url, bite.Config.SecretKey, bytes.NewBuffer(jsonRequest), resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}
