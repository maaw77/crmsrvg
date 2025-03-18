package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maaw77/crmsrvg/internal/models"
)

func subtAddEntryBadReq(t *testing.T) {
	gT := newGsmTable(crmDB)
	user := models.User{Username: "kjdlk", Password: "jldkdj"}

	payload, _ := json.Marshal(user)
	r := httptest.NewRequest("POST", "/gsm", bytes.NewReader(payload))
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(gT.addEntryGsm)
	handler(w, r)

	t.Log(w.Result().Status, w.Body.String())

}
