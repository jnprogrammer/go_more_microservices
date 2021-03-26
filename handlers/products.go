package handlers

import (
	"N/data"
	"log"
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

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.l.Println("PUT")
		//expect ID of uri
		reg := regexp.MustCompile(`/([0-9]+)`) //capture bwtween 1 and 9
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URI more than 1 ID")
			http.Error(rw, "Invaid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than 1 capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI unable to convert NUMBER ")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updateProducts(id, rw, r)
		return
	}
	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// Instead of having seporate funcs for GETS and POSTS, one func to do it all.
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
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

func (p Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("PUT:updating a product")
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)

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
