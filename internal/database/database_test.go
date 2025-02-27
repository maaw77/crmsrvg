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

		switch {
		case err != nil:
			t.Logf("GetIdOrCreateSites-> %s != nil", err)
		case id.ID != v.ID:
			t.Logf("GetIdOrCreateSites-> %d != %d", id.ID, v.ID)
		}

	}

	for k, v := range entries {
		s, err := crmDB.DelRowSites(context.Background(), v.ID)

		switch {
		case err != nil:
			t.Logf("DelRowSites-> %s != nil. For %s", err, k)
		case s == false:
			t.Logf("DelRowSites-> %v != true. For %s", s, k)
		}

	}

	for k, v := range entries {
		s, err := crmDB.DelRowSites(context.Background(), v.ID)

		switch {
		case err != nil:
			t.Logf("DelRowSites-> %s != nil. For %s", err, k)
		case s == true:
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

		switch {
		case err != nil:
			t.Logf("GetIdOrCreateStatuses-> %s != nil", err)
		case id.ID != v.ID:
			t.Logf("GetIdOrCreateStatuses-> %d != %d", id.ID, v.ID)
		}

	}

	for k, v := range entries {
		s, err := crmDB.DelRowStatuses(context.Background(), v.ID)

		switch {
		case err != nil:
			t.Logf("DelRowStatuses-> %s != nil. For %s", err, k)
		case s == false:
			t.Logf("DelRowSites-> %v != true. For %s", s, k)
		}

	}

	for k, v := range entries {
		s, err := crmDB.DelRowStatuses(context.Background(), v.ID)

		switch {
		case err != nil:
			t.Logf("DelRowStatuses-> %s != nil. For %s", err, k)
		case s == true:
			t.Logf("DelRowStatuses-> %v != false. For %s", s, k)
		}

	}
}
