package main

import (
	"net/http"

	"loja/routes"
)

func main() {

	routes.SetupRoutes()
	http.ListenAndServe(":8000", nil)

}
