package models

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string
	Err     error
	Data    interface{}
}

func (r *Response) HttpResponse(w http.ResponseWriter) {
	enconder := json.NewEncoder(w)
	enconder.Encode(r)
}
