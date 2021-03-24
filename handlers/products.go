package handlers

import (
	"N/data"
	"encoding/json"
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
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Can't marshal jason", http.StatusInternalServerError)
	}
	rw.Write(d)
}
