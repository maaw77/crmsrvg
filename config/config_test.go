package config

import (
	"os"
	"testing"
)

func TestInitConnStringNonPath(t *testing.T) {
	waitConnSting := "postgres://postgres:crmpassword@localhost:5433/postgres?sslmode=disable&pool_max_conns=10"

	connString := InitConnString("")
	if connString != waitConnSting {
		t.Errorf("%s != %s", connString, waitConnSting)
	}
}

func TestInitConnStringWithPath(t *testing.T) {
	waitConnSting := "postgres://postgres:crmpassword@localhost:5433/postgres?sslmode=disable&pool_max_conns=7"
	currDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	connString := InitConnString(currDir + "/config.yaml")
	if connString != waitConnSting {
		t.Errorf("%s != %s", connString, waitConnSting)
	}
}
