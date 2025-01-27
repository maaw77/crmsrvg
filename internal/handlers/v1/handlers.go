package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type crmService struct {
	db string
}

func (c *crmService) helloHandeler(w http.ResponseWriter, r *http.Request) {
	log.Println("helloHandler Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<h1>%s</h1>", c.db)

}

// getEtryGsm receives an entry with a specified ID from the GSM table
func (c *crmService) getEntryGsm(w http.ResponseWriter, r *http.Request) {
	log.Println("getEntryGsm Serving:", r.URL.Path, "from", r.Host)

	id, ok := mux.Vars(r)["id"]
	if !ok {
		log.Println("ID value not set!")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<h1>ID = %s</h1>\n", id)
}
func newCrmService() *crmService {
	return &crmService{db: "Hello Hugo boss!!!!!"}
}

// RegHanlders registers handlers according to their URLs.
func RegHanlders(r *mux.Router) {
	log.Println("Starting registration of URLs and handlers")
	crmS := newCrmService()

	apiR := r.PathPrefix("/api/v1").Subrouter()

	gsmR := apiR.PathPrefix("/gsm_table").Subrouter()
	gsmR.HandleFunc("/", crmS.helloHandeler).Methods("GET")
	gsmR.HandleFunc("/{id:[0-9]+}", crmS.getEntryGsm)

}
