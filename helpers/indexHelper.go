package helpers

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func initDBConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/book_store")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func ReadBooks() []Book {
	db := initDBConnection()
	defer db.Close()

	data, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	var books []Book
	for data.Next() {
		var book Book
		err := data.Scan(&book.Id, &book.Name, &book.Author)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	return books
}

func ReadOneData(id int) (*Book, error) {
	db := initDBConnection()
	defer db.Close()

	row, err := db.Query("SELECT * FROM books WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var book Book
	for row.Next() {
		err := row.Scan(&book.Id, &book.Name, &book.Author)
		if err != nil {
			log.Fatal(err)
		}
	}

	if book.Id == 0 {
		err := errors.New("invalid id")
		return nil, err
	}

	return &book, nil
}

func CreateNewBook(book Book) error {
	db := initDBConnection()
	defer db.Close()

	_, err := db.Exec("INSERT INTO books(name,author) VALUES (?,?)", book.Name, book.Author)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func UpdateBook(book Book, id int) error {
	db := initDBConnection()
	defer db.Close()

	existingBook, err := ReadOneData(id)

	if err != nil || existingBook == nil {
		log.Fatal(err)
	}

	if book.Id != 0 {
		existingBook.Id = book.Id
	}
	if book.Name != "" {
		existingBook.Name = book.Name
	}
	if book.Author != "" {
		existingBook.Author = book.Author
	}

	_, err = db.Exec("UPDATE books SET name=?, author=? WHERE id=?", existingBook.Name, existingBook.Author, id)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func DeleteBook(id int) error {
	db := initDBConnection()
	defer db.Close()

	existingBook, err := ReadOneData(id)

	if err != nil || existingBook == nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM books WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
