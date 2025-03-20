package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/maaw77/crmsrvg/internal/models"
)

func gsmEntriesEqual(a, b models.GsmTableEntry) bool {

	switch {
	case a.ID != b.ID:
		return false
	case a.DtReceiving != b.DtReceiving:
		return false
	case a.DtCrch != b.DtCrch:
		return false
	case a.Site != b.Site:
		return false
	case fmt.Sprintf("%.3f", a.IncomeKg) != fmt.Sprintf("%.3f", b.IncomeKg):
		return false
	case a.Operator != b.Operator:
		return false
	case a.Provider != b.Provider:
		return false
	case a.Contractor != b.Contractor:
		return false
	case a.LicensePlate != b.LicensePlate:
		return false
	case a.Status != b.Status:
		return false
	case a.BeenChanged != b.BeenChanged:
		return false
	case a.GUID != b.GUID:
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
		if k > maxId {
			maxId = k + 1
		}
	}
	_, err := crmDB.GetRowGsmTableId(context.Background(), maxId)
	if err != ErrNotExist {
		t.Logf("%s != %s", err, ErrNotExist)
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
			t.Errorf("%s != nil", err)
			continue
		}
		t.Log(len(gsmEntries), v)
		if len(gsmEntries) != v {
			t.Errorf("%d != %d", len(gsmEntries), v)
		}
	}
}

func subtDelRowGsmTable(t *testing.T) {
	for k := range gsmEntriesMap {
		b, err := crmDB.DelRowGsmTable(context.Background(), k)
		if err != nil {
			t.Errorf("%s != nil", err)

		}
		if !b {
			t.Errorf("%v != true", b)
		}
	}

	for k := range gsmEntriesMap {
		b, err := crmDB.DelRowGsmTable(context.Background(), k)
		if err != nil {
			t.Errorf("%s != nil", err)

		}
		if b {
			t.Errorf("%v != false", b)
		}
	}
}

func subtUpdateRowGsmTable(t *testing.T) {
	gsmE := models.GsmTableEntry{}
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

	id, err := crmDB.UpdateRowGsmTable(context.Background(), gsmE)
	// if !errors.Is(err, ErrNotExist) {
	// 	t.Errorf("%s != %s", err, ErrNotExist)
	// }

	t.Log(id, err)
}
