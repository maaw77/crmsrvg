package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	db "github.com/maaw77/crmsrvg/internal/database"
	"github.com/maaw77/crmsrvg/internal/models"
)

type UsersTable struct {
	storage *db.CrmDatabase
}

// getEntryGsm receives an entry with a specified date from the GSM table.
func (u *UsersTable) regUser(w http.ResponseWriter, r *http.Request) {
	log.Println("regUserServing:", r.URL.Path, "from", r.Host)
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	enc := json.NewEncoder(w)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		if err := enc.Encode(ErrorMessage{Details: err.Error()}); err != nil {
			log.Println(err)
		}
		return
	}

	user.Admin = false

	id, err := u.storage.AddUser(r.Context(), user)
	if errors.Is(err, db.ErrExists) {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		if err := enc.Encode(ErrorMessage{Details: fmt.Sprintf("user (Usrname=%s, ID=%d) already exists", user.Username, id.ID)}); err != nil {
			log.Println(err)
		}
		return
	} else if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	enc.Encode(id)
	log.Printf("%s is registered", user.Username)
}

// newGsmTable allocates and returns a new gsmTable.
func newUsersTable(srg *db.CrmDatabase) *UsersTable {
	return &UsersTable{storage: srg}
}
