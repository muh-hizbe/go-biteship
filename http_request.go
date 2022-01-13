package biteship

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpRequest interface {
	Call(method string, url string, secretKey string, body io.Reader, result interface{}) *Error
	//doRequest(request *http.Request, result interface{}) *Error
}

type HttpRequestImpl struct {
	Config *ConfigOption
}

func (client *HttpRequestImpl) Call(method string, url string, secretKey string, body io.Reader, result interface{}) *Error {
	//req := *http.Request
	//var errMarshal, errNewReq error
	//jsonRequest := []byte("")
	//log.Println("method :: ", method)
	//log.Println("url :: ", url)
	//log.Println("secretKey :: ", secretKey)
	//log.Println("Body :: ", body)

	if secretKey == "" {
		return &Error{
			Status:  http.StatusUnauthorized,
			Message: "missing/invalid secret key",
		}
	}

	//isBodyNil := body == nil || (reflect.ValueOf(body).Kind() == reflect.Ptr && reflect.ValueOf(body).IsNil())
	//isBodyNil := body == nil
	//log.Println("Body :: ", isBodyNil)
	//
	//if !isBodyNil {
	//	fmt.Println("masuk")
	//	jsonRequest, errMarshal = json.Marshal(body)
	//	if errMarshal != nil {
	//		log.Println(errMarshal)
	//		return &Error{
	//			Status:   http.StatusInternalServerError,
	//			Message:  "something went wrong",
	//			RawError: errMarshal,
	//		}
	//	}
	//
	//	req, errNewReq = http.NewRequest(method, url, bytes.NewBuffer(jsonRequest))
	//} else {
	//	req, errNewReq = http.NewRequest(method, url, body)
	//}

	req, errNewReq := http.NewRequest(method, url, body)
	if errNewReq != nil {
		return &Error{
			Status:   http.StatusInternalServerError,
			Message:  "Cannot create request",
			RawError: errNewReq.Error(),
		}
	}

	//log.Println(secretKey)
	req.Header.Add("Authorization", secretKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	//log.Println("================ Request ================")
	//log.Println(fmt.Sprintf("%v Request %v %v", req.Method, req.URL, req.Proto))

	//log.Println("result :: ", result)
	return client.doRequest(req, result)
}

func (client *HttpRequestImpl) doRequest(req *http.Request, result interface{}) *Error {
	httpClient := &http.Client{}

	response, errRequest := httpClient.Do(req)
	log.Println("response code :: ", response.StatusCode)
	if errRequest != nil {
		return ErrorGo(errRequest)
	}
	defer response.Body.Close()

	respBody, errRead := ioutil.ReadAll(response.Body)
	if errRead != nil {
		return ErrorGo(errRead)
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return ErrorHttp(response.StatusCode, respBody)
	}

	errUnmarshall := json.Unmarshal(respBody, &result)
	if errUnmarshall != nil {
		return ErrorGo(errUnmarshall)
	}

	return nil
}
