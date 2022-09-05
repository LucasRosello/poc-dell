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
		products, err := e.productService.Hello() //cambiar nombre
		if err != nil {
			fmt.Fprintf(resp, "%s", err)
			return
		}
		if len(products) == 0 {
			fmt.Fprintf(resp, "%s", err)
			return
		}

		fmt.Fprintf(resp, "%s", products)
	}
}
