package main

import "net/http"

type vehicle struct {
}

func (s *server) handleGetVehicles() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		s.respond(rw, r, nil, http.StatusOK)
	}
}

func (s *server) handleCreateVehicle() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		s.respond(rw, r, nil, http.StatusOK)
	}
}
