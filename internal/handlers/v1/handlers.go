package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

type GsmTableEntry struct {
	Dt_receiving  time.Time `json:"dt_receiving"`      //     dt_receiving: datetime.date | str  # Data priemki
	Dt_crch       time.Time `json:"dt_crch,omitempty"` //     dt_crch: datetime.date | str  # Data sozdaniya ili posledney pravki
	Site          string    `json:"site"`              //     site: str  # Uchastok
	Income_kg     float64   `json:"income_kg"`         //     income_kg: float   # Prinyato v kg
	Operator      string    `json:"operator"`          //     operator: str  # Operator
	Provider      string    `json:"provider"`          //     provider: str  # Postavshikdt_receiving
	Contractor    string    `json:"contractor"`        //     contractor: str  # Perevozshik
	License_plate string    `json:"license_plate"`     //     license_plate: str   # GOS nomer
	Status        string    `json:"status"`            //     status: str  # Zagruzgen
	Been_changed  bool      `json:"been_changed"`      //     been_changed: bool   # table_color = '#f7fcc5' = T
}

type gsmTable struct {
	db string
}

func (c *gsmTable) helloHandeler(w http.ResponseWriter, r *http.Request) {
	log.Println("helloHandler Serving:", r.URL.Path, "from", r.Host)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<h1>%s</h1>", c.db)

}

// getEtryGsm receives an entry with a specified ID from the GSM table
func (c *gsmTable) getGsmEntryId(w http.ResponseWriter, r *http.Request) {
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

// getEtryGsm receives an entry with a specified date from the GSM table
func (c *gsmTable) getGsmEntryDate(w http.ResponseWriter, r *http.Request) {
	log.Println("getEntryGsm Serving:", r.URL.Path, "from", r.Host)

	dateEntry, ok := mux.Vars(r)["date"]
	if !ok {
		log.Println("Date value not set!")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	dt, err := time.Parse(time.DateOnly, dateEntry)
	if err != nil {
		log.Println(err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		body := fmt.Sprintf(`{"details": "%s"}`, err)
		fmt.Fprintf(w, "%s", body)
		return
	}

	w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "<h1>date = %s</h1>\n", dt)

	gsmEntry := GsmTableEntry{
		Dt_receiving: dt, //time.Date(2024, time.December, 15, 0, 0, 0, 0, time.UTC),
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
	encod := json.NewEncoder(w)
	encod.Encode(&gsmEntry)
}

// newGsmTable() allocates and returns a new gsmTable.
func newGsmTable() *gsmTable {
	return &gsmTable{db: "Hello Hugo boss!!!!!"}
}

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
func RegHanlders(r *mux.Router) {
	log.Println("Starting registration of URLs and handlers")

	r.NotFoundHandler = http.HandlerFunc(defaultHandler)
	r.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowed)
	gsmT := newGsmTable()

	apiR := r.PathPrefix("/api/v1").Subrouter()

	gsmR := apiR.PathPrefix("/gsm_table").Subrouter()
	gsmR.HandleFunc("/", gsmT.helloHandeler).Methods("GET")
	gsmR.HandleFunc("/{id:[0-9]+}", gsmT.getGsmEntryId).Methods("GET")
	gsmR.HandleFunc("/{date:[0-9]{4}-[0-9]{2}-[0-9]{2}}", gsmT.getGsmEntryDate).Methods("GET")

}
