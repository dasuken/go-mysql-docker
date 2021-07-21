package database

import (
	"fmt"
	"github.com/noguchidaisuke/go-mysql-docker/api/database/config"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dns := config.BuildDNS()
	fmt.Println("dns: ", dns)

	database, err := gorm.Open(dns)
	if err != nil {
		panic("failed to connect database")
	}

	return database
}