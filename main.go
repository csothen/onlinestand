package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	var (
		port = flags.Int("port", 8080, "Port to listen on")
	)

	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	srv, err := newServer()
	if err != nil {
		return err
	}

	db, err := startDatabase()
	if err != nil {
		return err
	}
	defer db.Close()
	srv.db = db

	fmt.Printf("Online Stand API listening on :%d\n", *port)
	return http.ListenAndServe(addr, srv)
}

type server struct {
	db     *sql.DB
	router *mux.Router
}

func newServer() (*server, error) {
	srv := &server{
		router: mux.NewRouter(),
	}

	srv.routes()
	return srv, nil
}

func (s *server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(rw, r)
}

func (s *server) respond(rw http.ResponseWriter, r *http.Request, data interface{}, status int) {
	rw.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(rw).Encode(data)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (s *server) decode(rw http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
