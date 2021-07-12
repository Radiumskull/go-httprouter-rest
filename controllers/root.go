package controllers

import (
	"backend/controllers/auth"
	"backend/controllers/middlewares"
	"backend/controllers/user"
	"backend/repositories"
	"database/sql"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type BaseHandler struct {
	DB *sql.DB
}

type Route struct {
	Method  string
	Path    string
	Handler httprouter.Handle
}

func NewRoute(method string, path string, handler httprouter.Handle) Route {
	return Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}

func (h *BaseHandler) AuthHandler(router *httprouter.Router) {
	controller := auth.AuthController{
		UserRepo:    repositories.NewUserRepo(h.DB),
		Encrypt:     bcrypt.GenerateFromPassword,
		CompareHash: bcrypt.CompareHashAndPassword,
	}

	routes := []Route{
		NewRoute("POST", "/auth/login", controller.Login),
		NewRoute("POST", "/auth/register", controller.Register),
	}

	for _, route := range routes {
		router.Handle(route.Method, route.Path, route.Handler)
	}
}

func (h *BaseHandler) UserHandler(router *httprouter.Router) {
	controller := user.UserController{
		UserRepo: repositories.NewUserRepo(h.DB),
	}

	routes := []Route{
		NewRoute("GET", "/user/:id", controller.GetUser),
	}

	for _, route := range routes {
		router.Handle(route.Method, route.Path, middlewares.IsAuthenticated(route.Handler, controller.UserRepo))
	}
}
