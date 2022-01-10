package biteship

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

func (bite *BiteshipImpl) CreateOrder(request *CreateOrderRequestParam) (*ResponseCreateOrder, *Error) {

	var client = &http.Client{}
	var resp *ResponseCreateOrder
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

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequest))
	req.Header.Add("Authorization", bite.Config.SecretKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	//req.Header.Add("Biteship-lib", "go")
	//req.Header.Add("Biteship-lib-ver", "v0")

	response, errRequest := client.Do(req)
	if errRequest != nil {
		log.Println(errRequest)
		return resp, ErrorGo(errRequest)
	}
	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return resp, ErrorGo(err)
	}

	fmt.Printf("status %d", response.StatusCode)
	fmt.Printf("data %s", respBody)

	errUnmarshal := json.Unmarshal(respBody, &resp)
	if errUnmarshal != nil {
		log.Println(errUnmarshal)
		return resp, ErrorGo(errUnmarshal)
	}

	return resp, nil
}
