package routes

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Neel278/test3/helpers"
)

// @Summary Fetch all books list from book-store
// @Description Fetches all the books from book store.
// @Accept json
// @Produce json
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Router / [get]
func GetAllData(w http.ResponseWriter, r *http.Request) {
	data := helpers.ReadBooks()

	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(jsonData))
}

// @Summary Fetch single book by id
// @Description Fetches single book by it's id from the book store.
// @Accept json
// @Produce json
// @Param id query string true "id of book to fetch"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /book [get]
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

// @Summary Create a new book in book-store
// @Description Adds a new book to the book store.
// @Accept json
// @Produce json
// @Param book body helpers.Book true "book body param"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /create [post]
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

// @Summary Updates an existing book in book-store
// @Description Updates an existing book in book store.
// @Accept json
// @Produce json
// @Param id query string true "id of the book that you want to update to"
// @Param book body helpers.Book true "new info that you want to change in current book info"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /update [post]
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

// @Summary Delete single book by id
// @Description Delete single book by it's id from the book store.
// @Accept json
// @Produce json
// @Param id query string true "id of book to delete"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /delete [delete]
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
