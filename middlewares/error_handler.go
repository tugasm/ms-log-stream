package middlewares

import (
	"ms-briapi-log-stream/models/contract"
	"ms-briapi-log-stream/usecase"

	"github.com/gin-gonic/gin"
)

type ErrorHandler struct {
	ErrorHandlerUsecase usecase.ErrorHandlerUsecase
}

func NewErrorHandler(r *gin.RouterGroup, ehus usecase.ErrorHandlerUsecase) {
	handler := &ErrorHandler{
		ErrorHandlerUsecase: ehus,
	}

	r.Use(handler.errorHandler)
}

func (eh *ErrorHandler) errorHandler(c *gin.Context) {
	c.Next()
	errorToPrint := c.Errors.Last()
	if errorToPrint != nil {
		c.JSON(eh.ErrorHandlerUsecase.ResponseError(errorToPrint))
		c.Abort()
		return
	}
}

func TraceHeader(c *gin.Context) map[string]string {
	traceHeader := map[string]string{
		contract.HeaderParentReqID: c.GetHeader(contract.HeaderParentReqID),
		contract.HeaderReqID:       c.GetHeader(contract.HeaderReqID),
	}
	return traceHeader
}
