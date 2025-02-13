// Package handlers for the CRM Server
//
// Documentation for REST API
//
//	Schemes: http
//	BasePath: /api/v1
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
//
//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger@latest generate spec --scan-models -o ./swagger1.yaml
//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger@latest validate ./swagger1.yaml
package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ErrorResponse represents the details of the error message.
// swagger:response ErrorResponse
type ErrorResponse struct {
	// The error message
	// example: An error occurred
	// in: body
	Error string `json:"error"`
}

// swagger:route POST / defaultHandler
// defaultHandler for everything that is not a match.
// Works with all HTTP methods
//
// responses:
// 404: ErrorResponse

// defaultHandler is executed when no route matches.
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DefaultHandler Serving:", r.URL.Path, "from", r.Host, "with method", r.Method)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	encod := json.NewEncoder(w)
	encod.Encode(&ErrorResponse{Error: fmt.Sprintf(`%s is not supported`, r.URL.Path)})

	// body := fmt.Sprintf(`%s is not supported`, r.URL.Path)
	// fmt.Fprintf(w, "%s", body)
}

// swagger:route GET /* methodNotAllowed
// methodNotAllowed for endpoints used with incorrect HTTP request method
//
// responses:
//	404: ErrorResponse

// methodNotAllowed is execode when the request method does not match the route.
func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	log.Println("methodNotAllowed (Serving:", r.URL.Path, "from", r.Host, "with method", r.Method)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)

	// body := fmt.Sprintf(`{"details": "%s is NOT supported"}`, r.Method)
	// fmt.Fprintf(w, "%s", body)

	encod := json.NewEncoder(w)
	encod.Encode(&ErrorResponse{Error: fmt.Sprintf(`%s is not supported`, r.Method)})

}

// RegHanlders registers handlers according to their URLs.
func RegHanlders(rMux *mux.Router) {
	log.Println("Starting registration of URLs and handlers")

	rMux.NotFoundHandler = http.HandlerFunc(defaultHandler)
	rMux.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowed)

	gsmT := newGsmTable()

	apiR := rMux.PathPrefix("/api/v1").Subrouter()

	gsmR := apiR.PathPrefix("/gsm_table").Subrouter()
	gsmR.HandleFunc("/", gsmT.addEntryGsm).Methods("POST")
	gsmR.HandleFunc("/id/{id:[0-9]+}", gsmT.getGsmEntryId).Methods("GET")
	gsmR.HandleFunc("/date/{date:[0-9]{4}-[0-9]{2}-[0-9]{2}}", gsmT.getGsmEntryDate).Methods("GET")

}
