package main

import (
	routes "api/routes"
	"net/http"
)

func main() {
	router := routes.NewRouter()         // create a mux Router
	http.ListenAndServe(":3000", router) // start server
}
