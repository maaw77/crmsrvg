package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	// "github.com/maaw77/crmsrvg/internal/handlers/v1"
	"github.com/maaw77/crmsrvg/internal/auth"
	"github.com/maaw77/crmsrvg/internal/models"
)

func subtRegUeser(t *testing.T) {

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

	// If the user already exist.
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

	// If the username and/or password is not specified(empty).
	uesers = []models.User{{Username: "User_3", Password: ""},
		{Username: "", Password: "Password_dedde", Admin: true},
	}
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
		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
			return
		}

		errD := ErrorMessage{}
		// t.Log(w.Body.String())
		json.NewDecoder(w.Body).Decode(&errD)
		t.Log(errD)
	}
}

func subtLoginUsers(t *testing.T) {
	uT := newUsersTable(crmDB)

	// If everything is ok.
	uesers := []models.User{{Username: "User_1", Password: "Password_1"},
		{Username: "User_2", Password: "Password_2", Admin: true},
	}
	for _, v := range uesers {
		b, err := json.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		// t.Logf("%s", b)
		r := httptest.NewRequest("POST", "/login", bytes.NewReader(b))
		w := httptest.NewRecorder()
		handlers := http.HandlerFunc(uT.loginUser)
		handlers.ServeHTTP(w, r)

		// Check the HTTP status code is what we expect.
		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
			return
		}

		tkn := AccessToken{}
		json.NewDecoder(w.Body).Decode(&tkn)

		token, err := auth.VerifyToken(tkn.Token)
		if err != nil {
			t.Fatalf("%s != nil", err)
		}

		if !token.Valid {
			t.Fatal("token.Valid != true")
		}

		username, err := token.Claims.GetSubject()
		// t.Logf("username=%s", username)
		if err != nil {
			t.Fatalf("%s != nil", err)
		}

		if v.Username != username {
			t.Errorf("%s != %s", v.Username, username)

		}

		expectedIss := "crm-app"
		iss, err := token.Claims.GetIssuer()
		if err != nil {
			t.Fatalf("%s != nil", err)
		}

		if expectedIss != iss {
			t.Errorf("%s != %s", expectedIss, iss)
		}

	}

	// If the user is not registered.
	uesers = []models.User{{Username: "User_3", Password: "Password_3"},
		{Username: "User_4", Password: "Password_4", Admin: true},
	}
	for _, v := range uesers {
		b, err := json.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		// t.Logf("%s", b)
		r := httptest.NewRequest("POST", "/login", bytes.NewReader(b))
		w := httptest.NewRecorder()
		handlers := http.HandlerFunc(uT.loginUser)
		handlers.ServeHTTP(w, r)

		// Check the HTTP status code is what we expect.
		if status := w.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
			return
		}
		errorAnswer := ErrorMessage{}
		json.NewDecoder(w.Body).Decode(&errorAnswer)
		expected := fmt.Sprintf("user %s was not found", v.Username)
		if expected != errorAnswer.Details {
			t.Errorf("%s != %s", expected, errorAnswer.Details)
		}

	}

	// If an incorrect password is entered
	uesers = []models.User{{Username: "User_1", Password: "Password_8"},
		{Username: "User_2", Password: "Password_", Admin: true},
	}

	for _, v := range uesers {
		b, err := json.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		// t.Logf("%s", b)
		r := httptest.NewRequest("POST", "/login", bytes.NewReader(b))
		w := httptest.NewRecorder()
		handlers := http.HandlerFunc(uT.loginUser)
		handlers.ServeHTTP(w, r)

		// Check the HTTP status code is what we expect.
		if status := w.Code; status != http.StatusUnauthorized {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
			return
		}
		errorAnswer := ErrorMessage{}
		json.NewDecoder(w.Body).Decode(&errorAnswer)

		if errorAnswer.Details != "incorrect password entered" {
			t.Errorf(`"incorrect password entered" != %s`, errorAnswer.Details)
		}

	}

	// If the username and/or password is not specified(empty).
	uesers = []models.User{{Username: "User_3", Password: ""},
		{Username: "", Password: "Password_dedde", Admin: true},
	}
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
		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
			return
		}

		errD := ErrorMessage{}
		// t.Log(w.Body.String())
		json.NewDecoder(w.Body).Decode(&errD)
		t.Log(errD)
	}

}
