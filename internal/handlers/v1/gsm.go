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

// class GsmTableValidator(BaseModel):
//     """Priem"""
//     dt_receiving: datetime.date | str  # Data priemki
//     dt_crch: datetime.date | str  # Data sozdaniya ili posledney pravki
//     site: str  # Uchastok
//     income_kg: float   # Prinyato v kg
//     operator: str  # Operator
//     provider: str  # Postavshik
//     contractor: str  # Perevozshik
//     license_plate: str   # GOS nomer
//     status: str  # Zagruzgen
//     been_changed: bool   # table_color = '#f7fcc5' = Tru
//
//  [{"dt_receiving": "2024-06-01", "dt_crch": "", "site": "SITE_1", "income_kg": 26442.70, "operator": "OPERATOR_1", "provider": "PROVIDER_1", "contractor": "CONTRACTOR_1", "license_plate": "A902RUS", "status": "Uploaded", "been_changed": false},
// {"dt_receiving": "2024-06-01",
// "dt_crch": "",
// "site": "SITE_2",
// "income_kg": 562.20,
// "operator": "OPERATOR_2",
// "provider": "PROVIDER_2",
// "contractor": "CONTRACTOR_2",
// "license_plate": "A342RUS",
// "status": "Uploaded",
// "been_changed": false},
// ...]

// {"dt_receiving": "2026-01-02", "dt_crch": "1988-02-22", "site": "SITE_1", "income_kg": 26442.70, "operator": "OPERATOR_1", "provider": "PROVIDER_1", "contractor": "CONTRACTOR_1", "license_plate": "A902RUS", "status": "Uploaded", "been_changed": false}
// curl -v -X POST -d '{"dt_receiving": "2026-01-02", "dt_crch": "1988-02-22", "site": "SITE_1", "income_kg": 26442.70, "operator": "OPERATOR_1", "provider": "PROVIDER_1", "contractor": "CONTRACTOR_1", "license_plate": "A902RUS", "status": "Uploaded", "been_changed": false}'  localhost:8080/api/v1/gsm_table/

// swagger:parameters idParam
type _ struct {
	// The entry id
	// in: path
	// min:1
	// required: true
	ID int `json:"id"`
}

// swagger:parameters dateParam
type _ struct {
	// Fuel receiving date
	// in: path
	// required: true
	Date crmDate `json:"date"`
}

// swagger:parameters gsmTableEntryParam
type _ struct {
	// The entry in the GSM table
	// in:body
	// required: true
	GsmTableEntry
}

// swagger:response GsmTableEntryResponse
type _ struct {
	// The entry in the GSM table
	// in:body
	GsmTableEntry
}

// swagger:response idResponse
type _ struct {
	// The entry id
	// in:body
	// min:1
	ID int `json:"id"`
}

// crmDate is a date  based on the specified layout (time.DateOnly = "2006-01-02")
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

// GsmTableEntry defines the structure for the entry in the GSM table
//
// swagger:model
type GsmTableEntry struct {
	// Fuel receiving date
	// in: body
	//
	// required: true
	// example: 2026-01-02
	Dt_receiving crmDate `json:"dt_receiving"` //     dt_receiving: datetime.date | str  # Data priemki

	// Fuel receiving  date
	// in: body
	//
	// required: false
	// example: 2026-01-02
	Dt_crch crmDate `json:"dt_crch,omitempty"` //     dt_crch: datetime.date | str  # Data sozdaniya ili posledney pravki

	// Name of the mining site
	// in: body
	//
	// required: true
	// example: Some Name
	Site string `json:"site"` //     site: str  # Uchastok

	// The amount of fuel received at the warehouse in kilograms
	// in: body
	//
	// required: true
	// example: 362.20
	Income_kg float64 `json:"income_kg"` //     income_kg: float   # Prinyato v kg

	// Last name of the operator who took the fuel to the warehouse
	// in: body
	//
	// required: true
	// example: Some Last name
	Operator string `json:"operator"` //     operator: str  # Operator

	// Name of the fuel provider
	// in: body
	//
	// required: true
	// example: Some Name
	Provider string `json:"provider"` //     provider: str  # Postavshik

	// Name of the fuel carrier
	// in: body
	//
	// required: true
	// example: Some Name
	Contractor string `json:"contractor"` //     contractor: str  # Perevozshik

	// The state number of the transport that delivered the fuel
	// in: body
	//
	// required: true
	// example: A902RUS
	License_plate string `json:"license_plate"` //     license_plate: str   # GOS nomer

	// Fuel loading status
	// in: body
	//
	// required: true
	// example: Uploaded
	Status string `json:"status"` //     status: str  # Zagruzgen

	// The status of the fuel intake record in the database (changed or not)
	// in: body
	//
	// required: true
	// example: false
	Been_changed bool `json:"been_changed"` //     been_changed: bool   # table_color = '#f7fcc5' = T

}

// func (g GsmTableEntry) String() string {

// 	return fmt.Sprintf("{Dt_receiving: %v,  Dt_crch: %v, Site: %v, Income_kg: %v, Operator: %v, Provider: %v, Contractor: %v, License_plate: %v, Status: %v, Been_changed: %v}",
// 		g.Dt_receiving, g.Dt_crch, g.Site, g.Income_kg, g.Operator, g.Provider, g.Contractor, g.License_plate, g.Status, g.Been_changed)
// }

// gsmTable this defines a structure with configured data storage and request handlers.
type gsmTable struct {
	db string
}

// swagger:route GET /gsm_table/ getEntryGsm idParam
// getEntryGsm retrieves an entry with a specified ID from the GSM table
//
// responses:
// 200: GsmTableEntryResponse
// 404: ErrorResponse

// getEntryGsm receives an entry with a specified ID from the GSM table
func (g *gsmTable) getGsmEntryId(w http.ResponseWriter, r *http.Request) {
	log.Println("getEntryGsm Serving:", r.URL.Path, "from", r.Host)

	w.Header().Set("Content-Type", "application/json")
	encod := json.NewEncoder(w)

	id, ok := mux.Vars(r)["id"]
	if !ok {
		log.Println("ID value not set!")
		w.WriteHeader(http.StatusNotFound)
		encod.Encode(&ErrorResponse{Error: "ID value not set!"})
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<h1>ID = %s</h1>\n", id)
}

// swagger:route GET /gsm_table getEtryGsm dateParam
// getEntryGsm retrieves an entry with a specified date from the GSM table
//
// responses:
// 200: GsmTableEntryResponse
// 404: ErrorResponse
// 400: ErrorResponse

// getEntryGsm receives an entry with a specified date from the GSM table
func (g *gsmTable) getGsmEntryDate(w http.ResponseWriter, r *http.Request) {
	log.Println("getEntryGsm Serving:", r.URL.Path, "from", r.Host)

	w.Header().Set("Content-Type", "application/json")
	encod := json.NewEncoder(w)

	dateEntry, ok := mux.Vars(r)["date"]
	if !ok {
		log.Println("Date value not set!")
		w.WriteHeader(http.StatusNotFound)
		encod.Encode(&ErrorResponse{Error: "IDate value not set!"})
		return
	}

	dt, err := time.Parse(time.DateOnly, dateEntry)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		encod.Encode(&ErrorResponse{Error: fmt.Sprintf(`{"details": "%s"}`, err)})
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

// swagger:route POST /gsm_table addEntryGsm gsmTableEntryParam
// addEntryGsm adds an entry to the GSM table
//
// responses:
// 200: GsmTableEntryResponse
// 400: idResponse

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

// newGsmTable() allocates and returns a new gsmTable.
func newGsmTable() *gsmTable {
	return &gsmTable{db: "Hello Hugo boss!!!!!"}
}
