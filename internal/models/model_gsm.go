package models

import (
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// idEntry represents the ID of the entry in the database.
type IdEntry struct {
	// ID of the entry in the database
	ID int `json:"id" minimum:"1"`
}

// GsmEntryResponse defines the structure for the entry in the GSM table.
type GsmEntryResponse struct {
	// ID of the database entry
	ID int `json:"id,omitempty" db:"id" minimum:"1"`
	// Fuel receiving date
	DtReceiving pgtype.Date `json:"dt_receiving" db:"dt_receiving" validate:"required" swaggertype:"string" format:"date" example:"2024-11-15"`
	// Fuel receiving  date
	DtCrch pgtype.Date `json:"dt_crch,omitempty" db:"dt_crch" swaggertype:"string" format:"date" example:"2025-01-02"`
	// Name of the mining site
	Site string `json:"site" db:"site" validate:"required" example:"Name of the mining site"`
	// The amount of fuel received at the warehouse in kilograms
	IncomeKg float64 `json:"income_kg" db:"income_kg" validate:"required" example:"362.20"`
	// Last name of the operator who took the fuel to the warehouse
	Operator string `json:"operator" db:"operator" validate:"required" example:"Last name of the operator"`
	// Name of the fuel provider
	Provider string `json:"provider" db:"provider" validate:"required" example:"Name of the fuel provider"`
	// Name of the fuel carrier
	Contractor string `json:"contractor" db:"contractor" validate:"required" example:"Name of the fuel carrier"`
	// The state number of the transport that delivered the fuel
	LicensePlate string `json:"license_plate" db:"license_plate" validate:"required" example:" A902RUS"`
	// Fuel loading status
	Status string `json:"status" db:"status" validate:"required"  example:"Uploaded"`
	// The status of the fuel intake record in the database (changed or not)
	BeenChanged bool `json:"been_changed" db:"been_changed" example:"false"`
	// The global unique identifier of the record
	GUID string `json:"guid" db:"guid" validate:"required,uuid" example:"593ff941-405e-4afd-9eec-f8605a14351a"`
}

// GsmTableEntry
type GsmeEntryRequest struct {
	// ID of the database entry
	ID int `json:"id,omitempty" db:"id" minimum:"1" swaggerignore:"true"`
	// Fuel receiving date
	DtReceiving pgtype.Date `json:"dt_receiving" db:"dt_receiving" validate:"required" swaggertype:"string" format:"date" example:"2024-11-15"`
	// Fuel receiving  date
	DtCrch pgtype.Date `json:"dt_crch,omitempty" db:"dt_crch" swaggertype:"string" format:"date" example:"2025-01-02"`
	// Name of the mining site
	Site string `json:"site" db:"site" validate:"required" example:"Name of the mining site"`
	// The amount of fuel received at the warehouse in kilograms
	IncomeKg float64 `json:"income_kg" db:"income_kg" validate:"required" example:"362.20"`
	// Last name of the operator who took the fuel to the warehouse
	Operator string `json:"operator" db:"operator" validate:"required" example:"Last name of the operator"`
	// Name of the fuel provider
	Provider string `json:"provider" db:"provider" validate:"required" example:"Name of the fuel provider"`
	// Name of the fuel carrier
	Contractor string `json:"contractor" db:"contractor" validate:"required" example:"Name of the fuel carrier"`
	// The state number of the transport that delivered the fuel
	LicensePlate string `json:"license_plate" db:"license_plate" validate:"required" example:" A902RUS"`
	// Fuel loading status
	Status string `json:"status" db:"status" validate:"required"  example:"Uploaded"`
	// The status of the fuel intake record in the database (changed or not)
	BeenChanged bool `json:"been_changed" db:"been_changed" example:"true"`
	// The global unique identifier of the record
	GUID string `json:"guid" db:"guid" validate:"required,uuid" example:"593ff941-405e-4afd-9eec-f8605a14351a"`
}

// It's Stringer interface (https://pkg.go.dev/fmt@go1.24.0#Stringer).
func (g GsmEntryResponse) String() string {
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
