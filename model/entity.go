package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type Person struct {
	ID         string `json:"id"`
	PersonType string `json:"person_type"`
	Name       string `json:"name"`
}

func CreatePersonDB(conn *sql.DB, person *Person) (*Person, error) {
	newPerson := &Person{
		ID:         uuid.New().String(),
		Name:       person.Name,
		PersonType: person.PersonType,
	}

	_, err := conn.Exec("INSERT INTO pessoa(ID, Name, PersonType) VALUES(?, ?, ?)", person.ID, person.Name, person.PersonType)
	if err != nil {
		return nil, err
	}

	return newPerson, nil
}
