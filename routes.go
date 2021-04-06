package main

import "net/http"

func (s *server) routes() {
	s.router.HandleFunc("/api/vehicles", s.handleGetVehicles()).Methods(http.MethodGet)
	s.router.HandleFunc("/api/vehicles", s.handleCreateVehicle()).Methods(http.MethodPost)
	s.router.HandleFunc("/api/", s.handleAPI()).Methods(http.MethodGet)
	s.router.HandleFunc("/", s.handleIndex())
}
