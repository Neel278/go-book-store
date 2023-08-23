package main

import (
	"log"
	"net/http"

	"github.com/Neel278/test3/routes"
)

func main() {
	router := routes.RouterInitializer()
	err := http.ListenAndServe("localhost:3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
