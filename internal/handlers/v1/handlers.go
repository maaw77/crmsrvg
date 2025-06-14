// Package handlers CRM Server

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
type ErrorMessage struct {
	// Description of the situation
	Details string `json:"details" example:"An error occurred"`
}

// DefaultHandler is executed when no route matches.
//
//	@Summary		A custom path handler is executed if no route matches
//	@Description	Supports GET/POST/PUT/DELETE for any URL
//	@Failure		404	{object}	ErrorMessage
//
//	@Tags			default
//	@Produce		json
//
//	@Router			/{path} [get]
//	@Param			path	path	string	false	"An arbitrary path"
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DefaultHandler Serving:", r.URL.Path, "from", r.Host, "with method", r.Method)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	encod := json.NewEncoder(w)
	encod.Encode(&ErrorMessage{Details: fmt.Sprintf(`%s is not supported`, r.URL.Path)})

	// body := fmt.Sprintf(`%s is not supported`, r.URL.Path)
	// fmt.Fprintf(w, "%s", body)
}

// MethodNotAllowed is execode when the request method does not match the route.
//
//	@Summary		A custom path handler is executed if the request method does not match the route.
//	@Description	Supports GET/POST/PUT/DELETE for any URL
//
//	@Tags			default
//	@Produce		json
//
//	@Failure		404	{object}	ErrorMessage
//	@Router			/{path} [post]
//	@Param			path	path	string	false	"An arbitrary path"
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

	rMux.HandleFunc("/users", usersT.regUser).Methods(http.MethodPost)
	rMux.HandleFunc("/users/login", usersT.loginUser).Methods(http.MethodPost)
}
