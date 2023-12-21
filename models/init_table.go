package models

import "github.com/jinzhu/gorm"

func InitTableLogStream(db *gorm.DB) {
	db.Debug().AutoMigrate(&LogStream{})
}
