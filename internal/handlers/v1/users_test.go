package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/maaw77/crmsrvg/config"
	"github.com/maaw77/crmsrvg/internal/database"

	// "github.com/maaw77/crmsrvg/internal/handlers/v1"
	"github.com/maaw77/crmsrvg/internal/models"
)

func TestRegUeser(t *testing.T) {
	idUsers := []int{}
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
	uT := newUsersTable(crmDB)

	uesers := []models.User{{Username: "User_1", Password: "Password_1"},
		{Username: "User_2", Password: "Password_2", Admin: true},
	}
	// t.Log(uesers)
	for _, v := range uesers {
		b, err := json.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		// t.Logf("%s", b)
		r := httptest.NewRequest("POST", "/reguser", bytes.NewReader(b))
		w := httptest.NewRecorder()
		handlers := http.HandlerFunc(uT.regUser)
		handlers.ServeHTTP(w, r)

		// Check the HTTP status code is what we expect.
		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
			return
		}

		id := models.IdEntry{}
		json.NewDecoder(w.Body).Decode(&id)
		idUsers = append(idUsers, id.ID)
		// t.Log("idUsers =", idUsers)

	}

	for i, v := range uesers {
		b, err := json.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		// t.Logf("%s", b)
		r := httptest.NewRequest("POST", "/reguser", bytes.NewReader(b))
		w := httptest.NewRecorder()
		handlers := http.HandlerFunc(uT.regUser)
		handlers.ServeHTTP(w, r)

		// Check the HTTP status code is what we expect.
		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
			return
		}

		errD := ErrorMessage{}
		// t.Log(w.Body.String())
		json.NewDecoder(w.Body).Decode(&errD)
		expected := fmt.Sprintf("user (Usrname=%s, ID=%d) already exists", v.Username, idUsers[i])
		if errD.Details != expected {
			t.Errorf("%s != %s", errD.Details, expected)
		}
		// t.Log("idUsers =", idUsers)

	}

	for _, id := range idUsers {
		// t.Log("id =", id)

		crmDB.DelUser(context.Background(), id)
	}
}
