package contract

var ServiceCode string
var CaseCode string
var Rc string
var FieldErr string

const (
	ErrInvalidFieldFormat    = "400001"
	ErrInvalidFieldMandatory = "400002"
	ErrBadRequest            = "400003"
	ErrDataNotFound          = "400004"
	ErrFailedUpdate          = "400005"
	ErrCreateData            = "400006"
	ErrFailedDelete          = "400007"
	ErrGeneralError          = "5000000"
	ErrUnexpectedError       = "500001"
)

const (
	InvalidBadRequest     = "00"
	InvalidFieldFormat    = "01"
	InvalidMandatoryField = "02"
	Unauthorized          = "00"
	Conflict              = "00"
	GeneralError          = "00"
	InternalServerError   = "02"
)
