package database

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/maaw77/crmsrvg/internal/models"
)

func gsmEntriesEqual(a, b models.GsmEntryResponse) bool {

	switch {
	case a.ID != b.ID:
		log.Println("!ID")
		return false
	case a.DtReceiving.Time.Format(time.DateOnly) != b.DtReceiving.Time.Format(time.DateOnly):
		log.Println("!DtReceiving")

		return false
	case a.DtCrch.Time.Format(time.DateOnly) != b.DtCrch.Time.Format(time.DateOnly):
		log.Println("!DtCrch")

		return false
	case a.Site != b.Site:
		log.Println("!Site")

		return false
	case (a.IncomeKg - b.IncomeKg) > 1:
		log.Println("!IncomeKg")

		return false
	case a.Operator != b.Operator:
		log.Println("!Operator")

		return false
	case a.Provider != b.Provider:
		log.Println("!Provider")

		return false
	case a.Contractor != b.Contractor:
		log.Println("!Contractor")

		return false
	case a.LicensePlate != b.LicensePlate:
		log.Println("!.LicensePlate")

		return false
	case a.Status != b.Status:
		log.Println("!Status")

		return false
	case a.BeenChanged != b.BeenChanged:
		log.Println("!BeenChanged")

		return false
	case a.GUID != b.GUID:
		log.Println("!GUID")

		return false
	}

	return true
}

func changGsmEnry(gsmE *models.GsmEntryResponse) {
	gsmE.DtReceiving = pgtype.Date{Time: time.Now(), Valid: true}
	gsmE.DtCrch = pgtype.Date{Time: time.Now(), Valid: true}
	gsmE.Site = gsmE.Site + "100"
	gsmE.IncomeKg *= 1000
	gsmE.BeenChanged = !(gsmE.BeenChanged)
}

func readData(fileName string) (gsmEntries []models.GsmEntryResponse, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer f.Close()

	dec := json.NewDecoder(f)

	for {
		gsmE := models.GsmEntryResponse{}
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
func subtInserGsmTable(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	// t.Log(pwd)

	gsmEntriesArr, err := readData(pwd + "/testdata/gsm.data")
	if err != nil {
		t.Fatalf("%s != nil", err)
	}

	// t.Log(gsmEntriesEqual(gsmEntriesArr[0], gsmEntriesArr[2]))

	for _, gE := range gsmEntriesArr {
		id, err := crmDB.InsertGsmTable(context.Background(), gE)
		if err != nil {
			t.Errorf("%s != nil", err)
			continue
		}
		gE.ID = id.ID
		gsmEntriesMap[id.ID] = gE

	}
	// t.Log(gsmEntriesMap)

	for k, v := range gsmEntriesMap {

		id, err := crmDB.InsertGsmTable(context.Background(), v)
		if k != id.ID {
			t.Errorf("%d != %d", k, id.ID)
		}
		if !errors.Is(err, ErrExist) {
			t.Errorf("%v != %v", err, ErrExist)
		}

	}
	// t.Log(gsmEntriesMap)
}

func subtGetRowGsmTableId(t *testing.T) {
	var maxId int
	for k, v := range gsmEntriesMap {

		gE, err := crmDB.GetRowGsmTableId(context.Background(), k)
		if err != nil {
			t.Errorf("%s != nil", err)
			continue
		}

		if !gsmEntriesEqual(v, gE) {
			t.Errorf("%v != %v", v, gE)
		}
		if k >= maxId {
			maxId = k + 1
		}
	}

	t.Logf("maxId= %d", maxId)
	_, err := crmDB.GetRowGsmTableId(context.Background(), maxId)
	if err != ErrNotExist {
		t.Errorf("%s != %s", err, ErrNotExist)
	}

}

func subtGetRowGsmTableDtReceiving(t *testing.T) {
	dR := map[string]int{}
	for _, v := range gsmEntriesMap {
		dR[v.DtReceiving.Time.Format(time.DateOnly)]++
	}

	for k, v := range dR {
		tm, _ := time.Parse(time.DateOnly, k)
		gsmEntries, err := crmDB.GetRowsGsmTableDtReceiving(context.Background(), pgtype.Date{Time: tm, Valid: true})
		if err != nil {
			t.Errorf(`the received error message "%s" != nil`, err)
			continue
		}
		t.Log(len(gsmEntries), v)
		if len(gsmEntries) != v {
			t.Errorf("%d != %d", len(gsmEntries), v)
		}
	}

	_, err := crmDB.GetRowsGsmTableDtReceiving(context.Background(), pgtype.Date{Time: time.Now(), Valid: true})
	if err != ErrNotExist {
		t.Errorf(`the received error message "%s" != %s`, err, ErrNotExist)

	}

}

func subtDelRowGsmTable(t *testing.T) {
	for k := range gsmEntriesMap {
		err := crmDB.DelRowGsmTable(context.Background(), k)
		if err != nil {
			t.Errorf(`the received error message "%s" != nil`, err)
		}

	}

	for k := range gsmEntriesMap {
		err := crmDB.DelRowGsmTable(context.Background(), k)
		if err != ErrNotExist {
			t.Errorf(`the received error message "%s" != %s`, err, ErrNotExist)
		}

	}
}

func subtUpdateRowGsmTable(t *testing.T) {
	gsmE := models.GsmEntryResponse{}
	b := []byte(`{"dt_receiving": "2023-12-11",

				"dt_crch": "0001-01-01",

				"site": "Site_5",

				"income_kg": 720.9102379582451,

				"operator": "Operator_1",

				"provider": "Provider_3",

				"contractor": "Contractor_3",

				"license_plate": "LicensePlate_2",

				"status": "Status_1",

				"been_changed": false,

				"guid": "593ff941-405e-4afd-9eec-f99999999999999"}`)
	json.Unmarshal(b, &gsmE)

	_, err := crmDB.UpdateRowGsmTable(context.Background(), gsmE)
	if !errors.Is(err, ErrNotExist) {
		t.Errorf("%s != %s", err, ErrNotExist)
	}

	for k, v := range gsmEntriesMap {
		changGsmEnry(&v) // Changes v=gsmEntry.
		gsmEntriesMap[k] = v
		idOut, err := crmDB.UpdateRowGsmTable(context.Background(), v)
		if err != nil {
			t.Errorf("%s != nil", err)
		}
		if idOut.ID != k {
			t.Errorf("%d != %d", idOut, k)
		}
	}

	for k, v := range gsmEntriesMap {

		gE, err := crmDB.GetRowGsmTableId(context.Background(), k)
		if err != nil {
			t.Errorf("%s != nil", err)
			continue
		}

		if !gsmEntriesEqual(v, gE) {
			t.Errorf("%v != %v", gE, v)
		}
	}
}
