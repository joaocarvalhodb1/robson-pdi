package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"poc-pessoa/model"

	"github.com/google/uuid"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// return error json structure
	}

	var person *model.Person
	if err = json.Unmarshal(body, &person); err != nil {
		// return error json structure
		return
	}

	// logica de gravar no banco
	person.ID = uuid.New().String()

	// depois de gravado no banco, eu tenho o retorno

	pJson, _ := json.Marshal(person)
	w.Write(pJson)
}
