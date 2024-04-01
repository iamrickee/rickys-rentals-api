package db

import (
	"database/sql"
	"errors"

	"iamricky.com/truck-rental/config"
)

type DB interface {
	GetConnection() (*sql.DB, error)
}

func GetConn() (*sql.DB, error) {
	db := getDriver()
	if db != nil {
		return db.GetConnection()
	}
	return nil, errors.New("invalid database type")
}

func getDriver() DB {
	var db DB
	dbDriver := config.Load("DB_TYPE")
	switch dbDriver {
	case "mysql":
		db = MysqlDB{}
	}
	return db
}
