package handlers

import (
	"log"
	"main/data"
	"net/http"
	"regexp"
	"strconv"
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

	if req.Method == http.MethodPut {
		r := regexp.MustCompile("/([0-9]+)")
		group := r.FindAllStringSubmatch(req.URL.Path, -1)

		if len(group) != 1 {
			http.Error(resp, "Invalid URL", http.StatusBadRequest)
			return
		}

		if len(group[0]) != 2 {
			http.Error(resp, "Invalid URL", http.StatusBadRequest)
			return
		}

		idStr := group[0][1]
		id, _ := strconv.Atoi(idStr)

		p.updateProducts(id, resp, req)

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

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to parse the payload", http.StatusBadRequest)
		return
	}

	data.AddProduct(prod)
}

func (p *Products) updateProducts(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("Updating product")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to parse the payload", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
}
