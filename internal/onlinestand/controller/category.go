package controller

import (
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

}

func (cc *CategoryController) GetCategoryById(rw http.ResponseWriter, r *http.Request) {

}

func (cc *CategoryController) CreateCategory(rw http.ResponseWriter, r *http.Request) {

}

func (cc *CategoryController) UpdateCategoryById(rw http.ResponseWriter, r *http.Request) {

}

func (cc *CategoryController) DeleteCategoryById(rw http.ResponseWriter, r *http.Request) {

}
