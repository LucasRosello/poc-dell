package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LucasRosello/poc-dell/internal/product"
	"github.com/gorilla/mux"
)

type Product struct {
	productService product.Service
}

func NewProduct(e product.Service) *Product {
	return &Product{
		productService: e,
	}
}

func (e *Product) GetAll() func(resp http.ResponseWriter, req *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {
		products, err := e.productService.GetAll()
		if err != nil {
			fmt.Fprintf(resp, "%s", err)
			return
		}
		if len(products) == 0 {
			fmt.Fprintf(resp, "%s", err)
			return
		}

		json.NewEncoder(resp).Encode(&products)
	}
}

func (p *Product) Get() func(resp http.ResponseWriter, req *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {
		product_code, _ := mux.Vars(req)["product_code"]
		product, err := p.productService.Get(product_code)
		if err != nil {
			return
		}

		json.NewEncoder(resp).Encode(&product)
	}
}
