package auth

import (
	"backend/models"
	"backend/repositories"
	"backend/utils"
	"encoding/json"
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

func (h *AuthController) Login(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	user := &models.User{Username: "Aritra", Hash: "12345"}
	// h.UserRepo.Save()
	utils.SuccessResponse(w, user)
}

func (h *AuthController) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var body AuthRequest
	decoder := json.NewDecoder(r.Body)
	jsonerr := decoder.Decode(&body)

	if jsonerr != nil {
		utils.ErrorResponse(w, jsonerr)
		return
	}

	hashedPass, err := h.Encrypt([]byte(body.Password), 10)
	if err != nil {
		utils.ErrorResponse(w, jsonerr)
		return
	}

	user := &models.User{
		Username: body.Username,
		Hash:     string(hashedPass),
	}

	dberr := h.UserRepo.Save(user)
	if dberr != nil {
		utils.ErrorResponse(w, dberr)
		return
	}

	utils.SuccessResponse(w, user)

}
