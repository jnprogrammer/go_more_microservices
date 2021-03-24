package handlers

import (
	"N/data"
	"log"
	"net/http"
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

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// Instead of having seporate funcs for GETS and POSTS, one func to do it all.
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProcuts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Can't marshal jason", http.StatusInternalServerError)
	}

}
