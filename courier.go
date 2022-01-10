package biteship

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (bite *BiteshipImpl) GetCourier() (ResponseListCourier, error) {
	var client = http.Client{}
	var resp = ResponseListCourier{}
	var url = fmt.Sprintf("%s/v1/couriers", bite.Config.BiteshipUrl)

	req, _ := http.NewRequest("GET", url, nil)
	//req.SetBasicAuth(bite.Config.SecretKey, "")
	req.Header.Add("Authorization", bite.Config.SecretKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	response, errRequest := client.Do(req)
	if errRequest != nil {
		log.Println(errRequest)
		return resp, errRequest
	}
	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	fmt.Printf("status %d", response.StatusCode)
	fmt.Printf("data %s", respBody)

	errUnmarshal := json.Unmarshal(respBody, &resp)
	if errUnmarshal != nil {
		log.Println(errUnmarshal)
		return resp, errUnmarshal
	}

	return resp, nil
}
