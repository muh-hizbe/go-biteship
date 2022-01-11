package biteship

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"reflect"
)

func (bite *BiteshipImpl) GetRatesCouriers(request *RequestCourierRates) (*ResponseListRatesCouriers, *Error) {
	validate = validator.New()
	errValidate := validate.Struct(request)
	if errValidate != nil {
		return nil, ErrorRequestParam(errValidate)
	}

	var resp = &ResponseListRatesCouriers{}
	var url = fmt.Sprintf("%s/v1/rates/couriers", bite.Config.BiteshipUrl)

	var errMarshal error
	jsonRequest := []byte("")

	isParamsNil := reflect.ValueOf(request).Kind() == reflect.Ptr && reflect.ValueOf(request).IsNil()

	if !isParamsNil {
		jsonRequest, errMarshal = json.Marshal(request)
		if errMarshal != nil {
			log.Println(errMarshal)
			return nil, ErrorGo(errMarshal)
		}
	}

	errRequest := bite.HttpRequest.Call(http.MethodPost, url, bite.Config.SecretKey, bytes.NewBuffer(jsonRequest), resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return resp, nil
}
