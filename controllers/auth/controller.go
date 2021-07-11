package auth

import (
	"backend/models"
	"backend/repositories"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController struct {
	UserRepo *repositories.UserRepo
	Encrypt  func(password []byte, cost int) ([]byte, error)
}

func (h *AuthController) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.UserRepo.Save(&models.User{Name: "Aritra"})
	fmt.Fprintf(w, "Logged In!/n")
}

func (h *AuthController) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	hashPass, err := h.Encrypt([]byte("Hello"), 10)
	if err == nil {
		log.Fatal("Error Hashing the Password!")
	}

	fmt.Fprintf(w, string(hashPass))
}
