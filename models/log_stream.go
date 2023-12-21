package models

// import "github.com/gofrs/uuid"

type ErrMeta struct {
	ServiceCode string
	FieldErr    string
}

type ResponseCustomErr struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

const (
	CreateLogServiceCode = "82"
)

type LogStream struct {
	// gorm.Model `json:"-"`
	Id                    uint   `gorm:"primary_key"`
	Date                  string `gorm:"type:varchar(25)" json:"date"`
	Time                  string `gorm:"type:timestamp" json:"time"`
	Devapps               string `gorm:"type:varchar(100)" json:"devapps"`
	Product               string `gorm:"type:varchar(100)" json:"product"`
	Xforwarder            string `gorm:"type:varchar(20)" json:"xforwarder"`
	PathClient            string `gorm:"type:varchar(100)" json:"pathClient"`
	PathBackend           string `gorm:"type:varchar(100)" json:"pathBackend"`
	Method                string `gorm:"type:varchar(7)" json:"method"`
	HttpStatus            string `gorm:"type:varchar(3)" json:"httpStatus"`
	ClientHeaderRequest   string `json:"clientHeaderRequest"`
	ClientHeaderResponse  string `json:"clientHeaderResponse"`
	ClientRequest         string `json:"clientRequest"`
	ClientResponse        string `json:"clientResponse"`
	BackendHeaderRequest  string `json:"backendHeaderRequest"`
	BackendHeaderResponse string `json:"backendHeaderResponse"`
	BackendRequest        string `json:"backendRequest"`
	BackendResponse       string `json:"backendResponse"`
	ResponseTime          string `gorm:"type:varchar(10)" json:"responseTime"`
	Metadata              string `json:"metadata"`
}

type LogStreamRequest struct {
	Date                  string `gorm:"type:varchar(25)" json:"date"`
	Time                  string `gorm:"type:timestamp" json:"time"`
	Devapps               string `gorm:"type:varchar(100)" json:"devapps"`
	Product               string `gorm:"type:varchar(100)" json:"product"`
	Xforwarder            string `gorm:"type:varchar(20)" json:"xforwarder"`
	PathClient            string `gorm:"type:varchar(100)" json:"pathClient"`
	PathBackend           string `gorm:"type:varchar(100)" json:"pathBackend"`
	Method                string `gorm:"type:varchar(7)" json:"method"`
	HttpStatus            string `gorm:"type:varchar(3)" json:"httpStatus" validate:"min=1,max=3"`
	ClientHeaderRequest   string `json:"clientHeaderRequest"`
	ClientHeaderResponse  string `json:"clientHeaderResponse"`
	ClientRequest         string `json:"clientRequest"`
	ClientResponse        string `json:"clientResponse"`
	BackendHeaderRequest  string `json:"backendHeaderRequest"`
	BackendHeaderResponse string `json:"backendHeaderResponse"`
	BackendRequest        string `json:"backendRequest"`
	BackendResponse       string `json:"backendResponse"`
	ResponseTime          string `gorm:"type:varchar(10)" json:"responseTime"`
	Metadata              string `json:"metadata"`

	// ClientHeaderRequest  string `json:"clientHeaderRequest" gorm:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}
type ClientHeaderRequestDetail struct {
	// PartnerId string `json:"partnerId"`
}

type MetadataDetail struct {
	Amount string `json:"amount"`
	Fee    string `json:"fee"`
}

// TableName ..
func (s LogStream) TableName() string {
	return "log"
}
