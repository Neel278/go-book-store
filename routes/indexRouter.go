package routes

import (
	"github.com/gorilla/mux"
)

func RouterInitializer() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/create", CreateNewData)
	router.HandleFunc("/update", UpdateData)
	router.HandleFunc("/delete", DeleteData)
	router.HandleFunc("/book", GetOneData)
	router.HandleFunc("/", GetAllData)
	return router
}
