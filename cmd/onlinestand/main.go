package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/csothen/onlinestand/internal/onlinestand/controller"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Online Stand Web Application")

	r := configRouter()
	s, l := configServer(r)

	runServer(s, l)
}

func configRouter() *mux.Router {
	router := mux.NewRouter()

	gets := router.Methods(http.MethodGet).Subrouter()
	posts := router.Methods(http.MethodPost).Subrouter()
	puts := router.Methods(http.MethodPut).Subrouter()
	deletes := router.Methods(http.MethodDelete).Subrouter()

	mapUserUrls(gets, posts, puts, deletes)
	mapVehicleUrls(gets, posts, puts, deletes)
	mapFeatureUrls(gets, posts, puts, deletes)
	mapLocationUrls(gets, posts, puts, deletes)
	mapCategoryUrls(gets, posts, puts, deletes)

	return router
}

func configServer(r *mux.Router) (*http.Server, *log.Logger) {
	l := log.New(os.Stdout, "[ Server ] - ", log.LstdFlags)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return server, l
}

func runServer(s *http.Server, l *log.Logger) error {
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)

	sig := <-ch
	l.Printf("Recieved terminate, graceful shutdown\nSignal recieved: %v", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)

	return nil
}

func mapUserUrls(gets, posts, puts, deletes *mux.Router) {
	ul := log.New(os.Stdout, "[ User API ] - ", log.LstdFlags)
	uc := controller.NewUserController(ul)

	// User GET endpoints
	gets.HandleFunc("/api/users", uc.GetAllUsers)
	gets.HandleFunc("/api/users/{id:[0-9]+}", uc.GetUserById)
	gets.HandleFunc("/api/users/{username}", uc.GetUserByUsername)

	// User POST endpoints
	posts.HandleFunc("/api/users", uc.CreateUser)

	// User PUT endpoints
	puts.HandleFunc("/api/users/{id:[0-9]+}", uc.UpdateUserById)
	puts.HandleFunc("/api/users/{username}", uc.UpdateUserByUsername)

	// User DELETE endpoints
	deletes.HandleFunc("/api/users/{id:[0-9]+}", uc.DeleteUserById)
	deletes.HandleFunc("/api/users/{username}", uc.DeleteUserByUsername)
}

func mapVehicleUrls(gets, posts, puts, deletes *mux.Router) {
	vl := log.New(os.Stdout, "[ Vehicle API ] - ", log.LstdFlags)
	vc := controller.NewVehicleController(vl)

	// Vehicle GET endpoints
	gets.HandleFunc("/api/vehicles/available", vc.GetAllAvailableVehicles)
	gets.HandleFunc("/api/vehicles/{id:[0-9]+}", vc.GetVehicleById)
	gets.HandleFunc("/api/vehicles", vc.GetAllVehicles)

	// Vehicle POST endpoints
	posts.HandleFunc("/api/vehicles", vc.CreateVehicle)

	// Vehicle PUT endpoints
	puts.HandleFunc("/api/vehicles/{id:[0-9]+}", vc.UpdateVehicleById)

	// Vehicle DELETE endpoints
	deletes.HandleFunc("/api/vehicles/{id:[0-9]+}", vc.DeleteVehicleById)
}

func mapFeatureUrls(gets, posts, puts, deletes *mux.Router) {
	fl := log.New(os.Stdout, "[ Feature API ] - ", log.LstdFlags)
	fc := controller.NewFeatureController(fl)

	// Feature GET endpoints
	gets.HandleFunc("/api/features", fc.GetAllFeatures)
	gets.HandleFunc("/api/features/{id:[0-9]+}", fc.GetFeatureById)

	// Feature POST endpoints
	posts.HandleFunc("/api/features", fc.CreateFeature)

	// Feature PUT endpoints
	puts.HandleFunc("/api/features/{id:[0-9]+}", fc.UpdateFeatureById)

	// Feature DELETE endpoints
	deletes.HandleFunc("/api/features/{id:[0-9]+}", fc.DeleteFeatureById)
}

func mapLocationUrls(gets, posts, puts, deletes *mux.Router) {
	ll := log.New(os.Stdout, "[ Location API ] - ", log.LstdFlags)
	lc := controller.NewLocationController(ll)

	// Location GET endpoints
	gets.HandleFunc("/api/locations", lc.GetAllLocations)
	gets.HandleFunc("/api/locations/{id:[0-9]+}", lc.GetLocationById)

	// Location POST endpoints
	posts.HandleFunc("/api/locations", lc.CreateLocation)

	// Location PUT endpoints
	puts.HandleFunc("/api/locations/{id:[0-9]+}", lc.UpdateLocationById)

	// Location DELETE endpoints
	deletes.HandleFunc("/api/locations/{id:[0-9]+}", lc.DeleteLocationById)
}

func mapCategoryUrls(gets, posts, puts, deletes *mux.Router) {
	cl := log.New(os.Stdout, "[ Category API ] - ", log.LstdFlags)
	cc := controller.NewCategoryController(cl)

	// Category GET endpoints
	gets.HandleFunc("/api/categories", cc.GetAllCategories)
	gets.HandleFunc("/api/categories/{id:[0-9]+}", cc.GetCategoryByID)

	// Category POST endpoints
	posts.HandleFunc("/api/categories", cc.CreateCategory)

	// Category PUT endpoints
	puts.HandleFunc("/api/categories/{id:[0-9]+}", cc.UpdateCategoryByID)

	// Category DELETE endpoints
	deletes.HandleFunc("/api/categories/{id:[0-9]+}", cc.DeleteCategoryByID)
}
