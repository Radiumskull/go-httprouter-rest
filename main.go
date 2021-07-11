package main

import (
	"backend/database"
	"backend/repositories"
	"database/sql"
	"log"
	"net/http"

	"backend/controllers/auth"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type BaseHandler struct {
	db *sql.DB
}

func main() {
	router := httprouter.New()

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	h := &BaseHandler{db: db}
	h.AuthHandler(router)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func (h *BaseHandler) AuthHandler(router *httprouter.Router) {
	controller := auth.AuthController{
		UserRepo: repositories.NewUserRepo(h.db),
		Encrypt:  bcrypt.GenerateFromPassword,
	}
	router.POST("/auth/login", controller.Login)
	router.POST("/auth/register", controller.Register)
}
