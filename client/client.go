package client

//
//import (
//	"bytes"
//	"encoding/json"
//	"github.com/muh-hizbe/biteship"
//	"io/ioutil"
//	"log"
//	"net/http"
//)
//
//type Client interface {
//	CreateOrder(biteship.RequestParam) (*biteship.ResponseCreateOrder, *biteship.Error)
//}
//
//type ClientImplementation struct {
//	Config *biteship.ConfigOption
//}
//
//func New() *ClientImplementation {
//	return &ClientImplementation{
//		Config: &biteship.Config,
//	}
//}
//
//func (c *ClientImplementation) CreateOrder(request biteship.RequestParam) (*biteship.ResponseCreateOrder, *biteship.Error) {
//	var err error
//	var client = &http.Client{}
//	var req *http.Request
//	var resp biteship.ResponseCreateOrder
//	var url = c.Config.BiteshipUrl + "/v1/orders"
//	//var inputReq bytes.Buffer
//	jsonRequest, errMarshal := json.Marshal(request)
//	if errMarshal != nil {
//		panic(errMarshal)
//	}
//
//	//req, err := http.NewRequest("POST", url+"/users", nil)
//	//response, err := client.Post(url, "application/json", &inputReq)
//	req.Header.Set("Authorization", "Bearer "+string(c.Config.SecretKey))
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("biteship-lib", "go")
//	req.Header.Set("biteship-lib-ver", "v0")
//
//	response, err := client.Do(req)
//	response, err = client.Post(url, "application/json", bytes.NewBuffer(jsonRequest))
//	response.Body.Close()
//
//	respBody, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	errUnmarshal := json.Unmarshal(respBody, &resp)
//	if errUnmarshal != nil {
//		log.Fatal(errUnmarshal)
//	}
//	return &resp, nil
//}
//
////func (c Client) Config(config *biteship.ConfigOption) *Client {
////	return &Client{
////		Config: *config,
////	}
////}
