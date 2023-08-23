package main

import (
	"log"
	"net/http"

	"github.com/Neel278/test3/routes"
)

func main() {
	routes.RouterInitializer()
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
