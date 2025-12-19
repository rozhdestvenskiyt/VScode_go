package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	users string = "/users/"
	user  string = "/users/:uuid"
)

type Handler_user struct {
}

func (h *Handler_user) Register(router *httprouter.Router) {
	router.GET(users, h.GET)
}

func (h *Handler_user) GET(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("This main page"))
}
