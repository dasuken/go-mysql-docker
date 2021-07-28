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
		{
			Path: "/categories",
			Method: http.MethodGet,
			Handler: c.GetAllCategories,
		},
		{
			Path: "/categories/{category_id}",
			Method: http.MethodGet,
			Handler: c.GetCategory,
		},
		{
			Path: "/categories/{category_id}",
			Method: http.MethodPut,
			Handler: c.PutCategory,
		},
		{
			Path: "/categories/{category_id}",
			Method: http.MethodDelete,
			Handler: c.DeleteCategory,
		},
	}
}