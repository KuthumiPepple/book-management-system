package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kuthumipepple/book-management-system/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:5000", r))
}
