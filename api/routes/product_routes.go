package routes

import (
	"github.com/noguchidaisuke/go-mysql-docker/api/controllers"
	"net/http"
)

func NewProductRoutes(c controllers.ProductsController) []*Route {
	return []*Route {
		{
			Path: "/products",
			Method: http.MethodPost,
			Handler: c.PostProduct,
		},
	}
}