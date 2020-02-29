package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type User struct {
	Email    string
	Password string
}

type RegisterHandler struct {
	db *sql.DB
}

func NewRegisterHandler(db *sql.DB) *RegisterHandler {
	return &RegisterHandler{db: db}
}

func (h RegisterHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if checkRequestMethod(*req, "POST") {
		var person User

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(req.Body).Decode(&person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//Здесь буду регать юзера
		//TODO Register User
		h.db.Exec("INSERT INTO users", person.Email, person.Password)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("This route must run only by POST request"))
		if err != nil {
			panic(err)
		}
	}
}

func checkRequestMethod(req http.Request, need string) bool {
	if req.Method != need {
		return false
	}
	return true
}
