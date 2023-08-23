package main

import (
	"log"
	"net/http"

	_ "github.com/Neel278/test3/docs"
	"github.com/Neel278/test3/routes"
)

// @title Book store API
// @version 1.0
// @description This is an online book store CRUD web app.

// @host localhost:3000
// @BasePath /v1
func main() {
	router := routes.RouterInitializer()
	err := http.ListenAndServe("localhost:3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
