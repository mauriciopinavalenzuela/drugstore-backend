package main

import (
	"drugstore-backend/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.InitializeRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}
