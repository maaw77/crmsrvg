package database

import (
	"context"
	"errors"
	"testing"

	"github.com/maaw77/crmsrvg/config"
	"github.com/maaw77/crmsrvg/internal/models"
)

var (
	POSTGRESQL_URL     = config.InitConnString("")
	POSTGRESQL_URL_BAD = "postgres://postgres:crmpasswordocalhost:5433/postgres?sslmode=disable&pool_max_conns=10"

	// crmDB is the current instance of the CrmDatabase.
	crmDB         *CrmDatabase
	gsmEntriesMap = make(map[int]models.GsmEntryResponse)
)

func subtNewCrmDatabaseEmty(t *testing.T) {
	if _, err := NewCrmDatabase(context.Background(), ""); !errors.Is(err, ErrConStrEmty) {
		t.Fatalf("%s != %s", err, ErrConStrEmty)
	}

}

func subtPoolDbPing(t *testing.T) {
	var err error

	crmDB, err = NewCrmDatabase(context.Background(), POSTGRESQL_URL_BAD)
	if err == nil {
		t.Fatalf("%s == nil", err)
	}

	crmDB, err = NewCrmDatabase(context.Background(), POSTGRESQL_URL)
	if err != nil {
		t.Fatalf("%s != nil", err)
	}
	if err := crmDB.DBpool.Ping(context.Background()); err != nil {
		t.Fatalf("%s != nil", err)
	}
	crmDB.DBpool.Close()
	t.Logf("%#v\n", crmDB.DBpool.Config().ConnString())
}

func TestStorage(t *testing.T) {
	t.Run("EmtyConnStr", subtNewCrmDatabaseEmty)
	t.Run("ping", subtPoolDbPing)
}

func TestAuxTables(t *testing.T) {
	var err error
	crmDB, err = NewCrmDatabase(context.Background(), POSTGRESQL_URL)
	if err != nil {
		t.Fatalf("%s != nil", err)
	}
	defer crmDB.DBpool.Close()

	if err := crmDB.DBpool.Ping(context.Background()); err != nil {
		t.Fatalf("%s != nil", err)
	}

	t.Run("Sites", subtAuxilTableSites)
	t.Run("Statuses", subtAuxilTableStatuses)
	t.Run("Contractors", subtAuxilTableContractors)
	t.Run("LicensePlates", subtAuxilTableLicensePlates)
	t.Run("Operators", subtAuxilTableOperators)
	t.Run("Providers", subtAuxilTableProviders)
}

func TestGsmTable(t *testing.T) {
	var err error
	crmDB, err = NewCrmDatabase(context.Background(), POSTGRESQL_URL)
	if err != nil {
		t.Fatalf("%s != nil", err)
	}
	defer crmDB.DBpool.Close()

	if err := crmDB.DBpool.Ping(context.Background()); err != nil {
		t.Fatalf("%s != nil", err)
	}

	t.Run("InsertRow", subtInserGsmTable)
	t.Run("GetRowId", subtGetRowGsmTableId)
	t.Run("GetRowDtRec", subtGetRowGsmTableDtReceiving)
	t.Run("UpdateRow", subtUpdateRowGsmTable)
	t.Run("DelRow", subtDelRowGsmTable)

	gsmEntriesMap = make(map[int]models.GsmEntryResponse)

	// crmDB.DBpool.Exec(context.Background(), "DELETE FROM gsm_table;")

}
func BenchmarkGsm(b *testing.B) {
	var err error

	crmDB, err = NewCrmDatabase(context.Background(), POSTGRESQL_URL)
	if err != nil {
		b.Fatalf("%s != nil", err)
	}
	defer crmDB.DBpool.Close()

	if err := crmDB.DBpool.Ping(context.Background()); err != nil {
		b.Fatalf("%s != nil", err)

	}

	b.Run("Update", subBenchmarkGetIdOrCreateSites)

}

func TestUsersTable(t *testing.T) {
	var err error
	crmDB, err = NewCrmDatabase(context.Background(), POSTGRESQL_URL)
	if err != nil {
		t.Fatalf("%s != nil", err)
	}
	defer crmDB.DBpool.Close()

	if err := crmDB.DBpool.Ping(context.Background()); err != nil {
		t.Fatalf("%s != nil", err)
	}
	t.Run("Users", subtUserTable)
}
