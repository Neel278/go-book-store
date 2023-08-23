package routes

import (
	"github.com/gorilla/mux"
)

func RouterInitializer() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/create", CreateNewData).Methods("POST")
	router.HandleFunc("/update", UpdateData).Methods("POST")
	router.HandleFunc("/delete", DeleteData).Methods("DELETE")
	router.HandleFunc("/book", GetOneData).Methods("GET")
	router.HandleFunc("/", GetAllData).Methods("GET")
	return router
}
