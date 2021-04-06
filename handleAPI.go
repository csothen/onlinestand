package main

import (
	"fmt"
	"net/http"
)

func (s *server) handleAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Welcome to the Online Stand API")
	}
}
