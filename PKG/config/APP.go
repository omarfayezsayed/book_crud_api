package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB
var server = "localhost"
var port = "1433"
var user = "sa"
var password = "omar@1234"
var database = "books_cruds"

func Connect() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", server, user, password, port, database)
	con, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal(err)
	}
	db = con
}
func GetDB() *sql.DB {
	return db
}
