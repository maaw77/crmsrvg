// Package handlers CRM Server
//
// Documentation for REST API
//
//		Schemes: http
//		Host:
//		BasePath: /api/v1
//		Version: 1.0.0
//
//		TermsOfService: http://swagger.io/terms/
//		Contact: maaw@mai.ru
//		License: MIT http://opensource.org/licenses/MIT
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
//	     SecurityDefinitions:
//	         api_key:
//	             type: apiKey
//	             name: Authorization
//	             in: header
//
// swagger:meta
package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

// ErrorMessage is a detailed error message.
// swagger:model ErrorMessage
type ErrorMessage struct {
	// Description of the situation
	// example: An error occurred
	Details string `json:"details"`
}

// swagger:route GET / defaultHandler
//
// defaultHandler
//
// Default Handler for everything that is not a match.
// Works with all HTTP methods
//
// security:
// - api_key: []
// responses:
//	404: ErrorMessage

// defaultHandler is executed when no route matches.
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DefaultHandler Serving:", r.URL.Path, "from", r.Host, "with method", r.Method)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	encod := json.NewEncoder(w)
	encod.Encode(&ErrorMessage{Details: fmt.Sprintf(`%s is not supported`, r.URL.Path)})

	// body := fmt.Sprintf(`%s is not supported`, r.URL.Path)
	// fmt.Fprintf(w, "%s", body)
}

// swagger:route GET /* methodNotAllowed
//
// methodNotAllowed
//
// Default Handler for endpoints used with incorrect HTTP request method.
// Works with all paths and HTTP methods
//
// responses:
//	405: ErrorMessage

// methodNotAllowed is execode when the request method does not match the route.
func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	log.Println("methodNotAllowed (Serving:", r.URL.Path, "from", r.Host, "with method", r.Method)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)

	// body := fmt.Sprintf(`{"details": "%s is NOT supported"}`, r.Method)
	// fmt.Fprintf(w, "%s", body)

	encod := json.NewEncoder(w)
	encod.Encode(&ErrorMessage{Details: fmt.Sprintf(`%s is not supported`, r.Method)})

}

// RegHanlders registers handlers according to their URLs.
func RegHanlders(rMux *mux.Router) {
	log.Println("Starting registration of URLs and handlers")

	rMux.NotFoundHandler = http.HandlerFunc(defaultHandler)
	rMux.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowed)

	apiR := rMux.PathPrefix("/api/v1").Subrouter()

	docR := apiR.Methods(http.MethodGet).Subrouter()
	opts := middleware.RedocOpts{BasePath: "/api/v1", SpecURL: "/api/v1/docs/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	docR.Handle("/docs/", sh)
	docR.Handle("/docs/swagger.yaml", http.StripPrefix("/api/v1", http.FileServer(http.Dir("."))))

	gsmT := newGsmTable()

	gsmR := apiR.PathPrefix("/gsm_table").Subrouter()
	gsmR.HandleFunc("/", gsmT.addEntryGsm).Methods("POST")
	gsmR.HandleFunc("/id/{id:[0-9]+}", gsmT.getGsmEntryId).Methods("GET")
	gsmR.HandleFunc("/date/{date:[0-9]{4}-[0-9]{2}-[0-9]{2}}", gsmT.getGsmEntryDate).Methods("GET")

}
