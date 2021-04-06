package main

import (
	"net/http"
)

func (s *server) handleAPI() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		s.respond(rw, r, "Welcome to the Online Stand API", http.StatusOK)
	}
}
