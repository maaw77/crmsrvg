package handlers

import (
	"context"
	"testing"
	"time"

	"github.com/maaw77/crmsrvg/config"
	"github.com/maaw77/crmsrvg/internal/database"
	"github.com/maaw77/crmsrvg/internal/models"
)

func TestRegUeser(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	crmDB, err := database.NewCrmDatabase(ctx, config.InitConnString(""))
	if err != nil {
		t.Fatal(err)
	}
	if err := crmDB.DBpool.Ping(ctx); err != nil {
		t.Fatal(err)
	}
	defer crmDB.DBpool.Close()

	uesers := []models.User{{Username: "User_1", Password: "Password_1"},
		{Username: "User_2", Password: "Password_2", Admin: true},
	}

	t.Log(uesers)
}
