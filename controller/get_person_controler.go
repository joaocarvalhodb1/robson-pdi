package controller

import (
	"encoding/json"
	"net/http"
	"poc-pessoa/model"
)

func GetPerson(w http.ResponseWriter, r *http.Request) {
	person := model.Person{
		ID:         "1",
		Name:       "ROBSON",
		PersonType: "CUSTOMER",
	}
	pJson, _ := json.Marshal(person)
	w.Write(pJson)
}
