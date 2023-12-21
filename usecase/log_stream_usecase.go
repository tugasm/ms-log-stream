package usecase

import (
	"fmt"
	"ms-briapi-log-stream/models"
	"ms-briapi-log-stream/repo"
	"ms-briapi-log-stream/utils"
)

type LogStreamUsecaseStruct struct {
	repoLogStream repo.LogStreamRepoInterface
}

func CreateLogStreamUsecaseImpl(repoLogStream repo.LogStreamRepoInterface) LogStreamUsecaseInterface {
	return LogStreamUsecaseStruct{repoLogStream}
}

type LogStreamUsecaseInterface interface {
	CreateLogStreamData(dataReq models.LogStreamRequest) ([]*models.LogStream, *utils.APIErrors)
}

func (u LogStreamUsecaseStruct) CreateLogStreamData(dataReq models.LogStreamRequest) ([]*models.LogStream, *utils.APIErrors) {
	var dataArray []*models.LogStream
	var data models.LogStream

	// Marshall Data
	// clientheaderrequest, err := json.Marshal(dataReq.ClientHeaderRequest)
	// metadata, err := json.Marshal(dataReq.Metadata)

	// data.Id = dataReq.Id
	data.Date = dataReq.Date
	data.Time = dataReq.Time
	data.Devapps = dataReq.Devapps
	data.Product = dataReq.Product
	data.Xforwarder = dataReq.Xforwarder
	data.PathClient = dataReq.PathClient
	data.PathBackend = dataReq.PathBackend
	data.Method = dataReq.Method
	data.HttpStatus = dataReq.HttpStatus
	data.ClientHeaderRequest = dataReq.ClientHeaderRequest
	data.ClientHeaderResponse = dataReq.ClientHeaderResponse
	data.ClientRequest = dataReq.ClientRequest
	data.ClientResponse = dataReq.ClientResponse
	data.BackendHeaderRequest = dataReq.BackendHeaderRequest
	data.BackendHeaderResponse = dataReq.BackendHeaderResponse
	data.BackendRequest = dataReq.BackendRequest
	data.BackendResponse = dataReq.BackendResponse
	data.ResponseTime = dataReq.ResponseTime
	data.Metadata = dataReq.Metadata
	// data.Metadata = string(metadata)
	// data.ClientHeaderRequest = string(clientheaderrequest)

	// metadatax := models.MetadataDetail{}
	// clientx := models.ClientHeaderRequestDetail{}

	// fmt.Printf("meta : %v\n", dataReq.Metadata)
	// fmt.Printf("meta2 : %v\n", metadatax)

	dataResp, err := u.repoLogStream.CreateLogStreamData(data)

	// fmt.Printf("cek errir: %v\n", dataResp)
	// unmarshal JSON
	// if err := json.Unmarshal([]byte(dataResp.ClientHeaderRequest), &clientx); err != nil {
	// 	return nil, &utils.ErrBindingJSON
	// }
	// if err := json.Unmarshal([]byte(dataResp.Metadata), &metadatax); err != nil {
	// 	return nil, &utils.ErrBindingJSON
	// }

	if err != nil {
		fmt.Printf("[LogStreamUsecaseStruct.CreateLogStreamData] error when CreateLogStreamData with error : %v\n", err)
		return nil, &utils.ErrBindingJSON
	}
	dataArray = append(dataArray, dataResp)

	return dataArray, nil
}
