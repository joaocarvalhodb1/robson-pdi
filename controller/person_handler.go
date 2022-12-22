package controller

import "net/http"

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetPerson(w, r)
	}
	if r.Method == http.MethodPost {
		CreatePerson(w, r)
	}
}
