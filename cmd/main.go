package main

import (
	"database/sql"
	"log"
	"net/http"
	"register/api"
	"register/database"
)

func register(db *sql.DB) http.Handler {
	return api.NewRegisterHandler(db)
}

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	http.Handle("/register", register(db))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
