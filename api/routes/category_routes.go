package routes

import (
	"github.com/noguchidaisuke/go-mysql-docker/api/controllers"
	"net/http"
)

func NewCategoryRoutes(c controllers.CategoriesController) []*Route {
	return []*Route {
		{
			Path: "/categories",
			Method: http.MethodPost,
			Handler: c.PostCategory,
		},
	}
}