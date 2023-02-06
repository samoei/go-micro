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
	if req.Method == http.MethodGet {
		p.getProducts(resp, req)
		return
	}

	if req.Method == http.MethodPost {
		p.createProduct(resp, req)
		return
	}

	resp.WriteHeader(http.StatusMethodNotAllowed)
}
func (p *Products) getProducts(resp http.ResponseWriter, req *http.Request) {
	productList := data.GetProducts()
	err := productList.ToJSON(resp)

	if err != nil {
		http.Error(resp, "Unable to fetch products", http.StatusInternalServerError)
	}
}

func (p *Products) createProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Creating new product")
}
