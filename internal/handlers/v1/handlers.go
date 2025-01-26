package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandeler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1>", "Hello Hugo boss!")

}

// RegHanlders registers handlers according to their URLs.
func RegHanlders(r *mux.Router) {
	log.Println("Starting registration of URLs and handlers")
	apiR := r.PathPrefix("/api/v1").Subrouter()

	gsmR := apiR.PathPrefix("/gsm_table").Subrouter()

	gsmR.Methods("GET").HandlerFunc(helloHandeler)

}
