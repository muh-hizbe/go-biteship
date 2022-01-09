package client

//
//import (
//	"bytes"
//	"context"
//	"encoding/json"
//	"github.com/muh-hizbe/biteship"
//	"io/ioutil"
//	"net/http"
//	"reflect"
//	//"time"
//)
//
//type HttpRequest interface {
//	Call(ctx context.Context, method string, url string, secretKey string, header *http.Header, body interface{}, result interface{}) *biteship.Error
//	doRequest(req *http.Request, result interface{}) *biteship.Error
//}
//
//type HttpRequestImplementation struct {
//	HttpClient *http.Client
//}
//
//func (a *HttpRequestImplementation) Call(ctx context.Context, method string, url string, secretKey string, header *http.Header, body interface{}, result interface{}) *biteship.Error {
//	reqBody := []byte("")
//	var req *http.Request
//	var err error
//
//	isParamsNil := body == nil || (reflect.ValueOf(body).Kind() == reflect.Ptr && reflect.ValueOf(body).IsNil())
//
//	if !isParamsNil {
//		reqBody, err = json.Marshal(body)
//		if err != nil {
//			return biteship.ErrorGo(err)
//		}
//	}
//
//	req, err = http.NewRequestWithContext(
//		ctx,
//		method,
//		url,
//		bytes.NewBuffer(reqBody),
//	)
//	if err != nil {
//		return biteship.ErrorGo(err)
//	}
//
//	if header != nil {
//		req.Header = *header
//	}
//	req.Header.Set("Authorization", "Bearer "+string(secretKey))
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("biteship-lib", "go")
//	req.Header.Set("biteship-lib-ver", "v0")
//
//	return a.doRequest(req, result)
//}
//
//func (a *HttpRequestImplementation) doRequest(req *http.Request, result interface{}) *biteship.Error {
//	resp, err := a.HttpClient.Do(req)
//	if err != nil {
//		return biteship.ErrorGo(err)
//	}
//	defer resp.Body.Close()
//	respBody, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return biteship.ErrorGo(err)
//	}
//
//	if resp.StatusCode < 200 || resp.StatusCode > 299 {
//		return biteship.ErrorHttp(resp.StatusCode, respBody)
//	}
//
//	if errUnmarshal := json.Unmarshal(respBody, &result); errUnmarshal != nil {
//		return biteship.ErrorGo(errUnmarshal)
//	}
//
//	return nil
//}
//
//// ApiResponse : is a structs that may come from Midtrans API endpoints
//type ApiResponse struct {
//	Status     string // e.g. "200 OK"
//	StatusCode int    // e.g. 200
//	Proto      string // e.g. "HTTP/1.0"
//
//	// response Header contain a map of all HTTP header keys to values.
//	Header http.Header
//	// response body
//	RawBody []byte
//	// request that was sent to obtain the response
//	Request *http.Request
//}
//
//// newHTTPResponse : internal function to set HTTP Raw response return to ApiResponse
//func newHTTPResponse(res *http.Response, responseBody []byte) *ApiResponse {
//	return &ApiResponse{
//		Status:     res.Status,
//		StatusCode: res.StatusCode,
//		Proto:      res.Proto,
//		Header:     res.Header,
//		RawBody:    responseBody,
//		Request:    res.Request,
//	}
//}
//
////// logHttpHeaders : internal function to perform log from headers
////func logHttpHeaders(log LoggerInterface, header http.Header, isReq bool) {
////	// Loop through headers to perform log
////	for name, headers := range header {
////		name = strings.ToLower(name)
////		for _, h := range headers {
////			if name == "authorization" {
////				log.Debug("%v: %v", name, h)
////			} else {
////				if isReq {
////					log.Info("%v: %v", name, h)
////				} else {
////					log.Debug("%v: %v", name, h)
////				}
////			}
////		}
////	}
////}
//
////HasOwnProperty : Convert HTTP raw response body to map and check if the body has own field
//func HasOwnProperty(key string, body []byte) (bool, map[string]interface{}) {
//	data := make(map[string]interface{})
//	_ = json.Unmarshal(body, &data)
//	if _, found := data[key].(string); found {
//		return found, data
//	} else {
//		return found, data
//	}
//}
