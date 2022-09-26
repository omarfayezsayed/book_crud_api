package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/omarfayezsayed/book-crud-api/PKG/routes"
)

func main() {
	r := mux.NewRouter()
	routes.BookRoutes(r)
	handler := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	fmt.Println("connecting to the server port 8000")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handler, methods, origins)(r)))
}
