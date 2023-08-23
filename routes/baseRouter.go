package routes

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Neel278/test3/helpers"
)

func GetAllData(w http.ResponseWriter, r *http.Request) {
	data := helpers.ReadBooks()

	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(jsonData))
}

func GetOneData(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		err := errors.New("invalid id")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	idInInt, err := strconv.Atoi(id)

	if err != nil {
		err := errors.New("invalid id")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	data, err := helpers.ReadOneData(idInInt)

	if err != nil {
		err := errors.New(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		err := errors.New(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(jsonData))
}

func CreateNewData(w http.ResponseWriter, r *http.Request) {

	var book helpers.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		log.Fatal(err)
	}

	err = helpers.CreateNewBook(book)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Success"))
}

func UpdateData(w http.ResponseWriter, r *http.Request) {

	var book helpers.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		log.Fatal(err)
	}

	id := r.URL.Query().Get("id")

	if id == "" {
		err := errors.New("invalid id")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	idInInt, err := strconv.Atoi(id)

	if err != nil {
		err := errors.New("invalid id")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = helpers.UpdateBook(book, idInInt)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func DeleteData(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idInInt, err := strconv.Atoi(id)

	if err != nil {
		err := errors.New("invalid id")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = helpers.DeleteBook(idInInt)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
