package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/maaw77/crmsrvg/config"
	"github.com/maaw77/crmsrvg/internal/models"
)

var (
	POSTGRESQL_URL     = config.InitConnString("")
	POSTGRESQL_URL_BAD = "postgres://postgres:crmpasswordocalhost:5433/postgres?sslmode=disable&pool_max_conns=10"

	// crmDB is the current instance of the CrmDatabase.
	crmDB         *CrmDatabase
	gsmEntriesMap = make(map[int]models.GsmTableEntry)
)

func gsmEntriesEqual(a, b models.GsmTableEntry) bool {

	switch {
	case a.ID != b.ID:
		return false
	case a.DtReceiving != b.DtReceiving:
		return false
	case a.DtCrch != b.DtCrch:
		return false
	case a.Site != b.Site:
		return false
	case fmt.Sprintf("%.3f", a.IncomeKg) != fmt.Sprintf("%.3f", b.IncomeKg):
		return false
	case a.Operator != b.Operator:
		return false
	case a.Provider != b.Provider:
		return false
	case a.Contractor != b.Contractor:
		return false
	case a.LicensePlate != b.LicensePlate:
		return false
	case a.Status != b.Status:
		return false
	case a.BeenChanged != b.BeenChanged:
		return false
	case a.GUID != b.GUID:
		return false
	}

	return true
}
func readData(fileName string) (gsmEntries []models.GsmTableEntry, err error) {
	// fmt.Println("###############################")
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer f.Close()

	dec := json.NewDecoder(f)

	for {
		gsmE := models.GsmTableEntry{}
		err = dec.Decode(&gsmE)

		if err == io.EOF {
			break
		} else if err != nil {
			return
		}

		// log.Println(gsmE)
		gsmEntries = append(gsmEntries, gsmE)
	}

	return gsmEntries, nil
}

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
	if err := crmDB.DBpool.Ping(context.Background()); err != nil {
		t.Errorf("%s != nil", err)
	}

	t.Logf("%#v\n", crmDB.DBpool.Config().ConnString())
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
	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	// t.Log(pwd)

	gsmEntriesArr, err := readData(pwd + "/testdata/gsm.data")
	if err != nil {
		t.Fatalf("%s != nil", err)
	}

	// t.Log(gsmEntriesEqual(gsmEntriesArr[0], gsmEntriesArr[2]))

	for _, gE := range gsmEntriesArr {
		id, err := crmDB.InsertGsmTable(context.Background(), gE)
		if err != nil {
			t.Errorf("%s != nil", err)
			continue
		}
		gE.ID = id.ID
		gsmEntriesMap[id.ID] = gE

	}
	// t.Log(gsmEntriesMap)

	for k, v := range gsmEntriesMap {

		id, err := crmDB.InsertGsmTable(context.Background(), v)
		if k != id.ID {
			t.Errorf("%d != %d", k, id.ID)
		}
		if !errors.Is(err, ErrExists) {
			t.Errorf("%v != %v", err, ErrExists)
		}

	}
	// t.Log(gsmEntriesMap)
}

func TestGetRowGsmTableId(t *testing.T) {
	for k, v := range gsmEntriesMap {

		gE, err := crmDB.GetRowGsmTableId(context.Background(), k)
		if err != nil {
			t.Errorf("%s != nil", err)
			continue
		}

		if !gsmEntriesEqual(v, gE) {
			t.Errorf("%v != %v", v, gE)
		}

	}

}

func TestGetRowGsmTableDtReceiving(t *testing.T) {
	dR := map[string]int{}
	for _, v := range gsmEntriesMap {
		dR[v.DtReceiving.Time.Format(time.DateOnly)]++
	}

	for k, v := range dR {
		tm, _ := time.Parse(time.DateOnly, k)
		gsmEntries, err := crmDB.GetRowsGsmTableDtReceiving(context.Background(), pgtype.Date{Time: tm, Valid: true})
		if err != nil {
			t.Errorf("%s != nil", err)
			continue
		}
		t.Log(len(gsmEntries), v)
		if len(gsmEntries) != v {
			t.Errorf("%d != %d", len(gsmEntries), v)
		}
	}
}

func TestDelRowGsmTable(t *testing.T) {
	for k := range gsmEntriesMap {
		b, err := crmDB.DelRowGsmTable(context.Background(), k)
		if err != nil {
			t.Errorf("%s != nil", err)

		}
		if !b {
			t.Errorf("%v != true", b)
		}
	}

	for k := range gsmEntriesMap {
		b, err := crmDB.DelRowGsmTable(context.Background(), k)
		if err != nil {
			t.Errorf("%s != nil", err)

		}
		if b {
			t.Errorf("%v != false", b)
		}
	}
}
