package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// {"dt_receiving": "2026-01-02", "dt_crch": "1988-02-22", "site": "SITE_1", "income_kg": 26442.70, "operator": "OPERATOR_1", "provider": "PROVIDER_1", "contractor": "CONTRACTOR_1", "license_plate": "A902RUS", "status": "Uploaded", "been_changed": false}
// curl -v -X POST -d '{"dt_receiving": "2026-01-02", "dt_crch": "1988-02-22", "site": "SITE_1", "income_kg": 26442.70, "operator": "OPERATOR_1", "provider": "PROVIDER_1", "contractor": "CONTRACTOR_1", "license_plate": "A902RUS", "status": "Uploaded", "been_changed": false}'  localhost:8080/api/v1/gsm_table/

// swagger:model
type IdEntry struct {
	// ID of the database entry
	//
	// required: true
	// min: 1
	ID int `json:"id"`
}

// swagger:parameters paramIdEntry
type _ struct {
	// ID of the database entry
	// in:path
	// min: 1
	ID int `json:"id"`
}

// crmDate is a date  based on the specified layout (time.DateOnly = "2006-01-02")
// swagger:strfmt date
type crmDate struct {
	time.Time
}

func (c *crmDate) UnmarshalJSON(b []byte) (err error) {
	inpS := strings.Trim(string(b), `"`)
	if inpS == "" || inpS == "null" {
		return nil
	}
	c.Time, err = time.Parse(time.DateOnly, inpS)
	if err != nil {
		return err
	}
	return nil
}

func (c crmDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Format(time.DateOnly))
}

// swagger:parameters paramCrmDate
type _ struct {
	// It is a date  based on the specified layout ("2006-01-02")
	//
	// required: true
	// in:path
	DATE crmDate `json:"date"`
}

// GsmTableEntry defines the structure for the entry in the GSM table
//
// swagger:model
type GsmTableEntry struct {
	// ID of the database entry
	//
	// required: true
	// min: 1
	ID int `json:"id"`

	// Fuel receiving date
	//
	// required: true
	// example: 2026-01-02
	Dt_receiving crmDate `json:"dt_receiving"` //     dt_receiving: datetime.date | str  # Data priemki

	// Fuel receiving  date
	//
	// required: false
	// example: 2026-01-02
	Dt_crch crmDate `json:"dt_crch,omitempty"` //     dt_crch: datetime.date | str  # Data sozdaniya ili posledney pravki

	// Name of the mining site
	//
	// required: true
	// example: Some Name
	Site string `json:"site"` //     site: str  # Uchastok

	// The amount of fuel received at the warehouse in kilograms
	//
	// required: true
	// example: 362.20
	Income_kg float64 `json:"income_kg"` //     income_kg: float   # Prinyato v kg

	// Last name of the operator who took the fuel to the warehouse
	//
	// required: true
	// example: Some Last name
	Operator string `json:"operator"` //     operator: str  # Operator

	// Name of the fuel provider
	//
	// required: true
	// example: Some Name
	Provider string `json:"provider"` //     provider: str  # Postavshik

	// Name of the fuel carrier
	//
	// required: true
	// example: Some Name
	Contractor string `json:"contractor"` //     contractor: str  # Perevozshik

	// The state number of the transport that delivered the fuel
	//
	// required: true
	// example: A902RUS
	License_plate string `json:"license_plate"` //     license_plate: str   # GOS nomer

	// Fuel loading status
	//
	// required: true
	// example: Uploaded
	Status string `json:"status"` //     status: str  # Zagruzgen

	// The status of the fuel intake record in the database (changed or not)
	//
	// required: true
	// example: false
	Been_changed bool `json:"been_changed"` //     been_changed: bool   # table_color = '#f7fcc5' = T

}

// swagger:parameters paramGsmTableEntry
type _ struct {
	// The structure for the entry in the GSM table
	// in:body
	Body GsmTableEntry
}

// func (g GsmTableEntry) String() string {

// 	return fmt.Sprintf("{Dt_receiving: %v,  Dt_crch: %v, Site: %v, Income_kg: %v, Operator: %v, Provider: %v, Contractor: %v, License_plate: %v, Status: %v, Been_changed: %v}",
// 		g.Dt_receiving, g.Dt_crch, g.Site, g.Income_kg, g.Operator, g.Provider, g.Contractor, g.License_plate, g.Status, g.Been_changed)
// }

// gsmTable this defines a structure with configured data storage and request handlers.
type gsmTable struct {
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
func (g *gsmTable) addEntryGsm(w http.ResponseWriter, r *http.Request) {
	log.Println("addEntryGsm Serving:", r.URL.Path, "from", r.Host)

	var gsmEntry GsmTableEntry
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
func (g *gsmTable) getGsmEntryId(w http.ResponseWriter, r *http.Request) {
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
func (g *gsmTable) getGsmEntryDate(w http.ResponseWriter, r *http.Request) {
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

	gsmEntry := GsmTableEntry{
		Dt_receiving: crmDate{dt},
		// Dt_crch : "",
		Site:          "SITE_2",
		Income_kg:     562.20,
		Operator:      "OPERATOR_2",
		Provider:      "PROVIDER_2",
		Contractor:    "CONTRACTOR_2",
		License_plate: "A342RUS",
		Status:        "Uploaded",
		Been_changed:  false,
	}

	encod.Encode(&gsmEntry)
}

// newGsmTable() allocates and returns a new gsmTable.
func newGsmTable() *gsmTable {
	return &gsmTable{db: "Hello Hugo boss!!!!!"}
}
