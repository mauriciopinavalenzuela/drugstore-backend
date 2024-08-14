package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()
	// Define your routes here
	router.HandleFunc("/example", ExampleHandler).Methods("GET")
	return router
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
