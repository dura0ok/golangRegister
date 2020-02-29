package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"register/database"
)

type User struct {
	Email string
	Password string
}

func checkRequestMethod(req http.Request, need string) bool {
	if req.Method != need{
		return false
	}
	return true
}

func register(w http.ResponseWriter, req *http.Request, db *sql.DB){
	if checkRequestMethod(*req, "POST"){

		var person User

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(req.Body).Decode(&person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//TODO Register User
		//Здесь буду регать юзера


	}else{
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("This route must run only by POST request"))
		if err != nil {
			panic(err)
		}
	}
}


func main() {
	db, err := database.Connect()
	if err != nil{
		log.Fatalln(err)
	}
	http.HandleFunc("/register", register())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
