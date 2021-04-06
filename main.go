package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
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

	fmt.Printf("Online Stand API listening on :%d\n", *port)
	return http.ListenAndServe(addr, srv)
}

type server struct {
	db     *sql.DB
	router *http.ServeMux
}

func newServer() (*server, error) {
	srv := &server{
		router: http.NewServeMux(),
	}

	srv.routes()
	return srv, nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
