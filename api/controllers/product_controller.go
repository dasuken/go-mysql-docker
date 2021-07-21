package controllers

import (
	"encoding/json"
	"github.com/noguchidaisuke/go-mysql-docker/api/models"
	"github.com/noguchidaisuke/go-mysql-docker/api/repository"
	"github.com/noguchidaisuke/go-mysql-docker/api/utils"
	"io"
	"net/http"
)

type ProductsController interface {
	PostProduct(w http.ResponseWriter, r *http.Request)
}

type productsControllerImpl struct {
	productsRepository repository.ProductsRepository
}

func NewProductsController(productsRepository repository.ProductsRepository) ProductsController {
	return &productsControllerImpl{productsRepository:productsRepository}
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

