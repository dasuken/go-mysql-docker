package api

import (
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
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
	resetTables = flag.Bool("rt", false, "test mode or not")
)

func Run() {
	flag.Parse()

	db := database.Connect()
	fmt.Println("Database connected...")

	err := migrateTables()
	if err != nil {
		log.Fatalf("DB Migration Error: %v", err)
	}

	// repo
	productsRepo := repository.NewProductRepository(db)
	categoriesRepo := repository.NewCategoriesRepository(db)

	// controller
	productsCon := controllers.NewProductsController(productsRepo)
	categoriesCon := controllers.NewCategoriesController(categoriesRepo)

	// route
	var allRoutes []*routes.Route
	allRoutes = routes.NewProductRoutes(productsCon)
	allRoutes = append(allRoutes, routes.NewCategoryRoutes(categoriesCon)...)

	router := mux.NewRouter()
	routes.Install(router, allRoutes)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Request", "Location", "Entity", "Accept"})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("API Listening on", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handlers.CORS(headers, methods, origins)(router)))
}

func migrateTables() error {
	db := database.Connect()
	return db.AutoMigrate(&models.Category{},&models.Product{})
}
