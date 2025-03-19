package models

import (
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// // CrmDate is a date  based on the specified layout (time.DateOnly = "2006-01-02")
// // swagger:strfmt date
// type CrmDate struct {
// 	time.Time
// }

// func (c *CrmDate) UnmarshalJSON(b []byte) (err error) {
// 	inpS := strings.Trim(string(b), `"`)
// 	if inpS == "" || inpS == "null" {
// 		return nil
// 	}
// 	c.Time, err = time.Parse(time.DateOnly, inpS)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (c CrmDate) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(c.Time.Format(time.DateOnly))
// }

// func (c CrmDate) String() string {
// 	return c.Time.Format(time.DateOnly)
// }

// func (c *CrmDate) UnmarshalBinary(data []byte) error {
// 	return c.Time.UnmarshalBinary(data)
// }

// func (c CrmDate) MarshalBinary() ([]byte, error) {
// 	return c.Time.MarshalBinary()
// }

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
	// required:false
	// min:1
	ID int `json:"id,omitempty" db:"id" `

	// Fuel receiving date
	//
	// required: true
	// example: 2024-01-02
	DtReceiving pgtype.Date `json:"dt_receiving" db:"dt_receiving" validate:"required"` //     dt_receiving: datetime.date | str  # Data priemki

	// Fuel receiving  date
	//
	// required: false
	// example: 2025-01-02
	DtCrch pgtype.Date `json:"dt_crch,omitempty" db:"dt_crch"` //     dt_crch: datetime.date | str  # Data sozdaniya ili posledney pravki

	// Name of the mining site
	//
	// required: true
	// example: Some Name
	Site string `json:"site" db:"site" validate:"required"` //     site: str  # Uchastok

	// The amount of fuel received at the warehouse in kilograms
	//
	// required: true
	// example: 362.20
	IncomeKg float64 `json:"income_kg" db:"income_kg" validate:"required"` //     income_kg: float   # Prinyato v kg

	// Last name of the operator who took the fuel to the warehouse
	//
	// required: true
	// example: Some Last name
	Operator string `json:"operator" db:"operator" validate:"required"` //     operator: str  # Operator

	// Name of the fuel provider
	//
	// required: true
	// example: Some Name
	Provider string `json:"provider" db:"provider" validate:"required"` //     provider: str  # Postavshik

	// Name of the fuel carrier
	//
	// required: true
	// example: Some Name
	Contractor string `json:"contractor" db:"contractor" validate:"required"` //     contractor: str  # Perevozshik

	// The state number of the transport that delivered the fuel
	//
	// required: true
	// example: A902RUS
	LicensePlate string `json:"license_plate" db:"license_plate" validate:"required"` //     license_plate: str   # GOS nomer

	// Fuel loading status
	//
	// required: true
	// example: Uploaded
	Status string `json:"status" db:"status" validate:"required"` //     status: str  # Zagruzgen

	// The status of the fuel intake record in the database (changed or not)
	//
	// required: true
	// example: false
	BeenChanged bool `json:"been_changed" db:"been_changed"` //     been_changed: bool   # table_color = '#f7fcc5' = T

	// The global unique identifier of the record
	//
	// required: true
	// example: 6F9619FF-8B86-D011-B42D-00CF4FC964F
	GUID string `json:"guid" db:"guid" validate:"required,uuid"`
}

// It's Stringer interface (https://pkg.go.dev/fmt@go1.24.0#Stringer).
func (g GsmTableEntry) String() string {
	return fmt.Sprintf("{%d, %s, %s, %s, %.3f, %s, %s, %s, %s, %s, %v, %s}",
		g.ID,
		g.DtReceiving.Time.Format(time.DateOnly),
		g.DtCrch.Time.Format(time.DateOnly),
		g.Site,
		g.IncomeKg,
		g.Operator,
		g.Provider,
		g.Contractor,
		g.LicensePlate,
		g.Status,
		g.BeenChanged,
		g.GUID,
	)
}
