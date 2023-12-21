package repo

import (
	"fmt"
	"ms-briapi-log-stream/models"

	"github.com/jinzhu/gorm"
)

type LogStreamRepoStruct struct {
	db *gorm.DB
}

func CreateLogStreamRepoImpl(db *gorm.DB) LogStreamRepoInterface {
	return &LogStreamRepoStruct{db}
}

type LogStreamRepoInterface interface {
	CreateLogStreamData(data models.LogStream) (*models.LogStream, error)
}

func (i *LogStreamRepoStruct) CreateLogStreamData(data models.LogStream) (*models.LogStream, error) {
	err := i.db.Debug().Create(&data).Error
	if err != nil {
		return nil, fmt.Errorf("[LogStreamRepoStruct.CreateLogStreamData] Error when query save data with : %v", err)
	}

	return &data, nil
}
