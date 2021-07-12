package main

import (
	"backend/controllers"
	"backend/database"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {
	router := httprouter.New()

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	h := &controllers.BaseHandler{DB: db}

	h.AuthHandler(router)
	h.UserHandler(router)

	log.Fatal(http.ListenAndServe(":8000", router))
}
