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

	//Database
	db, err := sql.Open("pgx", "postgres://postgres:secret@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Domain Driven Design
	r := mux.NewRouter()
	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProduct(productService)

	// Routes
	r.HandleFunc("/products", productHandler.GetAll()).Methods("GET")
	r.HandleFunc("/products/{product_code}", productHandler.Get()).Methods("GET")
	r.HandleFunc("/products", productHandler.GetAll()).Methods("POST")
	http.ListenAndServe(":8080", r)
}
