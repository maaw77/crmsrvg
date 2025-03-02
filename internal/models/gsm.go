package models

import (
	"encoding/json"
	"strings"
	"time"
)

// CrmDate is a date  based on the specified layout (time.DateOnly = "2006-01-02")
// swagger:strfmt date
type CrmDate struct {
	time.Time
}

func (c *CrmDate) UnmarshalJSON(b []byte) (err error) {
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

func (c CrmDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Format(time.DateOnly))
}

// swagger:model
type IdEntry struct {
	// ID of the database entry
	//
	// required: true
	// min: 1
	ID int `json:"id"`
}

// GsmTableEntry defines the structure for the entry in the GSM table
//
// swagger:model
type GsmTableEntry struct {
	// ID of the database entry
	//
	// required:true
	// min:1
	ID int `json:"id"`

	// Fuel receiving date
	//
	// required: true
	// example: 2024-01-02
	DtReceiving CrmDate `json:"dt_receiving"` //     dt_receiving: datetime.date | str  # Data priemki

	// Fuel receiving  date
	//
	// required: false
	// example: 2025-01-02
	DtCrch CrmDate `json:"dt_crch,omitempty"` //     dt_crch: datetime.date | str  # Data sozdaniya ili posledney pravki

	// Name of the mining site
	//
	// required: true
	// example: Some Name
	Site string `json:"site"` //     site: str  # Uchastok

	// The amount of fuel received at the warehouse in kilograms
	//
	// required: true
	// example: 362.20
	IncomeKg float64 `json:"income_kg"` //     income_kg: float   # Prinyato v kg

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
	LicensePlate string `json:"license_plate"` //     license_plate: str   # GOS nomer

	// Fuel loading status
	//
	// required: true
	// example: Uploaded
	Status string `json:"status"` //     status: str  # Zagruzgen

	// The status of the fuel intake record in the database (changed or not)
	//
	// required: true
	// example: false
	BeenChanged bool `json:"been_changed"` //     been_changed: bool   # table_color = '#f7fcc5' = T

}

// func (g GsmTableEntry) String() string {

// 	return fmt.Sprintf("{Dt_receiving: %v,  Dt_crch: %v, Site: %v, Income_kg: %v, Operator: %v, Provider: %v, Contractor: %v, License_plate: %v, Status: %v, Been_changed: %v}",
// 		g.Dt_receiving, g.Dt_crch, g.Site, g.Income_kg, g.Operator, g.Provider, g.Contractor, g.License_plate, g.Status, g.Been_changed)
// }
