package auth

import (
	"backend/models"
	"backend/repositories"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController struct {
	UserRepo *repositories.UserRepo
	Encrypt  func(password []byte, cost int) ([]byte, error)
}

type AuthRequest struct {
	Username string
	Password string
}

func (h *AuthController) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	h.UserRepo.Save(&models.User{Username: "Aritra"})
	fmt.Fprintf(w, "Logged In!/n")
}

func (h *AuthController) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var body AuthRequest
	jsonerr := decoder.Decode(&body)

	if jsonerr != nil {
		fmt.Fprintf(w, jsonerr.Error())
		return
	}
	hashedPass, err := h.Encrypt([]byte(body.Password), 10)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	user := &models.User{
		Username: body.Username,
		Hash:     string(hashedPass),
	}

	dberr := h.UserRepo.Save(user)
	if dberr != nil {
		fmt.Fprintf(w, dberr.Error())
		return
	}

	fmt.Fprintf(w, "User Created")

}
