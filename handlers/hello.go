package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hey what have you been doing ?")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oopps? ?", http.StatusBadRequest)
		return
	}

	//log.Printf("Data %s", d)

	fmt.Fprintf(rw, "You sent the data: %s \n", d)
}
