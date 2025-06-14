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
	Token string `json:"token" validate:"required" example:"Some kind of JWT"`
}

type UsersTable struct {
	storage  *db.CrmDatabase
	validate *validator.Validate
}

// regUser creates a new user.
//
//	@Summary		Create a user
//	@Description	Create a new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.UserRequest	true	"User data"
//	@Success		201		{object}	models.IdEntry
//	@Failure		400		{object}	ErrorMessage
//	@Failure		500		{object}	ErrorMessage
//	@Router			/users [post]
func (u *UsersTable) regUser(w http.ResponseWriter, r *http.Request) {
	log.Println("regUser Serving:", r.URL.Path, "from", r.Host)
	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)

	var user models.UserResponse
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		if err := enc.Encode(ErrorMessage{Details: err.Error()}); err != nil {
			log.Println(err)
		}
		return
	}

	if err := u.validate.Struct(&user); err != nil {
		log.Println("error: ", err)
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
		log.Println("error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		if err := enc.Encode(ErrorMessage{Details: fmt.Sprintf("user (username=%s, id=%d) already exists", user.Username, id.ID)}); err != nil {
			log.Println(err)
		}
		return
	} else if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		if err := enc.Encode(ErrorMessage{Details: "something happened to the server"}); err != nil {
			log.Println("error: ", err)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := enc.Encode(id); err != nil {
		log.Println("error: ", err)
	}
	log.Printf("%s is registered", user.Username)
}

// loginUser identifies and authenticates the user. If successful, it returns an access token.
//
//	@Summary		User identification and authentication
//	@Description	User identification and authentication. If successful, it returns an access token.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.UserRequest	true	"User data"
//	@Success		200		{object}	AccessToken
//	@Failure		400		{object}	ErrorMessage
//	@Failure		404		{object}	ErrorMessage
//	@Failure		401		{object}	ErrorMessage
//	@Failure		500		{object}	ErrorMessage
//	@Router			/users/login [post]
func (u *UsersTable) loginUser(w http.ResponseWriter, r *http.Request) {
	log.Println("loginUser Serving:", r.URL.Path, "from", r.Host)
	w.Header().Set("Content-Type", "application/json")

	var user models.UserResponse
	enc := json.NewEncoder(w)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("error:", err)
		w.WriteHeader(http.StatusBadRequest)
		if err := enc.Encode(ErrorMessage{Details: err.Error()}); err != nil {
			log.Println("error: ", err)
		}
		return
	}

	if err := u.validate.Struct(&user); err != nil {
		log.Println("error:", err)
		w.WriteHeader(http.StatusBadRequest)
		if err := enc.Encode(ErrorMessage{Details: err.Error()}); err != nil {
			log.Println("error: ", err)
		}
		return
	}

	userReg, err := u.storage.GetUser(r.Context(), user.Username)
	switch {
	case errors.Is(err, db.ErrNotExist):
		w.WriteHeader(http.StatusNotFound)
		log.Printf("error: user %s was not found", user.Username)
		if err := enc.Encode(ErrorMessage{Details: fmt.Sprintf("user %s was not found", user.Username)}); err != nil {
			log.Println("error: ", err)
		}
		return
	case err != nil:
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error:", err)
		if err := enc.Encode(ErrorMessage{Details: err.Error()}); err != nil {
			log.Println("error: ", err)
		}
		return
	case userReg.Password != user.Password: // incorrect password entered
		w.WriteHeader(http.StatusUnauthorized)
		log.Printf("error: the password for %s was entered incorrectly", user.Username)
		if err := enc.Encode(ErrorMessage{Details: "incorrect password entered"}); err != nil {
			log.Println("error: ", err)
		}
		return
	}

	tokenA, err := auth.GetToken(userReg.Username, time.Now().Add(time.Hour).Unix())
	if err != nil {
		log.Print("error:", err)
	}
	w.WriteHeader(http.StatusOK)
	if err := enc.Encode(AccessToken{Token: tokenA}); err != nil {
		log.Print("error:", err)

	}
}

// newGsmTable allocates and returns a new gsmTable.
func newUsersTable(srg *db.CrmDatabase) *UsersTable {
	return &UsersTable{storage: srg,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}
