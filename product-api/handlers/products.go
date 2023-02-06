package handlers

import (
	"log"
	"main/data"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	productList := data.GetProducts()
	err := productList.ToJSON(resp)

	if err != nil {
		http.Error(resp, "Unable to fetch products", http.StatusInternalServerError)
	}

}
