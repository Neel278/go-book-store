package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func RouterInitializer() *mux.Router {
	router := mux.NewRouter()

	// Serve Swagger JSON
	router.HandleFunc("/v1/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/swagger.json")
	})

	router.PathPrefix("/v1/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/v1/swagger.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	router.HandleFunc("/v1/create", CreateNewData).Methods("POST")
	router.HandleFunc("/v1/update", UpdateData).Methods("POST")
	router.HandleFunc("/v1/delete", DeleteData).Methods("DELETE")
	router.HandleFunc("/v1/book", GetOneData).Methods("GET")
	router.HandleFunc("/v1/", GetAllData).Methods("GET")
	return router
}
