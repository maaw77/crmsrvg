package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/maaw77/crmsrvg/internal/auth"
	db "github.com/maaw77/crmsrvg/internal/database"
	"github.com/maaw77/crmsrvg/internal/models"
)

// AccessToken is a string that the OAuth client uses to make requests
// to the resource server.(https://oauth.net/2/access-tokens/).
type AccessToken struct {
	Token string `json:"token"`
}

type UsersTable struct {
	storage  *db.CrmDatabase
	validate *validator.Validate
}

// getEntryGsm receives an entry with a specified date from the GSM table.
func (u *UsersTable) regUser(w http.ResponseWriter, r *http.Request) {
	log.Println("regUser Serving:", r.URL.Path, "from", r.Host)
	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("error= ", err)
		w.WriteHeader(http.StatusBadRequest)
		if err := enc.Encode(ErrorMessage{Details: err.Error()}); err != nil {
			log.Println(err)
		}
		return
	}

	if err := u.validate.Struct(&user); err != nil {
		log.Println("error= ", err)
		w.WriteHeader(http.StatusBadRequest)
		if err := enc.Encode(ErrorMessage{Details: err.Error()}); err != nil {
			log.Println(err)
		}
		return
	}

	// There should always be a false here.
	user.Admin = false

	id, err := u.storage.AddUser(r.Context(), user)
	if errors.Is(err, db.ErrExist) {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		if err := enc.Encode(ErrorMessage{Details: fmt.Sprintf("user (Usrname=%s, ID=%d) already exists", user.Username, id.ID)}); err != nil {
			log.Println(err)
		}
		return
	} else if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		if err := enc.Encode(ErrorMessage{Details: "something happened to the server"}); err != nil {
			log.Println(err)
		}
		return
	}

	enc.Encode(id)
	log.Printf("%s is registered", user.Username)
}

// loginUser identifies and authenticates the user. If successful, it returns an access token.
func (u *UsersTable) loginUser(w http.ResponseWriter, r *http.Request) {
	log.Println("loginUser Serving:", r.URL.Path, "from", r.Host)
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	enc := json.NewEncoder(w)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("error = ", err)
		w.WriteHeader(http.StatusBadRequest)
		if err := enc.Encode(ErrorMessage{Details: err.Error()}); err != nil {
			log.Println(err)
		}
		return
	}

	if err := u.validate.Struct(&user); err != nil {
		log.Println("error = ", err)
		w.WriteHeader(http.StatusBadRequest)
		if err := enc.Encode(ErrorMessage{Details: err.Error()}); err != nil {
			log.Println(err)
		}
		return
	}

	userReg, err := u.storage.GetUser(r.Context(), user.Username)
	switch {
	case errors.Is(err, db.ErrNotExist):
		w.WriteHeader(http.StatusNotFound)
		log.Printf("user %s was not found", user.Username)
		if err := enc.Encode(ErrorMessage{Details: fmt.Sprintf("user %s was not found", user.Username)}); err != nil {
			log.Println(err)
		}
		return
	case err != nil:
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		if err := enc.Encode(ErrorMessage{Details: err.Error()}); err != nil {
			log.Println(err)
		}
		return
	case userReg.Password != user.Password: // incorrect password entered
		w.WriteHeader(http.StatusUnauthorized)
		log.Printf("the password for %s was entered incorrectly", user.Username)
		if err := enc.Encode(ErrorMessage{Details: "incorrect password entered"}); err != nil {
			log.Println(err)
		}
		return
	}

	tokenA, err := auth.GetToken(userReg.Username, time.Now().Add(time.Hour).Unix())
	if err != nil {
		log.Print(err)
	}
	enc.Encode(AccessToken{Token: tokenA})
}

// newGsmTable allocates and returns a new gsmTable.
func newUsersTable(srg *db.CrmDatabase) *UsersTable {
	return &UsersTable{storage: srg,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}
