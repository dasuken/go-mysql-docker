package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func BuildDNS() gorm.Dialector {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3307"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "root"
	}

	pass := os.Getenv("DB_PASS")
	if pass == "" {
		pass = ""
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "supertest"
	}

	return mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		 user, pass, host, port, dbname))
}