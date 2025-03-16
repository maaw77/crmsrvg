package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/maaw77/crmsrvg/internal/models"
)

// {"dt_receiving": "2026-01-02", "dt_crch": "1988-02-22", "site": "SITE_1", "income_kg": 26442.70, "operator": "OPERATOR_1", "provider": "PROVIDER_1", "contractor": "CONTRACTOR_1", "license_plate": "A902RUS", "status": "Uploaded", "been_changed": false}
// curl -v -X POST -d '{"dt_receiving": "2026-01-02", "dt_crch": "1988-02-22", "site": "SITE_1", "income_kg": 26442.70, "operator": "OPERATOR_1", "provider": "PROVIDER_1", "contractor": "CONTRACTOR_1", "license_plate": "A902RUS", "status": "Uploaded", "been_changed": false}'  localhost:8080/api/v1/gsm_table/

// swagger:parameters paramIdEntry
type _ struct {
	// ID of the database entry
	// in:path
	// min: 1
	ID int `json:"id"`
}

// swagger:parameters paramCrmDate
type _ struct {
	// It is a date  based on the specified layout ("2006-01-02")
	//
	// required: true
	// in:path
	DATE pgtype.Date `json:"date"`
}

// swagger:parameters paramGsmTableEntry
type _ struct {
	// The structure for the entry in the GSM table
	// in:body
	Body models.GsmTableEntry
}

// GsmTable this defines a structure with configured data storage and request handlers.
type GsmTable struct {
	db string
}

// swagger:route POST /gsm_table GSM paramGsmTableEntry
//
// addEntryGsm
//
// Adds an entry to the GSM table
//
//	produces:
//	- application/json
//
// responses:
//		200:IdEntry
// 		400:ErrorMessage

// addEntryGsm adds an entry to the GSM table.
func (g *GsmTable) addEntryGsm(w http.ResponseWriter, r *http.Request) {
	log.Println("addEntryGsm Serving:", r.URL.Path, "from", r.Host)

	var gsmEntry models.GsmTableEntry
	dec := json.NewDecoder(r.Body)

	w.Header().Set("Content-Type", "application/json")

	if err := dec.Decode(&gsmEntry); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		body := fmt.Sprintf(`{"details": "%s"}`, err)
		fmt.Fprintf(w, "%s", body)
		return
	}

	fmt.Fprintf(w, "%+v\n", gsmEntry)
}

// swagger:route GET /gsm_table/id/{id} GSM paramIdEntry
//
// getIdEntryGsm
//
//  Receives an entry with a specified ID from the GSM table
//
//	produces:
//	- application/json
//
// responses:
//		200:GsmTableEntry
// 		404:ErrorMessage

// getEntryGsm receives an entry with a specified ID from the GSM table.
func (g *GsmTable) getGsmEntryId(w http.ResponseWriter, r *http.Request) {
	log.Println("getEntryGsm Serving:", r.URL.Path, "from", r.Host)

	w.Header().Set("Content-Type", "application/json")
	encod := json.NewEncoder(w)

	id, ok := mux.Vars(r)["id"]
	if !ok {
		log.Println("ID value not set!")
		w.WriteHeader(http.StatusNotFound)
		encod.Encode(&ErrorMessage{Details: "ID value not set!"})
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<h1>ID = %s</h1>\n", id)
}

// swagger:route GET /gsm_table/date/{date} GSM paramCrmDate
//
// getDateEntryGsm
//
//  Receives aan entry with a specified date from the GSM table
//
//	produces:
//	- application/json
//
// responses:
//		200:GsmTableEntry
// 		404:ErrorMessage

// getEntryGsm receives an entry with a specified date from the GSM table.
func (g *GsmTable) getGsmEntryDate(w http.ResponseWriter, r *http.Request) {
	log.Println("getEntryGsm Serving:", r.URL.Path, "from", r.Host)

	w.Header().Set("Content-Type", "application/json")
	encod := json.NewEncoder(w)

	dateEntry, ok := mux.Vars(r)["date"]
	if !ok {
		log.Println("Date value not set!")
		w.WriteHeader(http.StatusNotFound)
		encod.Encode(&ErrorMessage{Details: "IDate value not set!"})
		return
	}

	dt, err := time.Parse(time.DateOnly, dateEntry)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		encod.Encode(&ErrorMessage{Details: fmt.Sprintf(`{"details": "%s"}`, err)})
		// body := fmt.Sprintf(`{"details": "%s"}`, err)
		// fmt.Fprintf(w, "%s", body)
		return
	}

	w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "<h1>date = %s</h1>\n", dt)

	gsmEntry := models.GsmTableEntry{
		ID:          12,
		DtReceiving: pgtype.Date{Time: dt, Valid: true},
		// Dt_crch : "",
		Site:         "SITE_2",
		IncomeKg:     562.20,
		Operator:     "OPERATOR_2",
		Provider:     "PROVIDER_2",
		Contractor:   "CONTRACTOR_2",
		LicensePlate: "A342RUS",
		Status:       "Uploaded",
		BeenChanged:  false,
	}

	encod.Encode(&gsmEntry)
}

// newGsmTable allocates and returns a new gsmTable.
func newGsmTable() *GsmTable {
	return &GsmTable{db: "Hello Hugo boss!!!!!"}
}
