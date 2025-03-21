package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/maaw77/crmsrvg/internal/models"
)

func gsmEntriesEqual(a, b models.GsmTableEntry) bool {
	// fmt.Println(a, b)
	switch {
	case a.ID != b.ID:
		fmt.Println("a.ID != b.ID")
		return false
	case a.DtReceiving != b.DtReceiving:
		fmt.Println("!DtReceiving")
		return false
	case a.DtCrch != b.DtCrch:
		fmt.Println("!DtCrch")
		return false
	case a.Site != b.Site:
		return false
	case fmt.Sprintf("%.3f", a.IncomeKg) != fmt.Sprintf("%.3f", b.IncomeKg):
		fmt.Println("!IncomeKg")
		return false
	case a.Operator != b.Operator:
		fmt.Println("!Operator")
		return false
	case a.Provider != b.Provider:
		fmt.Println("!Provider")
		return false
	case a.Contractor != b.Contractor:
		fmt.Println("!Contractor")
		return false
	case a.LicensePlate != b.LicensePlate:
		fmt.Println("!LicensePlate")
		return false
	case a.Status != b.Status:
		fmt.Println("!Status")
		return false
	case a.BeenChanged != b.BeenChanged:
		fmt.Println("!BeenChanged")
		return false
	case a.GUID != b.GUID:
		fmt.Println("!GUID")
		return false
	}

	return true
}

func readData(fileName string) (gsmEntries []models.GsmTableEntry, err error) {
	// fmt.Println("###############################")
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer f.Close()

	dec := json.NewDecoder(f)

	for {
		gsmE := models.GsmTableEntry{}
		err = dec.Decode(&gsmE)

		if err == io.EOF {
			break
		} else if err != nil {
			return
		}

		// log.Println(gsmE)
		gsmEntries = append(gsmEntries, gsmE)
	}

	return gsmEntries, nil
}

func subtAddEntryBadReq(t *testing.T) {
	gT := newGsmTable(crmDB)
	user := models.User{Username: "kjdlk", Password: "jldkdj"}

	payload, _ := json.Marshal(user)
	r := httptest.NewRequest("POST", "/gsm", bytes.NewReader(payload))
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(gT.addEntryGsm)
	handler(w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusBadRequest)
	}

	gsmB := []byte(`{"dt_receiving": "2023-12-11",

				"dt_crch": "0001-01-01",

				"site": "Site_5",

				"income_kg": 720.9102379582451,

				"operator": "Operator_1",

				"provider": "Provider_3",

				"contractor": "Contractor_3",

				"license_plate": "LicensePlate_2",

				"status": "Status_1",

				"been_changed": false,

				"guid": "593ff941"}`)

	// gsmE := models.GsmTableEntry{}
	// json.Unmarshal(gsmB, &gsmE)
	// t.Log(gsmE)
	r = httptest.NewRequest("POST", "/gsm", bytes.NewReader(gsmB))
	w = httptest.NewRecorder()
	handler = http.HandlerFunc(gT.addEntryGsm)
	handler(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			w.Code, http.StatusBadRequest)
	}
	errD := ErrorMessage{}
	// t.Log(w.Body.String())
	json.NewDecoder(w.Body).Decode(&errD)
	// t.Log(errD)
	if errD.Details != "data validation error" {
		t.Errorf(`%s != "data validation error"`, errD.Details)
	}
}

func subtAddEntryGoodReq(t *testing.T) {
	gT := newGsmTable(crmDB)

	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	// t.Log(pwd)

	gsmEntriesArr, err := readData(pwd + "/testdata/gsm.data")
	if err != nil {
		t.Fatalf("%s != nil", err)
	}
	handler := http.HandlerFunc(gT.addEntryGsm)
	for _, gsmEntry := range gsmEntriesArr {
		b, _ := json.Marshal(gsmEntry)

		r := httptest.NewRequest("POST", "/gsm", bytes.NewReader(b))
		w := httptest.NewRecorder()
		handler(w, r)

		if w.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				w.Code, http.StatusOK)
		}

		var id models.IdEntry
		json.NewDecoder(w.Body).Decode(&id)
		gsmEntry.ID = id.ID
		idGsmMap[id.ID] = gsmEntry

	}

	// If the user already exist.
	for _, gsmEntry := range idGsmMap {
		b, _ := json.Marshal(gsmEntry)

		r := httptest.NewRequest("POST", "/gsm", bytes.NewReader(b))
		w := httptest.NewRecorder()
		handler(w, r)

		if w.Code != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusBadRequest)
		}

		errD := ErrorMessage{}
		// t.Log(w.Body.String())
		json.NewDecoder(w.Body).Decode(&errD)
		expected := fmt.Sprintf("entry (guid=%s, id=%d) already exists", gsmEntry.GUID, gsmEntry.ID)
		if errD.Details != expected {
			t.Errorf("%s != %s", errD.Details, expected)
		}
		// t.Log(errD)

	}
}

func subtGetGsmEntryId(t *testing.T) {
	gT := newGsmTable(crmDB)

	var maxId int
	for k := range idGsmMap {
		if k > maxId {
			maxId = k
		}
	}

	router := mux.NewRouter()
	router.HandleFunc("/id/{id:[0-9]+}", gT.getGsmEntryId).Methods("GET")

	// r := httptest.NewRequest("GET", fmt.Sprintf("/id/%d", maxId), nil)
	r := httptest.NewRequest("GET", "/id/0", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	if w.Code != http.StatusBadRequest && w.Body.String() != `{"details":"data validation error"}` {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusBadRequest)
	}

	r = httptest.NewRequest("GET", fmt.Sprintf("/id/%d", maxId+1), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, r)
	if w.Code != http.StatusBadRequest && w.Body.String() != `{"details":"it doesn't exist"}` {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusBadRequest)
	}

	for k, v := range idGsmMap {
		var inGsmEntry models.GsmTableEntry
		r = httptest.NewRequest("GET", fmt.Sprintf("/id/%d", k), nil)
		w = httptest.NewRecorder()
		router.ServeHTTP(w, r)
		if w.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
		}

		json.NewDecoder(w.Body).Decode(&inGsmEntry)
		if !gsmEntriesEqual(inGsmEntry, v) {
			t.Errorf("inGsmEntry[%d] != entry[%d]", inGsmEntry.ID, k)
		}

	}

}
