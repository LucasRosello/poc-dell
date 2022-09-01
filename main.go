package main

import (
	"net/http"

	"github.com/LucasRosello/poc-dell/cmd/server/handler"
	"github.com/LucasRosello/poc-dell/internal/product"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	productService := product.NewService()
	productHandler := handler.NewProduct(productService)
	r.HandleFunc("/hello", productHandler.Hello())

	http.ListenAndServe(":8080", r)
}
