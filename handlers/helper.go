package handlers

import (
	"encoding/json"
	"julo/constants"
	"net/http"
)

type ResponseHTTP struct {
	StatusCode int
	Response   ResponseData
}

type ResponseData struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

// Response is the new type for define all of the response from service
type Response interface{}

var (
	ErrRespServiceMaintance = ResponseHTTP{
		StatusCode: http.StatusServiceUnavailable,
		Response:   ResponseData{Status: constants.Fail}}
	ErrRespUnauthorize = ResponseHTTP{
		StatusCode: http.StatusUnauthorized,
		Response:   ResponseData{Status: constants.Fail}}
	ErrRespAuthInvalid = ResponseHTTP{
		StatusCode: http.StatusUnauthorized,
		Response:   ResponseData{Status: constants.Fail}}
	ErrRespBadRequest = ResponseHTTP{
		StatusCode: http.StatusBadRequest,
		Response:   ResponseData{Status: constants.Fail}}
	ErrRespInternalServer = ResponseHTTP{
		StatusCode: http.StatusServiceUnavailable,
		Response:   ResponseData{Status: constants.Fail}}
)

func writeResponse(res http.ResponseWriter, resp Response, code int, err error) {
	res.Header().Set("Content-Type", "application/json")

	if err != nil {
		errJSON := NewError("901", "404", err.Error())
		respErr, _ := json.Marshal(errJSON)
		res.WriteHeader(code)
		res.Write(respErr)
		return
	}

	r, _ := json.Marshal(resp)

	res.WriteHeader(code)
	res.Write(r)
	return
}

type Error struct {
	Id     string `json:"id"`
	Status string `json:"status"`
	Title  string `json:"title"`
}

func NewError(id string, status string, title string) *Error {
	return &Error{
		Id:     id,
		Status: status,
		Title:  title,
	}
}
