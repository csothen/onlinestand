package controller

import (
	"log"
	"net/http"

	"github.com/csothen/onlinestand/internal/onlinestand/service"
	"github.com/csothen/onlinestand/pkg/responsehandler"
	"github.com/gorilla/mux"
)

// CategoryController : Handles the HTTP requests relative to Categories
type CategoryController struct {
	l       *log.Logger
	service *service.CategoryService
}

// NewCategoryController : Creates and initializes a new CategoryController
func NewCategoryController(l *log.Logger) *CategoryController {
	service := service.NewCategoryService()
	return &CategoryController{l, service}
}

// GetAllCategories : Handles the retrieval of all categories
func (cc *CategoryController) GetAllCategories(rw http.ResponseWriter, r *http.Request) {
	cc.l.Println("GET Categories")

	res, err := cc.service.GetAll()
	responsehandler.Respond(cc.l, rw, res, err)
}

// GetCategoryByID : Handles the retrieval of a category based on an ID
func (cc *CategoryController) GetCategoryByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	cc.l.Println("GET Category by ID", id)

	res, err := cc.service.GetByID(id)
	responsehandler.Respond(cc.l, rw, res, err)
}

// CreateCategory : Handles the persistence of a new category
func (cc *CategoryController) CreateCategory(rw http.ResponseWriter, r *http.Request) {
	cc.l.Println("POST New Category")

	res, err := cc.service.Create(r.Body)
	responsehandler.Respond(cc.l, rw, res, err)
}

// UpdateCategoryByID : Handles the update of an existing category based on an ID
func (cc *CategoryController) UpdateCategoryByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	cc.l.Println("PUT Category by ID", id)

	res, err := cc.service.UpdateByID(id, r.Body)
	responsehandler.Respond(cc.l, rw, res, err)
}

// DeleteCategoryByID : Handles the deletion of an existing category based on an ID
func (cc *CategoryController) DeleteCategoryByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	cc.l.Println("DELETE Category by ID", id)

	res, err := cc.service.DeleteByID(id)
	responsehandler.Respond(cc.l, rw, res, err)
}
