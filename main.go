package main

import (
	routes "api/routes"
	"net/http"
)

func main() {
	router := routes.NewRouter()
	http.ListenAndServe(":3000", router)
}
