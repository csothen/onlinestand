package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/csothen/onlinestand/internal/onlinestand/service"
)

type FeatureController struct {
	l       *log.Logger
	service *service.FeatureService
}

func NewFeatureController(l *log.Logger) *FeatureController {
	service := service.NewFeatureService()
	return &FeatureController{l, service}
}

func (fc *FeatureController) GetAllFeatures(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "GetAllFeatures")
}

func (fc *FeatureController) GetFeatureById(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "GetFeatureById")
}

func (fc *FeatureController) CreateFeature(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "CreateFeature")
}

func (fc *FeatureController) UpdateFeatureById(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "UpdateFeatureById")
}

func (fc *FeatureController) DeleteFeatureById(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "DeleteFeatureById")
}
