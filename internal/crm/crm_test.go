package crm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/maaw77/crmsrvg/config"
	"github.com/maaw77/crmsrvg/internal/database"
	"github.com/maaw77/crmsrvg/internal/handlers/v1"
	"github.com/maaw77/crmsrvg/internal/models"
)

func TestCrm(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	crmDB, err := database.NewCrmDatabase(ctx, config.InitConnString(""))
	if err != nil {
		log.Fatalln(err)
	}
	if err := crmDB.DBpool.Ping(ctx); err != nil {
		log.Fatal(err)
	}
	defer crmDB.DBpool.Close()

	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(handlers.DefaultHandler)
	r.MethodNotAllowedHandler = http.HandlerFunc(handlers.MethodNotAllowed)

	apiR := r.PathPrefix("/api/v1").Subrouter()

	// Uesrs
	handlers.RegUsersHanlders(apiR, crmDB)
	// GSM
	handlers.RegGsmHanlders(apiR, crmDB)

	ts := httptest.NewServer(r)
	defer ts.Close()

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	// First request.
	payload := []byte(`{"username":"Kolya", "password":"something"}`)
	req, err := http.NewRequest(http.MethodPost, ts.URL+"/api/v1"+"/users", bytes.NewReader(payload))
	if err != nil {
		t.Fatal("Error:", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Error:", err)
	}

	if status := resp.StatusCode; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
	idU := models.IdEntry{}
	json.NewDecoder(resp.Body).Decode(&idU)
	t.Log("idUser:", idU)
	resp.Body.Close()

	// Next request.
	payload = []byte(`{"username":"Kolya", "password":"something"}`)
	req, err = http.NewRequest(http.MethodPost, ts.URL+"/api/v1"+"/users/login", bytes.NewReader(payload))
	if err != nil {
		t.Fatal("Error:", err)
	}

	resp, err = client.Do(req)
	if err != nil {
		t.Fatal("Error:", err)
	}

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
	token := handlers.AccessToken{}
	json.NewDecoder(resp.Body).Decode(&token)
	resp.Body.Close()

	// Next request.
	payload = []byte(`{"dt_receiving": "2024-02-23",

					"dt_crch": "0001-01-01",

					"site": "Site_4",

					"income_kg": 239.33,

					"operator": "Operator_4",

					"provider": "Provider_3",

					"contractor": "Contractor_3",

					"license_plate": "LicensePlate_5",

					"status": "Status_2",

					"been_changed": false,

					"guid": "96935cf8-06ee-421a-af26-8f720eb9bc39"
					}`)

	req, err = http.NewRequest(http.MethodPost, ts.URL+"/api/v1"+"/gsm", bytes.NewReader(payload))
	if err != nil {
		t.Fatal("Error:", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))

	resp, err = client.Do(req)
	if err != nil {
		t.Fatal("Error:", err)
	}
	if status := resp.StatusCode; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	idE := models.IdEntry{}
	json.NewDecoder(resp.Body).Decode(&idE)
	resp.Body.Close()
	t.Log("idEntry:", idE)

	// Next request.
	req, err = http.NewRequest(http.MethodGet, ts.URL+"/api/v1"+"/gsm/id/"+strconv.Itoa(idE.ID), nil)
	if err != nil {
		t.Fatal("Error:", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))

	resp, err = client.Do(req)
	if err != nil {
		t.Fatal("Error:", err)
	}
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	gsmE := models.GsmEntryResponse{}
	json.NewDecoder(resp.Body).Decode(&gsmE)
	if gsmE.GUID != "96935cf8-06ee-421a-af26-8f720eb9bc39" {
		t.Error("Invalid GUID")
	}
	t.Log("gsmE= ", gsmE)
	resp.Body.Close()

	// Next request.
	req, err = http.NewRequest(http.MethodGet, ts.URL+"/api/v1"+"/gsm/date/"+"2024-02-23", nil)
	if err != nil {
		t.Fatal("Error:", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))

	resp, err = client.Do(req)
	if err != nil {
		t.Fatal("Error:", err)
	}
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	gsmEntries := []models.GsmEntryResponse{}
	json.NewDecoder(resp.Body).Decode(&gsmEntries)
	guid := gsmEntries[0].GUID
	if guid != "96935cf8-06ee-421a-af26-8f720eb9bc39" {
		t.Error("Invalid GUID")
	}
	t.Log("gsmEntries= ", gsmEntries)
	resp.Body.Close()

	// Next request.
	req, err = http.NewRequest(http.MethodDelete, ts.URL+"/api/v1"+"/gsm/id/"+strconv.Itoa(idE.ID), nil)
	if err != nil {
		t.Fatal("Error:", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))

	resp, err = client.Do(req)
	if err != nil {
		t.Fatal("Error:", err)
	}
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	idE = models.IdEntry{}
	json.NewDecoder(resp.Body).Decode(&idE)
	resp.Body.Close()
	t.Log("Del idEntry:", idE)
	resp.Body.Close()

	// Cleanup.
	crmDB.DelUser(context.Background(), idU.ID)
	// crmDB.DelRowGsmTable(context.Background(), idE.ID)
}
