package controllers

import (
	"fmt"
	"ms-briapi-log-stream/models"
	"ms-briapi-log-stream/models/contract"
	"ms-briapi-log-stream/usecase"
	"ms-briapi-log-stream/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type LogStreamController struct {
	usecaseLogStream usecase.LogStreamUsecaseInterface
	errr             usecase.ErrorHandlerUsecase
}

func CreateLogStreamController(router *gin.RouterGroup, lsu usecase.LogStreamUsecaseInterface, errr usecase.ErrorHandlerUsecase) {
	handler := LogStreamController{lsu, errr}

	router.POST("/log", handler.InsertLogStream)
}

func (i *LogStreamController) InsertLogStream(c *gin.Context) {
	var dataReq models.LogStreamRequest
	var reqBindName models.LogStream
	v := validator.New()

	errr := v.Struct(&dataReq)
	var reqBodyTemp map[string]interface{}
	contract.ServiceCode = models.CreateLogServiceCode
	field, err := i.errr.ValidatorJsonName(reqBodyTemp, reqBindName, "optionalreq")
	if field, errr = i.errr.ValidateRequest(dataReq); errr != nil {
		c.Error(errr).SetMeta(models.ErrMeta{
			ServiceCode: contract.ServiceCode,
			FieldErr:    field,
		})
		c.Abort()
		return
	}

	err2 := c.ShouldBindJSON(&dataReq)
	if err2 != nil {
		utils.NewErrorMessage(c, &utils.ErrBindingJSON)
		fmt.Printf("[LogStreamController.InsertLogStream] error when encode data enkripsi : %v\n", err2)
		return
	}

	dataResp, errs := i.usecaseLogStream.CreateLogStreamData(dataReq)
	if errs != nil {
		utils.NewErrorMessage(c, errs)
		fmt.Printf("[LogStreamController.InsertLogStream] error when encode create log  with error : %v\n", err)
		return
	}

	utils.SuccessData(c, http.StatusOK, dataResp)
}
