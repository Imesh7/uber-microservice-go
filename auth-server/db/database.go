package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnDataBase struct {
	DB *gorm.DB
}

var ConnectedDataBase *ConnDataBase

func ConnectDatabase() (*ConnDataBase, error) {
	dsn := "host=auth-server-db user=postgres password=1234 dbname=authdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	ConnectedDataBase = &ConnDataBase{
		DB: db,
	}
	return ConnectedDataBase, nil
}
