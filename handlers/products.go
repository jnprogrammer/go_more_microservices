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

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProcuts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Can't marshal jason", http.StatusInternalServerError)
	}
}
