package utils

import (
	"gorm.io/gorm"
	"log"
)

func LogDBError(result *gorm.DB) {
	defer recoverFromDBError()
	if result.Error != nil {
		panic(result.Error)
	}
}

func LogError(err error) {
	// logs.GetLogger("ORM").Println("dev mode")
	if err != nil {
		log.Fatal(err)
	}
}

func recoverFromDBError() {
	if r := recover(); r != nil {
		log.Print("recovered from ", r)
	}
}
