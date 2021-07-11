package user

import (
	"backend/repositories"
	"backend/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	UserRepo *repositories.UserRepo
}

func (r *UserController) GetUser(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	userId, _ := strconv.ParseInt(params.ByName("id"), 10, 4)
	user, err := r.UserRepo.FindByID(int(userId))
	if err != nil {
		utils.ErrorResponse(w, err)
		return
	}

	utils.SuccessResponse(w, user)
}
