package config

import (
	"fmt"
	"log"
	"ms-briapi-log-stream/models"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	if DB != nil {
		return DB
	}

	var err error
	dbConfig := Config.DB
	fmt.Printf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConfig.UserDB, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	if dbConfig.Adapter == "postgres" {
		DB, err = gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbConfig.UserDB, dbConfig.Password, dbConfig.Host, dbConfig.Name))
		log.Println("Connected to Database Local postgressql")
	}

	if err != nil {
		log.Println("[Driver.ConnectDB] error when connect to database")
		log.Fatal("[Driver.ConnectDB] error when connect to database")
	} else {
		log.Println("SUCCES CONNECT TO DATABASE")
	}

	go doEvery(6*time.Minute, pingDb, DB)

	// Database Pooling
	DB.DB().SetMaxIdleConns(20)
	DB.DB().SetMaxOpenConns(200)
	DB.DB().SetConnMaxLifetime(45 * time.Second)

	models.InitTableLogStream(DB)

	return DB
}

func doEvery(d time.Duration, f func(*gorm.DB), y *gorm.DB) {
	for _ = range time.Tick(d) {
		f(y)
	}
}

func pingDb(db *gorm.DB) {
	log.Println("PING CONNECTION")
	err := db.DB().Ping()
	if err != nil {
		log.Println("PING CONNECTION FAILURE")
		return
	}
}
