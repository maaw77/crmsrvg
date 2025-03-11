package handlers

import (
	"encoding/json"
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
	log.Printf("Context = %v\n", r.Context())
	id, err := u.storage.AddUser(r.Context(), user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	enc.Encode(id)
	log.Printf("%s is registered", user.Username)
}
