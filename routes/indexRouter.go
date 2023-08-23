package routes

import "net/http"

func RouterInitializer() {
	http.HandleFunc("/create", CreateNewData)
	http.HandleFunc("/update", UpdateData)
	http.HandleFunc("/delete", DeleteData)
	http.HandleFunc("/book", GetOneData)
	http.HandleFunc("/", GetAllData)
}
