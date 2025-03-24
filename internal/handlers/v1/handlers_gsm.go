package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"

	db "github.com/maaw77/crmsrvg/internal/database"
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
	storage  *db.CrmDatabase
	validate *validator.Validate
}

// addEntryGsm adds an entry to the GSM table.
func (g *GsmTable) addEntryGsm(w http.ResponseWriter, r *http.Request) {
	log.Println("addEntryGsm Serving:", r.URL.Path, "from", r.Host)
	w.Header().Set("Content-Type", "application/json")

	encod := json.NewEncoder(w)

	var gsmEntry models.GsmTableEntry
	if err := json.NewDecoder(r.Body).Decode(&gsmEntry); err != nil {
		log.Println("error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		if err := encod.Encode(ErrorMessage{Details: err.Error()}); err != nil {
			log.Println("error: ", err)
		}
		return

	}

	if err := g.validate.Struct(&gsmEntry); err != nil {
		log.Println("error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		if err := encod.Encode(ErrorMessage{Details: "data validation error"}); err != nil {
			log.Println("error: ", err)
		}
		return
	}

	id, err := g.storage.InsertGsmTable(r.Context(), gsmEntry)
	switch {
	case errors.Is(err, db.ErrExist):
		log.Println("error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		if err := encod.Encode(ErrorMessage{Details: fmt.Sprintf("entry (guid=%s, id=%d) already exists", gsmEntry.GUID, id.ID)}); err != nil {
			log.Println(err)
		}
		return
	case err != nil:
		log.Println("error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		if err := encod.Encode(ErrorMessage{Details: "something happened to the server"}); err != nil {
			log.Println(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := encod.Encode(id); err != nil {
		log.Println("error: ", err)
	}

}

// updateEntryGsm updates an entry in the GSM table with the specified GUID.
func (g *GsmTable) updateEntryGsm(w http.ResponseWriter, r *http.Request) {
	log.Println("updateEntryGsm Serving:", r.URL.Path, "from", r.Host)
	w.Header().Set("Content-Type", "application/json")

	encod := json.NewEncoder(w)

	var gsmEntry models.GsmTableEntry
	if err := json.NewDecoder(r.Body).Decode(&gsmEntry); err != nil {
		log.Println("error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		if err := encod.Encode(ErrorMessage{Details: err.Error()}); err != nil {
			log.Println("error: ", err)
		}
		return

	}

	if err := g.validate.Struct(&gsmEntry); err != nil {
		log.Println("error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		if err := encod.Encode(ErrorMessage{Details: "data validation error"}); err != nil {
			log.Println("error: ", err)
		}
		return
	}

	id, err := g.storage.UpdateRowGsmTable(r.Context(), gsmEntry)
	switch {
	case errors.Is(err, db.ErrNotExist):
		log.Println("error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		if err := encod.Encode(ErrorMessage{Details: fmt.Sprintf("entry (guid=%s) not exists", gsmEntry.GUID)}); err != nil {
			log.Println(err)
		}
		return
	case err != nil:
		log.Println("error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		if err := encod.Encode(ErrorMessage{Details: "something happened to the server"}); err != nil {
			log.Println(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := encod.Encode(id); err != nil {
		log.Println("error: ", err)
	}

}

// getEntryGsm receives an entry with a specified ID from the GSM table.
func (g *GsmTable) getGsmEntryId(w http.ResponseWriter, r *http.Request) {
	log.Println("getEntryGsm Serving:", r.URL.Path, "from", r.Host)

	w.Header().Set("Content-Type", "application/json")

	encod := json.NewEncoder(w)

	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		log.Println("error: id value not set!")
		w.WriteHeader(http.StatusNotFound)
		encod.Encode(ErrorMessage{Details: "id value not set!"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		if err == nil {
			log.Println("error: id<1")
		} else {
			log.Println("error: ", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		if err := encod.Encode(ErrorMessage{Details: "data validation error"}); err != nil {
			log.Println("error: ", err)
		}
		return
	}

	gsmEntry, err := g.storage.GetRowGsmTableId(r.Context(), id)
	switch {
	case errors.Is(err, db.ErrNotExist):
		w.WriteHeader(http.StatusBadRequest)
		if err := encod.Encode(ErrorMessage{Details: db.ErrNotExist.Error()}); err != nil {
			log.Println("error: ", err)
		}
		return
	case err != nil:
		log.Println("error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		if err := encod.Encode(ErrorMessage{Details: "something happened to the server"}); err != nil {
			log.Println(err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	encod.Encode(gsmEntry)
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

	dateStr, ok := mux.Vars(r)["date"]
	if !ok {
		log.Println("error: date value not set!")
		w.WriteHeader(http.StatusNotFound)
		encod.Encode(ErrorMessage{Details: "date value not set!"})
		return
	}

	date, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := encod.Encode(ErrorMessage{Details: "data validation error"}); err != nil {
			log.Println("error: ", err)
		}
		return
	}

	gsmEntries, err := g.storage.GetRowsGsmTableDtReceiving(r.Context(), pgtype.Date{Time: date, Valid: true})

	switch {
	case errors.Is(err, db.ErrNotExist):
		w.WriteHeader(http.StatusBadRequest)
		if err := encod.Encode(ErrorMessage{Details: db.ErrNotExist.Error()}); err != nil {
			log.Println("error: ", err)
		}
		return
	case err != nil:
		log.Println("error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		if err := encod.Encode(ErrorMessage{Details: "something happened to the server"}); err != nil {
			log.Println(err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	encod.Encode(gsmEntries)
}

// delGsmEntryId deletess an entry with a specified ID from the GSM table.
func (g *GsmTable) delGsmEntryId(w http.ResponseWriter, r *http.Request) {
	log.Println("getEntryGsm Serving:", r.URL.Path, "from", r.Host)

	w.Header().Set("Content-Type", "application/json")

	encod := json.NewEncoder(w)

	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		log.Println("error: id value not set!")
		w.WriteHeader(http.StatusNotFound)
		encod.Encode(ErrorMessage{Details: "id value not set!"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		if err == nil {
			log.Println("error: id<1")
		} else {
			log.Println("error: ", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		if err := encod.Encode(ErrorMessage{Details: "data validation error"}); err != nil {
			log.Println("error: ", err)
		}
		return
	}

	err = g.storage.DelRowGsmTable(r.Context(), id)
	switch {
	case errors.Is(err, db.ErrNotExist):
		w.WriteHeader(http.StatusBadRequest)
		if err := encod.Encode(ErrorMessage{Details: db.ErrNotExist.Error()}); err != nil {
			log.Println("error: ", err)
		}
		return
	case err != nil:
		log.Println("error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		if err := encod.Encode(ErrorMessage{Details: "something happened to the server"}); err != nil {
			log.Println(err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	encod.Encode(models.IdEntry{ID: id})
}

// newGsmTable allocates and returns a new gsmTable.
func newGsmTable(srg *db.CrmDatabase) *GsmTable {
	return &GsmTable{storage: srg,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}
