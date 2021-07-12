package middlewares

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func IsAuthenticated(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		logger := log.Default()
		token := r.Header.Get("Authorization")

		logger.Output(1, token)

		h(w, r, ps)
	}
}
