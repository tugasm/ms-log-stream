package usecase

import (
	"ms-briapi-log-stream/models"
	"ms-briapi-log-stream/utils"
)

type LogStreamUsecase interface {
	CreateLogStreamData(dataReq models.LogStreamRequest) ([]*models.LogStream, *utils.APIErrors)
}

type ErrorHandlerUsecase interface {
	ResponseError(A interface{}) (int, interface{})
	ValidateRequest(error interface{}) (string, error)
	ValidatorJsonName(data interface{}, model interface{}, status string) (string, error)
}
