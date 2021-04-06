package main

import (
	"net/http"
	"path/filepath"
)

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("public", "index.html"))
	}
}
