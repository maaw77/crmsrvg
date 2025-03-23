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

	"github.com/gorilla/mux"
	"github.com/maaw77/crmsrvg/internal/database"
	"github.com/maaw77/crmsrvg/internal/middleware"
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

// DefaultHandler is executed when no route matches.
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
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

// MethodNotAllowed is execode when the request method does not match the route.
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	log.Println("methodNotAllowed (Serving:", r.URL.Path, "from", r.Host, "with method", r.Method)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)

	// body := fmt.Sprintf(`{"details": "%s is NOT supported"}`, r.Method)
	// fmt.Fprintf(w, "%s", body)

	encod := json.NewEncoder(w)
	encod.Encode(&ErrorMessage{Details: fmt.Sprintf(`%s is not supported`, r.Method)})

}

// RegGsmHanlders registers handlers according to their URLs.
func RegGsmHanlders(rMux *mux.Router, srg *database.CrmDatabase) {
	log.Println("starting registration of URLs and handlers for GSM")

	gsmT := newGsmTable(srg)

	gsmR := rMux.PathPrefix("/gsm").Subrouter()
	gsmR.Use(middleware.AuthMiddleware)

	gsmR.Methods(http.MethodPost).HandlerFunc(gsmT.addEntryGsm)
	gsmR.Methods(http.MethodPut).HandlerFunc(gsmT.updateEntryGsm)

	gsmR.HandleFunc("/id/{id:[0-9]+}", gsmT.getGsmEntryId).Methods(http.MethodGet)
	gsmR.HandleFunc("/id/{id:[0-9]+}", gsmT.delGsmEntryId).Methods(http.MethodDelete)
	gsmR.HandleFunc("/date/{date:[0-9]{4}-[0-9]{2}-[0-9]{2}}", gsmT.getGsmEntryDate).Methods(http.MethodGet)

}

// RegUsersHanlders registers handlers according to their URLs.
func RegUsersHanlders(rMux *mux.Router, srg *database.CrmDatabase) {
	log.Println("starting registration of URLs and handlers for Users")

	usersT := newUsersTable(srg)

	rMux.HandleFunc("/reguser", usersT.regUser).Methods(http.MethodPost)
	rMux.HandleFunc("/login", usersT.loginUser).Methods(http.MethodPost)
}
