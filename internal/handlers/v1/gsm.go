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

type crmDate time.Time

func (c *crmDate) UnmarshalJSON(b []byte) error {
	inpS := strings.Trim(string(b), `"`)
	if inpS == "" || inpS == "null" {
		return nil
	}
	t, err := time.Parse(time.DateOnly, inpS)
	// log.Println("t, err =", t, err)
	if err != nil {
		return err
	}
	// log.Println("crmDate= ", crmDate(t))
	*c = crmDate(t)
	// log.Printf("*c= %#v", c)
	return nil
}

func (c crmDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format(time.DateOnly) + `"`), nil
}

func (c *crmDate) String() string {
	// log.Println(time.Time(*c).Format(time.DateOnly))
	return time.Time(*c).Format(time.DateOnly)
}

func (c *crmDate) GoString() string {
	return time.Time(*c).Format(time.DateOnly)
}

type GsmTableEntry struct {
	Dt_receiving  crmDate `json:"dt_receiving"`      //     dt_receiving: datetime.date | str  # Data priemki
	Dt_crch       crmDate `json:"dt_crch,omitempty"` //     dt_crch: datetime.date | str  # Data sozdaniya ili posledney pravki
	Site          string  `json:"site"`              //     site: str  # Uchastok
	Income_kg     float64 `json:"income_kg"`         //     income_kg: float   # Prinyato v kg
	Operator      string  `json:"operator"`          //     operator: str  # Operator
	Provider      string  `json:"provider"`          //     provider: str  # Postavshikdt_receiving
	Contractor    string  `json:"contractor"`        //     contractor: str  # Perevozshik
	License_plate string  `json:"license_plate"`     //     license_plate: str   # GOS nomer
	Status        string  `json:"status"`            //     status: str  # Zagruzgen
	Been_changed  bool    `json:"been_changed"`      //     been_changed: bool   # table_color = '#f7fcc5' = T

}

func (g GsmTableEntry) String() string {

	return fmt.Sprintf("{Dt_receiving: %v,  Dt_crch: %v, Site: %v, Income_kg: %v, Operator: %v, Provider: %v, Contractor: %v, License_plate: %v, Status: %v, Been_changed: %v}",
		g.Dt_receiving.String(), g.Dt_crch.String(), g.Site, g.Income_kg, g.Operator, g.Provider, g.Contractor, g.License_plate, g.Status, g.Been_changed)
}

type gsmTable struct {
	db string
}

func (g *gsmTable) helloHandeler(w http.ResponseWriter, r *http.Request) {
	log.Println("helloHandler Serving:", r.URL.Path, "from", r.Host)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<h1>%s</h1>", g.db)

}

// getEtryGsm receives an entry with a specified ID from the GSM table
func (g *gsmTable) getGsmEntryId(w http.ResponseWriter, r *http.Request) {
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
func (g *gsmTable) getGsmEntryDate(w http.ResponseWriter, r *http.Request) {
	log.Println("getEntryGsm Serving:", r.URL.Path, "from", r.Host)

	dateEntry, ok := mux.Vars(r)["date"]
	if !ok {
		log.Println("Date value not set!")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	dt, err := time.Parse(time.DateOnly, dateEntry)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		body := fmt.Sprintf(`{"details": "%s"}`, err)
		fmt.Fprintf(w, "%s", body)
		return
	}

	w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "<h1>date = %s</h1>\n", dt)

	gsmEntry := GsmTableEntry{
		Dt_receiving: crmDate(dt), //time.Date(2024, time.December, 15, 0, 0, 0, 0, time.UTC),
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

// addEntryGsm adds an entry to the GSM table
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

	fmt.Fprintf(w, "%v", gsmEntry)
}

// newGsmTable() allocates and returns a new gsmTable.
func newGsmTable() *gsmTable {
	return &gsmTable{db: "Hello Hugo boss!!!!!"}
}
