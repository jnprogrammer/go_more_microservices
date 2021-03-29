package handlers

import (
	"N/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// Instead of having seporate funcs for GETS and POSTS, one func to do it all.
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("GET:returning a list of products")

	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Can't marshal jason", http.StatusInternalServerError)
	}

}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("POST:Adding a product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Can't unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", prod)

	data.AddProduct(prod)
}

func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Can't convert id", http.StatusBadRequest)
	}

	p.l.Println("PUT:updating product: ", id)

	prod := &data.Product{}
	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Can't unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
