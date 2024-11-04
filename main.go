package main

import (
	"go-web-application/routes"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
