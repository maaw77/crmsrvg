package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// defaultHandler is executed when no route matches.
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DefaultHandler Serving:", r.URL.Path, "from", r.Host, "with method", r.Method)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	body := fmt.Sprintf(`{"details": "%s is not supported"}`, r.URL.Path)
	fmt.Fprintf(w, "%s", body)
}

// methodNotAllowed is execode when the request method does not match the route.
func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	log.Println("methodNotAllowed (Serving:", r.URL.Path, "from", r.Host, "with method", r.Method)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)

	body := fmt.Sprintf(`{"details": "%s is NOT supported"}`, r.Method)
	fmt.Fprintf(w, "%s", body)

}

// RegHanlders registers handlers according to their URLs.
func RegHanlders(rMux *mux.Router) {
	log.Println("Starting registration of URLs and handlers")

	rMux.NotFoundHandler = http.HandlerFunc(defaultHandler)
	rMux.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowed)

	gsmT := newGsmTable()

	apiR := rMux.PathPrefix("/api/v1").Subrouter()

	gsmR := apiR.PathPrefix("/gsm_table").Subrouter()
	gsmR.HandleFunc("/", gsmT.helloHandeler).Methods("GET")
	gsmR.HandleFunc("/", gsmT.addEntryGsm).Methods("POST")
	gsmR.HandleFunc("/{id:[0-9]+}", gsmT.getGsmEntryId).Methods("GET")
	gsmR.HandleFunc("/{date:[0-9]{4}-[0-9]{2}-[0-9]{2}}", gsmT.getGsmEntryDate).Methods("GET")

}
