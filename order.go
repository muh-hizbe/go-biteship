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

func (bite *BiteshipImpl) CreateOrder(request *RequestParam) (*ResponseCreateOrder, error) {
	var client = &http.Client{}
	var req *http.Request
	var resp *ResponseCreateOrder
	var url = fmt.Sprintf("%s/v1/orders", bite.Config.BiteshipUrl)
	var errMarshal error
	jsonRequest := []byte("")
	//var inputReq bytes.Buffer

	isParamsNil := reflect.ValueOf(request).Kind() == reflect.Ptr && reflect.ValueOf(request).IsNil()

	if !isParamsNil {
		log.Println("tidak nil")
		jsonRequest, errMarshal = json.Marshal(request)
		if errMarshal != nil {
			log.Println(errMarshal)
			return resp, errMarshal
		}
	}

	//req, err := http.NewRequest("POST", url+"/users", nil)
	//response, err := client.Post(url, "application/json", &inputReq)
	log.Println("sampai sini")
	log.Println(bite.Config.SecretKey)
	//token := fmt.Sprintf("Bearer %s", bite.Config.SecretKey)
	//req.Header.Set("Authorization", token)
	req.SetBasicAuth(bite.Config.SecretKey, "")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Biteship-lib", "go")
	req.Header.Set("Biteship-lib-ver", "v0")
	log.Println("nggak sampai sini")

	response, errRequest := client.Do(req)
	if errRequest != nil {
		log.Println(errRequest)
		return resp, errRequest
	}
	response, err := client.Post(url, "application/json", bytes.NewBuffer(jsonRequest))
	//response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	errUnmarshal := json.Unmarshal(respBody, &resp)
	if errUnmarshal != nil {
		log.Println(errUnmarshal)
		return resp, errUnmarshal
	}
	return resp, nil
}
