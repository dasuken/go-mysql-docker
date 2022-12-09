package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/noguchidaisuke/go-mysql-docker/api/models"
	"github.com/noguchidaisuke/go-mysql-docker/api/repository"
	"github.com/noguchidaisuke/go-mysql-docker/api/utils"
	"io"
	"net/http"
	"strconv"
)

type ProductsController interface {
	PostProduct(w http.ResponseWriter, r *http.Request)
	GetProduct(w http.ResponseWriter, r *http.Request)
	GetAllProducts(w http.ResponseWriter, r *http.Request)
	PutProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
}

type productsControllerImpl struct {
	productsRepository repository.ProductsRepository
	paginationBuilder  repository.PaginationBuilderRepository
}

func NewProductsController(productsRepository repository.ProductsRepository) ProductsController {
	return &productsControllerImpl{productsRepository: productsRepository}
}

func (c *productsControllerImpl) PostProduct(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	product := &models.Product{}
	err = json.Unmarshal(bytes, product)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = product.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	product.CheckStatus()

	product, err = c.productsRepository.Save(product)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	utils.WriteAsJson(w, product)
}

func (c *productsControllerImpl) GetProduct(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	product_id, err := strconv.Atoi(v["product_id"])
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
	}

	product, err := c.productsRepository.Find(uint64(product_id))
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
	}

	utils.WriteAsJson(w, product)
}

func (c *productsControllerImpl) GetAllProducts(w http.ResponseWriter, r *http.Request) {

	products, err := c.productsRepository.FindAll()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
	}

	utils.WriteAsJson(w, products)
}

func (c *productsControllerImpl) PutProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	product_id, err := strconv.ParseUint(params["product_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
	}

	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
	}

	product := models.Product{}
	err = json.Unmarshal(bytes, &product)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
	}

	product.ID = product_id
	product.CheckStatus()

	err = c.productsRepository.Update(&product)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
	}

	utils.WriteAsJson(w, map[string]bool{"success": err == nil})
}

func (c *productsControllerImpl) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	product_id, err := strconv.ParseUint(params["product_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.productsRepository.Delete(product_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
	}

	w.Header().Set("Entity", fmt.Sprint(product_id))
	w.WriteHeader(http.StatusNoContent)

	utils.WriteAsJson(w, "{}")
}
