package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/csothen/onlinestand/internal/onlinestand/service"
)

type CategoryController struct {
	l       *log.Logger
	service *service.CategoryService
}

func NewCategoryController(l *log.Logger) *CategoryController {
	service := service.NewCategoryService()
	return &CategoryController{l, service}
}

func (cc *CategoryController) GetAllCategories(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "GetAllCategories")
}

func (cc *CategoryController) GetCategoryById(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "GetCategoryById")
}

func (cc *CategoryController) CreateCategory(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "CreateCategory")
}

func (cc *CategoryController) UpdateCategoryById(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "UpdateCategoryById")
}

func (cc *CategoryController) DeleteCategoryById(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "DeleteCategoryById")
}
