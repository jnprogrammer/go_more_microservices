package main

import (
	"github.com/jnprogrammer/go_more_microservices/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh = handlers.NewGoodbye(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/later", gh)

	http.ListenAndServe(":9090", nil)
}
