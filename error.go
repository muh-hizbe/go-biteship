package biteship

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Status         int          `json:"status,omitempty"`
	ErrorCode      string       `json:"error_code,omitempty"`
	Message        string       `json:"message,omitempty"`
	RawError       error        `json:"raw_error"`
	RawApiResponse *ApiResponse `json:"raw_api_response"`
}

//func (e *Error) Error() string {
//	return e.Message
//}

func (e *Error) GetStatus() int {
	return e.Status
}

func ErrorGo(err error) *Error {
	return &Error{
		Status:    http.StatusInternalServerError,
		ErrorCode: "Error Go",
		Message:   err.Error(),
	}
}

func ErrorHttp(status int, respBody []byte) *Error {
	var httpError *Error
	if err := json.Unmarshal(respBody, &httpError); err != nil {
		return ErrorGo(err)
	}
	httpError.Status = status

	return httpError
}
