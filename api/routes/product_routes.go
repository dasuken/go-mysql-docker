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
		{
			Path: "/products",
			Method: http.MethodGet,
			Handler: c.GetAllProducts,
		},
		{
			Path: "/products/{product_id}",
			Method: http.MethodGet,
			Handler: c.GetProduct,
		},
		{
			Path: "/products/{product_id}",
			Method: http.MethodPut,
			Handler: c.PutProduct,
		},
		{
			Path: "/products/{product_id}",
			Method: http.MethodDelete,
			Handler: c.DeleteProduct,
		},
	}
}