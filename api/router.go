package api

import (
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	for _, route := range routes {
		var curr Route = route
		router.Handle(curr.Method, curr.Path, curr.Handler)
	}
	return router
}
