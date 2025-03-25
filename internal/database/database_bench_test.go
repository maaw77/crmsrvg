package database

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/maaw77/crmsrvg/internal/models"
)

func subBenchmarkGetIdOrCreateSites(b *testing.B) {
	gsmE := models.GsmEntryResponse{}
	payload := []byte(`{"dt_receiving": "2023-12-11",

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
	json.Unmarshal(payload, &gsmE)

	for b.Loop() {
		crmDB.UpdateRowGsmTable(context.Background(), gsmE)
	}

}
