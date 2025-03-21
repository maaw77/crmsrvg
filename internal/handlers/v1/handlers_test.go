package handlers

import (
	"context"
	"testing"
	"time"

	"github.com/maaw77/crmsrvg/config"
	"github.com/maaw77/crmsrvg/internal/database"
	"github.com/maaw77/crmsrvg/internal/models"
)

var (
	idUsers  = []int{}
	idGsmMap = map[int]models.GsmTableEntry{}
	crmDB    *database.CrmDatabase
)

func TestUsers(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var err error

	crmDB, err = database.NewCrmDatabase(ctx, config.InitConnString(""))
	if err != nil {
		t.Fatal(err)
	}
	if err := crmDB.DBpool.Ping(ctx); err != nil {
		t.Fatal(err)
	}
	defer crmDB.DBpool.Close()

	t.Run("RegUsers", subtRegUeser)
	t.Run("LoginUsers", subtLoginUsers)

	for _, id := range idUsers {
		// t.Log("id =", id)

		crmDB.DelUser(context.Background(), id)
	}
}

func TestGsm(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var err error

	crmDB, err = database.NewCrmDatabase(ctx, config.InitConnString(""))
	if err != nil {
		t.Fatal(err)
	}
	if err := crmDB.DBpool.Ping(ctx); err != nil {
		t.Fatal(err)
	}
	defer crmDB.DBpool.Close()

	t.Run("BadReq", subtAddEntryBadReq)
	t.Run("GoodReq", subtAddEntryGoodReq)
	t.Run("Update", subtUpdateEntryGsm)
	t.Run("GetId", subtGetGsmEntryId)

	for k := range idGsmMap {
		crmDB.DelRowGsmTable(context.Background(), k)
	}
}
