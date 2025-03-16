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
	for k, v := range gsmEntriesMap {

		gE, err := crmDB.GetRowGsmTableId(context.Background(), k)
		if err != nil {
			t.Errorf("%s != nil", err)
			continue
		}

		if !gsmEntriesEqual(v, gE) {
			t.Errorf("%v != %v", v, gE)
		}

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
