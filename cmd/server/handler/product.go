package handler

import (
	"fmt"
	"net/http"

	"github.com/LucasRosello/poc-dell/internal/product"
)

type Product struct {
	productService product.Service
}

func NewProduct(e product.Service) *Product {
	return &Product{
		productService: e,
	}
}

func (e *Product) Hello() func(resp http.ResponseWriter, req *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {

		name := e.productService.Hello()

		fmt.Fprintf(resp, "%s", name)
	}
}
