package db

import (
	"database/sql"
	"errors"
)

func MysqlConnecttion(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.New("error connecting to mysql: " + err.Error())
	}
	return db, nil
}
