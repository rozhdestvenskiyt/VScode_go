package handler

import "github.com/julienschmidt/httprouter"

type handler interface {
	Register(router *httprouter.Router)
}