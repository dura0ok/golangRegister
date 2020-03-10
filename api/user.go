package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

//HashPassword generate hash from string to store to database
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//User struct which store json data from post(email, password)
type User struct {
	Email    string
	Password string
}

// RegisterHandler allow me add db to http handler
type RegisterHandler struct {
	db *sql.DB
}

// NewRegisterHandler create http handle with db connect
func NewRegisterHandler(db *sql.DB) *RegisterHandler {
	return &RegisterHandler{db: db}
}

func (h RegisterHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if checkRequestMethod(*req, "POST") {
		var person User

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		if err := json.NewDecoder(req.Body).Decode(&person); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(person)

		password, err := HashPassword(person.Password)
		if err != nil {
			panic(err)
		}
		result, err := h.db.Exec("INSERT INTO users (email, password) VALUES(?, ?)", person.Email, password)
		fmt.Println(result, err)
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
