package models

import (
	"database/sql"
	"log"

	"github.com/omarfayezsayed/book-crud-api/PKG/config"
)

var db *sql.DB

type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"book_name"`
	Author string `json:"book_author"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetAllBooks() []Book {
	var books []Book
	rows, err := db.Query("select * from books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var book Book
		err_2 := rows.Scan(&book.ID, &book.Name, &book.Author)
		if err_2 != nil {
			break
		}
		books = append(books, book)
	}
	return books
}

func GetBookByID(id int) (Book, error) {
	var book Book
	if err := db.QueryRow("select * from books where (id = ?)", id).Scan(&book.ID, &book.Name, &book.Author); err != nil {
		return book, err
	}
	return book, nil
}

func DeleteBook(id int) error {
	rows, err := db.Exec("delete from books where id =?", id)
	if err != nil {
		return err
	}
	num, err := rows.RowsAffected()
	if num == 0 {
		return sql.ErrNoRows
	}
	return nil
}
func CreateBook(title string, author string) {
	_, err := db.Exec("insert into books (title,author) values (?,?)", title, author)
	if err != nil {
		log.Fatal(err)
	}
}
func UpdateBook(id int, title string, author string) {
	_, err := db.Exec("update books set title =? , author = ? where id =?", title, author, id)
	if err != nil {
		log.Fatal(err)
	}
}
