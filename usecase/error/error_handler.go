package error

import (
	"encoding/json"
	"errors"
	"fmt"
	"ms-briapi-log-stream/models"
	"ms-briapi-log-stream/models/contract"
	"ms-briapi-log-stream/usecase"
	"ms-briapi-log-stream/utils"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/Saucon/errcntrct"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/lib/pq"
)

type errorHandlerUsecase struct {
}

func NewErrorHandlerUsecase() usecase.ErrorHandlerUsecase {
	return &errorHandlerUsecase{}
}

func (eh *errorHandlerUsecase) ResponseError(A interface{}) (int, interface{}) {
	var T interface{}
	var fieldNameErr string
	var serviceCode string

	if A.(*gin.Error).Meta != nil {
		fieldNameErr = A.(*gin.Error).Meta.(models.ErrMeta).FieldErr
		serviceCode = A.(*gin.Error).Meta.(models.ErrMeta).ServiceCode
	}

	// Check A is a correct error type and assign to T
	if A.(*gin.Error).Err != nil {
		T = A.(*gin.Error).Err
	}

	switch T.(type) {
	case error:
		if _, ok := T.(*pq.Error); ok {
			switch T.(*pq.Error).Code.Name() {
			case "unique_violation":
				return errcntrct.ErrorMessage(http.StatusBadRequest, "", errors.New(contract.ErrUnexpectedError))
			}
		}

		switch T.(error).Error() {
		case contract.ErrInvalidFieldFormat:
			return responseErrorAdapter(T.(error), http.StatusBadRequest, contract.ErrInvalidFieldFormat, serviceCode, fieldNameErr)
		case contract.ErrInvalidFieldMandatory:
			return responseErrorAdapter(T.(error), http.StatusBadRequest, contract.ErrInvalidFieldMandatory, serviceCode, fieldNameErr)
		case contract.ErrBadRequest:
			contract.CaseCode = "00"
			return responseErrorAdapter(T.(error), http.StatusBadRequest, "", serviceCode, "")
		case contract.ErrDataNotFound:
			contract.CaseCode = "16"
			return responseErrorAdapter(T.(error), http.StatusNotFound, "", serviceCode, "")
		case contract.ErrFailedUpdate:
			contract.CaseCode = "01"
			return responseErrorAdapter(T.(error), http.StatusBadRequest, "", serviceCode, "")
		case contract.ErrCreateData:
			contract.CaseCode = "01"
			return responseErrorAdapter(T.(error), http.StatusNotFound, "", serviceCode, "")
		case contract.ErrFailedDelete:
			contract.CaseCode = "01"
			return responseErrorAdapter(T.(error), http.StatusNotFound, "", serviceCode, "")
		default:
			contract.CaseCode = "01"
			return responseErrorAdapter(errors.New(contract.ErrUnexpectedError), http.StatusInternalServerError, "", serviceCode, "")
		}
	}

	return responseErrorAdapter(T.(error), http.StatusInternalServerError, "", serviceCode, "")
}

var Case string

func responseErrorAdapter(errHttpStatus interface{}, httpStatusCode int, ctr string, serviceCode string, fieldErr string) (int, models.ResponseCustomErr) {
	_, errData := errcntrct.ErrorMessage(httpStatusCode, "", errHttpStatus)
	var resp models.ResponseCustomErr
	errCase := strconv.Itoa(httpStatusCode)
	resp.ResponseCode = errCase + contract.ServiceCode + contract.CaseCode
	if strings.Contains(contract.FieldErr, " ") {
		resp.ResponseMessage = fmt.Sprintf(errData.Msg, contract.FieldErr)
	} else if ctr == "400001" || ctr == "400002" {
		resp.ResponseMessage = fmt.Sprintf(errData.Msg, utils.LowerCamelCase(fieldErr))
	} else {
		resp.ResponseMessage = fmt.Sprintf(errData.Msg)
	}
	return httpStatusCode, resp
}

func (eh *errorHandlerUsecase) ValidateRequest(T interface{}) (string, error) {
	v := validator.New()
	var errArr error
	var field string
	switch T.(type) {
	case models.LogStreamRequest:
		err := v.Struct(T)
		for _, e := range err.(validator.ValidationErrors) {
			if e.Value() != "" {
				switch e.Tag() {
				case "numeric", "max", "email", "lt", "gte", "len", "alpha", "min":
					contract.CaseCode = contract.InvalidFieldFormat
					field = e.Field()
					errArr = errors.New(contract.ErrInvalidFieldFormat)
				}
				break
			} else {
				switch e.Tag() {
				case "required":
					contract.CaseCode = contract.InvalidMandatoryField
					field = e.Field()
					errArr = errors.New(contract.ErrInvalidFieldMandatory)
				}
				break
			}
		}

		if errArr != nil {
			return field, errArr
		}

		return "", nil
	default:
		return "", errors.New(contract.ErrUnexpectedError)
	}

}

func (b *errorHandlerUsecase) ValidatorJsonName(data interface{}, model interface{}, status string) (string, error) {
	var modelMap map[string]interface{}
	var dataMap map[string]interface{}
	isTrue := true
	jsonModel, _ := json.Marshal(model)
	json.Unmarshal(jsonModel, &modelMap)

	jsonModelMap, _ := json.Marshal(data)
	json.Unmarshal(jsonModelMap, &dataMap)

	dt := reflect.ValueOf(modelMap)
	dt2 := reflect.ValueOf(dataMap)

	for _, key := range dt.MapKeys() {
		strct := dt.MapIndex(key)
		res := fmt.Sprintf("%s", strct.Interface())
		for _, key2 := range dt2.MapKeys() {
			strct2 := dt2.MapIndex(key2)
			res2 := fmt.Sprintf("%s", strct2.Interface())
			fmt.Println(res, res2, "lop")
			if key.String() != key2.String() {
				isTrue = false
			} else if key.String() == key2.String() {
				isTrue = true
				break
			}
		}
		if !isTrue {
			contract.CaseCode = "02"
			if status == "optionalreq" {
				return key.String(), nil
			}
			return key.String(), errors.New(contract.ErrInvalidFieldMandatory)
		} else if strings.Contains(res, "map[") {
			contract.CaseCode = "02"
			field, err := b.ValidatorJsonName(dataMap[key.String()], modelMap[key.String()], status)
			if status == "optionalreq" {
				return field, nil
			}
			if err != nil {
				return field, err
			}
		}
	}

	return "", nil
}

func Search(str string, key []reflect.Value) bool {
	for i := 0; i < len(key); i++ {
		if str == key[i].String() {
			return true
			break
		}
	}
	return false
}
