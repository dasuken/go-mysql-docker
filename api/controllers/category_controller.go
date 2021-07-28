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

type CategoriesController interface {
	PostCategory(w http.ResponseWriter, r *http.Request)
	GetCategory(w http.ResponseWriter, r *http.Request)
	GetAllCategories(w http.ResponseWriter, r *http.Request)
	PutCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
}

type categoriesControllerImpl struct {
	categoriesRepository repository.CategoriesRepository
}

func NewCategoriesController(categoriesRepository repository.CategoriesRepository) CategoriesController {
	return &categoriesControllerImpl{categoriesRepository: categoriesRepository}
}

func (c *categoriesControllerImpl) PostCategory(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	category := &models.Category{}
	err = json.Unmarshal(bytes, category)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = category.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	category, err = c.categoriesRepository.Save(category)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	utils.WriteAsJson(w, category)
}

func (c *categoriesControllerImpl) GetCategory(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	category_id, err := strconv.Atoi(v["category_id"])
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
	}

	category, err := c.categoriesRepository.Find(uint64(category_id))
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
	}

	utils.WriteAsJson(w, category)
}

func (c *categoriesControllerImpl) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := c.categoriesRepository.FindAll()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
	}

	utils.WriteAsJson(w, categories)
}

func (c *categoriesControllerImpl) PutCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	category_id, err := strconv.ParseUint(params["category_id"], 10, 64)
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

	category := models.Category{}
	err = json.Unmarshal(bytes, &category)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
	}

	category.ID = category_id
	err = category.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
	}

	err = c.categoriesRepository.Update(&category)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
	}

	utils.WriteAsJson(w, map[string]bool{"success": err == nil})
}

func (c *categoriesControllerImpl) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	category_id, err := strconv.ParseUint(params["category_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.categoriesRepository.Delete(category_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
	}

	w.Header().Set("Entity", fmt.Sprint(category_id))
	w.WriteHeader(http.StatusNoContent)

	utils.WriteAsJson(w, "{}")
}