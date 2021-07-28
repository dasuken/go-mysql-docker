package routes

import (
	"github.com/gorilla/mux"
	"github.com/noguchidaisuke/go-mysql-docker/api/middlewares"
	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func Install(router *mux.Router, routeList []*Route) {
	for _, route := range routeList {
		handler := middlewares.Logger(route.Handler)
		router.HandleFunc(route.Path, handler).Methods(route.Method)
	}
}
