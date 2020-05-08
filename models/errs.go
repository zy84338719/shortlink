package models

import "net/http"

type Result struct {
	Code   string      `json:"code"`
	ErrMsg string      `json:"errMsg"`
	Data   interface{} `json:"data"`
}

type Response struct {
	Status int
	Result Result
}

var (
	ErrorRequestBodyParseFailed = Response{http.StatusBadRequest, Result{Code: "001", ErrMsg: "Request body is not correct"}}
	ErrorNotAuthUser            = Response{http.StatusUnauthorized, Result{Code: "002", ErrMsg: "User authentication failed."}}
	ErrorRedisError             = Response{http.StatusInternalServerError, Result{Code: "003", ErrMsg: "Redis ops failed"}}
	ErrorInternalFaults         = Response{http.StatusInternalServerError, Result{Code: "004", ErrMsg: "Internal service error"}}
	ErrorTooManyRequests        = Response{http.StatusTooManyRequests, Result{Code: "005", ErrMsg: "Too many Request"}}
	ErrorParameterValidate      = Response{http.StatusBadRequest, Result{Code: "006", ErrMsg: "validate parameters failed"}}
	ErrorStorageError           = Response{http.StatusInternalServerError, Result{Code: "007", ErrMsg: "storage error"}}
)
