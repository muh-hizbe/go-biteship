package biteship

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

func (bite *BiteshipImpl) GetRatesCouriers(request *RequestCourierRates) (*ResponseListRatesCouriers, *Error) {
	//var client = http.Client{}
	var resp = &ResponseListRatesCouriers{}
	var url = fmt.Sprintf("%s/v1/rates/couriers", bite.Config.BiteshipUrl)

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

	//
	//req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequest))
	//req.Header.Add("Authorization", bite.Config.SecretKey)
	//req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Accept", "application/json")
	//
	//response, errRequest := client.Do(req)
	//if errRequest != nil {
	//	log.Println(errRequest)
	//	return resp, errRequest
	//}
	//defer response.Body.Close()
	//
	//respBody, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	log.Println(err)
	//	return resp, err
	//}
	//
	//fmt.Printf("status %d", response.StatusCode)
	//fmt.Printf("data %s", respBody)
	//
	//errUnmarshal := json.Unmarshal(respBody, &resp)
	//if errUnmarshal != nil {
	//	log.Println(errUnmarshal)
	//	return resp, errUnmarshal
	//}

	return resp, nil
}
