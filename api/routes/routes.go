package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func Install(router *mux.Router, routeList []*Route) {
	for _, route := range routeList {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
}
