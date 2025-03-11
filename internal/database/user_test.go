package database

import (
	"context"
	"testing"

	"github.com/maaw77/crmsrvg/internal/models"
)

func TestUser(t *testing.T) {
	user := models.User{Username: "Vasya", Password: "abracadabra"}
	id, err := crmDB.AddUser(context.Background(), user)
	if err != nil {
		t.Errorf("%s != nil", err)
	}
	t.Logf("id = %d", id.ID)

	id, err = crmDB.AddUser(context.Background(), user)
	if err != ErrExists {
		t.Errorf("%s != %s", err, ErrExists)
	}

	userGet, err := crmDB.GetUser(context.Background(), user.Username)

	if err != nil {
		t.Errorf("%s != nil", err)
	}

	user.ID = id.ID

	if user != userGet {
		t.Errorf("%v != %v", user, userGet)
	}

	s, err := crmDB.DelUser(context.Background(), id.ID)

	if err != nil {
		t.Errorf("%s != nil", err)
	}

	if !s {
		t.Errorf("%v != true", s)
	}

	s, err = crmDB.DelUser(context.Background(), id.ID)

	if err != nil {
		t.Errorf("%s != nil", err)
	}

	if s {
		t.Errorf("%v != false", s)
	}
}
