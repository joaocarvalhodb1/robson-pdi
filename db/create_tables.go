package db

import (
	"database/sql"
	"errors"
)

func CreateTables(connection *sql.DB) error {
	// Create person
	_, err := connection.Exec("create table pessoa(id varchar(36), name varchar(60), person_type varchar(10), primary key (id))")
	if err != nil {
		return errors.New("error creating table pessoa: " + err.Error())
	}
	return nil
}
