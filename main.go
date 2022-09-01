package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/LucasRosello/poc-dell/cmd/server/handler"
	"github.com/LucasRosello/poc-dell/internal/product"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "postgres://postgres:secret@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	r := mux.NewRouter()
	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProduct(productService)
	r.HandleFunc("/hello", productHandler.Hello())

	http.ListenAndServe(":8080", r)
}
