package auth

import (
	"backend/models"
	"backend/repositories"
	"backend/utils"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController struct {
	UserRepo    *repositories.UserRepo
	Encrypt     func(password []byte, cost int) ([]byte, error)
	CompareHash func(hashedPassword []byte, password []byte) error
}

type AuthRequest struct {
	Username string
	Password string
}

type AuthResponse struct {
	Username string
	Token    string
}

func (h *AuthController) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var body AuthRequest
	decoder := json.NewDecoder(r.Body)
	jsonErr := decoder.Decode(&body)

	if jsonErr != nil {
		utils.ErrorResponse(w, jsonErr)
		return
	}

	user, dbErr := h.UserRepo.FindByUsername(body.Username)
	if dbErr != nil {
		utils.ErrorResponse(w, dbErr)
		return
	}

	hashErr := h.CompareHash([]byte(user.Hash), []byte(body.Password))

	if hashErr != nil {
		utils.ErrorResponse(w, errors.New("wrong credentials"))
		return
	} else {
		token, err := utils.CreateToken(user)
		if err != nil {
			utils.ErrorResponse(w, err)
			return
		}

		response := &AuthResponse{
			Username: user.Username,
			Token:    token,
		}

		utils.SuccessResponseWithMessage(w, response, "Successfully Logged In")
		return
	}

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
