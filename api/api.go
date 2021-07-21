package api

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/noguchidaisuke/go-mysql-docker/api/controllers"
	"github.com/noguchidaisuke/go-mysql-docker/api/database"
	"github.com/noguchidaisuke/go-mysql-docker/api/models"
	"github.com/noguchidaisuke/go-mysql-docker/api/repository"
	"github.com/noguchidaisuke/go-mysql-docker/api/routes"
	"log"
	"net/http"
)

var (
	port = flag.Int("p", 5000, "api server address")
)

func Run() {
	flag.Parse()

	conn := database.Connect()
	err := createSuperTestTables()
	if err != nil {
		log.Fatalf("DB Migration Error: %v", err)
	}

	fmt.Println("Database connected...")

	// repo
	productsRepo := repository.NewProductRepository(conn)
	categoriesRepo := repository.NewCategoriesRepository(conn)

	// controller
	productsCon := controllers.NewProductsController(productsRepo)
	categoriesCon := controllers.NewCategoriesController(categoriesRepo)

	// route
	var allRoutes []*routes.Route
	allRoutes = routes.NewProductRoutes(productsCon)
	allRoutes = append(allRoutes, routes.NewCategoryRoutes(categoriesCon)...)

	router := mux.NewRouter()
	routes.Install(router, allRoutes)

	fmt.Println("API Listening on", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}

func createSuperTestTables() error {
	conn := database.Connect()
	return conn.AutoMigrate(&models.Category{},&models.Product{})
}