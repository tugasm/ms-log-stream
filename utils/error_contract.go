package utils

import (
	"net/http"
)

var (
	ErrBindingJSON      = APIErrors{StatusCode: http.StatusBadRequest, ResponseCode: "0001", ResponseDescription: "Failed create log"}
	ErrDuplicateLogData = APIErrors{StatusCode: http.StatusBadRequest, ResponseCode: "0002", ResponseDescription: "log is already exist"}
	ErrInsertLogData    = APIErrors{StatusCode: http.StatusBadRequest, ResponseCode: "0003", ResponseDescription: "Failed"}
	ErrorGetLogData     = APIErrors{StatusCode: http.StatusBadRequest, ResponseCode: "0004", ResponseDescription: "log does not exist"}
	ErrorDeleteLogData  = APIErrors{StatusCode: http.StatusBadRequest, ResponseCode: "0005", ResponseDescription: "error delete log"}
)
