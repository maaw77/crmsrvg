package database

import (
	"context"
	"errors"
	"strconv"
	"testing"

	"github.com/maaw77/crmsrvg/config"
	"github.com/maaw77/crmsrvg/internal/models"
)

var (
	POSTGRESQL_URL     = config.InitConnString("")
	POSTGRESQL_URL_BAD = "postgres://postgres:crmpasswordocalhost:5433/postgres?sslmode=disable&pool_max_conns=10"

	// crmDB is the current instance of the CrmDatabase.
	crmDB *CrmDatabase
)

func TestNewCrmDatabaseEmty(t *testing.T) {
	if _, err := NewCrmDatabase(context.Background(), ""); !errors.Is(err, ErrConStrEmty) {
		t.Errorf("%s != %s", err, ErrConStrEmty)
	}

}

func TestPoolDbPing(t *testing.T) {
	var err error

	crmDB, err = NewCrmDatabase(context.Background(), POSTGRESQL_URL_BAD)
	if err == nil {
		t.Errorf("%s == nil", err)
	}

	crmDB, err = NewCrmDatabase(context.Background(), POSTGRESQL_URL)
	if err != nil {
		t.Errorf("%s != nil", err)
	}
	if err := crmDB.dbpool.Ping(context.Background()); err != nil {
		t.Errorf("%s != nil", err)
	}

	t.Logf("%#v\n", crmDB.dbpool.Config().ConnString())
}

func TestAuxilTableSites(t *testing.T) {
	var err error
	entries := map[string]models.IdEntry{}

	for i := 0; i <= 10; i++ {
		entries["Site_"+strconv.Itoa(i)], err = crmDB.GetIdOrCreateSites(context.Background(), "Site_"+strconv.Itoa(i))
		if err != nil {
			t.Logf("GetIdOrCreateSites-> %s != nil", err)
		}
	}

	for k, v := range entries {
		id, err := crmDB.GetIdOrCreateSites(context.Background(), k)

		if err != nil {
			t.Logf("GetIdOrCreateSites-> %s != nil", err)
		}
		if id.ID != v.ID {
			t.Logf("GetIdOrCreateSites-> %d != %d", id.ID, v.ID)
		}

	}

	for k, v := range entries {
		s, err := crmDB.DelRowSites(context.Background(), v.ID)

		if err != nil {
			t.Logf("DelRowSites-> %s != nil. For %s", err, k)
		}
		if !s {
			t.Logf("DelRowSites-> %v != true. For %s", s, k)
		}

	}

	for k, v := range entries {
		s, err := crmDB.DelRowSites(context.Background(), v.ID)

		if err != nil {
			t.Logf("DelRowSites-> %s != nil. For %s", err, k)
		}
		if s {
			t.Logf("DelRowSites-> %v != false. For %s", s, k)
		}

	}
}

func TestAuxilTableStatuses(t *testing.T) {
	var err error
	entries := map[string]models.IdEntry{}

	for i := 0; i <= 10; i++ {
		entries["Statuses_"+strconv.Itoa(i)], err = crmDB.GetIdOrCreateStatuses(context.Background(), "Statuses_"+strconv.Itoa(i))
		if err != nil {
			t.Logf("GetIdOrCreateStatuses-> %s != nil", err)
		}
	}

	for k, v := range entries {
		id, err := crmDB.GetIdOrCreateStatuses(context.Background(), k)

		if err != nil {
			t.Logf("GetIdOrCreateStatuses-> %s != nil", err)
		}
		if id.ID != v.ID {
			t.Logf("GetIdOrCreateStatuses-> %d != %d", id.ID, v.ID)
		}

	}

	for k, v := range entries {
		s, err := crmDB.DelRowStatuses(context.Background(), v.ID)

		if err != nil {
			t.Logf("DelRowStatuses-> %s != nil. For %s", err, k)
		}
		if !s {
			t.Logf("DelRowStatuses-> %v != false. For %s", s, k)
		}
	}

	for k, v := range entries {
		s, err := crmDB.DelRowStatuses(context.Background(), v.ID)

		if err != nil {
			t.Logf("DelRowStatuses-> %s != nil. For %s", err, k)
		}
		if s {
			t.Logf("DelRowStatuses-> %v != false. For %s", s, k)
		}

	}
}

func TestAuxilTableContractors(t *testing.T) {
	var err error
	entries := map[string]models.IdEntry{}

	for i := 0; i <= 10; i++ {
		entries["Contractors_"+strconv.Itoa(i)], err = crmDB.GetIdOrCreateContractors(context.Background(), "Contractors_"+strconv.Itoa(i))
		if err != nil {
			t.Logf("GetIdOrCreateContractors-> %s != nil", err)
		}
	}

	for k, v := range entries {
		id, err := crmDB.GetIdOrCreateContractors(context.Background(), k)

		if err != nil {
			t.Logf("GetIdOrCreateContractors-> %s != nil", err)
		}
		if id.ID != v.ID {
			t.Logf("GetIdOrCreateContractors-> %d != %d", id.ID, v.ID)
		}

	}

	for k, v := range entries {
		s, err := crmDB.DelRowContractors(context.Background(), v.ID)

		if err != nil {
			t.Logf("DelRowContractors-> %s != nil. For %s", err, k)
		}
		if !s {
			t.Logf("DelRowContractors-> %v != false. For %s", s, k)
		}

	}

	for k, v := range entries {
		s, err := crmDB.DelRowContractors(context.Background(), v.ID)

		if err != nil {
			t.Logf("DelRowContractors-> %s != nil. For %s", err, k)
		}
		if s {
			t.Logf("DelRowContractors-> %v != false. For %s", s, k)
		}

	}
}

func TestAuxilTableLicensePlates(t *testing.T) {
	var err error
	entries := map[string]models.IdEntry{}

	for i := 0; i <= 10; i++ {
		entries["LicensePlates_"+strconv.Itoa(i)], err = crmDB.GetIdOrCreateLicensePlates(context.Background(), "LicensePlates_"+strconv.Itoa(i))
		if err != nil {
			t.Logf("GetIdOrLicensePlates-> %s != nil", err)
		}
	}

	for k, v := range entries {
		id, err := crmDB.GetIdOrCreateLicensePlates(context.Background(), k)

		if err != nil {
			t.Logf("GetIdOrCreateLicensePlates-> %s != nil", err)
		}
		if id.ID != v.ID {
			t.Logf("GetIdOrCreateLicensePlates-> %d != %d", id.ID, v.ID)
		}

	}

	for k, v := range entries {
		s, err := crmDB.DelRowLicensePlates(context.Background(), v.ID)

		if err != nil {
			t.Logf("DelRowLicensePlates-> %s != nil. For %s", err, k)
		}
		if !s {
			t.Logf("DelRowLicensePlates-> %v != false. For %s", s, k)
		}

	}

	for k, v := range entries {
		s, err := crmDB.DelRowLicensePlates(context.Background(), v.ID)

		if err != nil {
			t.Logf("DelRowLicensePlates-> %s != nil. For %s", err, k)
		}
		if s {
			t.Logf("DelRowLicensePlates-> %v != false. For %s", s, k)
		}

	}
}

func TestAuxilTableOperators(t *testing.T) {
	var err error
	entries := map[string]models.IdEntry{}

	for i := 0; i <= 10; i++ {
		entries["Operators_"+strconv.Itoa(i)], err = crmDB.GetIdOrCreateOperators(context.Background(), "Operators_"+strconv.Itoa(i))
		if err != nil {
			t.Logf("GetIdOrCreateOperators-> %s != nil", err)
		}
	}

	for k, v := range entries {
		id, err := crmDB.GetIdOrCreateOperators(context.Background(), k)

		if err != nil {
			t.Logf("GetIdOrCreateOperators-> %s != nil", err)
		}
		if id.ID != v.ID {
			t.Logf("GetIdOrCreateOperators-> %d != %d", id.ID, v.ID)
		}

	}

	for k, v := range entries {
		s, err := crmDB.DelRowOperators(context.Background(), v.ID)

		if err != nil {
			t.Logf("DelRowOperators-> %s != nil. For %s", err, k)
		}
		if !s {
			t.Logf("DelRowOperators-> %v != false. For %s", s, k)
		}
	}

	for k, v := range entries {
		s, err := crmDB.DelRowOperators(context.Background(), v.ID)

		if err != nil {
			t.Logf("DelRowOperators-> %s != nil. For %s", err, k)
		}
		if s {
			t.Logf("DelRowOperators-> %v != false. For %s", s, k)
		}

	}
}

func TestAuxilTableProviders(t *testing.T) {
	var err error
	entries := map[string]models.IdEntry{}

	for i := 0; i <= 10; i++ {
		entries["Providers_"+strconv.Itoa(i)], err = crmDB.GetIdOrCreateProviders(context.Background(), "Providers_"+strconv.Itoa(i))
		if err != nil {
			t.Logf("GetIdOrCreateProviders-> %s != nil", err)
		}
	}

	for k, v := range entries {
		id, err := crmDB.GetIdOrCreateProviders(context.Background(), k)

		if err != nil {
			t.Logf("GetIdOrCreateProviders-> %s != nil", err)
		}
		if id.ID != v.ID {
			t.Logf("GetIdOrCreateProviders-> %d != %d", id.ID, v.ID)
		}

	}

	for k, v := range entries {
		s, err := crmDB.DelRowProviders(context.Background(), v.ID)

		if err != nil {
			t.Logf("DelRowProviders-> %s != nil. For %s", err, k)
		}
		if !s {
			t.Logf("DelRowProviders-> %v != false. For %s", s, k)
		}
	}

	for k, v := range entries {
		s, err := crmDB.DelRowProviders(context.Background(), v.ID)

		if err != nil {
			t.Logf("DelRowProviders-> %s != nil. For %s", err, k)
		}
		if s {
			t.Logf("DelRowProviders-> %v != false. For %s", s, k)
		}

	}
}

func TestInserGsmTable(t *testing.T) {
	// dt, _ := time.Parse(time.DateOnly, "2024-01-02")
	// tReceiving := models.CrmDate{Time: dt}
	// gsmEntry := models.GsmTableEntry{
	// 	ID:          12,
	// 	DtReceiving: models.CrmDate{Time: dt},
	// 	// Dt_crch : "",
	// 	Site:         "SITE_2",
	// 	IncomeKg:     562.20,
	// 	Operator:     "OPERATOR_2",
	// 	Provider:     "PROVIDER_2",
	// 	Contractor:   "CONTRACTOR_2",
	// 	LicensePlate: "A342RUS",
	// 	Status:       "Uploaded",
	// 	BeenChanged:  false,
	// }
	// t.Logf("gsmEntr=%v,\n dt=%s\n", gsmEntry, tReceiving)
	// id, err := crmDB.InserGsmTable(context.Background(), gsmEntry)
	// t.Log(id, err)
	// t.Log(crmDB.DelRowGsmTable(context.Background(), 3))

	t.Log(crmDB.GetRowGsmTableId(context.Background(), 1))
}
