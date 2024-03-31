package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"iamricky.com/truck-rental/config"
)

type MysqlDB struct{}

func (db MysqlDB) Execute(query string) error {
	fmt.Println("hello MySQL")
	return nil
}

func (db MysqlDB) GetConnection() (*sql.DB, error) {
	dbHost := config.Load("DB_HOST")
	dbUser := config.Load("DB_USER")
	dbPassword := config.Load("DB_PASSWORD")
	dbPort := config.Load("DB_PORT")
	dbName := config.Load("DB_NAME")
	conn, err := sql.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?multiStatements=true")
	return conn, err
}
