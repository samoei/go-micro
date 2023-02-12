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

// implement the server http interface
func (p *Products) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet { //curl 127.0.0.1:9091/ | jq
		p.getProducts(resp, req)
		return
	}

	if req.Method == http.MethodPost { // curl 127.0.0.1:9091/ -v -d '{"id":3, "name":"Kenyan Tea", "description":"From the Kenyan Highlands"}' | jq
		p.createProduct(resp, req)
		return
	}

	if req.Method == http.MethodPut { //curl 127.0.0.1:9091/3 -XPUT -v -d '{"id":3, "name":"Ethiopian Green Tea", "description":"From the Ethiopian Plains"}' | jq
		p.l.Println("PUT", req.URL.Path)
		r := regexp.MustCompile(`/([0-9]+)`)
		group := r.FindAllStringSubmatch(req.URL.Path, -1)

		p.l.Println("Len: ", len(group))
		p.l.Println("Value: ", group)

		if len(group) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(resp, "Invalid URL", http.StatusBadRequest)
			return
		}

		if len(group[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(resp, "Invalid URL", http.StatusBadRequest)
			return
		}

		idStr := group[0][1]
		id, err := strconv.Atoi(idStr)

		if err != nil {
			p.l.Println("Invalid URI unable to convert to numer", idStr)
			http.Error(resp, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updateProducts(id, resp, req)
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

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to parse the payload", http.StatusBadRequest)
		return
	}

	data.AddProduct(prod)
}

func (p *Products) updateProducts(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")
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

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}
