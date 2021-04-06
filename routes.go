package main

import "net/http"

func (s *server) routes() {
	s.router.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	s.router.HandleFunc("/api/", s.handleAPI())
	s.router.HandleFunc("/", s.handleIndex())
}
