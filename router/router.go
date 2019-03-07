package router

import (
	"net/http"

	"github.com/septianw/bara/handler"

	"github.com/gorilla/pat"
)

func Router() *pat.Router {
	router := pat.New()
	router.NotFoundHandler = http.HandlerFunc(handler.NotFound)
	router.Get("/v1/version", handler.Version)
	router.Post("/v1/user", handler.PostUser)

	return router
}
