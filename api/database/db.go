package database

import (
	"github.com/noguchidaisuke/go-mysql-docker/api/database/config"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	database, err := gorm.Open(config.BuildDNS())
	if err != nil {
		panic("failed to connect database")
	}

	return database
}