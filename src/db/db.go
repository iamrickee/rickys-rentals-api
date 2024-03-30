package db

import (
	"database/sql"
	"errors"

	"iamricky.com/truck-rental/config"
)

type DB interface {
	Execute(query string) error
	GetConnection() (*sql.DB, error)
}

func Exec(query string) error {
	db := getDriver()
	if db != nil {
		return db.Execute(query)
	}
	return errors.New("invalid database type")
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
