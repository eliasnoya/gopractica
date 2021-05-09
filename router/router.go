package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	// Start Layout routes.go
	router.HandleFunc("/", Home).Methods("GET")
	// Assets
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	return router
}
