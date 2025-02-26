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
	entries := map[string]models.IdEntry{}

}
