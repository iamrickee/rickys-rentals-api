package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"iamricky.com/truck-rental/config"
)

type MysqlDB struct{}

func (db MysqlDB) GetConnection() (*sql.DB, error) {
	dbHost := config.Load("DB_HOST")
	dbUser := config.Load("DB_USER")
	dbPassword := config.Load("DB_PASSWORD")
	dbPort := config.Load("DB_PORT")
	dbName := config.Load("DB_NAME")
	conn, err := sql.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?multiStatements=true")
	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)
	return conn, err
}
